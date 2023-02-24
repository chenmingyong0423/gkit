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

type MapSet[T comparable] struct {
	mp map[T]struct{}
}

func NewMapSet[T comparable](size int) MapSet[T] {
	return MapSet[T]{
		mp: make(map[T]struct{}, size),
	}
}

func (ms *MapSet[T]) Add(val T) {
	ms.mp[val] = struct{}{}
}

func (ms *MapSet[T]) Remove(val T) {
	delete(ms.mp, val)
}

func (ms *MapSet[T]) Contains(val T) bool {
	_, ok := ms.mp[val]
	return ok
}

func (ms *MapSet[T]) IsEmpty() bool {
	return len(ms.mp) == 0
}

func (ms *MapSet[T]) Size() int {
	return len(ms.mp)
}

func (ms *MapSet[T]) Clear() {
	if len(ms.mp) != 0 {
		ms.mp = map[T]struct{}{}
	}
}
