package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMap(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	strArr := make([]string, 1000000)
	for i := range strArr {
		strArr[i] = fmt.Sprintf("%d", r.Intn(1000000))
	}
	f := func(s string) string {
		return s + s
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map(strArr, f)
	}
}

func BenchmarkMapParallel(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	strArr := make([]string, 1000000)
	for i := range strArr {
		strArr[i] = fmt.Sprintf("%d", r.Intn(1000000))
	}
	f := func(s string) string {
		return s + s
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapParallel(strArr, f)
	}
}

func BenchmarkMapParallel2(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	strArr := make([]string, 1000000)
	for i := range strArr {
		strArr[i] = fmt.Sprintf("%d", r.Intn(1000000))
	}
	f := func(s string) string {
		return s + s
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MapParallel2(strArr, f)
	}
}
