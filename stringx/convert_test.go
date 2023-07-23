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

package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCamelToSnake(t *testing.T) {
	testCases := []struct {
		name      string
		camelCase string

		want string
	}{
		{
			name:      "empty string",
			camelCase: "",

			want: "",
		},
		{
			name:      "Not camel string",
			camelCase: "User",

			want: "user",
		},
		{
			name:      "Camel string",
			camelCase: "UserAgent",

			want: "user_agent",
		},
		{
			name:      "Complex camel string",
			camelCase: "MyUserAgent",

			want: "my_user_agent",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			snake := CamelToSnake(tt.camelCase)
			assert.Equal(t, tt.want, snake)
		})
	}
}
