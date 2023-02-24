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

// Deduplicate 去除重复元素
func Deduplicate[T comparable](data []T) []T {
	// 转成 map，自动去重
	mp := toMap[T](data)
	res := make([]T, 0, len(mp))
	for val := range mp {
		res = append(res, val)
	}
	return res
}

// DeduplicateByEqFunc 如果切片的元素类型比较复杂（如结构体），可以使用这个方法进行去重，去重条件由调用方自定义
// 只保留第一个出现的元素
func DeduplicateByEqFunc[T any](data []T, equalFunc equalFunc[T]) []T {
	res := make([]T, 0, len(data))
	for i, val := range data {
		// 如果 res 没有相同元素，则添加
		if !ContainsByFunc[T](res[:i], val, equalFunc) {
			res = append(res, val)
		}
	}
	return res
}
