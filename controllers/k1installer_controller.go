/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"k1-installer-demo/pkg"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	k1installerv1 "k1-installer-demo/api/v1"
)

// K1InstallerReconciler reconciles a K1Installer object
type K1InstallerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=k1-installer.jplab.cloud,resources=k1installers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=k1-installer.jplab.cloud,resources=k1installers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=k1-installer.jplab.cloud,resources=k1installers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the K1Installer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *K1InstallerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	k8sLog := log.FromContext(ctx)

	var k1Inst k1installerv1.K1Installer

	if err := r.Get(ctx, req.NamespacedName, &k1Inst); err != nil {
		k8sLog.Info("error loading object", "name", req.NamespacedName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if k1Inst.Status.Status == "" {
		if err := pkg.CreateBucket(false, "joao-test-crd"); err != nil {
			k8sLog.Info("error creation s3 bucket ", "name", req.NamespacedName)
		}
		k1Inst.Status.Status = "S3 bucket created!"
	}

	if err := r.Status().Update(ctx, &k1Inst); err != nil {
		k8sLog.Info("error updating status", "name", req.NamespacedName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	k8sLog.Info("some logs: ", "name", req.NamespacedName)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *K1InstallerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&k1installerv1.K1Installer{}).
		Complete(r)
}
