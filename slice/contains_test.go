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

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		data []T
		item T
		want bool
	}
	tests := []testCase[int]{
		{
			name: "空切片",
			item: 2,
			want: false,
		},
		{
			name: "无元素的切片",
			data: []int{},
			item: 2,
			want: false,
		},
		{
			name: "元素不存在",
			data: []int{1, 3, 5, 7, 9},
			item: 2,
			want: false,
		},
		{
			name: "元素存在",
			data: []int{1, 3, 5, 7, 9},
			item: 1,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Contains(tt.data, tt.item))
		})
	}
}

func TestContainsByFunc(t *testing.T) {
	type testCase[T any] struct {
		name    string
		data    []T
		dstItem T
		want    bool
	}
	tests := []testCase[int]{
		{
			name:    "空切片",
			dstItem: 2,
			want:    false,
		},
		{
			name:    "无元素的切片",
			data:    []int{},
			dstItem: 2,
			want:    false,
		},
		{
			name:    "元素不存在",
			data:    []int{1, 3, 5, 7, 9},
			dstItem: 2,
			want:    false,
		},
		{
			name:    "元素存在",
			data:    []int{1, 3, 5, 7, 9},
			dstItem: 1,
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ContainsByFunc(tt.data, tt.dstItem, func(srcItem, dstItem int) bool {
				return srcItem == dstItem
			}))
		})
	}
}
