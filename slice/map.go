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

// Deduplicate removes duplicate elements from the given slice and returns a new slice.
// This function has a time complexity of O(n), where n is the length of the input slice.
// Parameters:
// - data: the slice to deduplicate
//
// Returns:
// - the deduplicated new slice
//
// toMap 将给定的 slice 转换为一个 map，其中键为元素，值为 struct{}{}（一个空结构体）。
// 使用 struct{}{} 可以将 map 中的内存使用降至最小，因为它不会分配任何空间。
// 参数：
// - data: 要转换为 map 的 slice
//
// 返回值：
// - 转换后的 map
func toMap[T comparable](data []T) map[T]struct{} {
	res := make(map[T]struct{}, len(data))
	for _, k := range data {
		res[k] = struct{}{}
	}
	return res
}
