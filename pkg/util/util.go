/*
Copyright © 2022 SUSE LLC

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

package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RemoveInvalidConditions(conditions []metav1.Condition) []metav1.Condition {
	newConditions := []metav1.Condition{}
	for _, cond := range conditions {
		if cond.Type == "" || cond.Status == "" || cond.LastTransitionTime.IsZero() || cond.Reason == "" {
			continue
		}
		newConditions = append(newConditions, cond)
	}
	return newConditions
}
