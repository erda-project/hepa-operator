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

	"github.com/pkg/errors"
	netv1 "k8s.io/api/networking/v1"

	"github.com/erda-project/hepa-operator/pkg/middleware"
	"github.com/erda-project/hepa-operator/pkg/utils/lines"
)

const (
	PeerIP       IPType = "x-peer-ip"
	XRealIP      IPType = "x-real-ip"
	XForwardedIP IPType = "x-forwarded-ip"
)

const (
	HeaderKVXPeerIPRemoteAddr = "x-peer-ip $remote_addr"
)

type SafetyIP struct {
	Switch `json:",inline"`

	IPType                     IPType `json:"ipType,omitempty"`
	WhiteListSourceRange       string `json:"whiteListSourceRange" annotation:"nginx.ingress.kubernetes.io/whitelist-source-range"`
	BlackListSourceRange       string `json:"blackListSourceRange" annotation:"mse.ingress.kubernetes.io/blacklist-source-range"`
	DomainWhiteListSourceRange string `json:"domainWhiteListSourceRange" annotation:"mse.ingress.kubernetes.io/domain-whitelist-source-range"`
	DomainBlackListSourceRange string `json:"domainBlackListSourceRange" annotation:"mse.ingress.kubernetes.io/domain-blacklist-source-range"`

	KeyRateLimitingValue string `json:"keyRateLimitingValue,omitempty"`
}

func (in *SafetyIP) SwitchOnForIngress(ctx context.Context, ingress *netv1.Ingress) error {
	if len(ingress.Annotations) == 0 {
		ingress.Annotations = make(map[string]string)
	}
	headers := ingress.Annotations[RequestHeaderControlAdd.String()]
	switch {
	case in.WhiteListSourceRange != "":
		delete(ingress.Annotations, BlackListSourceRange.String())
		ingress.Annotations[WhiteListSourceRange.String()] = in.WhiteListSourceRange // todo: 格式检查
		if in.KeyRateLimitingValue != "" {
			ingress.Annotations[RequestHeaderControlAdd.String()] = lines.From(headers).Set(HeaderKVXPeerIPRemoteAddr).String()
		} else {
			ingress.Annotations[RequestHeaderControlAdd.String()] = lines.From(headers).Delete(HeaderKVXPeerIPRemoteAddr).String()
		}
	case in.BlackListSourceRange != "":
		delete(ingress.Annotations, WhiteListSourceRange.String())
		ingress.Annotations[RequestHeaderControlAdd.String()] = lines.From(headers).Delete(HeaderKVXPeerIPRemoteAddr).String()
		ingress.Annotations[BlackListSourceRange.String()] = in.BlackListSourceRange // todo: 格式检查
	default:
		return errors.New("invalid white list and invalid black list")
	}
	if ingress.Annotations[RequestHeaderControlAdd.String()] == "" {
		delete(ingress.Annotations, RequestHeaderControlAdd.String())
	}
	return nil
}

func (in *SafetyIP) SwitchOffForIngress(ctx context.Context, ingress *netv1.Ingress) error {
	for _, annotation := range []Annotation{
		WhiteListSourceRange,
		BlackListSourceRange,
		DomainWhiteListSourceRange,
		DomainBlackListSourceRange,
	} {
		delete(ingress.Annotations, annotation.String())
	}
	if headers := ingress.Annotations[RequestHeaderControlAdd.String()]; headers != "" {
		ingress.Annotations[RequestHeaderControlAdd.String()] = lines.From(headers).Delete(HeaderKVXPeerIPRemoteAddr).String()
	}
	return nil
}

func (in *SafetyIP) SwitchOnForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}
func (in *SafetyIP) SwitcherOffForMiddleware(ctx context.Context, mw middleware.Middleware) error { /*todo: not implement*/
	return nil
}

type IPType string

func (ipType IPType) String() string {
	return string(ipType)
}
