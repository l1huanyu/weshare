package gateway

import (
	"sync"
)

type lightCache struct {
	sync.RWMutex
	HopMap     map[string]func(userID string, content string) string
	TypeMap    map[string]int
	ReadMap    map[string][]uint
	CurrentMap map[string]uint
}

func newlightCache() *lightCache {
	c := new(lightCache)
	c.HopMap = make(map[string]func(userID string, content string) string)
	c.TypeMap = make(map[string]int)
	c.ReadMap = make(map[string][]uint)
	c.CurrentMap = make(map[string]uint)
	return c
}

func (lc *lightCache) readHopMap(key string) (func(userID string, content string) string, bool) {
	lc.RLock()
	defer lc.RUnlock()
	value, ok := lc.HopMap[key]
	return value, ok
}

func (lc *lightCache) writeHopMap(key string, value func(userID string, content string) string) {
	lc.Lock()
	defer lc.Unlock()
	lc.HopMap[key] = value
}

func (lc *lightCache) deleteHopMap(key string) {
	lc.Lock()
	defer lc.Unlock()
	delete(lc.HopMap, key)
}

func (lc *lightCache) readTypeMap(key string) (int, bool) {
	lc.RLock()
	defer lc.RUnlock()
	value, ok := lc.TypeMap[key]
	return value, ok
}

func (lc *lightCache) writeTypeMap(key string, value int) {
	lc.Lock()
	defer lc.Unlock()
	lc.TypeMap[key] = value
}

func (lc *lightCache) deleteTypeMap(key string) {
	lc.Lock()
	defer lc.Unlock()
	delete(lc.TypeMap, key)
}

func (lc *lightCache) readReadMap(key string) []uint {
	lc.RLock()
	defer lc.RUnlock()
	value, ok := lc.ReadMap[key]
	if ok {
		return value
	}
	return nil
}

func (lc *lightCache) writeReadMap(key string, value uint) {
	lc.Lock()
	defer lc.Unlock()
	if lc.ReadMap[key] == nil {
		lc.ReadMap[key] = make([]uint, 0)
	}
	lc.ReadMap[key] = append(lc.ReadMap[key], value)
}

func (lc *lightCache) deleteReadMap(key string) {
	lc.Lock()
	defer lc.Unlock()
	delete(lc.ReadMap, key)
}

func (lc *lightCache) readCurrentMap(key string) (uint, bool) {
	lc.RLock()
	defer lc.RUnlock()
	value, ok := lc.CurrentMap[key]
	return value, ok
}

func (lc *lightCache) writeCurrentMap(key string, value uint) {
	lc.Lock()
	defer lc.Unlock()
	lc.CurrentMap[key] = value
}

func (lc *lightCache) deleteCurrentMap(key string) {
	lc.Lock()
	defer lc.Unlock()
	delete(lc.CurrentMap, key)
}
