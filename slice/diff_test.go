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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	testCases := []struct {
		name   string
		s1     []int
		s2     []int
		expect []int
	}{
		{
			name:   "nil slice",
			s1:     nil,
			s2:     nil,
			expect: []int{},
		},
		{
			name:   "empty slice",
			s1:     []int{},
			s2:     []int{},
			expect: []int{},
		},
		{
			name:   "s1 is nil",
			s1:     nil,
			s2:     []int{1, 2, 3},
			expect: []int{},
		},
		{
			name:   "s2 is nil",
			s1:     []int{1, 2, 3},
			s2:     nil,
			expect: []int{1, 2, 3},
		},
		{
			name:   "s1 is empty",
			s1:     []int{},
			s2:     []int{1, 2, 3},
			expect: []int{},
		},
		{
			name:   "s2 is empty",
			s1:     []int{1, 2, 3},
			s2:     []int{},
			expect: []int{1, 2, 3},
		},
		{
			name:   "s1 is subset of s2",
			s1:     []int{1, 2},
			s2:     []int{1, 2, 3},
			expect: []int{},
		},
		{
			name:   "s2 is subset of s1",
			s1:     []int{1, 2, 3},
			s2:     []int{1, 2},
			expect: []int{3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.expect, Diff(tc.s1, tc.s2))
		})
	}
}

func TestDiffFunc(t *testing.T) {
	type User struct {
		Id   int64
		Name string
	}
	testCases := []struct {
		name      string
		s1        []User
		s2        []User
		equalFunc equalFunc[User]
		expect    []User
	}{
		{
			name:      "nil slice",
			s1:        nil,
			s2:        nil,
			equalFunc: nil,
			expect:    []User{},
		},
		{
			name:      "empty slice",
			s1:        []User{},
			s2:        []User{},
			equalFunc: nil,
			expect:    []User{},
		},
		{
			name:      "s1 is nil",
			s1:        nil,
			s2:        []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}, {Id: 3, Name: "James"}},
			equalFunc: nil,
			expect:    []User{},
		},
		{
			name: "s2 is nil",
			s1:   []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}, {Id: 3, Name: "James"}},
			s2:   nil,
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			expect: []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}, {Id: 3, Name: "James"}},
		},
		{
			name:      "s1 is empty",
			s1:        []User{},
			s2:        []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}, {Id: 3, Name: "James"}},
			equalFunc: nil,
			expect:    []User{},
		},
		{
			name: "s2 is empty",
			s1:   []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}, {Id: 3, Name: "James"}},
			s2:   []User{},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			expect: []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}, {Id: 3, Name: "James"}},
		},
		{
			name: "s1 is subset of s2",
			s1:   []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}},
			s2:   []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}, {Id: 3, Name: "James"}},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			expect: []User{},
		},
		{
			name: "s2 is subset of s1",
			s1:   []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}, {Id: 3, Name: "James"}},
			s2:   []User{{Id: 1, Name: "Burt"}, {Id: 2, Name: "Jan"}},
			equalFunc: func(srcItem, dstItem User) bool {
				return srcItem.Name == dstItem.Name
			},
			expect: []User{{Id: 3, Name: "James"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ElementsMatch(t, tc.expect, DiffFunc(tc.s1, tc.s2, tc.equalFunc))
		})
	}
}
