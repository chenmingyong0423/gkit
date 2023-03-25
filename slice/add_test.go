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

func TestAppendDistinct(t *testing.T) {

	testCases := []struct {
		name  string
		s     []int
		items []int
		want  []int
	}{
		{
			name:  "nil s",
			s:     nil,
			items: []int{1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "nil items",
			s:     []int{1, 2},
			items: nil,
			want:  []int{1, 2},
		},
		{
			name:  "nil items and s",
			s:     nil,
			items: nil,
			want:  []int{},
		},
		{
			name:  "empty items",
			s:     []int{1, 2},
			items: []int{},
			want:  []int{1, 2},
		},
		{
			name:  "empty s",
			s:     []int{},
			items: []int{1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "empty items and s",
			s:     []int{},
			items: []int{},
			want:  []int{},
		},
		{
			name:  "normal s and items",
			s:     []int{1, 2},
			items: []int{3, 4},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "contains dup item",
			s:     []int{1, 2},
			items: []int{1, 2, 3, 4},
			want:  []int{1, 2, 3, 4},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AppendDistinct(tt.s, tt.items...))
		})
	}
}

func TestAddDistinctFunc(t *testing.T) {

	type user struct {
		Name      string
		Telephone string
	}

	testCases := []struct {
		name  string
		s     []user
		equal equalFunc[user]
		items []user

		want []user
	}{
		{
			name: "nil s",
			s:    nil,
			equal: func(srcItem, dstItem user) bool {
				return srcItem.Telephone == dstItem.Telephone
			},
			items: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
		},
		{
			name: "nil items",
			s: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.Telephone == dstItem.Telephone
			},
			items: nil,
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
		},
		{
			name: "nil items and s",
			s:    nil,
			equal: func(srcItem, dstItem user) bool {
				return srcItem.Telephone == dstItem.Telephone
			},
			items: nil,
			want:  []user{},
		},
		{
			name: "empty items",
			s: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.Telephone == dstItem.Telephone
			},
			items: []user{},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
		},
		{
			name: "empty s",
			s:    []user{},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.Telephone == dstItem.Telephone
			},
			items: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
		},
		{
			name: "empty items and s",
			s:    []user{},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.Telephone == dstItem.Telephone
			},
			items: []user{},
			want:  []user{},
		},
		{
			name: "normal s and items",
			s: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.Telephone == dstItem.Telephone
			},
			items: []user{
				{
					Name:      "u2",
					Telephone: "456",
				},
			},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "456",
				},
			},
		},
		{
			name: "contains dup item",
			s: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
			},
			equal: func(srcItem, dstItem user) bool {
				return srcItem.Telephone == dstItem.Telephone
			},
			items: []user{
				{
					Name:      "u2",
					Telephone: "456",
				},
				{
					Name:      "u3",
					Telephone: "123",
				},
			},
			want: []user{
				{
					Name:      "u1",
					Telephone: "123",
				},
				{
					Name:      "u2",
					Telephone: "456",
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AddDistinctFunc(tt.s, tt.equal, tt.items...))
		})
	}
}
