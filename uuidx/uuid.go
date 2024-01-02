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
	"strings"

	"github.com/google/uuid"
)

func UUID4() string {
	return uuid.New().String()
}

func RearrangeStrUUID4(uuid string) (string, error) {
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		return "", errors.New("invalid uuid")
	}
	return parts[2] + parts[1] + parts[0] + parts[3] + parts[4], nil
}

func RearrangeUUID4() (string, error) {
	return RearrangeStrUUID4(UUID4())
}
