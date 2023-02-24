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

// DeleteByIndex 删除下标为 idx 的元素
func DeleteByIndex[T any](data []T, index int) ([]T, error) {
	length := len(data)
	if index < 0 || index >= length {
		return nil, errors.NewIndexOutOfRange(length, index)
	}
	return DeleteByFilterFunc[T](data, func(idx int, item T) bool {
		return idx == index
	}), nil
}

// DeleteByItem 删除指定元素 item
func DeleteByItem[T comparable](data []T, dstItem T) []T {
	return DeleteByFilterFunc[T](data, func(idx int, srcItem T) bool {
		return dstItem == srcItem
	})
}

// DeleteByFilterFunc 删除符合条件的元素
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
