// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lint

import (
	"strings"

	"testing"
)

const badCompositeDir = "rules/testdata/badcomposite"
const goodCompositeDir = "rules/testdata/goodcomposite"

func TestMissingServicesConfig(t *testing.T) {
	m := All(badCompositeDir).Messages
	if len(m) != 1 {
		t.Errorf("All didn't fail with expected errors, got %#v", m)
	}
	if !strings.Contains(m[0].Err.Error(), "file does not exist") {
		t.Errorf("All didn't have the error for file does not exist")
	}
}

func TestGoodComposite(t *testing.T) {
	m := All(goodCompositeDir).Messages
	if len(m) != 0 {
		t.Errorf("All failed but shouldn't have: %#v", m)
	}
}
