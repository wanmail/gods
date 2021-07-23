package main

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/wanmail/gods/sets/rediset"
)

// RedisSetExample to demonstrate basic usage of RedisSet
func main() {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	client := redis.NewClient(
		&redis.Options{
			Addr: s.Addr(),
		},
	)
	set := rediset.New(client, "RedisSet") // empty (keys are of type int)
	set.Add(1)                             // 1
	set.Add(2, 2, 3, 4, 5)                 // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Remove(4)                          // 5, 3, 2, 1 (random order)
	set.Remove(2, 3)                       // 1, 5 (random order)
	set.Contains(1)                        // true
	set.Contains(1, 5)                     // true
	set.Contains(1, 6)                     // false
	_ = set.Values()                       // []int{5,1} (random order)
	set.Clear()                            // empty
	set.Empty()                            // true
	set.Size()                             // 0
}
