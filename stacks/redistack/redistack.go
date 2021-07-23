// Package redistack implements a stack backed by a redis list.
//
// Structure is not thread safe.
//
// You can use this stack to expand your application to distributed.
//
// Reference: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array
package redistack

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/wanmail/gods/stacks"
	"github.com/wanmail/gods/utils"
)

func assertStackImplementation() {
	var _ stacks.Stack = (*Stack)(nil)
}

// Stack holds elements in a redis list
type Stack struct {
	ctx    context.Context
	client *redis.Client
	key    string
}

// New instantiates a new empty stack
func New(client *redis.Client, key string) *Stack {
	return &Stack{
		ctx:    context.Background(),
		client: client,
		key:    key,
	}
}

// Push adds a value onto the top of the stack
func (stack *Stack) Push(value interface{}) {
	stack.client.LPush(stack.ctx, stack.key, value)
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
func (stack *Stack) Pop() (interface{}, bool) {
	element, err := stack.client.LPop(stack.ctx, stack.key).Result()
	if err != nil {
		return nil, false
	}
	return element, true
}

// Peek returns top element on the stack without removing it, or nil if stack is empty.
func (stack *Stack) Peek() (value interface{}, ok bool) {
	element, err := stack.client.LIndex(stack.ctx, stack.key, 0).Result()
	if err != nil {
		return nil, false
	}
	return element, true
}

// Empty returns true if stack does not contain any elements.
func (stack *Stack) Empty() bool {
	return stack.Size() == 0
}

// Size returns number of elements within the stack.
func (stack *Stack) Size() int {
	return int(stack.client.LLen(stack.ctx, stack.key).Val())
}

// Clear removes all elements from the stack.
func (stack *Stack) Clear() {
	stack.client.Del(stack.ctx, stack.key)
}

func (stack *Stack) values() []string {
	return stack.client.LRange(stack.ctx, stack.key, 0, -1).Val()
}

// Values returns all elements in the stack (LIFO order).
func (stack *Stack) Values() []interface{} {
	return utils.Strings2Interfaces(stack.values())
}

// String returns a string representation of container
func (stack *Stack) String() string {
	str := "RedisStack\n"
	str += strings.Join(stack.values(), ", ")
	return str
}
