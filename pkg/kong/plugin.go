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

type Plugin struct {
	Id        string                 `json:"id"`
	Name      string                 `json:"name"`
	CreatedAt int                    `json:"created_at"`
	Route     *PluginRoute           `json:"route,omitempty"`
	Service   *PluginService         `json:"service,omitempty"`
	Consumer  *PluginConsumer        `json:"consumer,omitempty"`
	Config    map[string]interface{} `json:"config"`
	Protocols []string               `json:"protocols"`
	Enabled   bool                   `json:"enabled"`
	Tags      []string               `json:"tags"`
	Ordering  PluginOrdering         `json:"ordering"`
}

type PluginRoute struct {
	ID string `json:"id,omitempty"`
}

type PluginService struct {
	ID string `json:"id,omitempty"`
}

type PluginConsumer struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type PluginOrdering struct {
	Before []string `json:"before"`
}
