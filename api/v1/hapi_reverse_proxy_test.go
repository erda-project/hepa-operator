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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	v1 "github.com/erda-project/hepa-operator/api/v1"
	"github.com/erda-project/hepa-operator/pkg/mse"
)

func TestReverseProxyRule_PatchExternal(t *testing.T) {
	var hapi = v1.Hapi{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Hapi",
			APIVersion: "hepa.erda.cloud/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "hapi-sample",
			Namespace: "hapi-operator-sample",
			Labels: map[string]string{
				"configZone": "hapi-operator-sample",
				"packageId":  "c82396e5fc13ef7bbf6bc078502a21e4",
			},
		},
		Spec: v1.HapiSpec{
			Hosts: []string{
				"hapi-sample.mse-daily.terminus.io",
			},
			Path: "/",
			Backend: v1.Backend{
				RedirectBy:    "url",
				ServiceName:   "go-httpbin",
				ServicePort:   80,
				UpstreamHost:  "baidu.com",
				RewriteTarget: "/s",
			},
			Policy: v1.Policy{
				Auth: v1.Auth{
					AuthType: "sign-auth",
					Switch: v1.Switch{
						Global: false,
						Switch: true,
					},
				},
				CORS:       v1.CORS{},
				Metric:     v1.Metric{},
				Proxy:      v1.Proxy{},
				SafetyCSRF: v1.SafetyCSRF{},
				SafetyIP: v1.SafetyIP{
					IPType:               v1.XRealIP,
					KeyRateLimitingValue: "12 query_per_second",
					Switch: v1.Switch{
						Global: true,
						Switch: false,
					},
					WhiteListSourceRange: "123.45.67.1/18",
				},
				SafetySBAC:  v1.SafetySBAC{},
				ServerGuard: v1.ServerGuard{},
			},
		},
		Status: v1.HapiStatus{},
	}
	var global = v1.Policy{
		Auth: v1.Auth{
			AuthType: "sign-auth",
			Switch: v1.Switch{
				Global: false,
				Switch: true,
			},
		},
		CORS:       v1.CORS{},
		Metric:     v1.Metric{},
		Proxy:      v1.Proxy{},
		SafetyCSRF: v1.SafetyCSRF{},
		SafetyIP: v1.SafetyIP{
			IPType:               v1.XRealIP,
			KeyRateLimitingValue: "12 query_per_second",
			Switch: v1.Switch{
				Global: true,
				Switch: false,
			},
			BlackListSourceRange: "123.45.67.1/18",
		},
		SafetySBAC:  v1.SafetySBAC{},
		ServerGuard: v1.ServerGuard{},
	}
	var in = v1.NewReverseProxyRule(&hapi, &global)
	var ingress netv1.Ingress
	if err := in.PatchExternal(context.Background(), &ingress, new(mse.PluginRequest)); err != nil {
		t.Fatal(err)
	}
	data, err := yaml.Marshal(ingress)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
