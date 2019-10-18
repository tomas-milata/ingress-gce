/*
Copyright 2019 The Kubernetes Authors.

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

package common

import (
	"k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
)

var (
	KeyFunc = cache.DeletionHandlingMetaNamespaceKeyFunc
)

// ToString returns namespaced name string of a given ingress.
// Note: This is used for logging.
func ToString(ing *v1beta1.Ingress) string {
	if ing == nil {
		return ""
	}
	return types.NamespacedName{Namespace: ing.Namespace, Name: ing.Name}.String()
}

// IngressKeyFunc returns ingress key for given ingress as generated by Ingress Store.
// This falls back to utility function in case of an error.
// Note: Ingress Store and ToString both return same key in general. But, Ingress Store
// returns <name> where as ToString returns /<name> when <namespace> is empty.
func IngressKeyFunc(ing *v1beta1.Ingress) string {
	ingKey, err := KeyFunc(ing)
	if err == nil {
		return ingKey
	}
	// An error is returned only if ingress object does not have a valid meta.
	// So, this should not happen in production, fall back on utility function to
	// get ingress key.
	klog.Errorf("Cannot get key for Ingress %v/%v: %v, using utility function", ing.Namespace, ing.Name, err)
	return ToString(ing)
}

// ToIngressKeys returns a list of ingress keys for given list of ingresses.
func ToIngressKeys(ings []*v1beta1.Ingress) []string {
	ingKeys := make([]string, 0, len(ings))
	for _, ing := range ings {
		ingKeys = append(ingKeys, IngressKeyFunc(ing))
	}
	return ingKeys
}