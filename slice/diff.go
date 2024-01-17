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

// Diff 计算两个切片的差集，该函数只支持 comparable 类型的切片元素
// 返回的新切片已去重且顺序不确定
func Diff[T comparable](s1 []T, s2 []T) []T {
	m := toMap[T](s1)
	for _, item := range s2 {
		delete(m, item)
	}
	result := make([]T, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// DiffFunc 计算两个切片的差集，由 equal 函数判断两个元素是否相等
// 返回的新切片已去重
func DiffFunc[T any](s1 []T, s2 []T, equal equalFunc[T]) []T {
	result := make([]T, 0, len(s1))
	for _, item := range s1 {
		if !ContainsByFunc[T](s2, item, equal) {
			result = append(result, item)
		}
	}
	return DeduplicateByEqFunc(result, equal)
}
