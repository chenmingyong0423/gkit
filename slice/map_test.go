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

func Test_toMap(t *testing.T) {
	testCases := []struct {
		name string
		data []int
		want map[int]struct{}
	}{
		{
			name: "nil slice 转换成 map",
			data: nil,
			want: map[int]struct{}{},
		},
		{
			name: "空 slice 转换成 map",
			data: []int{},
			want: map[int]struct{}{},
		},
		{
			name: "非空 slice（有重复值）转换成 map",
			data: []int{1, 2, 2, 3},
			want: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
		{
			name: "非空 slice 转换成 map",
			data: []int{1, 2, 3},
			want: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, toMap(tt.data))
		})
	}
}

func Test_mapKeyToSlice(t *testing.T) {
	testCases := []struct {
		name string
		data map[int]struct{}
		want []int
	}{
		{
			name: "make the nil map to slice",
			data: nil,
			want: []int{},
		},
		{
			name: "make the empty map to slice",
			data: map[int]struct{}{},
			want: []int{},
		},
		{
			name: "make the map to slice",
			data: map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}},
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, mapKeyToSlice(tt.data))
		})
	}
}

func TestMap(t *testing.T) {
	type user struct {
		id   int
		name string
	}

	testCases := []struct {
		name string
		src  []user
		fn   func(idx int, s user) int
		want []int
	}{
		{
			name: "nil slice",
			src:  nil,
			fn:   func(idx int, s user) int { return s.id },
			want: []int{},
		},
		{
			name: "empty slice",
			src:  []user{},
			fn:   func(idx int, s user) int { return s.id },
			want: []int{},
		},
		{
			name: "non-empty slice",
			src: []user{
				{id: 1, name: "陈明勇"},
				{id: 2, name: "Gopher"},
			},
			fn:   func(idx int, s user) int { return s.id },
			want: []int{1, 2},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, Map(tc.src, tc.fn))
		})
	}
}
