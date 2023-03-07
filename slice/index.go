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

// IndexOf searches for the given target element in the provided slice and return the index of its first occurrence.
// If the target element is not found, - 1 is returned.
//
// Parameters:
//   - data: The slice to search.
//   - dst: The target element to find.
//
// Returns:
//   - The index of the target element in the slice, or -1 if it is not found.
//
// Note: This function requires the element type to support the == operator.
// IndexOf 在提供的切片中查找给定目标元素，并返回其第一次出现的位置索引。如果未找到目标元素，则返回 -1。
//
// 参数：
//   - data: 要搜索的切片。
//   - dst: 要查找的目标元素。
//
// 返回：
//   - 目标元素在切片中的索引，如果未找到则返回 -1。
//
// 注意：该函数要求元素类型必须支持 == 操作符。
func IndexOf[T comparable](data []T, dst T) int {
	return IndexOfByFunc(data, dst, func(srcItem, dstItem T) bool {
		return srcItem == dstItem
	})
}

// IndexOfByFunc looks up the index of the first occurrence of a given target element in the provided slice based on the equality function.
// If the target element is found, its index in the slice is returned; Otherwise, it returns - 1.
//
// Parameters:
//   - data: The slice to search.
//   - dst: The target element to find.
//   - equalFunc: The custom function used to determine element equality.
//
// Returns:
//   - The index of the target element in the slice, or -1 if it is not found.
//
// IndexOfByFunc 根据相等性函数查找给定目标元素第一次出现在提供的切片中的索引。
// 如果找到目标元素，则返回其在切片中的索引；否则返回 -1。
//
// 参数：
//   - data: 要搜索的切片。
//   - dst: 要查找的目标元素。
//   - equalFunc: 用于判断元素相等性的自定义函数。
//
// 返回：
//   - 目标元素在切片中的索引，如果未找到则返回 -1。
func IndexOfByFunc[T any](data []T, dst T, equalFunc equalFunc[T]) int {
	for idx, item := range data {
		if equalFunc(item, dst) {
			return idx
		}
	}
	return -1
}

// LastIndexOf searches for the given target element in the provided slice, and returns the position index of its last occurrence.
// If the target element is not found, - 1 is returned.
//
// Parameters:
//   - data: The slice to search.
//   - dst: The target element to find.
//
// Returns:
//   - The index of the target element in the slice, or -1 if it is not found.
//
// Note: This function requires the element type to support the == operator.
// LastIndexOf 在提供的切片中查找给定目标元素，并返回其最后一次出现的位置索引。如果未找到目标元素，则返回 -1。
//
// 参数：
//   - data: 要搜索的切片。
//   - dst: 要查找的目标元素。
//
// 返回：
//   - 目标元素在切片中的索引，如果未找到则返回 -1。
//
// 注意：该函数要求元素类型必须支持 == 操作符。
func LastIndexOf[T comparable](data []T, dst T) int {
	return LastIndexOfByFunc(data, dst, func(srcItem, dstItem T) bool {
		return srcItem == dstItem
	})
}

// LastIndexOfByFunc finds the index of the last occurrence of the given target element in the provided slice according to the equality function.
// If the target element is found, its index in the slice is returned; Otherwise, it returns - 1.
//
// Parameter:
// -data: slices to be searched.
// -dst: The target element to look for.
// -equalFunc: A function used to compare elements for equality.
//
// Return:
// - Index of the target element in the slice, returns -1 if not found.
//
// LastIndexOfByFunc 根据相等性函数查找给定目标元素最后一次出现在提供的切片中的索引。
//
// 参数：
//   - data: 要搜索的切片。
//   - dst: 要查找的目标元素。
//   - equalFunc: 用于比较元素是否相等的函数。
//
// 返回：
//   - 目标元素在切片中的索引，如果未找到则返回 -1。
func LastIndexOfByFunc[T any](data []T, dst T, equalFunc equalFunc[T]) int {
	for i := len(data) - 1; i >= 0; i-- {
		if equalFunc(data[i], dst) {
			return i
		}
	}
	return -1
}
