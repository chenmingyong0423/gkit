// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package set

var _ set[int] = &MapSet[int]{}

// MapSet 使用一个 map 来实现 set 功能，map 的 value 使用一个空 struct 来占位
// MapSet is a generic set implementation using a map[T]struct{} to store the set elements.
type MapSet[T comparable] struct {
	mp map[T]struct{}
}

// NewMapSet returns a new MapSet with the given initial capacity.
// Params:
// - size: initial capacity of the map.
//
// Return:
// - a new MapSet.
//
// NewMapSet 创建一个新的 MapSet
// 参数：
// - size：map的初始大小
//
// 返回值：
// - 一个新的 MapSet
func NewMapSet[T comparable](size int) MapSet[T] {
	return MapSet[T]{
		mp: make(map[T]struct{}, size),
	}
}

// Add adds a value to the MapSet.
// Params:
// - val: value to add.
//
// Add 向 MapSet 中添加一个元素
// 参数：
// - val：待添加的元素
func (ms *MapSet[T]) Add(val T) {
	ms.mp[val] = struct{}{}
}

// Remove removes a value from the MapSet.
// Params:
// - val: value to remove.
//
// Remove 从 MapSet 中删除一个元素
// 参数：
// - val：待删除的元素
func (ms *MapSet[T]) Remove(val T) {
	delete(ms.mp, val)
}

// Contains checks if the MapSet contains a value.
// Params:
// - val: value to check.
//
// Return:
// - true if the MapSet contains the value, false otherwise.
//
// Contains 检查 MapSet 是否包含某个元素
// 参数：
// - val：待检查的元素
//
// 返回值：
// - 如果 MapSet 包含该元素，则返回 true；否则返回 false
func (ms *MapSet[T]) Contains(val T) bool {
	_, ok := ms.mp[val]
	return ok
}

// IsEmpty checks if the MapSet is empty.
// Return:
// - true if the MapSet is empty, false otherwise.
//
// IsEmpty 检查 MapSet 是否为空
// 返回值：
// - 如果 MapSet 为空，则返回 true；否则返回 false
func (ms *MapSet[T]) IsEmpty() bool {
	return len(ms.mp) == 0
}

// Size returns the size of the MapSet.
// Return:
// - the size of the MapSet.
//
// Size 返回 MapSet 中元素的个数
// 返回值：
// - MapSet 中元素的个数
func (ms *MapSet[T]) Size() int {
	return len(ms.mp)
}

// Clear removes all values from the MapSet.
// Clear 清空 MapSet 中所有元素
func (ms *MapSet[T]) Clear() {
	if len(ms.mp) != 0 {
		ms.mp = map[T]struct{}{}
	}
}
