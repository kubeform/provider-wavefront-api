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

	cloudv1alpha1 "kubeform.dev/provider-wavefront-api/apis/cloud/v1alpha1"
	versioned "kubeform.dev/provider-wavefront-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-wavefront-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-wavefront-api/client/listers/cloud/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IntegrationEc2Informer provides access to a shared informer and lister for
// IntegrationEc2s.
type IntegrationEc2Informer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.IntegrationEc2Lister
}

type integrationEc2Informer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIntegrationEc2Informer constructs a new informer for IntegrationEc2 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIntegrationEc2Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIntegrationEc2Informer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIntegrationEc2Informer constructs a new informer for IntegrationEc2 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIntegrationEc2Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().IntegrationEc2s(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().IntegrationEc2s(namespace).Watch(context.TODO(), options)
			},
		},
		&cloudv1alpha1.IntegrationEc2{},
		resyncPeriod,
		indexers,
	)
}

func (f *integrationEc2Informer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIntegrationEc2Informer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *integrationEc2Informer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloudv1alpha1.IntegrationEc2{}, f.defaultInformer)
}

func (f *integrationEc2Informer) Lister() v1alpha1.IntegrationEc2Lister {
	return v1alpha1.NewIntegrationEc2Lister(f.Informer().GetIndexer())
}