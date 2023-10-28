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

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name string
		data []int
		dst  int
		want int
	}{
		{
			name: "nil slice",
			data: nil,
			dst:  1,
			want: -1,
		},
		{
			name: "empty slice",
			data: []int{},
			dst:  1,
			want: -1,
		},
		{
			name: "exist",
			data: []int{1, 2, 3, 4},
			dst:  2,
			want: 1,
		},
		{
			name: "find_first_of_two_elements",
			data: []int{1, 2, 2, 3},
			dst:  2,
			want: 1,
		},
		{
			name: "non-existent",
			data: []int{1, 2, 3, 4},
			dst:  0,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IndexOf(tt.data, tt.dst))
		})
	}
}

func TestIndexOfByFunc(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	tests := []struct {
		name      string
		data      []User
		dst       User
		equalFunc equalFunc[User]

		want int
	}{
		{
			name: "nil slice",
			data: nil,
			dst:  User{},
			want: -1,
		},
		{
			name: "empty slice",
			data: []User{},
			dst:  User{},
			want: -1,
		},
		{
			name: "exist",
			data: []User{
				{
					Id:   1,
					Name: "Burt",
				},
				{
					Id:   2,
					Name: "Jan",
				},
				{
					Id:   3,
					Name: "James",
				},
			},
			dst: User{
				Id:   2,
				Name: "James",
			},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			want: 2,
		},
		{
			name: "find_first_of_two_elements",
			data: []User{
				{
					Id:   1,
					Name: "Burt",
				},
				{
					Id:   2,
					Name: "James",
				},
				{
					Id:   3,
					Name: "James",
				},
			},
			dst: User{
				Id:   2,
				Name: "James",
			},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			want: 1,
		},
		{
			name: "non-existent",
			data: []User{
				{
					Id:   1,
					Name: "Burt",
				},
				{
					Id:   2,
					Name: "James",
				},
				{
					Id:   3,
					Name: "James",
				},
			},
			dst: User{
				Id:   4,
				Name: "Fred",
			},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IndexOfByFunc(tt.data, tt.dst, tt.equalFunc))
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	tests := []struct {
		name string
		data []int
		dst  int
		want int
	}{
		{
			name: "nil slice",
			data: nil,
			dst:  1,
			want: -1,
		},
		{
			name: "empty slice",
			data: []int{},
			dst:  1,
			want: -1,
		},
		{
			name: "exist",
			data: []int{1, 2, 3, 4},
			dst:  2,
			want: 1,
		},
		{
			name: "find_first_of_two_elements",
			data: []int{1, 2, 2, 3},
			dst:  2,
			want: 2,
		},
		{
			name: "non-existent",
			data: []int{1, 2, 3, 4},
			dst:  0,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LastIndexOf(tt.data, tt.dst))
		})
	}
}

func TestLastIndexOfByFunc(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	tests := []struct {
		name      string
		data      []User
		dst       User
		equalFunc equalFunc[User]

		want int
	}{
		{
			name: "nil slice",
			data: nil,
			dst:  User{},
			want: -1,
		},
		{
			name: "empty slice",
			data: []User{},
			dst:  User{},
			want: -1,
		},
		{
			name: "exist",
			data: []User{
				{
					Id:   1,
					Name: "Burt",
				},
				{
					Id:   2,
					Name: "Jan",
				},
				{
					Id:   3,
					Name: "James",
				},
			},
			dst: User{
				Id:   2,
				Name: "James",
			},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			want: 2,
		},
		{
			name: "find_first_of_two_elements",
			data: []User{
				{
					Id:   1,
					Name: "Burt",
				},
				{
					Id:   2,
					Name: "James",
				},
				{
					Id:   3,
					Name: "James",
				},
			},
			dst: User{
				Id:   2,
				Name: "James",
			},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			want: 2,
		},
		{
			name: "non-existent",
			data: []User{
				{
					Id:   1,
					Name: "Burt",
				},
				{
					Id:   2,
					Name: "james",
				},
				{
					Id:   3,
					Name: "James",
				},
			},
			dst: User{
				Id:   4,
				Name: "Fred",
			},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LastIndexOfByFunc(tt.data, tt.dst, tt.equalFunc))
		})
	}
}
