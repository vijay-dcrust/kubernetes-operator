/*

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
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"log"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	batchv1 "example/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	errors "k8s.io/apimachinery/pkg/api/errors"
	types "k8s.io/apimachinery/pkg/types"
)

// GenericDaemonReconciler reconciles a GenericDaemon object
type GenericDaemonReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=batch.tutorial.kubebuilder.io,resources=genericdaemons,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.tutorial.kubebuilder.io,resources=genericdaemons/status,verbs=get;update;patch

func (r *GenericDaemonReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()

	// your logic here
	//var gd batchv1.GenericDaemon
	gd := &batchv1.GenericDaemon{}
	if err := r.Get(ctx, req.NamespacedName, gd); err != nil {
		return ctrl.Result{}, err
	}
	// Define the desired Daemonset object
	daemonset := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gd.Name + "-daemonset",
			Namespace: gd.Namespace,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"daemonset": gd.Name + "-daemonset"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"daemonset": gd.Name + "-daemonset"},
				},
				Spec: corev1.PodSpec{
					NodeSelector: map[string]string{"daemon": gd.Spec.Label},
					Containers: []corev1.Container{
						{
							Name:  "genericdaemon",
							Image: gd.Spec.Image,
						},
					},
				},
			},
		},
	}
	if err := ctrl.SetControllerReference(gd, daemonset, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}
	found := &appsv1.DaemonSet{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: daemonset.Name, Namespace: daemonset.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		log.Printf("Creating Daemonset %s/%s\n", daemonset.Namespace, daemonset.Name)
		err := r.Create(context.TODO(), daemonset)
		if err != nil {
			return ctrl.Result{}, err
		}
	} else if err != nil {
		return ctrl.Result{}, err
	}
	// Get the number of Ready daemonsets and set the Count status
	if found.Status.NumberReady != gd.Status.Count {
		log.Printf("Updating Status %s/%s\n", gd.Namespace, gd.Name)
		gd.Status.Count = found.Status.NumberReady
		err = r.Update(context.TODO(), gd)
		if err != nil {
			return ctrl.Result{}, err
		}
	}
	// Update the found object and write the result back if there are any changes
	if !reflect.DeepEqual(daemonset.Spec, found.Spec) {
		found.Spec = daemonset.Spec
		log.Printf("Updating Daemonset %s/%s\n", daemonset.Namespace, daemonset.Name)
		err = r.Update(context.TODO(), found)
		if err != nil {
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

func (r *GenericDaemonReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.GenericDaemon{}).
		Complete(r)
}
