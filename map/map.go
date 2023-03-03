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

package maps

// Keys extracts all the keys from the given map and returns them as a slice.
// Parameters:
// - mp: the map to extract keys from
//
// Returns:
// - a slice containing all the keys in the given map
//
// Keys 从给定的 map 中提取所有的键，并将它们以 slice 的形式返回
// 参数：
// - mp: 要提取键的 map
// - Note: The order of keys is random
//
// 返回值：
// - 一个 slice，其中包含了给定 map 中的所有键
// - 注意：键的顺序是随机的
func Keys[K comparable, V any](mp map[K]V) []K {
	keys := make([]K, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	return keys
}

// Values extracts all the values from the given map and returns them as a slice.
// Parameters:
// - mp: the map to extract values from
//
// Returns:
// - a slice containing all the values in the given map
// - Note: The order of values is random
//
// Values 从给定的 map 中提取所有的值，并将它们以 slice 的形式返回
// 参数：
// - mp: 要提取值的 map
//
// 返回值：
// - 一个 slice，其中包含了给定 map 中的所有值
// - 注意：值的顺序是随机的
func Values[K comparable, V any](mp map[K]V) []V {
	values := make([]V, 0, len(mp))
	for _, v := range mp {
		values = append(values, v)
	}
	return values
}

// KeyValues extracts all the keys and values from the given map and returns them as separate slices.
// Parameters:
// - mp: the map to extract key-value pairs from
//
// Returns:
// - a slice containing all the keys in the given map, and a slice containing all the values in the given map
// - Note: The order of keys and values is random
//
// KeyValues 从给定的 map 中提取所有的键和值，并将它们分别以 slice 的形式返回
// 参数：
// - mp: 要提取键值对的 map
//
// 返回值：
// - 一个包含所有键的 slice，一个包含所有值的 slice
// - 注意：键和值的顺序是随机的
func KeyValues[K comparable, V any](mp map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(mp))
	values := make([]V, 0, len(mp))
	for k, v := range mp {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
