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
	v1alpha1 "kubeform.dev/provider-wavefront-api/apis/user/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GroupLister helps list Groups.
// All objects returned here must be treated as read-only.
type GroupLister interface {
	// List lists all Groups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Group, err error)
	// Groups returns an object that can list and get Groups.
	Groups(namespace string) GroupNamespaceLister
	GroupListerExpansion
}

// groupLister implements the GroupLister interface.
type groupLister struct {
	indexer cache.Indexer
}

// NewGroupLister returns a new GroupLister.
func NewGroupLister(indexer cache.Indexer) GroupLister {
	return &groupLister{indexer: indexer}
}

// List lists all Groups in the indexer.
func (s *groupLister) List(selector labels.Selector) (ret []*v1alpha1.Group, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Group))
	})
	return ret, err
}

// Groups returns an object that can list and get Groups.
func (s *groupLister) Groups(namespace string) GroupNamespaceLister {
	return groupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GroupNamespaceLister helps list and get Groups.
// All objects returned here must be treated as read-only.
type GroupNamespaceLister interface {
	// List lists all Groups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Group, err error)
	// Get retrieves the Group from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Group, error)
	GroupNamespaceListerExpansion
}

// groupNamespaceLister implements the GroupNamespaceLister
// interface.
type groupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Groups in the indexer for a given namespace.
func (s groupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Group, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Group))
	})
	return ret, err
}

// Get retrieves the Group from the indexer for a given namespace and name.
func (s groupNamespaceLister) Get(name string) (*v1alpha1.Group, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("group"), name)
	}
	return obj.(*v1alpha1.Group), nil
}
