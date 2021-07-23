package main

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/wanmail/gods/stacks/redistack"
)

// RedisStackExample to demonstrate basic usage of RedisStack
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
	stack := redistack.New(client, "RedisStack") // empty (keys are of type int)
	stack.Push(1)                                // 1
	stack.Push(2)                                // 1, 2
	stack.Values()                               // 2, 1 (LIFO order)
	_, _ = stack.Peek()                          // 2,true
	_, _ = stack.Pop()                           // 2, true
	_, _ = stack.Pop()                           // 1, true
	_, _ = stack.Pop()                           // nil, false (nothing to pop)
	stack.Push(1)                                // 1
	stack.Clear()                                // empty
	stack.Empty()                                // true
	stack.Size()                                 // 0
}
