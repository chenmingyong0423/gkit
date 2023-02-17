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

// Contains 判断 data 里是否包含 item
func Contains[T comparable](data []T, item T) bool {
	return ContainsByFunc(data, item, func(srcItem, dstItem T) bool {
		return srcItem == dstItem
	})
}

// ContainsByFunc 根据函数返回值判断 data 里是否包含 item
func ContainsByFunc[T any](data []T, dstItem T, equalFunc equalFunc[T]) bool {
	for _, srcItem := range data {
		if equalFunc(srcItem, dstItem) {
			return true
		}
	}
	return false
}
