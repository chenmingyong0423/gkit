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

// AppendDistinct appends elements of a given slice to the target slice, If s or items are nil, both are processed as empty slices.
// Parameter:
// -s: indicates the target slice.
// -items: items to append.
//
// Return value:
// - A new slice that contains elements of the original slice and the new slice, with no duplicate
//
// AppendDistinct  将给定的切片的元素追加到目标切片中，如果 s 或 items 为 nil，都按照 empty 切片来处理。
// 参数：
// - s：目标切片。
// - items：要追加的元素。
//
// 返回值：
// - 一个新的切片，其中包含原始切片和新切片的元素，元素无重复。
func AppendDistinct[T comparable](s []T, items ...T) []T {
	result := make([]T, 0, len(s)+len(items))
	for _, item := range s {
		result = append(result, item)
	}
	for _, item := range items {
		result = append(result, item)
	}
	return Deduplicate[T](result)
}

// AddDistinctFunc appends the element of the given slice to the target slice. If s or items are nil, they are processed as empty slices.
// Parameter:
// -s: indicates the target slice.
// -equal: A function used to determine whether two elements are equal, used when determining duplicate elements.
// -items: items to append.
//
// Return value:
// - A new slice that contains elements of the original slice and the new slice, with no duplicate elements.
//
// AddDistinctFunc  将给定的切片的元素追加到目标切片中，如果 s 或 items 为 nil，都按照 empty 切片来处理。
// 参数：
// - s：目标切片。
// - equal：用于判断两个元素是否相等的函数，判断重复元素时使用。
// - items：要追加的元素。
//
// 返回值：
// - 一个新的切片，其中包含原始切片和新切片的元素，元素无重复。
func AddDistinctFunc[T any](s []T, equal equalFunc[T], items ...T) []T {
	result := make([]T, 0, len(s)+len(items))
	for _, item := range s {
		result = append(result, item)
	}
	for _, item := range items {
		if !ContainsByFunc[T](result, item, equal) {
			result = append(result, item)
		}
	}
	return result
}
