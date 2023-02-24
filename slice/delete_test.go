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

	"github.com/chenmingyong0423/gkit/internal/errors"

	"github.com/stretchr/testify/assert"
)

func TestDeleteByFilterFunc(t *testing.T) {

	type testCase[T any] struct {
		name       string
		data       []T
		filterFunc filterFunc[T]

		want []T
	}
	tests := []testCase[int]{
		{
			name: "空切片",
			data: []int{},
			filterFunc: func(idx int, src int) bool {
				return false
			},
			want: []int{},
		},
		{
			name: "不删除元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return false
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "删除首位元素",
			data: []int{0, 1, 2, 3, 4, 5, 6},
			filterFunc: func(idx int, src int) bool {
				return idx == 0
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "删除前面两个元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return idx == 0 || idx == 1
			},
			want: []int{2, 3, 4, 5, 6, 7},
		},
		{
			name: "删除中间单个元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return idx == 3
			},
			want: []int{0, 1, 2, 4, 5, 6, 7},
		},
		{
			name: "删除中间多个不连续元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return idx == 2 || idx == 4
			},
			want: []int{0, 1, 3, 5, 6, 7},
		},
		{
			name: "删除中间多个连续元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return idx == 3 || idx == 4
			},
			want: []int{0, 1, 2, 5, 6, 7},
		},
		{
			name: "删除中间多个元素，第一部分为一个元素，第二部分为连续元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return idx == 2 || idx == 4 || idx == 5
			},
			want: []int{0, 1, 3, 6, 7},
		},
		{
			name: "删除中间多个元素，第一部分为连续元素，第二部分为一个元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return idx == 2 || idx == 3 || idx == 5
			},
			want: []int{0, 1, 4, 6, 7},
		},
		{
			name: "删除后面两个元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return idx == 6 || idx == 7
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "删除末尾元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return idx == 7
			},
			want: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "删除所有元素",
			data: []int{0, 1, 2, 3, 4, 5, 6, 7},
			filterFunc: func(idx int, src int) bool {
				return true
			},

			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DeleteByFilterFunc(tt.data, tt.filterFunc))
		})
	}
}

func TestDeleteByIndex(t *testing.T) {
	type testCase[T any] struct {
		name    string
		data    []T
		index   int
		want    []T
		wantErr error
	}
	testCases := []testCase[int]{
		{
			name:    "下标越界，小于 0 的情况",
			data:    []int{1, 2, 3, 4},
			index:   -1,
			want:    nil,
			wantErr: errors.NewIndexOutOfRange(4, -1),
		},
		{
			name:    "下标越界，等于长度的情况",
			data:    []int{1, 2, 3, 4, 5},
			index:   5,
			want:    nil,
			wantErr: errors.NewIndexOutOfRange(5, 5),
		},
		{
			name:    "切片为空的情况",
			data:    []int{},
			index:   0,
			want:    nil,
			wantErr: errors.NewIndexOutOfRange(0, 0),
		},
		{
			name:    "从仅有一个元素的切片中删除一个元素",
			data:    []int{1},
			index:   0,
			want:    []int{},
			wantErr: nil,
		},
		{
			name:    "删除下标为 0 的元素",
			data:    []int{1, 2, 3},
			index:   0,
			want:    []int{2, 3},
			wantErr: nil,
		},
		{
			name:    "删除下标为 2 的元素",
			data:    []int{1, 2, 3},
			index:   2,
			want:    []int{1, 2},
			wantErr: nil,
		},
		{
			name:    "删除下标为 1 的元素",
			data:    []int{1, 2, 3},
			index:   1,
			want:    []int{1, 3},
			wantErr: nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res, err := DeleteByIndex(tt.data, tt.index)
			if tt.wantErr != nil || err != nil {
				assert.Equal(t, err, tt.wantErr)
			}
			assert.Equal(t, tt.want, res)
		})
	}
}
