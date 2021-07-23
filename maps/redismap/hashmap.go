// Package redismap implements a map backed by a redis hash.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// You can use this hash to expand your application to distributed.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package redismap

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/wanmail/gods/maps"
	"github.com/wanmail/gods/utils"
)

func assertMapImplementation() {
	var _ maps.Map = (*Map)(nil)
}

// Map holds the elements in redis hash
type Map struct {
	ctx    context.Context
	client *redis.Client
	key    string
}

// New instantiates a hash map.
func New(client *redis.Client, key string) *Map {
	return &Map{
		ctx:    context.Background(),
		client: client,
		key:    key,
	}
}

// Put inserts element into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	m.client.HSet(m.ctx, m.key, utils.ToString(key), value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
func (m *Map) Get(key interface{}) (interface{}, bool) {
	element, err := m.client.HGet(m.ctx, m.key, utils.ToString(key)).Result()
	if err != nil {
		return nil, false
	}
	return element, true
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key interface{}) {
	m.client.HDel(m.ctx, m.key, utils.ToString(key))
}

// Empty returns true if map does not contain any elements
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return int(m.client.HLen(m.ctx, m.key).Val())
}

// Keys returns all keys (random order).
func (m *Map) Keys() []interface{} {
	return utils.Strings2Interfaces(m.client.HKeys(m.ctx, m.key).Val())
}

// Values returns all values (random order).
func (m *Map) Values() []interface{} {
	return utils.Strings2Interfaces(m.client.HVals(m.ctx, m.key).Val())
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.client.Del(m.ctx, m.key)
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "RedisMap\n"
	str += fmt.Sprintf("%v", m.client.HGetAll(m.ctx, m.key).Val())
	return str
}
