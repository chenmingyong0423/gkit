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

// IndexStructsByKey 将给定的结构体切片转换为 map，其中键为结构体的某个字段的值，值为结构体本身
func IndexStructsByKey[T any, K comparable](data []T, keyExtractor func(T) K) map[K]T {
	result := make(map[K]T, len(data))
	for _, item := range data {
		result[keyExtractor(item)] = item
	}
	return result
}

// CombineAndDeduplicateNestedSlicesByEqFunc 从任何类型的切片中提取并组合嵌套的切片，然后使用自定义的比较函数去重
func CombineAndDeduplicateNestedSlicesByEqFunc[Src any, Dst any](src []Src, extractFunc func(idx int, s Src) []Dst, eqFunc equalFunc[Dst]) []Dst {
	result := make([]Dst, 0)
	for i, s := range src {
		result = append(result, extractFunc(i, s)...)
	}
	return DeduplicateByEqFunc(result, eqFunc)
}

// CombineAndDeduplicateNestedSlices 从任何类型的切片中提取并组合嵌套的切片，然后去重
func CombineAndDeduplicateNestedSlices[Src any, Dst comparable](src []Src, extractFunc func(idx int, s Src) []Dst) []Dst {
	result := make([]Dst, 0)
	for i, s := range src {
		result = append(result, extractFunc(i, s)...)
	}
	return Deduplicate(result)
}

// CombineNestedSlices 从任何类型的切片中提取并组合嵌套的切片
func CombineNestedSlices[Src any, Dst any](src []Src, extractFunc func(idx int, s Src) []Dst) []Dst {
	result := make([]Dst, 0)
	for i, s := range src {
		result = append(result, extractFunc(i, s)...)
	}
	return result
}

// FilterMap 将给定的切片转换成一个新的切片，其中每个元素都是通过给定的函数 fn 转换得到的。如果 fn 返回的是 bool，则元素将被忽略。
func FilterMap[Src any, Dst any](src []Src, fn func(idx int, s Src) (Dst, bool)) []Dst {
	dst := make([]Dst, 0, len(src))
	for i, s := range src {
		if item, ok := fn(i, s); ok {
			dst = append(dst, item)
		}
	}
	return dst
}

// Map 将给定的切片转换成一个新的切片，其中每个元素都是通过给定的函数 fn 转换得到的。
func Map[Src any, Dst any](src []Src, fn func(idx int, s Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = fn(i, s)
	}
	return dst
}

// Deduplicate removes duplicate elements from the given slice and returns a new slice.
// This function has a time complexity of O(n), where n is the length of the input slice.
// Parameters:
// - data: the slice to deduplicate
//
// Returns:
// - the deduplicated new slice
//
// toMap 将给定的 slice 转换为一个 map，其中键为元素，值为 struct{}{}（一个空结构体）。
// 使用 struct{}{} 可以将 map 中的内存使用降至最小，因为它不会分配任何空间。
// 参数：
// - data: 要转换为 map 的 slice
//
// 返回值：
// - 转换后的 map
func toMap[T comparable](data []T) map[T]struct{} {
	res := make(map[T]struct{}, len(data))
	for _, k := range data {
		res[k] = struct{}{}
	}
	return res
}

// mapKeyToSlice 将给定的 map 的 key 转成切片
func mapKeyToSlice[T comparable, V any](data map[T]V) []T {
	result := make([]T, 0, len(data))
	for key := range data {
		result = append(result, key)
	}
	return result
}
