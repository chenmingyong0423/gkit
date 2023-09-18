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

package reflectx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPtr(t *testing.T) {
	testCases := []struct {
		name  string
		value any

		want bool
	}{
		{
			name:  "整数",
			value: 1,
			want:  false,
		},
		{
			name: "整数指针",
			value: func() any {
				value := 1
				return &value
			}(),
			want: true,
		},
		{
			name: "整数指针的指针",
			value: func() any {
				value := 1
				value2 := &value
				return &value2
			}(),
			want: true,
		},
		{
			name:  "nil",
			value: nil,
			want:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, IsPtr(tc.value))
		})
	}
}
