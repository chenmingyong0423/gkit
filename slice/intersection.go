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

// IntersectionSet 从给定的两个切片中取交集，返回的新切片已去重，该函数只支持 comparable 类型的切片元素
func IntersectionSet[T comparable](slice1 []T, slice2 []T) []T {
	ms := toMap[T](slice1)
	result := make([]T, 0, len(slice1))
	for _, item := range slice2 {
		if _, ok := ms[item]; ok {
			result = append(result, item)
		}
	}
	return Deduplicate[T](result)
}

// IntersectionSetFunc 从给定的两个切片中取交集，由 equal 函数判断两个元素是否相等，返回的新切片已去重
func IntersectionSetFunc[T any](slice1 []T, slice2 []T, equal equalFunc[T]) []T {
	result := make([]T, 0, len(slice1))
	for i := range slice1 {
		for j := range slice2 {
			if equal(slice1[i], slice2[j]) {
				result = append(result, slice1[i])
				break
			}
		}
	}
	return DeduplicateByEqFunc[T](result, equal)
}
