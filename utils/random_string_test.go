package utils

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func BenchmarkRandString(b *testing.B) {
	// create a set with sync.Map
	set := sync.Map{}
	var mu sync.Mutex // 创建一个互斥锁
	var tmp string
	count := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			count++
			mu.Lock() // 在访问 set 映射之前锁定互斥锁
			tmp = RandString(10)
			mu.Unlock() // 在访问 set 映射之后解锁互斥锁

			_, ok := set.Load(tmp)
			assert.False(b, ok, "At %d times, duplicate string: %s", count, tmp)
			//if ok {
			//	panic(fmt.Sprintf("At %d times, duplicate string: %s", count, tmp))
			//}
			set.Store(tmp, struct{}{})
		}
	})
}
