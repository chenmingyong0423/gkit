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

// Filter filters the slice element according to the filter function and returns a new slice.
//
// Parameters:
//   - data: A slice of type T that contains the elements to be filtered.
//   - filterFunc: A function indicating whether the element should be included in the filtered slice.
//
// Returns:
//   - A new slice of type T containing the elements that pass the filter function.
//
// # Filter 根据过滤函数对切片元素进行过滤，并返回一个新切片。
//
// 参数：
//   - data：T 类型的切片，其中包含要过滤的元素。
//   - filterFunc：一个函数，表示是否应将该元素包含在过滤后的切片中。
//
// 返回值：
//   - 一个新的 T 类型切片，其中包含通过过滤函数的元素。
func Filter[T any](data []T, filterFunc filterFunc[T]) []T {
	res := make([]T, 0)
	for idx, item := range data {
		if filterFunc(idx, item) {
			res = append(res, item)
		}
	}
	return res
}
