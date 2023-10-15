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

func TestReverse(t *testing.T) {
	testCase := []struct {
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
			name: "empty slice",
			data: []int{},

			want: []int{},
		},
		{
			name: "slice with one element",
			data: []int{1},

			want: []int{1},
		},
		{
			name: "slice with many element",
			data: []int{1, 2, 3},

			want: []int{3, 2, 1},
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Reverse(tt.data))
		})
	}
}

func TestReverseInplace(t *testing.T) {
	testCase := []struct {
		name string
		data []int

		want []int
	}{
		{
			name: "nil slice",
			data: nil,

			want: nil,
		},
		{
			name: "empty slice",
			data: []int{},

			want: []int{},
		},
		{
			name: "slice with one element",
			data: []int{1},

			want: []int{1},
		},
		{
			name: "slice with many element",
			data: []int{1, 2, 3},

			want: []int{3, 2, 1},
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			ReverseInplace(tt.data)
			assert.Equal(t, tt.want, tt.data)
		})
	}
}
