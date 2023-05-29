package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func BenchmarkRandString(b *testing.B) {
	b.N = 1000 * 1000
	set := make(map[string]struct{})
	var mu sync.Mutex // 创建一个互斥锁
	var tmp string
	count := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			count++
			tmp = RandString(10)
			mu.Lock() // 在访问 set 映射之前锁定互斥锁
			_, ok := set[tmp]
			mu.Unlock() // 在访问 set 映射之后解锁互斥锁
			assert.False(b, ok, fmt.Sprintf("At %d times, duplicate string: %s", count, tmp))
			//if ok {
			//	panic(fmt.Sprintf("At %d times, duplicate string: %s", count, tmp))
			//}
			mu.Lock() // 在访问 set 映射之前锁定互斥锁
			set[RandString(10)] = struct{}{}
			mu.Unlock() // 在访问 set 映射之后解锁互斥锁
		}
	})
}
