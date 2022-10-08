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

package kong

type Service struct {
	Id                string            `json:"id"`
	CreatedAt         int               `json:"created_at"`
	UpdatedAt         int               `json:"updated_at"`
	Name              string            `json:"name"`
	Retries           int               `json:"retries"`
	Protocol          string            `json:"protocol"`
	Host              string            `json:"host"`
	Port              int               `json:"port"`
	Path              string            `json:"path"`
	ConnectTimeout    int               `json:"connect_timeout"`
	WriteTimeout      int               `json:"write_timeout"`
	ReadTimeout       int               `json:"read_timeout"`
	Tags              []string          `json:"tags"`
	ClientCertificate ClientCertificate `json:"client_certificate"`
	TlsVerify         bool              `json:"tls_verify"`
	CaCertificates    []string          `json:"ca_certificates"`
	Enabled           bool              `json:"enabled"`
}

type ClientCertificate struct {
	Id string `json:"id"`
}
