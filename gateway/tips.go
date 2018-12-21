package gateway

import (
	"sync"
)

type safeMap struct {
	sync.RWMutex
	Map      map[string]func(userID string, content string) string
	StateMap map[string]int
}

func newSafaMap() *safeMap {
	sm := new(safeMap)
	sm.Map = make(map[string]func(userID string, content string) string)
	sm.StateMap = make(map[string]int)
	return sm
}

func (sm *safeMap) read(key string) (func(userID string, content string) string, bool) {
	sm.RLock()
	value, ok := sm.Map[key]
	sm.RUnlock()
	return value, ok
}

func (sm *safeMap) write(key string, value func(userID string, content string) string) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}

func (sm *safeMap) delete(key string) {
	sm.Lock()
	delete(sm.Map, key)
	sm.Unlock()
}

func (sm *safeMap) readState(key string) (int, bool) {
	sm.RLock()
	value, ok := sm.StateMap[key]
	sm.RUnlock()
	return value, ok
}

func (sm *safeMap) writeState(key string, value int) {
	sm.Lock()
	sm.StateMap[key] = value
	sm.Unlock()
}

func (sm *safeMap) deleteState(key string) {
	sm.Lock()
	delete(sm.StateMap, key)
	sm.Unlock()
}
