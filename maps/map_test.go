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

package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	testCases := []struct {
		name string
		mp   map[int]int
		want []int
	}{
		{
			name: "nil",
			mp:   nil,
			want: make([]int, 0),
		},
		{
			name: "empty",
			mp:   make(map[int]int, 0),
			want: make([]int, 0),
		},
		{
			name: "3 element",
			mp:   map[int]int{1: 1, 2: 2, 3: 3},
			want: []int{1, 2, 3},
		},
		{
			name: "4 element",
			mp:   map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Keys[int, int](tt.mp)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestValues(t *testing.T) {
	testCases := []struct {
		name string
		mp   map[int]string
		want []string
	}{
		{
			name: "nil",
			mp:   nil,
			want: make([]string, 0),
		},
		{
			name: "empty",
			mp:   make(map[int]string, 0),
			want: make([]string, 0),
		},
		{
			name: "3 element",
			mp:   map[int]string{1: "a", 2: "b", 3: "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "4 element",
			mp:   map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
			want: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Values[int, string](tt.mp)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestKeyValues(t *testing.T) {
	testCases := []struct {
		name       string
		mp         map[int]string
		wantKeys   []int
		wantValues []string
	}{
		{
			name:       "nil",
			mp:         nil,
			wantKeys:   make([]int, 0),
			wantValues: make([]string, 0),
		},
		{
			name:       "empty",
			mp:         make(map[int]string, 0),
			wantKeys:   make([]int, 0),
			wantValues: make([]string, 0),
		},
		{
			name:       "3 element",
			mp:         map[int]string{1: "a", 2: "b", 3: "c"},
			wantKeys:   []int{1, 2, 3},
			wantValues: []string{"a", "b", "c"},
		},
		{
			name:       "4 element",
			mp:         map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
			wantKeys:   []int{1, 2, 3, 4},
			wantValues: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			keys, values := KeyValues[int, string](tt.mp)
			assert.ElementsMatch(t, tt.wantKeys, keys)
			assert.ElementsMatch(t, tt.wantValues, values)
		})
	}
}
