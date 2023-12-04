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

import "sync"

type Locker interface {
	Lock(key string)
	Unlock(key string)
	RLock(key string)
	RUnLock(key string)
	TryLock(key string) bool
	TryRLock(key string) bool
}

type MapKeyLock struct {
	locks sync.Map
}

func (l *MapKeyLock) Lock(key string) {
	mu, _ := l.locks.LoadOrStore(key, &sync.RWMutex{})
	mu.(*sync.RWMutex).Lock()
}

func (l *MapKeyLock) Unlock(key string) {
	if mu, ok := l.locks.Load(key); ok {
		mu.(*sync.RWMutex).Unlock()
	}
}

func (l *MapKeyLock) RLock(key string) {
	mu, _ := l.locks.LoadOrStore(key, &sync.RWMutex{})
	mu.(*sync.RWMutex).RLock()
}

func (l *MapKeyLock) RUnLock(key string) {
	if mu, ok := l.locks.Load(key); ok {
		mu.(*sync.RWMutex).RUnlock()
	}
}

func (l *MapKeyLock) TryLock(key string) bool {
	mu, _ := l.locks.LoadOrStore(key, &sync.RWMutex{})
	return mu.(*sync.RWMutex).TryLock()
}

func (l *MapKeyLock) TryRLock(key string) bool {
	mu, _ := l.locks.LoadOrStore(key, &sync.RWMutex{})
	return mu.(*sync.RWMutex).TryRLock()
}
