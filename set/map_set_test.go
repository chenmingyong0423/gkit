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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSet_Add(t *testing.T) {
	testCases := []struct {
		name string
		ms   MapSet[int]
		val  int

		expected MapSet[int]
	}{
		{
			name: "空 MapSet 添加元素",
			ms: MapSet[int]{
				mp: make(map[int]struct{}),
			},
			val: 1,
			expected: MapSet[int]{
				mp: map[int]struct{}{
					1: {},
				},
			},
		},
		{
			name: "非空 MapSet 添加元素",
			ms: MapSet[int]{
				mp: map[int]struct{}{
					1: {},
				},
			},
			val: 2,
			expected: MapSet[int]{
				mp: map[int]struct{}{
					1: {},
					2: {},
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.ms.Add(tt.val)
			assert.Equal(t, tt.ms, tt.expected)
		})
	}
}

func TestMapSet_Remove(t *testing.T) {
	testCases := []struct {
		name string
		ms   MapSet[int]
		val  int

		expected MapSet[int]
	}{
		{
			name: "空 MapSet 删除元素",
			ms: MapSet[int]{
				mp: make(map[int]struct{}),
			},
			val: 1,
			expected: MapSet[int]{
				mp: map[int]struct{}{},
			},
		},
		{
			name: "非空 MapSet 删除元素",
			ms: MapSet[int]{
				mp: map[int]struct{}{
					1: {},
				},
			},
			val: 1,
			expected: MapSet[int]{
				mp: map[int]struct{}{},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.ms.Remove(tt.val)
			assert.Equal(t, tt.ms, tt.expected)
		})
	}
}

func TestMapSet_Contains(t *testing.T) {
	testCases := []struct {
		name string
		ms   MapSet[int]
		val  int

		expected bool
	}{
		{
			name: "空 MapSet 判断",
			ms: MapSet[int]{
				mp: make(map[int]struct{}),
			},
			val:      1,
			expected: false,
		},
		{
			name: "非空 MapSet 判断（不存在）",
			ms: MapSet[int]{
				mp: map[int]struct{}{
					1: {},
				},
			},
			val:      0,
			expected: false,
		},
		{
			name: "非空 MapSet 判断（存在）",
			ms: MapSet[int]{
				mp: map[int]struct{}{
					1: {},
				},
			},
			val:      1,
			expected: true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.ms.Contains(tt.val))
		})
	}
}

func TestMapSet_IsEmpty(t *testing.T) {
	testCases := []struct {
		name string
		ms   MapSet[int]

		expected bool
	}{
		{
			name: "空 MapSet",
			ms: MapSet[int]{
				mp: make(map[int]struct{}),
			},
			expected: true,
		},
		{
			name: "非空 MapSet",
			ms: MapSet[int]{
				mp: map[int]struct{}{
					1: {},
				},
			},
			expected: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.ms.IsEmpty())
		})
	}
}

func TestMapSet_Size(t *testing.T) {
	testCases := []struct {
		name string
		ms   MapSet[int]

		expected int
	}{
		{
			name: "空 MapSet",
			ms: MapSet[int]{
				mp: make(map[int]struct{}),
			},
			expected: 0,
		},
		{
			name: "非空 MapSet",
			ms: MapSet[int]{
				mp: map[int]struct{}{
					1: {},
				},
			},
			expected: 1,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.ms.Size())
		})
	}
}

func TestMapSet_Clear(t *testing.T) {
	testCases := []struct {
		name string
		ms   MapSet[int]

		expected MapSet[int]
	}{
		{
			name: "空 MapSet 清空",
			ms: MapSet[int]{
				mp: make(map[int]struct{}),
			},
			expected: MapSet[int]{
				mp: make(map[int]struct{}),
			},
		},
		{
			name: "非空 MapSet 清空",
			ms: MapSet[int]{
				mp: map[int]struct{}{
					1: {},
				},
			},
			expected: MapSet[int]{
				mp: make(map[int]struct{}),
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.ms.Clear()
			assert.Equal(t, tt.expected, tt.ms)
		})
	}
}

func TestNewMapSet(t *testing.T) {
	tests := []struct {
		name string
		size int
		want MapSet[int]
	}{
		{
			name: "初始化长度为 0",
			size: 1,
			want: MapSet[int]{
				mp: make(map[int]struct{}, 0),
			},
		},
		{
			name: "初始化长度为 1",
			size: 1,
			want: MapSet[int]{
				mp: make(map[int]struct{}, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewMapSet[int](tt.size))
		})
	}
}
