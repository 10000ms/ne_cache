package node

import (
	"github.com/shirou/gopsutil/mem"
	grpcService "ne_cache/node/grpc"
	"neko_server_go/utils"
	"sync"
	"time"
)

type SingleCache struct {
	Value  []byte
	Expire int64
	front  *SingleCache
	back   *SingleCache
}

type cacheManage struct {
	Cache          map[string]*SingleCache
	EndSingleCache *SingleCache
	Lock sync.RWMutex // 读写需要加锁
}

var CacheManager = cacheManage{
	Cache: make(map[string]*SingleCache),
}

func (c *cacheManage) Add(key string, cache SingleCache) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	// 如果已经有值，则更新
	if v, ok := CacheManager.Cache[key]; ok {
		CacheManager.Cache[key] = &cache
		cache.front = v.front
		cache.back = v.back
	} else {
		CacheManager.Cache[key] = &cache
		// 如果是不是第一个cache
		if CacheManager.EndSingleCache != nil {
			CacheManager.EndSingleCache.back = &cache
			cache.front = CacheManager.EndSingleCache
		}
		CacheManager.EndSingleCache = &cache
	}
}

func (c *cacheManage) Get(key string) ([]byte, grpcService.GetValueResponse_Status) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	var v []byte
	var s grpcService.GetValueResponse_Status
	if c, ok := CacheManager.Cache[key]; ok {
		v = c.Value
		s = grpcService.GetValueResponse_OK
	} else {
		v = make([]byte, 0)
		s = grpcService.GetValueResponse_FAIL
	}
	return v, s

}

// 检测过期key
func ExpireChecker() {
	// TODO
	// 加锁
}

// 内存检测
func MemChecker() {
	// TODO
	// 加锁
	v, _ := mem.VirtualMemory()
	utils.LogInfo(v)
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
