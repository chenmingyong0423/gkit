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

func TestUnion(t *testing.T) {
	testCases := []struct {
		name   string
		slices [][]int
		want   []int
	}{
		{
			name: "Takes the union set in two nil slices",
			slices: [][]int{
				nil,
				nil,
			},
			want: []int{},
		},
		{
			name: "Takes the union set in two empty slices",
			slices: [][]int{
				{},
				{},
			},
			want: []int{},
		},
		{
			name: "Takes the union of two slices where one slice contains a duplicate element",
			slices: [][]int{
				{1, 2, 3, 4},
				{1, 2, 2, 3, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Takes the union of two slices",
			slices: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},

		{
			name: "Takes the union set in three nil slices",
			slices: [][]int{
				nil,
				nil,
				nil,
			},
			want: []int{},
		},
		{
			name: "Takes the union set in three empty slices",
			slices: [][]int{
				{},
				{},
				{},
			},
			want: []int{},
		},
		{
			name: "Takes the union of three slices where one slice contains a duplicate element",
			slices: [][]int{
				{1, 2, 3, 4},
				{1, 2, 2, 3, 5},
				{5, 5, 6, 7},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "Takes the union of three slices",
			slices: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, Union(tt.slices...))
		})
	}
}

func TestUnionFunc(t *testing.T) {
	type user struct {
		id   int
		name string
	}
	testCases := []struct {
		name   string
		slices [][]user
		equal  equalFunc[user]

		want []user
	}{
		{
			name: "Takes the union set in two nil slices",
			slices: [][]user{
				nil,
				nil,
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.id == dstItem.id
			},

			want: []user{},
		},
		{
			name: "Takes the union set in two empty slices",
			slices: [][]user{
				{},
				{},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.id == dstItem.id
			},

			want: []user{},
		},
		{
			name: "Takes the union of two slices where one slice contains a duplicate element",
			slices: [][]user{
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}},
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 5, name: "5"}},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.id == dstItem.id
			},

			want: []user{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}, {id: 5, name: "5"}},
		},
		{
			name: "Takes the union of two slices",
			slices: [][]user{
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}},
				{{id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.id == dstItem.id
			},

			want: []user{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}, {id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}},
		},

		{
			name: "Takes the union set in three nil slices",
			slices: [][]user{
				nil,
				nil,
				nil,
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.id == dstItem.id
			},

			want: []user{},
		},
		{
			name: "Takes the union set in three empty slices",
			slices: [][]user{
				{},
				{},
				{},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.id == dstItem.id
			},

			want: []user{},
		},
		{
			name: "Takes the union of three slices where one slice contains a duplicate element",
			slices: [][]user{
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}},
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 5, name: "5"}},
				{{id: 5, name: "5"}, {id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.id == dstItem.id
			},

			want: []user{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}, {id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}},
		},
		{
			name: "Takes the union of three slices",
			slices: [][]user{
				{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}},
				{{id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}},
				{{id: 9, name: "9"}, {id: 10, name: "10"}},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.id == dstItem.id
			},

			want: []user{{id: 1, name: "1"}, {id: 2, name: "2"}, {id: 3, name: "3"}, {id: 4, name: "4"}, {id: 5, name: "5"}, {id: 6, name: "6"}, {id: 7, name: "7"}, {id: 8, name: "8"}, {id: 9, name: "9"}, {id: 10, name: "10"}},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, UnionFunc[user](tt.equal, tt.slices...))
		})
	}
}
