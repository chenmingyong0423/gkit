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

// Contains checks whether the given slice contains the specified element.
// Parameters:
// - data: the slice to search
// - item: the element to search for
//
// Returns:
// - true if the specified element is found; false otherwise
//
// Contains 检查给定的 slice 中是否包含指定元素
// 参数：
// - data: 要搜索的 slice
// - item: 要查找的元素
//
// 返回值：
// - 如果找到了指定的元素，则为 true；否则为 false
func Contains[T comparable](data []T, item T) bool {
	return ContainsByFunc(data, item, func(srcItem, dstItem T) bool {
		return srcItem == dstItem
	})
}

// ContainsByFunc checks whether the given slice contains the specified element using the provided equality function.
// Parameters:
// - data: the slice to search
// - dstItem: the element to search for
// - equalFunc: the function used to compare element equality
//
// Returns:
// - true if the specified element is found; false otherwise
//
// ContainsByFunc 通过提供的相等函数，检查给定的 slice 中是否包含指定元素
// 参数：
// - data: 要搜索的 slice
// - dstItem: 要查找的元素
// - equalFunc: 用于比较元素相等性的函数
//
// 返回值：
// - 如果找到了指定的元素，则为 true；否则为 false
func ContainsByFunc[T any](data []T, dstItem T, equalFunc equalFunc[T]) bool {
	for _, srcItem := range data {
		if equalFunc(srcItem, dstItem) {
			return true
		}
	}
	return false
}
