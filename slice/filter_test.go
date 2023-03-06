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

func TestFilter(t *testing.T) {
	testCases := []struct {
		name       string
		data       []int
		filterFunc filterFunc[int]
		want       []int
	}{
		{
			name: "'Odd number'",
			data: []int{1, 2, 3, 4, 5, 6},
			filterFunc: func(idx int, item int) bool {
				return item%2 == 1
			},
			want: []int{1, 3, 5},
		},
		{
			name: "Even number",
			data: []int{1, 2, 3, 4, 5, 6},
			filterFunc: func(idx int, item int) bool {
				return item%2 == 0
			},
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Filter(tt.data, tt.filterFunc))
		})
	}
}
