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

	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "k8s.io/api/core/v1"
	susql "github.com/metalcycling/susql/api/v1"
)

// Functions to get data from the cluster
func (r *LabelGroupReconciler) GetPodNamesMatchingLabels(ctx context.Context, labelGroup *susql.LabelGroup) ([]string, error) {
	pods := &v1.PodList{}

	if err := r.List(ctx, pods, client.UnsafeDisableDeepCopy, (client.MatchingLabels)(labelGroup.Spec.Labels)); err != nil {
		return nil, err
	}

	var podNames []string

	for _, pod := range pods.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames, nil
}