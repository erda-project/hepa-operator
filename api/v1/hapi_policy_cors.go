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

type CORS struct {
	Switch `json:",inline"`

	EnableCORS           bool   `json:"enableCORS,omitempty"`
	CORSAllowOrigin      string `json:"corsAllowOrigin,omitempty"`
	CORSAllowMethods     string `json:"corsAllowMethods,omitempty"`
	CORSAllowHeaders     string `json:"corsAllowHeaders,omitempty"`
	CORSExposeHeaders    string `json:"corsExposeHeaders,omitempty"`
	CORSAllowCredentials bool   `json:"corsAllowCredentials,omitempty"`
	CORSMaxAge           int    `json:"corsMaxAge,omitempty"`
}

func (in *CORS) SwitchOnForIngress(ctx context.Context, ingress *netv1.Ingress) error {
	if len(ingress.Annotations) == 0 {
		ingress.Annotations = make(map[string]string)
	}
	ingress.Annotations[EnableCORS.String()] = strconv.FormatBool(in.EnableCORS)
	ingress.Annotations[CORSAllowOrigin.String()] = in.CORSAllowOrigin
	ingress.Annotations[CORSAllowMethods.String()] = in.CORSAllowMethods
	ingress.Annotations[CORSAllowHeaders.String()] = in.CORSExposeHeaders
	ingress.Annotations[CORSExposeHeaders.String()] = in.CORSExposeHeaders
	ingress.Annotations[CORSAllowCredentials.String()] = strconv.FormatBool(in.CORSAllowCredentials)
	ingress.Annotations[CORSMaxAge.String()] = strconv.Itoa(in.CORSMaxAge)
	return nil
}

func (in *CORS) SwitchOffForIngress(ctx context.Context, ingress *netv1.Ingress) error {
	if len(ingress.Annotations) == 0 {
		return nil
	}
	for _, annotation := range []Annotation{
		EnableCORS,
		CORSAllowOrigin,
		CORSAllowMethods,
		CORSAllowHeaders,
		CORSExposeHeaders,
		CORSAllowCredentials,
		CORSMaxAge,
	} {
		delete(ingress.Annotations, annotation.String())
	}
	return nil
}

func (in *CORS) SwitchOnForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
func (in *CORS) SwitcherOffForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
