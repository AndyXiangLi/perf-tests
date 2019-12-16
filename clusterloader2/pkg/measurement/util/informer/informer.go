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

package informer

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// NewInformer creates a new informer
// for given kind, namespace, fieldSelector and labelSelector.
func NewInformer(
	c clientset.Interface,
	kind string,
	namespace, fieldSelector, labelSelector string,
	handleObj func(interface{}, interface{}),
) cache.SharedInformer {
	optionsModifier := func(options *metav1.ListOptions) {
		options.FieldSelector = fieldSelector
		options.LabelSelector = labelSelector
	}
	listerWatcher := cache.NewFilteredListWatchFromClient(c.CoreV1().RESTClient(), kind, namespace, optionsModifier)
	informer := cache.NewSharedInformer(listerWatcher, nil, 0)
	addEventHandler(informer, handleObj)

	return informer
}

// NewDynamicInformer creates a new dynamic informer
// for given namespace, fieldSelector and labelSelector.
func NewDynamicInformer(
	c dynamic.Interface,
	gvr schema.GroupVersionResource,
	namespace, fieldSelector, labelSelector string,
	handleObj func(interface{}, interface{}),
) cache.SharedInformer {
	optionsModifier := func(options *metav1.ListOptions) {
		options.FieldSelector = fieldSelector
		options.LabelSelector = labelSelector
	}
	tweakListOptions := dynamicinformer.TweakListOptionsFunc(optionsModifier)
	dInformerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(c, 0, namespace, tweakListOptions)

	informer := dInformerFactory.ForResource(gvr).Informer()
	addEventHandler(informer, handleObj)
	return informer
}

func addEventHandler(i cache.SharedInformer,
	handleObj func(interface{}, interface{}),
) {
	i.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			handleObj(nil, obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			handleObj(oldObj, newObj)
		},
		DeleteFunc: func(obj interface{}) {
			if tombstone, ok := obj.(cache.DeletedFinalStateUnknown); ok {
				handleObj(tombstone.Obj, nil)
			} else {
				handleObj(obj, nil)
			}
		},
	})
}
