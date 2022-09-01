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

// https://help.aliyun.com/document_detail/424813.htm
const (
	AnnotationUpstreamVHost Annotation = "nginx.ingress.kubernetes.io/upstream-vhost"
	AnnotationRewriteTarget Annotation = "nginx.ingress.kubernetes.io/rewrite-target"

	WhiteListSourceRange       Annotation = "nginx.ingress.kubernetes.io/whitelist-source-range"
	BlackListSourceRange       Annotation = "mse.ingress.kubernetes.io/blacklist-source-range"
	DomainWhiteListSourceRange Annotation = "mse.ingress.kubernetes.io/domain-whitelist-source-range"
	DomainBlackListSourceRange Annotation = "mse.ingress.kubernetes.io/domain-blacklist-source-range"
	RequestHeaderControlAdd    Annotation = "mse.ingress.kubernetes.io/request-header-control-add"

	RouteLimitRpm             Annotation = "mse.ingress.kubernetes.io/route-limit-rpm"
	RouteLimitRps             Annotation = "mse.ingress.kubernetes.io/route-limit-rps"
	RouteLimitBurstMultiplier Annotation = "mse.ingress.kubernetes.io/route-limit-burst-multiplier"

	ForceSSLRedirect Annotation = "nginx.ingress.kubernetes.io/force-ssl-redirect"
	Timeout          Annotation = "mse.ingress.kubernetes.io/timeout"

	EnableCORS           Annotation = "nginx.ingress.kubernetes.io/enable-cors"            // Ingress开启或关闭跨域。
	CORSAllowOrigin      Annotation = "nginx.ingress.kubernetes.io/cors-allow-origin"      // Ingress	允许的第三方站点。
	CORSAllowMethods     Annotation = "nginx.ingress.kubernetes.io/cors-allow-methods"     // Ingress	允许的请求方法，如GET、POST、PUT等。
	CORSAllowHeaders     Annotation = "nginx.ingress.kubernetes.io/cors-allow-headers"     // Ingress	允许的请求Header。
	CORSExposeHeaders    Annotation = "nginx.ingress.kubernetes.io/cors-expose-headers"    // Ingress	允许的暴露给浏览器的响应Header。
	CORSAllowCredentials Annotation = "nginx.ingress.kubernetes.io/cors-allow-credentials" // Ingress	是否允许携带凭证信息。
	CORSMaxAge           Annotation = "nginx.ingress.kubernetes.io/cors-max-age"           // Ingress 预检结果的最大缓存时间。
)

type Annotation string

func (in Annotation) String() string {
	return string(in)
}
