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

package v1

import (
	"context"
	"strconv"

	netv1 "k8s.io/api/networking/v1"

	"github.com/erda-project/hepa-operator/pkg/middleware"
)

type Proxy struct {
	Switch `json:",inline"`

	ProxyTimeout int  `json:"proxyTimeout,omitempty"`
	SslRedirect  bool `json:"sslRedirect,omitempty"`
}

func (in *Proxy) SwitchOnForIngress(ctx context.Context, ingress *netv1.Ingress) error {
	if len(ingress.Annotations) == 0 {
		ingress.Annotations = make(map[string]string)
	}
	ingress.Annotations[Timeout.String()] = strconv.Itoa(in.ProxyTimeout)
	ingress.Annotations[ForceSSLRedirect.String()] = strconv.FormatBool(in.SslRedirect)
	return nil
}

func (in *Proxy) SwitchOffForIngress(ctx context.Context, ingress *netv1.Ingress) error {
	if len(ingress.Annotations) == 0 {
		return nil
	}
	for _, annotation := range []Annotation{
		Timeout,
		ForceSSLRedirect,
	} {
		delete(ingress.Annotations, annotation.String())
	}
	return nil
}

func (in *Proxy) SwitchOnForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
func (in *Proxy) SwitcherOffForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
