package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExpire(t *testing.T) {
	k := "aaa"
	o := []byte("111")
	c := SingleCache{
		Key:     k,
		Value:   o,
		Expire:  int64(100 * time.Millisecond),
		SetTime: time.Now().UnixNano(),
	}
	CacheManager.Add(k, c)

	var (
		v []byte
	)

	// 测试取值

	// 立即
	v, _ = CacheManager.Get(k)
	assert.Equal(t, o, v)

	// 延迟时间内
	time.Sleep(20 * time.Millisecond)
	v, _ = CacheManager.Get(k)
	assert.Equal(t, o, v)

	// 延迟时间后
	time.Sleep(100 * time.Millisecond)
	v, _ = CacheManager.Get(k)
	assert.Equal(t, make([]byte, 0), v)
}
