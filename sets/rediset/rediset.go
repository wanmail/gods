// Package rediset implements a set backed by a redis set.
//
// Structure is thread safe maybe.
//
// You can use this set to expand your application to distributed.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package rediset

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/wanmail/gods/sets"
	"github.com/wanmail/gods/utils"
)

func assertSetImplementation() {
	var _ sets.Set = (*Set)(nil)
}

// Set holds elements in a redis set
type Set struct {
	ctx    context.Context
	client *redis.Client
	key    string
}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New(client *redis.Client, key string, values ...interface{}) *Set {
	set := &Set{
		ctx:    context.Background(),
		client: client,
		key:    key,
	}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	set.client.SAdd(set.ctx, set.key, items...)
}

// Remove removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	set.client.SRem(set.ctx, set.key, items...)
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		if contains, _ := set.client.SIsMember(set.ctx, set.key, item).Result(); !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return int(set.client.SCard(set.ctx, set.key).Val())
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.client.Del(set.ctx, set.key)
}

func (set *Set) values() []string {
	return set.client.SMembers(set.ctx, set.key).Val()
}

// Values returns all items in the set.
func (set *Set) Values() []interface{} {
	return utils.Strings2Interfaces(set.values())
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "RedisSet\n"
	str += strings.Join(set.values(), ", ")
	return str
}
