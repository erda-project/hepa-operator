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

package stringsutils

import (
	"os"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

func LoadEnv(v interface{}) error {
	if v == nil {
		return errors.New("The argument is nil")
	}
	valueOf := reflect.ValueOf(v)
	if !valueOf.IsValid() {
		return errors.New("Parameter invalid")
	}
	typeOf := reflect.TypeOf(v)
	if typeOf.Kind() != reflect.Pointer {
		return errors.New("The parameters must be of type pointer")
	}
	if valueOf.IsNil() {
		return errors.New("The argument is nil")
	}
	for i := 0; i < typeOf.Elem().NumField(); i++ {
		if t := typeOf.Elem().Field(i).Type.Kind(); t != reflect.String {
			return errors.Errorf("field type must be String, got %s", t)
		}
		envTag := typeOf.Elem().Field(i).Tag.Get("env")
		envTags := strings.Split(envTag, ",")
		key := envTags[0]
		if key == "" {
			continue
		}
		value := os.Getenv(key)
		if value == "" && strings.Contains(envTag, "required") {
			return errors.Errorf("%s required but not found", key)
		}
		valueOf.Elem().Field(i).SetString(value)
	}
	return nil
}
