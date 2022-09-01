// Generated by hapi-operator tools

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
	"reflect"
	"strings"

	"k8s.io/apimachinery/pkg/api/equality"
)

func (in Auth) DeepEqual(i interface{}) ([]string, bool) {
	var v Auth
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case Auth:
		v = t
	case *Auth:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in Backend) DeepEqual(i interface{}) ([]string, bool) {
	var v Backend
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case Backend:
		v = t
	case *Backend:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in BaseStat) DeepEqual(i interface{}) ([]string, bool) {
	var v BaseStat
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case BaseStat:
		v = t
	case *BaseStat:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in CORS) DeepEqual(i interface{}) ([]string, bool) {
	var v CORS
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case CORS:
		v = t
	case *CORS:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in ConfigZone) DeepEqual(i interface{}) ([]string, bool) {
	var v ConfigZone
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case ConfigZone:
		v = t
	case *ConfigZone:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in ConfigZoneList) DeepEqual(i interface{}) ([]string, bool) {
	var v ConfigZoneList
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case ConfigZoneList:
		v = t
	case *ConfigZoneList:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in ConfigZoneSpec) DeepEqual(i interface{}) ([]string, bool) {
	var v ConfigZoneSpec
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case ConfigZoneSpec:
		v = t
	case *ConfigZoneSpec:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in ConfigZoneStatus) DeepEqual(i interface{}) ([]string, bool) {
	var v ConfigZoneStatus
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case ConfigZoneStatus:
		v = t
	case *ConfigZoneStatus:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in Consumer) DeepEqual(i interface{}) ([]string, bool) {
	var v Consumer
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case Consumer:
		v = t
	case *Consumer:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in Hapi) DeepEqual(i interface{}) ([]string, bool) {
	var v Hapi
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case Hapi:
		v = t
	case *Hapi:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in HapiList) DeepEqual(i interface{}) ([]string, bool) {
	var v HapiList
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case HapiList:
		v = t
	case *HapiList:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in HapiSpec) DeepEqual(i interface{}) ([]string, bool) {
	var v HapiSpec
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case HapiSpec:
		v = t
	case *HapiSpec:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in HapiStatus) DeepEqual(i interface{}) ([]string, bool) {
	var v HapiStatus
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case HapiStatus:
		v = t
	case *HapiStatus:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in HapiStatusResourceVersion) DeepEqual(i interface{}) ([]string, bool) {
	var v HapiStatusResourceVersion
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case HapiStatusResourceVersion:
		v = t
	case *HapiStatusResourceVersion:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in KongPlugin) DeepEqual(i interface{}) ([]string, bool) {
	var v KongPlugin
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case KongPlugin:
		v = t
	case *KongPlugin:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in KongRoute) DeepEqual(i interface{}) ([]string, bool) {
	var v KongRoute
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case KongRoute:
		v = t
	case *KongRoute:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in KongService) DeepEqual(i interface{}) ([]string, bool) {
	var v KongService
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case KongService:
		v = t
	case *KongService:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in Kongress) DeepEqual(i interface{}) ([]string, bool) {
	var v Kongress
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case Kongress:
		v = t
	case *Kongress:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in KongressList) DeepEqual(i interface{}) ([]string, bool) {
	var v KongressList
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case KongressList:
		v = t
	case *KongressList:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in KongressSpec) DeepEqual(i interface{}) ([]string, bool) {
	var v KongressSpec
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case KongressSpec:
		v = t
	case *KongressSpec:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in KongressStatus) DeepEqual(i interface{}) ([]string, bool) {
	var v KongressStatus
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case KongressStatus:
		v = t
	case *KongressStatus:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in Metric) DeepEqual(i interface{}) ([]string, bool) {
	var v Metric
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case Metric:
		v = t
	case *Metric:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in Policy) DeepEqual(i interface{}) ([]string, bool) {
	var v Policy
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case Policy:
		v = t
	case *Policy:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in Proxy) DeepEqual(i interface{}) ([]string, bool) {
	var v Proxy
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case Proxy:
		v = t
	case *Proxy:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in RedirectByServiceStat) DeepEqual(i interface{}) ([]string, bool) {
	var v RedirectByServiceStat
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case RedirectByServiceStat:
		v = t
	case *RedirectByServiceStat:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in RedirectByUrlStat) DeepEqual(i interface{}) ([]string, bool) {
	var v RedirectByUrlStat
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case RedirectByUrlStat:
		v = t
	case *RedirectByUrlStat:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in ReverseProxyRule) DeepEqual(i interface{}) ([]string, bool) {
	var v ReverseProxyRule
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case ReverseProxyRule:
		v = t
	case *ReverseProxyRule:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in SafetyCSRF) DeepEqual(i interface{}) ([]string, bool) {
	var v SafetyCSRF
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case SafetyCSRF:
		v = t
	case *SafetyCSRF:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in SafetyIP) DeepEqual(i interface{}) ([]string, bool) {
	var v SafetyIP
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case SafetyIP:
		v = t
	case *SafetyIP:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in SafetySBAC) DeepEqual(i interface{}) ([]string, bool) {
	var v SafetySBAC
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case SafetySBAC:
		v = t
	case *SafetySBAC:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in ServerGuard) DeepEqual(i interface{}) ([]string, bool) {
	var v ServerGuard
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case ServerGuard:
		v = t
	case *ServerGuard:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func (in Switch) DeepEqual(i interface{}) ([]string, bool) {
	var v Switch
	switch t := i.(type) {
	case nil:
		return []string{"(invalid)"}, false
	case Switch:
		v = t
	case *Switch:
		if t == nil {
			return []string{"(invalid)"}, false
		}
		v = *t
	default:
		return []string{"(invalid)"}, false
	}

	v1 := reflect.ValueOf(in)
	t1 := reflect.TypeOf(in)
	v2 := reflect.ValueOf(v)
	for i := 0; i < v1.NumField(); i++ {
		if !v1.Field(i).IsValid() && !v2.Field(i).IsValid() {
			continue
		}
		if !v1.Field(i).IsValid() || !v2.Field(i).IsValid() {
			return getTag(t1.Field(i)), false
		}

		if de, ok := v1.Field(i).Interface().(interface {
			DeepEqual(interface{}) ([]string, bool)
		}); ok {
			if field, ok := de.DeepEqual(v2.Field(i).Interface()); !ok {
				return append(getTag(t1.Field(i)), field...), ok
			}
			continue
		}
		if ok := equality.Semantic.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()); !ok {
			return getTag(t1.Field(i)), false
		}
	}
	return nil, true
}

func getTag(field reflect.StructField) []string {
	jTag := field.Tag.Get("json")
	if jTag != "" {
		return []string{strings.Split(jTag, ",")[0]}
	}
	if strings.HasPrefix(jTag, ",inline") {
		return nil
	}
	return []string{field.Name}
}
