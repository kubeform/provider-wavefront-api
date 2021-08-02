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

// IntegrationNewrelicLister helps list IntegrationNewrelics.
// All objects returned here must be treated as read-only.
type IntegrationNewrelicLister interface {
	// List lists all IntegrationNewrelics in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationNewrelic, err error)
	// IntegrationNewrelics returns an object that can list and get IntegrationNewrelics.
	IntegrationNewrelics(namespace string) IntegrationNewrelicNamespaceLister
	IntegrationNewrelicListerExpansion
}

// integrationNewrelicLister implements the IntegrationNewrelicLister interface.
type integrationNewrelicLister struct {
	indexer cache.Indexer
}

// NewIntegrationNewrelicLister returns a new IntegrationNewrelicLister.
func NewIntegrationNewrelicLister(indexer cache.Indexer) IntegrationNewrelicLister {
	return &integrationNewrelicLister{indexer: indexer}
}

// List lists all IntegrationNewrelics in the indexer.
func (s *integrationNewrelicLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationNewrelic, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationNewrelic))
	})
	return ret, err
}

// IntegrationNewrelics returns an object that can list and get IntegrationNewrelics.
func (s *integrationNewrelicLister) IntegrationNewrelics(namespace string) IntegrationNewrelicNamespaceLister {
	return integrationNewrelicNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IntegrationNewrelicNamespaceLister helps list and get IntegrationNewrelics.
// All objects returned here must be treated as read-only.
type IntegrationNewrelicNamespaceLister interface {
	// List lists all IntegrationNewrelics in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationNewrelic, err error)
	// Get retrieves the IntegrationNewrelic from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IntegrationNewrelic, error)
	IntegrationNewrelicNamespaceListerExpansion
}

// integrationNewrelicNamespaceLister implements the IntegrationNewrelicNamespaceLister
// interface.
type integrationNewrelicNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IntegrationNewrelics in the indexer for a given namespace.
func (s integrationNewrelicNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationNewrelic, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationNewrelic))
	})
	return ret, err
}

// Get retrieves the IntegrationNewrelic from the indexer for a given namespace and name.
func (s integrationNewrelicNamespaceLister) Get(name string) (*v1alpha1.IntegrationNewrelic, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("integrationnewrelic"), name)
	}
	return obj.(*v1alpha1.IntegrationNewrelic), nil
}