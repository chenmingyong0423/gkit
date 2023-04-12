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
	"errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestOnce_Do(t *testing.T) {
	testCases := []struct {
		name      string
		resource  map[string]string
		makeError bool

		wantErr      error
		wantResource map[string]string
	}{
		{
			name:      "no error",
			resource:  nil,
			makeError: false,

			wantErr: nil,
			wantResource: map[string]string{
				"k": "v",
			},
		},
		{
			name:      "has error",
			resource:  nil,
			makeError: true,

			wantErr:      errors.New("error"),
			wantResource: nil,
		},
	}
	for _, tt := range testCases {
		once := Once{}
		t.Run(tt.name, func(t *testing.T) {
			var err error
			var wg sync.WaitGroup
			f := func() error {
				if tt.makeError {
					return errors.New("error")
				}
				tt.resource = map[string]string{
					"k": "v",
				}
				return nil
			}
			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					err = once.Do(f)
				}()
			}
			wg.Wait()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantResource, tt.resource)
		})
	}
}
