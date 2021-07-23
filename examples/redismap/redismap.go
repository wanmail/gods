package main

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/wanmail/gods/maps/redismap"
)

// RedisMapExample to demonstrate basic usage of RedisMap
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
	m := redismap.New(client, "RedisSet") // empty (keys are of type int)
	m.Put(1, "x")                         // 1->x
	m.Put(2, "b")                         // 2->b, 1->x  (random order)
	m.Put(1, "a")                         // 2->b, 1->a (random order)
	_, _ = m.Get(2)                       // b, true
	_, _ = m.Get(3)                       // nil, false
	_ = m.Values()                        // []interface {}{"b", "a"} (random order)
	_ = m.Keys()                          // []interface {}{1, 2} (random order)
	m.Remove(1)                           // 2->b
	m.Clear()                             // empty
	m.Empty()                             // true
	m.Size()                              // 0
}
