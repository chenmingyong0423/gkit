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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicate(t *testing.T) {
	testCases := []struct {
		name string
		data []int
		want []int
	}{
		{
			name: "nil slice",
			data: nil,
			want: []int{},
		},
		{
			name: "空 slice",
			data: []int{},
			want: []int{},
		},
		{
			name: "无重复元素的 slice",
			data: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "有重复元素的 slice",
			data: []int{1, 2, 2, 3},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Deduplicate(tt.data))
		})
	}
}

func TestDeduplicateByEqFunc(t *testing.T) {
	type User struct {
		Id   int
		Name string
		Age  int
	}
	testCases := []struct {
		name      string
		data      []User
		equalFunc equalFunc[User]
		want      []User
	}{
		{
			name: "nil slice",
			data: nil,
			want: []User{},
		},
		{
			name: "空 slice",
			data: []User{},
			want: []User{},
		},
		{
			name: "去除 User.Name 相等的元素（切片无变化）",
			data: []User{
				{
					Id:   1,
					Name: "Tom",
					Age:  18,
				},
				{
					Id:   2,
					Name: "Michal",
					Age:  18,
				},
			},
			equalFunc: func(srcItem, dstItem User) bool {
				fmt.Println(srcItem.Name, dstItem.Name)
				return srcItem.Name == dstItem.Name
			},
			want: []User{
				{
					Id:   1,
					Name: "Tom",
					Age:  18,
				},
				{
					Id:   2,
					Name: "Michal",
					Age:  18,
				},
			},
		},
		{
			name: "去除 User.Name 相等的元素",
			data: []User{
				{
					Id:   1,
					Name: "Tom",
					Age:  18,
				},
				{
					Id:   2,
					Name: "Michal",
					Age:  18,
				},
				{
					Id:   3,
					Name: "Michal",
					Age:  19,
				},
			},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			want: []User{
				{
					Id:   1,
					Name: "Tom",
					Age:  18,
				},
				{
					Id:   2,
					Name: "Michal",
					Age:  18,
				},
			},
		},
		{
			name: "去除 User.Age 相等的元素",
			data: []User{
				{
					Id:   1,
					Name: "Tom",
					Age:  18,
				},
				{
					Id:   2,
					Name: "Michal",
					Age:  18,
				},
				{
					Id:   3,
					Name: "Michal",
					Age:  19,
				},
			},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Age == dstItem.Age
			},
			want: []User{
				{
					Id:   1,
					Name: "Tom",
					Age:  18,
				},
				{
					Id:   3,
					Name: "Michal",
					Age:  19,
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DeduplicateByEqFunc(tt.data, tt.equalFunc))
		})
	}
}
