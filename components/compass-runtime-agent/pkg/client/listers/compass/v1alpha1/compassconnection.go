// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/compass-runtime-agent/pkg/apis/compass/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CompassConnectionLister helps list CompassConnections.
// All objects returned here must be treated as read-only.
type CompassConnectionLister interface {
	// List lists all CompassConnections in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CompassConnection, err error)
	// Get retrieves the CompassConnection from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CompassConnection, error)
	CompassConnectionListerExpansion
}

// compassConnectionLister implements the CompassConnectionLister interface.
type compassConnectionLister struct {
	indexer cache.Indexer
}

// NewCompassConnectionLister returns a new CompassConnectionLister.
func NewCompassConnectionLister(indexer cache.Indexer) CompassConnectionLister {
	return &compassConnectionLister{indexer: indexer}
}

// List lists all CompassConnections in the indexer.
func (s *compassConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.CompassConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CompassConnection))
	})
	return ret, err
}

// Get retrieves the CompassConnection from the index for a given name.
func (s *compassConnectionLister) Get(name string) (*v1alpha1.CompassConnection, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("compassconnection"), name)
	}
	return obj.(*v1alpha1.CompassConnection), nil
}
