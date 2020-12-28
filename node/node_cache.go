package node

import (
	"math/rand"
	grpcService "ne_cache/node/grpc"
	"neko_server_go/utils"
	"sync"
	"time"
)

type SingleCache struct {
	Key    string
	Value  []byte
	Expire int64
	front  *SingleCache
	back   *SingleCache
}

type cacheManage struct {
	Cache              map[string]*SingleCache
	EndSingleCache     *SingleCache
	Lock               sync.RWMutex // 读写需要加锁
	CacheSize          int64
	CacheSizeLimit     int64
	ExpireCheckRate    float64 // 所有key中，过期检查的比例
	ExpireAllCheckRate float64 // 过期检查的key中，多少比例的key过期会触发全体key过期检查
}

var CacheManager = cacheManage{
	Cache:              make(map[string]*SingleCache),
	CacheSizeLimit:     1024 * 1024 * 1024,
	ExpireCheckRate:    0.05,
	ExpireAllCheckRate: 0.5,
}

func (c *cacheManage) Add(key string, cache SingleCache) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	newSize := len(cache.Value)
	// 如果已经有值，则更新
	if v, ok := CacheManager.Cache[key]; ok {
		oldSize := len(v.Value)
		CacheManager.Cache[key] = &cache
		cache.front = v.front
		cache.back = v.back
		CacheManager.CacheSize += int64(newSize - oldSize)
	} else {
		CacheManager.Cache[key] = &cache
		// 如果是不是第一个cache
		if CacheManager.EndSingleCache != nil {
			CacheManager.EndSingleCache.back = &cache
			cache.front = CacheManager.EndSingleCache
		}
		CacheManager.EndSingleCache = &cache
		CacheManager.CacheSize += int64(newSize)
	}
}

func (c *cacheManage) Get(key string) ([]byte, grpcService.GetValueResponse_Status) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	var v []byte
	var s grpcService.GetValueResponse_Status
	if c, ok := CacheManager.Cache[key]; ok {
		// TODO 需要检查key是否过期
		v = c.Value
		s = grpcService.GetValueResponse_OK
	} else {
		v = make([]byte, 0)
		s = grpcService.GetValueResponse_FAIL
	}
	return v, s
}

func (c *cacheManage) GetKeys() []string {
	j := 0
	keys := make([]string, len(c.Cache))
	for k, _ := range c.Cache {
		keys[j] = k
		j++
	}
	return keys
}

func (c *cacheManage) PopEndSingleCache() *SingleCache {
	CacheManager.Lock.Lock()
	defer CacheManager.Lock.Unlock()
	if c.EndSingleCache != nil {
		r := c.EndSingleCache
		c.EndSingleCache = r.front
		c.EndSingleCache.back = nil
		c.CacheSize -= int64(len(r.Value))
		delete(c.Cache, r.Key)
		return r
	} else {
		return nil
	}
}

// 检测过期key
func ExpireChecker() {
	// TODO
	// 加锁
	if len(CacheManager.Cache) > 0 {
		utils.LogDebug("part key expire check")
		CacheManager.Lock.Lock()
		defer CacheManager.Lock.Unlock()

		rand.Seed(time.Now().Unix())
		k := CacheManager.GetKeys()


	}
}

// 容量检测
func MemChecker() {
	for CacheManager.CacheSize > CacheManager.CacheSizeLimit {
		CacheManager.PopEndSingleCache()
	}
}

func Checker() {
	var ch chan int
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			ExpireChecker()
			MemChecker()
		}
		ch <- 1
	}()
	<-ch
}
