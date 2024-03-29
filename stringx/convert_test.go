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
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestBigCamelToSmallCamel(t *testing.T) {

	testCases := []struct {
		name     string
		bigCamel string

		want string
	}{
		{
			name:     "empty string",
			bigCamel: "",

			want: "",
		},
		{
			name:     "Not Camel string",
			bigCamel: "User",

			want: "user",
		},
		{
			name:     "SmallCamel string",
			bigCamel: "userAgent",

			want: "userAgent",
		},
		{
			name:     "BigCamel string",
			bigCamel: "UserAgent",

			want: "userAgent",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, BigCamelToSmallCamel(tt.bigCamel))
		})
	}
}

func TestCapitalizeFirstLetter(t *testing.T) {
	testCases := []struct {
		name  string
		input string

		want string
	}{
		{
			name:  "empty string",
			input: "",

			want: "",
		},
		{
			name:  "Not capital string",
			input: "user",

			want: "User",
		},
		{
			name:  "Capital string",
			input: "User",

			want: "User",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CapitalizeFirstLetter(tt.input))
		})
	}
}
