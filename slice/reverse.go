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

// Reverse Reverses the slice. The new slice is returned
// Parameter:
// -data raw slice
//
// Return value:
// - New slice
//
// Reverse 反转切片，返回的是新切片
// 参数：
//   - data 原切片
//
// 返回值：
//   - 新切片
func Reverse[T any](data []T) []T {
	res := make([]T, 0, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		res = append(res, data[i])
	}
	return res
}
