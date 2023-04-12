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
	"sync"
	"sync/atomic"
)

// Once is similar to sync.Once and overrides the Do and doSlow methods with the error return value
// Once 类似 sync.Once，在其基础上重写 Do 和 doSlow 方法，新增 error 返回值
type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 0 {
		return o.doSlow(f)
	}
	return nil
}

func (o *Once) doSlow(f func() error) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if atomic.LoadUint32(&o.done) == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return
}
