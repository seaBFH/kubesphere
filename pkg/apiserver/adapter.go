package apiserver

import (
	"reflect"

	"k8s.io/client-go/dynamic/dynamicinformer"
	"kubesphere.io/kubesphere/pkg/informers"
)

var _ informers.GenericInformerFactory = &DynamicInformerGenericWrapper{}

// wrapper for adpating dynamic informer factory to generic informer factory(just for cache sync)
// WARN: DO NOT use this for other purpose for now
type DynamicInformerGenericWrapper struct {
	dynamicInformerFactory dynamicinformer.DynamicSharedInformerFactory
}

func (d *DynamicInformerGenericWrapper) Start(stopCh <-chan struct{}) {
	d.dynamicInformerFactory.Start(stopCh)
}

func (d *DynamicInformerGenericWrapper) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	// WARN: just for init cache, no need to return cache sync status
	// TODO: if the return value is needed, try implement this in the future
	d.dynamicInformerFactory.WaitForCacheSync(stopCh)
	return nil
}
