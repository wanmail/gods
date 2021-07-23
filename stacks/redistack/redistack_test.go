package redistack

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

func newMockClient() *redis.Client {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(
		&redis.Options{
			Addr: s.Addr(),
		},
	)
	return client
}

func TestStackPush(t *testing.T) {
	stack := New(newMockClient(), "RedisStack")
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	if actualValue := stack.Values(); actualValue[0] != "c" || actualValue[1] != "b" || actualValue[2] != "a" {
		t.Errorf("Got %v expected %v", actualValue, "[c,b,a]")
	}
	if actualValue := stack.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := stack.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := stack.Peek(); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestStackPeek(t *testing.T) {
	stack := New(newMockClient(), "RedisStack")
	if actualValue, ok := stack.Peek(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	if actualValue, ok := stack.Peek(); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestStackPop(t *testing.T) {
	stack := New(newMockClient(), "RedisStack")
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Pop()
	if actualValue, ok := stack.Peek(); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := stack.Pop(); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := stack.Pop(); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := stack.Pop(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := stack.Values(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}
}
