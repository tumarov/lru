package example

import (
	"fmt"
	"github.com/tumarov/lru"
	"strconv"
	"sync"
)

func example(size int) {
	cache := lru.NewLRUCache(size)

	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			cache.Add("str"+strconv.Itoa(i), "g")
			wg.Done()
		}(i)
	}

	for i := 0; i < 6; i++ {
		go func(i int) {
			fmt.Println(cache.Get("str" + strconv.Itoa(i)))
		}(i)
	}

	wg.Wait()

	fmt.Println("finish")
}
