// Copyright (c) 2022 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package v1_test

import (
	"testing"

	v1 "github.com/erda-project/hepa-operator/api/v1"
)

func TestPolicy_ListAll(t *testing.T) {
	var p v1.Policy
	m := p.ListAll()
	for k, v := range m {
		t.Logf("name: %s, global: %v, swtichOn: %v\n", k, v.GetGlobal(), v.GetSwitch())
	}
}
