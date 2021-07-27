package cache

import (
	"math"
	"math/rand"
	grpcService "ne_cache/grpc"
	"sync"
	"time"
)

type SingleCache struct {
	Key     string
	Value   []byte
	Expire  int64
	SetTime int64
	front   *SingleCache
	back    *SingleCache
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

func (s *SingleCache) Expired() bool {
	return s.Expire != 0 && time.Now().UnixNano() >= s.Expire + s.SetTime
}

func (c *cacheManage) Add(key string, cache SingleCache) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	newSize := len(cache.Value)
	// 如果已经有值，则更新
	if v, ok := c.Cache[key]; ok {
		oldSize := len(v.Value)
		c.Cache[key] = &cache
		cache.front = v.front
		cache.back = v.back
		c.CacheSize += int64(newSize - oldSize)
	} else {
		// 没有值则插入
		c.Cache[key] = &cache
		// 如果是不是第一个cache
		if c.EndSingleCache != nil {
			c.EndSingleCache.back = &cache
			cache.front = c.EndSingleCache
		}
		c.EndSingleCache = &cache
		c.CacheSize += int64(newSize)
	}
}

func (c *cacheManage) Get(key string) ([]byte, grpcService.GetValueResponse_Status) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	var v []byte
	var s grpcService.GetValueResponse_Status
	if ca, ok := c.Cache[key]; ok {
		if ca.Expired() {
			v = make([]byte, 0)
			s = grpcService.GetValueResponse_FAIL
		} else {
			v = ca.Value
			s = grpcService.GetValueResponse_OK
		}
	} else {
		v = make([]byte, 0)
		s = grpcService.GetValueResponse_FAIL
	}
	return v, s
}

func (c *cacheManage) Delete(key string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	if ca, ok := c.Cache[key]; ok {
		if ca.back == nil && ca.front == nil {
			c.EndSingleCache = nil
		} else if ca.back != nil {
			ca.back.front = ca.front
			ca.back = nil
			if ca.front != nil {
				ca.front.back = ca.back
				ca.front = nil
			}
		} else {
			ca.front.front = nil
			c.EndSingleCache = ca.front
			c.EndSingleCache.back = nil
		}
		c.CacheSize -= int64(len(ca.Value))
		delete(c.Cache, ca.Key)
	}
}

func (c *cacheManage) GetKeys() []string {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	j := 0
	keys := make([]string, len(c.Cache))
	for k := range c.Cache {
		keys[j] = k
		j++
	}
	return keys
}

func (c *cacheManage) PopEndSingleCache() {
	if c.EndSingleCache != nil {
		c.Delete(c.EndSingleCache.Key)
	}
}

var CacheManager = cacheManage{
	Cache:              make(map[string]*SingleCache),
	CacheSizeLimit:     1024 * 1024 * 1024,
	ExpireCheckRate:    0.05,
	ExpireAllCheckRate: 0.5,
}

/*
CheckExpire
判断一个key是否过期

返回的bool是标识这个key是否过期，true是过期，false是未过期
*/
func (c *cacheManage) CheckExpire(key string) bool {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	if ca, ok := c.Cache[key]; ok {
		if ca.Expired() {
			return true
		} else {
			return false
		}
	}
	// 不存在默认为过期了
	return true
}

// ExpireChecker 检测过期key
func ExpireChecker() {
	// 数量太少不进行过期检查
	if len(CacheManager.Cache) > 10 {
		rand.Seed(time.Now().Unix())
		k := CacheManager.GetKeys()
		cacheSize := len(k)
		checkCount := int(math.Round(float64(cacheSize) * CacheManager.ExpireCheckRate))
		allCheckCount := int(math.Round(float64(checkCount) * CacheManager.ExpireAllCheckRate))
		expireCount := 0
		for i := 0; i < checkCount; i++ {
			key := k[rand.Intn(cacheSize)]
			e := CacheManager.CheckExpire(key)
			if e == true {
				expireCount += 1
				CacheManager.Delete(key)
			}
		}
		// 确定是否需要全体检测
		if expireCount >= allCheckCount {
			ak := CacheManager.GetKeys()
			for _, sk := range ak {
				er := CacheManager.CheckExpire(sk)
				if er == true {
					CacheManager.Delete(sk)
				}
			}
		}
	}
}

// MemChecker 容量检测
func MemChecker() {
	for CacheManager.CacheSize > CacheManager.CacheSizeLimit {
		CacheManager.PopEndSingleCache()
	}
}

func Checker() {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			ExpireChecker()
			MemChecker()
		}
	}()
}
