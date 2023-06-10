/*
Copyright 2023.

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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"kangmingfa/kubernetes-container-checkpoint/internal/dockerCheckpoint"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kangmingfagov1 "kangmingfa/kubernetes-container-checkpoint/api/v1"
)

// K8sCRCheckpointReconciler reconciles a K8sCRCheckpoint object
type K8sCRCheckpointReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=kangmingfa.go.kangmingfa.go,resources=k8scrcheckpoints,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kangmingfa.go.kangmingfa.go,resources=k8scrcheckpoints/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kangmingfa.go.kangmingfa.go,resources=k8scrcheckpoints/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the K8sCRCheckpoint object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *K8sCRCheckpointReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	var k8sCRCheckpoint kangmingfagov1.K8sCRCheckpoint
	err := r.Get(ctx, req.NamespacedName, &k8sCRCheckpoint)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	fmt.Println("userId  is " + k8sCRCheckpoint.Spec.UserId)
	fmt.Println("chapterId  is " + k8sCRCheckpoint.Spec.ChapterId)
	fmt.Println("questionId  is " + k8sCRCheckpoint.Spec.QuestionId)
	fmt.Println("try to find which checkpoint to start")
	// Todo : use docker client to start the found checkpoint
	result := dockerCheckpoint.StartFromCheckpoint("")
	if result.Success {
		fmt.Println(result.Msg)
	} else {
		fmt.Println(result.Msg)
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *K8sCRCheckpointReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kangmingfagov1.K8sCRCheckpoint{}).
		Complete(r)
}
