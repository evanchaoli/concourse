package creds

import (
	"context"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagerctx"
	"encoding/json"
)

type VarSourcePool interface {
	FindOrCreate(lager.Logger, map[string]interface{}, ManagerFactory) (Secrets, error)
}

type inPoolManager struct {
	manager        Manager
	secretsFactory SecretsFactory
	lastUseTime    time.Time
}

func (m *inPoolManager) Close(logger lager.Logger) {
	m.manager.Close(logger)
}

func (m *inPoolManager) NewSecrets() Secrets {
	m.lastUseTime = time.Now()
	return m.secretsFactory.NewSecrets()
}

type varSourcePool struct {
	pool map[string]*inPoolManager
	lock sync.Mutex
	ttl  time.Duration
}

func (pool *varSourcePool) FindOrCreate(logger lager.Logger, config map[string]interface{}, factory ManagerFactory) (Secrets, error) {
	b, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	key := string(b)

	pool.lock.Lock()
	defer pool.lock.Unlock()

	if pool.pool[key] == nil {
		manager, err := factory.NewInstance(config)
		if err != nil {
			return nil, err
		}
		err = manager.Init(logger)
		if err != nil {
			return nil, err
		}
		secretsFactory, err := manager.NewSecretsFactory(logger)
		if err != nil {
			return nil, err
		}

		pool.pool[key] = &inPoolManager{
			manager:        manager,
			secretsFactory: secretsFactory,
			lastUseTime:    time.Now(),
		}
	} else {
		logger.Debug("found-existing-credential-manager")
	}

	return pool.pool[key].NewSecrets(), nil
}

func (pool *varSourcePool) Run(ctx context.Context) error {
	logger := lagerctx.FromContext(ctx).Session("var-source-reaper")

	pool.lock.Lock()
	defer pool.lock.Unlock()

	logger.Debug("start", lager.Data{"count": len(pool.pool)})
	defer func() {
		logger.Debug("done", lager.Data{"count": len(pool.pool)})
	}()

	toDeleteKeys := []string{}

	for key, manager := range pool.pool {
		if manager.lastUseTime.Add(pool.ttl).Before(time.Now()) {
			toDeleteKeys = append(toDeleteKeys, key)
			manager.Close(logger)
		}
	}

	for _, key := range toDeleteKeys {
		delete(pool.pool, key)
	}

	return nil
}

var pool = varSourcePool{
	pool: map[string]*inPoolManager{},
	lock: sync.Mutex{},
	ttl:  5 * time.Minute,
}

func VarSourcePoolInstance() *varSourcePool {
	return &pool
}
