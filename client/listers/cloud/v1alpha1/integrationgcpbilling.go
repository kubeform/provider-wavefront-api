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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-wavefront-api/apis/cloud/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IntegrationGcpBillingLister helps list IntegrationGcpBillings.
// All objects returned here must be treated as read-only.
type IntegrationGcpBillingLister interface {
	// List lists all IntegrationGcpBillings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationGcpBilling, err error)
	// IntegrationGcpBillings returns an object that can list and get IntegrationGcpBillings.
	IntegrationGcpBillings(namespace string) IntegrationGcpBillingNamespaceLister
	IntegrationGcpBillingListerExpansion
}

// integrationGcpBillingLister implements the IntegrationGcpBillingLister interface.
type integrationGcpBillingLister struct {
	indexer cache.Indexer
}

// NewIntegrationGcpBillingLister returns a new IntegrationGcpBillingLister.
func NewIntegrationGcpBillingLister(indexer cache.Indexer) IntegrationGcpBillingLister {
	return &integrationGcpBillingLister{indexer: indexer}
}

// List lists all IntegrationGcpBillings in the indexer.
func (s *integrationGcpBillingLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationGcpBilling, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationGcpBilling))
	})
	return ret, err
}

// IntegrationGcpBillings returns an object that can list and get IntegrationGcpBillings.
func (s *integrationGcpBillingLister) IntegrationGcpBillings(namespace string) IntegrationGcpBillingNamespaceLister {
	return integrationGcpBillingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IntegrationGcpBillingNamespaceLister helps list and get IntegrationGcpBillings.
// All objects returned here must be treated as read-only.
type IntegrationGcpBillingNamespaceLister interface {
	// List lists all IntegrationGcpBillings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationGcpBilling, err error)
	// Get retrieves the IntegrationGcpBilling from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IntegrationGcpBilling, error)
	IntegrationGcpBillingNamespaceListerExpansion
}

// integrationGcpBillingNamespaceLister implements the IntegrationGcpBillingNamespaceLister
// interface.
type integrationGcpBillingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IntegrationGcpBillings in the indexer for a given namespace.
func (s integrationGcpBillingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationGcpBilling, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationGcpBilling))
	})
	return ret, err
}

// Get retrieves the IntegrationGcpBilling from the indexer for a given namespace and name.
func (s integrationGcpBillingNamespaceLister) Get(name string) (*v1alpha1.IntegrationGcpBilling, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("integrationgcpbilling"), name)
	}
	return obj.(*v1alpha1.IntegrationGcpBilling), nil
}