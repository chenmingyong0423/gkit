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

package gkit

// ToPtr 取地址
func ToPtr[T any](t T) *T {
	return &t
}

// GetValueOrDefault 获取指针指向的值，如果为 nil，则返回默认值
// 注意：只会取第一层指针所指向的值
func GetValueOrDefault[T any](src *T) (t T) {
	if src != nil {
		t = *src
	}
	return
}
