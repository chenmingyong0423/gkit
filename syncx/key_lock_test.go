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

package syncx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapKeyLock_Lock(t *testing.T) {
	lock := NewMapKeyLock()
	key1, key2 := "key1", "key2"
	lock.Lock(key1)
	assert.False(t, lock.TryLock(key1))
	assert.False(t, lock.TryRLock(key1))

	assert.True(t, lock.TryLock(key2))

	lock.Unlock(key1)

	assert.True(t, lock.TryLock(key1))
	defer lock.Unlock(key1)

	lock.Unlock(key2)
}

func TestMapKeyLock_RLock(t *testing.T) {
	lock := NewMapKeyLock()
	key1, key2 := "key1", "key2"

	lock.RLock(key1)
	assert.False(t, lock.TryLock(key1))
	assert.True(t, lock.TryRLock(key1))

	lock.RLock(key2)

	lock.RUnLock(key1)

	assert.False(t, lock.TryLock(key1))
	lock.RUnLock(key1)

	assert.True(t, lock.TryLock(key1))
	defer lock.Unlock(key1)

	lock.RUnLock(key2)
}
