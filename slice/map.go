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

// 将切片转换成 map
// 转化之后方便切片后续的某些操作，如去重、取交集等。
func toMap[T comparable](data []T) map[T]struct{} {
	res := make(map[T]struct{}, len(data))
	for _, k := range data {
		res[k] = struct{}{}
	}
	return res
}
