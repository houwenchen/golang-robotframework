package server

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func DoRequest() {
	client := http.Client{}
	url := "http://localhost:8080/handle"

	request, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 串行--看一段时间内函数执行的次数，及每次执行时占用的时间和内存
// 21901	     53494 ns/op	    3558 B/op	      45 allocs/op  1.893s
func BenchmarkDoRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoRequest()
	}
}

// 并行--看b.parallelism * runtime.GOMAXPROCS(0)这么多进程下的并发情况
// 35497	    166022 ns/op	    9742 B/op	      79 allocs/op  6.405s
func BenchmarkDoRequestP(b *testing.B) {
	b.SetParallelism(5)

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			DoRequest()
		}
	})
}
