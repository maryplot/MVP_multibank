package storage

import (
	"sync"
)

// BalanceStorage хранит изменения балансов счетов
type BalanceStorage struct {
	balances map[string]float64 // accountID -> balance change
	mutex    sync.RWMutex
}

var instance *BalanceStorage
var once sync.Once

// GetInstance возвращает singleton экземпляр хранилища
func GetInstance() *BalanceStorage {
	once.Do(func() {
		instance = &BalanceStorage{
			balances: make(map[string]float64),
		}
	})
	return instance
}

// UpdateBalance обновляет баланс счета
func (bs *BalanceStorage) UpdateBalance(accountID string, change float64) {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()
	bs.balances[accountID] += change
}

// GetBalanceChange возвращает изменение баланса для счета
func (bs *BalanceStorage) GetBalanceChange(accountID string) float64 {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()
	return bs.balances[accountID]
}

// GetAllBalanceChanges возвращает все изменения балансов
func (bs *BalanceStorage) GetAllBalanceChanges() map[string]float64 {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()
	
	result := make(map[string]float64)
	for k, v := range bs.balances {
		result[k] = v
	}
	return result
}