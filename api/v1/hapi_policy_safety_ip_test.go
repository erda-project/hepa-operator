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
	"context"
	"testing"

	netv1 "k8s.io/api/networking/v1"
	"sigs.k8s.io/yaml"

	v1 "github.com/erda-project/hepa-operator/api/v1"
)

func TestSafetyIP_SwitchOnForIngress(t *testing.T) {
	var ingress netv1.Ingress
	var ipList = "12.12.12.1/18"
	var item = &v1.SafetyIP{
		Switch: v1.Switch{
			Global: false,
			Switch: true,
		},
		IPType:                     v1.PeerIP,
		WhiteListSourceRange:       "",
		BlackListSourceRange:       ipList,
		DomainWhiteListSourceRange: "",
		DomainBlackListSourceRange: "",
		KeyRateLimitingValue:       "12 query_per_second",
	}

	t.Run("blackList", func(t *testing.T) {
		if err := item.SwitchOnForIngress(context.Background(), &ingress); err != nil {
			t.Fatalf("failed to SwitchOnForIngress: %v", err)
		}
		data, err := yaml.Marshal(ingress)
		if err != nil {
			t.Fatalf("failed to yaml.Marshal(ingress): %s", err)
		}
		t.Log(string(data))
		ips, ok := ingress.Annotations[v1.BlackListSourceRange.String()]
		if !ok {
			t.Fatalf("error missing annotation %s", v1.BlackListSourceRange)
		}
		if ips != ipList {
			t.Fatalf("error annotation %s's value, expect: %s, got: %s", v1.BlackListSourceRange, ipList, ips)
		}
		_, ok = ingress.Annotations[v1.WhiteListSourceRange.String()]
		if ok {
			t.Fatalf("%s should be deleted", v1.WhiteListSourceRange)
		}
	})

	t.Run("whiteList", func(t *testing.T) {
		item.WhiteListSourceRange = ipList
		item.BlackListSourceRange = ""
		if err := item.SwitchOnForIngress(context.Background(), &ingress); err != nil {
			t.Fatalf("failed to SwitchOnForIngress: %v", err)
		}
		data, err := yaml.Marshal(ingress)
		if err != nil {
			t.Fatalf("failed to yaml.Marshal(ingress): %s", err)
		}
		t.Log(string(data))
		ips, ok := ingress.Annotations[v1.WhiteListSourceRange.String()]
		if !ok {
			t.Fatalf("error missing annotation %s", v1.WhiteListSourceRange)
		}
		if ips != ipList {
			t.Fatalf("error annotation %s's value, expect: %s, got: %s", v1.WhiteListSourceRange, ipList, ips)
		}
		_, ok = ingress.Annotations[v1.BlackListSourceRange.String()]
		if ok {
			t.Fatalf("%s should be deleted", v1.BlackListSourceRange)
		}
	})

	t.Run("blackList", func(t *testing.T) {
		item.WhiteListSourceRange = ""
		item.BlackListSourceRange = ipList
		if err := item.SwitchOnForIngress(context.Background(), &ingress); err != nil {
			t.Fatalf("failed to SwitchOnForIngress: %v", err)
		}
		data, err := yaml.Marshal(ingress)
		if err != nil {
			t.Fatalf("failed to yaml.Marshal(ingress): %s", err)
		}
		t.Log(string(data))
		ips, ok := ingress.Annotations[v1.BlackListSourceRange.String()]
		if !ok {
			t.Fatalf("error missing annotation %s", v1.BlackListSourceRange)
		}
		if ips != ipList {
			t.Fatalf("error annotation %s's value, expect: %s, got: %s", v1.BlackListSourceRange, ipList, ips)
		}
		_, ok = ingress.Annotations[v1.WhiteListSourceRange.String()]
		if ok {
			t.Fatalf("%s should be deleted", v1.WhiteListSourceRange)
		}
	})
}
