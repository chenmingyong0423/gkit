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
	"strings"
	"unicode"
)

// CamelToSnake 驼峰命名转成蛇形命名，如果不是驼峰命名，则转成对应小写字符串
// UserAgent → user_agent
// User → user
func CamelToSnake(camelCase string) string {
	var result strings.Builder
	for i, c := range camelCase {
		if unicode.IsUpper(c) && i > 0 {
			result.WriteByte('_')
		}
		result.WriteRune(unicode.ToLower(c))
	}
	return result.String()
}

// BigCamelToSmallCamel 大驼峰格式的字符串转小驼峰格式的字符串
// UserAgent → userAgent
func BigCamelToSmallCamel(bigCamel string) string {
	if len(bigCamel) == 0 {
		return ""
	}

	firstChar := strings.ToLower(string(bigCamel[0]))
	return firstChar + bigCamel[1:]
}

// CapitalizeFirstLetter 将首字母转换为大写
// user → User
func CapitalizeFirstLetter(input string) string {
	if len(input) == 0 {
		return input
	}

	firstChar := []rune(input)[:1]
	firstCharUpper := string(unicode.ToUpper(firstChar[0]))

	return firstCharUpper + input[1:]
}
