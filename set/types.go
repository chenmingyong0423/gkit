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

// set is a generic interface that defines basic operations for set-like data structures.
// set 是一个通用接口，用于定义集合类型的基本操作。
type set[T comparable] interface {
	// Add adds an element to the set.
	// Add 用于向集合中添加一个元素。
	Add(t T)

	// Remove removes an element from the set.
	// Remove 用于从集合中移除一个元素。
	Remove(t T)

	// Contains checks if an element is in the set.
	// Contains 用于判断集合中是否包含一个元素。
	Contains(t T) bool

	// IsEmpty checks if the set is empty.
	// IsEmpty 用于判断集合是否为空。
	IsEmpty() bool

	// Size returns the number of elements in the set.
	// Size 用于获取集合中元素的个数。
	Size() int

	// Clear removes all elements from the set.
	// Clear 用于清空集合中所有元素。
	Clear()
}
