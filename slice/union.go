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

// Union 计算多个切片的并集，该函数只支持 comparable 类型的切片元素
func Union[T comparable](slices ...[]T) []T {
	unionMap := make(map[T]struct{})
	for _, s := range slices {
		for _, item := range s {
			unionMap[item] = struct{}{}
		}
	}
	return mapKeyToSlice[T](unionMap)
}

// UnionFunc 计算多个切片的并集，支持任意类型的切片元素
// 由 equal 参数决定切片元素的相等规则
func UnionFunc[T any](equal equalFunc[T], slices ...[]T) []T {
	totalLength := 0
	for _, s := range slices {
		totalLength += len(s)
	}
	result := make([]T, 0, totalLength)
	for _, s := range slices {
		result = append(result, s...)
	}
	return DeduplicateByEqFunc[T](result, equal)
}
