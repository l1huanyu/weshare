package gateway

import (
	"sync"
)

type safeMap struct {
	sync.RWMutex
	Map map[string]func(userID string, content int) string
}

func newSafaMap() *safeMap {
	sm := new(safeMap)
	sm.Map = make(map[string]func(userID string, content int) string)
	return sm
}

func (sm *safeMap) read(key string) (func(userID string, content int) string, bool) {
	sm.RLock()
	value, ok := sm.Map[key]
	sm.RUnlock()
	return value, ok
}

func (sm *safeMap) write(key string, value func(userID string, content int) string) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}

func (sm *safeMap) delete(key string) {
	sm.Lock()
	delete(sm.Map, key)
	sm.Unlock()
}
