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

func TestCombineNestedSlices(t *testing.T) {
	type PostCategory struct {
		Tags []string
	}
	testCases := []struct {
		name        string
		slices      []PostCategory
		extractFunc func(idx int, s PostCategory) []string
		want        []string
	}{
		{
			name: "nil slice",
			want: []string{},
		},
		{
			name:   "empty slice",
			slices: []PostCategory{},
			want:   []string{},
		},
		{
			name: "non-empty slice with nil Tags",
			slices: []PostCategory{
				{},
			},
			extractFunc: func(idx int, s PostCategory) []string { return s.Tags },
			want:        []string{},
		},
		{
			name: "non-empty slice with empty Tags",
			slices: []PostCategory{
				{Tags: []string{}},
			},
			extractFunc: func(idx int, s PostCategory) []string { return s.Tags },
			want:        []string{},
		},
		{
			name: "non-empty slice with non-empty Tags",
			slices: []PostCategory{
				{Tags: []string{"Go", "Java"}},
				{Tags: []string{"Vue", "React", "Go"}},
			},
			extractFunc: func(idx int, s PostCategory) []string { return s.Tags },
			want:        []string{"Go", "Java", "Vue", "React", "Go"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, CombineNestedSlices(tc.slices, tc.extractFunc))
		})
	}
}

func TestCombineAndDeduplicateNestedSlices(t *testing.T) {
	type PostCategory struct {
		Tags []string
	}
	testCases := []struct {
		name        string
		slices      []PostCategory
		extractFunc func(idx int, s PostCategory) []string
		want        []string
	}{
		{
			name: "nil slice",
			want: []string{},
		},
		{
			name:   "empty slice",
			slices: []PostCategory{},
			want:   []string{},
		},
		{
			name: "non-empty slice with nil Tags",
			slices: []PostCategory{
				{},
			},
			extractFunc: func(idx int, s PostCategory) []string { return s.Tags },
			want:        []string{},
		},
		{
			name: "non-empty slice with empty Tags",
			slices: []PostCategory{
				{Tags: []string{}},
			},
			extractFunc: func(idx int, s PostCategory) []string { return s.Tags },
			want:        []string{},
		},
		{
			name: "non-empty slice with non-empty Tags",
			slices: []PostCategory{
				{Tags: []string{"Go", "Java"}},
				{Tags: []string{"Vue", "React", "Go"}},
			},
			extractFunc: func(idx int, s PostCategory) []string { return s.Tags },
			want:        []string{"Go", "Java", "Vue", "React"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, CombineAndDeduplicateNestedSlices(tc.slices, tc.extractFunc))
		})
	}
}

func TestCombineAndDeduplicateNestedSlicesByEqFunc(t *testing.T) {
	type Tag struct {
		Id   int
		Name string
	}
	type PostCategory struct {
		Tags []Tag
	}

	testCases := []struct {
		name        string
		slices      []PostCategory
		extractFunc func(idx int, s PostCategory) []Tag
		eqFunc      func(a, b Tag) bool
		want        []Tag
	}{
		{
			name: "nil slice",
			want: []Tag{},
		},
		{
			name:   "empty slice",
			slices: []PostCategory{},
			want:   []Tag{},
		},
		{
			name: "non-empty slice with nil Tags",
			slices: []PostCategory{
				{},
			},
			extractFunc: func(idx int, s PostCategory) []Tag { return s.Tags },
			eqFunc:      func(a, b Tag) bool { return a.Name == b.Name },
			want:        []Tag{},
		},
		{
			name: "non-empty slice with empty Tags",
			slices: []PostCategory{
				{Tags: []Tag{}},
			},
			extractFunc: func(idx int, s PostCategory) []Tag { return s.Tags },
			eqFunc:      func(a, b Tag) bool { return a.Name == b.Name },
			want:        []Tag{},
		},
		{
			name: "non-empty slice with non-empty Tags",
			slices: []PostCategory{
				{Tags: []Tag{{Id: 1, Name: "Go"}, {Id: 2, Name: "Java"}}},
				{Tags: []Tag{{Id: 3, Name: "Vue"}, {Id: 4, Name: "React"}, {Id: 5, Name: "Go"}}},
			},
			extractFunc: func(idx int, s PostCategory) []Tag { return s.Tags },
			eqFunc:      func(a, b Tag) bool { return a.Name == b.Name },
			want:        []Tag{{Id: 1, Name: "Go"}, {Id: 2, Name: "Java"}, {Id: 3, Name: "Vue"}, {Id: 4, Name: "React"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, CombineAndDeduplicateNestedSlicesByEqFunc(tc.slices, tc.extractFunc, tc.eqFunc))
		})
	}
}

func TestIndexStructsByKey(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}

	testCases := []struct {
		name         string
		data         []User
		keyExtractor func(s User) int
		want         map[int]User
	}{
		{
			name: "nil slice",
			want: map[int]User{},
		},
		{
			name: "empty slice",
			data: []User{},
			want: map[int]User{},
		},
		{
			name: "non-empty slice",
			data: []User{
				{Id: 1, Name: "陈明勇"},
				{Id: 2, Name: "Gopher"},
			},
			keyExtractor: func(s User) int { return s.Id },
			want: map[int]User{
				1: {Id: 1, Name: "陈明勇"},
				2: {Id: 2, Name: "Gopher"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, IndexStructsByKey[User, int](tc.data, tc.keyExtractor))
		})
	}
}

func TestFilterMap(t *testing.T) {
	type user struct {
		id   int
		name string
	}

	testCases := []struct {
		name string
		src  []user
		fn   func(idx int, s user) (int, bool)
		want []int
	}{
		{
			name: "nil slice",
			src:  nil,
			fn:   func(idx int, s user) (int, bool) { return s.id, true },
			want: []int{},
		},
		{
			name: "empty slice",
			src:  []user{},
			fn:   func(idx int, s user) (int, bool) { return s.id, true },
			want: []int{},
		},
		{
			name: "non-empty slice",
			src: []user{
				{id: 1, name: "陈明勇"},
				{id: 2, name: "Gopher"},
			},
			fn: func(idx int, s user) (int, bool) {
				if s.name == "陈明勇" {
					return s.id, true
				}
				return 0, false
			},
			want: []int{1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, FilterMap(tc.src, tc.fn))
		})
	}
}
