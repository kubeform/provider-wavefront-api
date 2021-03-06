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

// IntegrationAzureActivityLogLister helps list IntegrationAzureActivityLogs.
// All objects returned here must be treated as read-only.
type IntegrationAzureActivityLogLister interface {
	// List lists all IntegrationAzureActivityLogs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationAzureActivityLog, err error)
	// IntegrationAzureActivityLogs returns an object that can list and get IntegrationAzureActivityLogs.
	IntegrationAzureActivityLogs(namespace string) IntegrationAzureActivityLogNamespaceLister
	IntegrationAzureActivityLogListerExpansion
}

// integrationAzureActivityLogLister implements the IntegrationAzureActivityLogLister interface.
type integrationAzureActivityLogLister struct {
	indexer cache.Indexer
}

// NewIntegrationAzureActivityLogLister returns a new IntegrationAzureActivityLogLister.
func NewIntegrationAzureActivityLogLister(indexer cache.Indexer) IntegrationAzureActivityLogLister {
	return &integrationAzureActivityLogLister{indexer: indexer}
}

// List lists all IntegrationAzureActivityLogs in the indexer.
func (s *integrationAzureActivityLogLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationAzureActivityLog, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationAzureActivityLog))
	})
	return ret, err
}

// IntegrationAzureActivityLogs returns an object that can list and get IntegrationAzureActivityLogs.
func (s *integrationAzureActivityLogLister) IntegrationAzureActivityLogs(namespace string) IntegrationAzureActivityLogNamespaceLister {
	return integrationAzureActivityLogNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IntegrationAzureActivityLogNamespaceLister helps list and get IntegrationAzureActivityLogs.
// All objects returned here must be treated as read-only.
type IntegrationAzureActivityLogNamespaceLister interface {
	// List lists all IntegrationAzureActivityLogs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationAzureActivityLog, err error)
	// Get retrieves the IntegrationAzureActivityLog from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IntegrationAzureActivityLog, error)
	IntegrationAzureActivityLogNamespaceListerExpansion
}

// integrationAzureActivityLogNamespaceLister implements the IntegrationAzureActivityLogNamespaceLister
// interface.
type integrationAzureActivityLogNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IntegrationAzureActivityLogs in the indexer for a given namespace.
func (s integrationAzureActivityLogNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationAzureActivityLog, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationAzureActivityLog))
	})
	return ret, err
}

// Get retrieves the IntegrationAzureActivityLog from the indexer for a given namespace and name.
func (s integrationAzureActivityLogNamespaceLister) Get(name string) (*v1alpha1.IntegrationAzureActivityLog, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("integrationazureactivitylog"), name)
	}
	return obj.(*v1alpha1.IntegrationAzureActivityLog), nil
}
