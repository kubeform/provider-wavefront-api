/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	alertv1alpha1 "kubeform.dev/provider-wavefront-api/apis/alert/v1alpha1"
	versioned "kubeform.dev/provider-wavefront-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-wavefront-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-wavefront-api/client/listers/alert/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TargetInformer provides access to a shared informer and lister for
// Targets.
type TargetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TargetLister
}

type targetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTargetInformer constructs a new informer for Target type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTargetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTargetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTargetInformer constructs a new informer for Target type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTargetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertV1alpha1().Targets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertV1alpha1().Targets(namespace).Watch(context.TODO(), options)
			},
		},
		&alertv1alpha1.Target{},
		resyncPeriod,
		indexers,
	)
}

func (f *targetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTargetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *targetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&alertv1alpha1.Target{}, f.defaultInformer)
}

func (f *targetInformer) Lister() v1alpha1.TargetLister {
	return v1alpha1.NewTargetLister(f.Informer().GetIndexer())
}
