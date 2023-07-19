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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectionSet(t *testing.T) {
	testCases := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{
			name:   "Takes the intersection of two empty slices",
			slice1: []int{},
			slice2: []int{},

			want: []int{},
		},
		{
			name:   "Takes the intersection of two slices, one of which is an empty slice",
			slice1: []int{1, 2, 3},
			slice2: []int{},

			want: []int{},
		},
		{
			name:   "Takes the intersection of two slices and get empty slice",
			slice1: []int{0, 2, 3},
			slice2: []int{1, 4, 5},

			want: []int{},
		},
		{
			name:   "Takes the intersection of two slices",
			slice1: []int{1, 2, 3},
			slice2: []int{1, 4, 5},

			want: []int{1},
		},
		{
			name:   "Takes the intersection between two slices with duplicate elements",
			slice1: []int{1, 2, 2, 3},
			slice2: []int{1, 1, 2, 4, 5},

			want: []int{1, 2},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IntersectionSet(tt.slice1, tt.slice2)
			assert.ElementsMatch(t, tt.want, result)
		})
	}
}

func TestIntersectionSetFunc(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	testCases := []struct {
		name   string
		slice1 []user
		slice2 []user
		equal  equalFunc[user]

		want []user
	}{
		{
			name:   "Takes the intersection of two empty slices",
			slice1: []user{},
			slice2: []user{},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.name == dstItem.name && srcItem.age == srcItem.age
			},

			want: []user{},
		},
		{
			name:   "Takes the intersection of two slices, one of which is an empty slice",
			slice1: []user{{name: "陈明勇", age: 24}},
			slice2: []user{},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.name == dstItem.name && srcItem.age == srcItem.age
			},

			want: []user{},
		},
		{
			name:   "Takes the intersection of two slices and get empty slice",
			slice1: []user{{name: "陈明勇", age: 24}},
			slice2: []user{{name: "cmy", age: 24}},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.name == dstItem.name && srcItem.age == srcItem.age
			},

			want: []user{},
		},
		{
			name:   "Takes the intersection of two slices",
			slice1: []user{{name: "陈明勇", age: 24}, {name: "n", age: 24}},
			slice2: []user{{name: "陈明勇", age: 24}},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.name == dstItem.name && srcItem.age == srcItem.age
			},

			want: []user{{name: "陈明勇", age: 24}},
		},
		{
			name:   "Takes the intersection between two slices with duplicate elements",
			slice1: []user{{name: "陈明勇", age: 24}, {name: "n", age: 24}, {name: "n", age: 24}, {name: "o", age: 24}},
			slice2: []user{{name: "陈明勇", age: 24}, {name: "陈明勇", age: 24}, {name: "n", age: 24}, {name: "m", age: 24}},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.name == dstItem.name && srcItem.age == srcItem.age
			},

			want: []user{{name: "陈明勇", age: 24}, {name: "n", age: 24}},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IntersectionSetFunc(tt.slice1, tt.slice2, tt.equal)
			assert.ElementsMatch(t, tt.want, result)
		})
	}
}
