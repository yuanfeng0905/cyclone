/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubernetes "github.com/caicloud/cyclone/kubernetes"
	v1alpha1 "github.com/caicloud/cyclone/listers/cyclone/v1alpha1"
	cyclonev1alpha1 "github.com/caicloud/cyclone/pkg/apis/cyclone/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// WorkflowRunInformer provides access to a shared informer and lister for
// WorkflowRuns.
type WorkflowRunInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.WorkflowRunLister
}

type workflowRunInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWorkflowRunInformer constructs a new informer for WorkflowRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkflowRunInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkflowRunInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWorkflowRunInformer constructs a new informer for WorkflowRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkflowRunInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CycloneV1alpha1().WorkflowRuns(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CycloneV1alpha1().WorkflowRuns(namespace).Watch(options)
			},
		},
		&cyclonev1alpha1.WorkflowRun{},
		resyncPeriod,
		indexers,
	)
}

func (f *workflowRunInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWorkflowRunInformer(client.(kubernetes.Interface), f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *workflowRunInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cyclonev1alpha1.WorkflowRun{}, f.defaultInformer)
}

func (f *workflowRunInformer) Lister() v1alpha1.WorkflowRunLister {
	return v1alpha1.NewWorkflowRunLister(f.Informer().GetIndexer())
}
