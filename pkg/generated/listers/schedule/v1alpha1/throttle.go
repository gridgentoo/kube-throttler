// Licensed to Shingo Omura under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Shingo Omura licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/everpeace/kube-throttler/pkg/apis/schedule/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ThrottleLister helps list Throttles.
// All objects returned here must be treated as read-only.
type ThrottleLister interface {
	// List lists all Throttles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Throttle, err error)
	// Throttles returns an object that can list and get Throttles.
	Throttles(namespace string) ThrottleNamespaceLister
	ThrottleListerExpansion
}

// throttleLister implements the ThrottleLister interface.
type throttleLister struct {
	indexer cache.Indexer
}

// NewThrottleLister returns a new ThrottleLister.
func NewThrottleLister(indexer cache.Indexer) ThrottleLister {
	return &throttleLister{indexer: indexer}
}

// List lists all Throttles in the indexer.
func (s *throttleLister) List(selector labels.Selector) (ret []*v1alpha1.Throttle, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Throttle))
	})
	return ret, err
}

// Throttles returns an object that can list and get Throttles.
func (s *throttleLister) Throttles(namespace string) ThrottleNamespaceLister {
	return throttleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ThrottleNamespaceLister helps list and get Throttles.
// All objects returned here must be treated as read-only.
type ThrottleNamespaceLister interface {
	// List lists all Throttles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Throttle, err error)
	// Get retrieves the Throttle from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Throttle, error)
	ThrottleNamespaceListerExpansion
}

// throttleNamespaceLister implements the ThrottleNamespaceLister
// interface.
type throttleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Throttles in the indexer for a given namespace.
func (s throttleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Throttle, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Throttle))
	})
	return ret, err
}

// Get retrieves the Throttle from the indexer for a given namespace and name.
func (s throttleNamespaceLister) Get(name string) (*v1alpha1.Throttle, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("throttle"), name)
	}
	return obj.(*v1alpha1.Throttle), nil
}
