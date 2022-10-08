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

package v2

import (
	"context"
	"os"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/erda-project/hepa-operator/pkg/utils/stringsutils"
)

// Reconciler is not implemented
type Reconciler struct {
	client        client.Client
	Scheme        *runtime.Scheme
	Log           logr.Logger
	EventRecorder record.EventRecorder
	Config        Config
}

func New(cli client.Client, scheme *runtime.Scheme, log logr.Logger, recorder record.EventRecorder) *Reconciler {
	var r = &Reconciler{
		client:        cli,
		Scheme:        scheme,
		Log:           log.WithName("reconciler/nginx-kong.v2"),
		EventRecorder: recorder,
		Config:        Config{},
	}
	// todo: use erda pkg
	if err := stringsutils.LoadEnv(&r.Config); err != nil {
		log.Error(err, "failed to LoadEnv")
		os.Exit(1)
	}
	return r
}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (ctrl.Result, error) {
	// todo: implement nothing yet
	return ctrl.Result{}, nil
}

type Config struct {
	KongAdminOpenapi        string `json:"kongAdminOpenapi,omitempty" env:"KONG_ADMIN_OPENAPI,required"`
	KongNamespace           string `json:"kongNamespace,omitempty" env:"KONG_NAMESPACE,required"`
	KongServiceName         string `json:"kongServiceName,omitempty" env:"KONG_SERVICE_NAME,required"`
	NginxNamespace          string `json:"nginxNamespace,omitempty" env:"NGINX_NAMESPACE,required"`
	NginxConfigMapNamespace string `json:"nginxConfigMapNamespace,omitempty" env:"NGINX_CONFIG_MAP_NAMESPACE,required"`
	NginxConfigMapName      string `json:"nginxConfigMapName,omitempty" env:"nginxConfigMapName,required"`
}
