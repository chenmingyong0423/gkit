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

// equalFunc is a function type used to compare whether two values of a given type are equal.
// Parameters:
// - srcItem: The first value to be compared.
// - dstItem: The second value to be compared.
// Returns:
// - If the two values are equal, return true; otherwise, return false.
//
// equalFunc 是一个函数类型，用于比较给定类型的两个值是否相等。
// 参数：
// - srcItem：待比较的第一个值。
// - dstItem：待比较的第二个值。
// 返回值：
// - 如果两个值相等，返回 true；否则返回 false。
type equalFunc[T any] func(srcItem, dstItem T) bool

// filterFunc is a function type used to determine whether an element of a given type meets a specific condition.
// Parameters:
// - idx: The index value of the element to be judged in the slice.
// - item: The value of the element to be judged.
// Returns:
// - If the element to be judged meets the condition, return true; otherwise, return false.
//
// filterFunc 是一个函数类型，用于判断给定类型的元素是否符合特定条件。
// 参数：
// - idx：待判断元素在切片中的索引值。
// - item：待判断元素的值。
// 返回值：
// - 如果待判断元素符合条件，返回true；否则返回false。
type filterFunc[T any] func(idx int, item T) bool
