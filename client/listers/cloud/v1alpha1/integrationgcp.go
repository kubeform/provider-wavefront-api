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

// IntegrationGcpLister helps list IntegrationGcps.
// All objects returned here must be treated as read-only.
type IntegrationGcpLister interface {
	// List lists all IntegrationGcps in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationGcp, err error)
	// IntegrationGcps returns an object that can list and get IntegrationGcps.
	IntegrationGcps(namespace string) IntegrationGcpNamespaceLister
	IntegrationGcpListerExpansion
}

// integrationGcpLister implements the IntegrationGcpLister interface.
type integrationGcpLister struct {
	indexer cache.Indexer
}

// NewIntegrationGcpLister returns a new IntegrationGcpLister.
func NewIntegrationGcpLister(indexer cache.Indexer) IntegrationGcpLister {
	return &integrationGcpLister{indexer: indexer}
}

// List lists all IntegrationGcps in the indexer.
func (s *integrationGcpLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationGcp, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationGcp))
	})
	return ret, err
}

// IntegrationGcps returns an object that can list and get IntegrationGcps.
func (s *integrationGcpLister) IntegrationGcps(namespace string) IntegrationGcpNamespaceLister {
	return integrationGcpNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IntegrationGcpNamespaceLister helps list and get IntegrationGcps.
// All objects returned here must be treated as read-only.
type IntegrationGcpNamespaceLister interface {
	// List lists all IntegrationGcps in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationGcp, err error)
	// Get retrieves the IntegrationGcp from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IntegrationGcp, error)
	IntegrationGcpNamespaceListerExpansion
}

// integrationGcpNamespaceLister implements the IntegrationGcpNamespaceLister
// interface.
type integrationGcpNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IntegrationGcps in the indexer for a given namespace.
func (s integrationGcpNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationGcp, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationGcp))
	})
	return ret, err
}

// Get retrieves the IntegrationGcp from the indexer for a given namespace and name.
func (s integrationGcpNamespaceLister) Get(name string) (*v1alpha1.IntegrationGcp, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("integrationgcp"), name)
	}
	return obj.(*v1alpha1.IntegrationGcp), nil
}
