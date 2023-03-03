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
// Deduplicate 从给定的 slice 中去除重复的元素，并返回一个新的 slice。
// 参数：
// - data: 要去重的 slice
//
// 返回值：
// - 去重后的新 slice
func Deduplicate[T comparable](data []T) []T {
	// 转成 map，自动去重
	mp := toMap[T](data)
	res := make([]T, 0, len(mp))
	for val := range mp {
		res = append(res, val)
	}
	return res
}

// DeduplicateByEqFunc function takes a generic slice "data" and an equivalence comparison function "equalFunc[T]", and returns a new slice with duplicate elements removed.
// Parameters:
// - data: a slice of any type "[]T" that needs to be deduplicated.
// - equalFunc: a function of type "equalFunc[T]" used to compare whether two elements are equal, where "T" can be any type.
//
// Return:
// - a new deduplicated slice of type "[]T".
//
// DeduplicateByEqFunc 函数接受一个泛型切片 data 和一个等价比较函数 equalFunc[T]，并返回一个新的切片，其中重复的元素已经被去除了。
// 参数：
// - data: 待去重的切片，类型为 []T，T 为任意类型；
// - equalFunc: 用于比较两个元素是否相等的函数，类型为 equalFunc[T]，T 为任意类型。
//
// 返回值：
// - 去重后的新切片，类型为 []T
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
