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

import (
	"github.com/chenmingyong0423/gkit/internal/errors"
)

// DeleteByIndex is used to delete a specific index of a given type in a slice.
// Parameters:
// - data: The slice to operate on.
// - index: The index of the item to be deleted.
//
// Returns:
// - If the index value is within the valid range, a new slice containing only the elements that were not deleted is returned.
// - If the index value is out of range, a nil slice and an IndexOutOfRange error are returned.
//
// DeleteByIndex 用于删除给定类型的切片中特定索引的元素。
// 参数：
// - data：待操作的切片。
// - index：待删除元素的索引。
//
// 返回值：
// - 如果索引值在有效范围内，则返回一个新的切片，其中仅包含没有被删除的元素。
// - 如果索引值不在有效范围内，则返回一个nil切片和一个IndexOutOfRange错误。
func DeleteByIndex[T any](data []T, index int) ([]T, error) {
	length := len(data)
	if index < 0 || index >= length {
		return nil, errors.NewIndexOutOfRange(length, index)
	}
	return DeleteByFilterFunc[T](data, func(idx int, item T) bool {
		return idx == index
	}), nil
}

// DeleteByItem is used to delete specified elements in a slice of a given type.
// Parameters:
// - data: The slice to operate on.
// - dstItem: The item to be deleted.
//
// Returns:
// - A new slice containing only the elements that were not deleted.
//
// DeleteByItem 用于删除给定类型的切片中指定的元素。
// 参数：
// - data：待操作的切片。
// - dstItem：待删除元素。
//
// 返回值：
// - 一个新的切片，其中仅包含没有被删除的元素。
func DeleteByItem[T comparable](data []T, dstItem T) []T {
	return DeleteByFilterFunc[T](data, func(idx int, srcItem T) bool {
		return dstItem == srcItem
	})
}

// DeleteByFilterFunc is a generic function used to delete elements in a slice of a given type that meet a specific condition.
// Parameters:
// - data: The slice to operate on.
// - filterFunc: The function that represents the filter condition.
//
// Returns:
// - A new slice containing only the elements that were not filtered out by the filter function.
//
// DeleteByFilterFunc 是一个通用函数，用于删除给定类型的切片中符合特定条件的元素。
// 参数：
// - data：待操作的切片。
// - filterFunc：过滤条件的函数。
//
// 返回值：
// - 一个新的切片，其中仅包含没有被过滤器函数过滤掉的元素。
func DeleteByFilterFunc[T any](data []T, filterFunc filterFunc[T]) []T {
	pos := 0
	for idx := range data {
		if filterFunc(idx, data[idx]) {
			continue
		}
		data[pos] = data[idx]
		pos++
	}
	return data[:pos]
}
