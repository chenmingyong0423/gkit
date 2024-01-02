// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uuidx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID4(t *testing.T) {
	t.Run("uuid4", func(t *testing.T) {
		uuid := UUID4()
		assert.NotEmpty(t, uuid)
	})
}

func TestRearrangeUUID(t *testing.T) {
	testCases := []struct {
		name      string
		uuid      string
		want      string
		wantError error
	}{
		{
			name:      "invalid uuid",
			uuid:      "58e0a7d7-eebc-11d8-9669-0800200c9a66-1234",
			want:      "",
			wantError: errors.New("invalid uuid"),
		},
		{
			name: "valid uuid",
			uuid: "58e0a7d7-eebc-11d8-9669-0800200c9a66",
			want: "11d8eebc58e0a7d796690800200c9a66",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := RearrangeStrUUID4(tc.uuid)
			if tc.wantError != nil {
				assert.Equal(t, err, tc.wantError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestRearrangeUUID4(t *testing.T) {
	t.Run("rearrange uuid4", func(t *testing.T) {
		uuid, err := RearrangeUUID4()
		assert.NoError(t, err)
		assert.NotEmpty(t, uuid)
	})
}
