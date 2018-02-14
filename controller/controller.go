package controller

import (
	"fmt"
	"time"

	"github.com/golang/glog"

	dsv1a1 "github.com/ckatsak/crdgentest/pkg/apis/tests.ckatsak/v1alpha1"
	dsclientset "github.com/ckatsak/crdgentest/pkg/client/clientset/versioned"
	dsiv1a1 "github.com/ckatsak/crdgentest/pkg/client/informers/externalversions/tests.ckatsak/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type DisttateController struct {
	Controller cache.Controller
	Indexer    cache.Indexer
	queue      workqueue.RateLimitingInterface
}

func NewDisttateController(klient dsclientset.Interface, namespace string) *DisttateController {
	// If needed, create a Disttate client set.
	if klient == nil {
		konfig, err := rest.InClusterConfig()
		if err != nil {
			glog.Fatalf("error in rest.InClusterConfig: %v", err)
		}
		klient, err = dsclientset.NewForConfig(konfig)
		if err != nil {
			glog.Fatalf("error in dsclientset.NewForConfig: %v", err)
		}
	}

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer := dsiv1a1.NewFilteredDisttateInformer(
		klient,
		namespace,
		30*time.Second,
		cache.Indexers{},
		func(opts *metav1.ListOptions) {
			opts.LabelSelector = "app=disttate"
		},
	)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			glog.Infof("ADD caught!")
			if key, err := cache.MetaNamespaceKeyFunc(obj); err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			glog.Infof("UPDATE caught!")
			if key, err := cache.MetaNamespaceKeyFunc(new); err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			glog.Infof("DELETE caught!")
			if key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj); err == nil {
				queue.Add(key)
			}
		},
	})

	return &DisttateController{
		Controller: informer,
		Indexer:    informer.GetIndexer(),
		queue:      queue,
	}
}

func (dsc *DisttateController) Run(stopChan <-chan struct{}) {
	defer runtime.HandleCrash()
	defer dsc.queue.ShutDown()

	glog.Infof("DisttateController: Running...")
	go dsc.Controller.Run(stopChan)

	if !cache.WaitForCacheSync(stopChan, dsc.Controller.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	go wait.Until(dsc.runWorker, time.Second, stopChan)

	<-stopChan
	glog.Infof("DisttateController: Stopping...")
}

func (dsc *DisttateController) runWorker() {
	for dsc.processNextItem() {
	}
}

func (dsc *DisttateController) processNextItem() bool {
	key, quit := dsc.queue.Get() // <-- NOTE: blocks for as long as the queue is empty
	if quit {
		return false
	}
	defer dsc.queue.Done(key)

	err := dsc.BusinessLogic(key.(string))
	dsc.handleErr(err, key)

	return true
}

func (dsc *DisttateController) BusinessLogic(key string) error {
	obj, exists, err := dsc.Indexer.GetByKey(key)
	switch {
	case err != nil:
		glog.Errorf("BusinessLogic: error in dsc.Indexer.GetByKey for key %q: %v", err, key)
		return err
	case !exists:
		glog.Warningf("BusinessLogic: object %q does not appear to be in the cache!", key)
	default:
		disttate := obj.(*dsv1a1.Disttate)
		glog.Infof("BusinessLogic: Spec.Name: %q, Spec.RingSize: %d, Spec.BitSet: %#v\ncool name: %q, cool ring size: %d\n%#v",
			disttate.Spec.Name, disttate.Spec.RingSize, disttate.Spec.Bitset,
			disttate.Spec.GetCoolName(), disttate.Spec.GetCoolRingSize(),
			disttate)
		//dsc.countdown(disttate.Spec.RingSize, key)
	}
	return nil
}

// call it from within DisttateController.BusinessLogic() to see how subsequent
// updates during the handling of a previous update, get collapsed into the
// latest version.
func (dsc *DisttateController) countdown(x int, key string) {
	for i := x; i > 0; i-- {
		v, _, _ := dsc.Indexer.GetByKey(key)
		glog.Infof("%d... (%d)", i, v.(*dsv1a1.Disttate).Spec.GetCoolRingSize())
		time.Sleep(time.Second)
	}
}

func (dsc *DisttateController) handleErr(err error, key interface{}) {
	if err == nil {
		dsc.queue.Forget(key)
		return
	}

	glog.Infof("(reties for key %q: %d)", key.(string), dsc.queue.NumRequeues(key))

	if dsc.queue.NumRequeues(key) < 5 {
		glog.Infof("error in BusinessLogic for %v: %v", key, err)

		glog.Infof("re-enqueuing")
		dsc.queue.AddRateLimited(key)
		return
	}

	glog.Infof("forgetting!")
	dsc.queue.Forget(key)
	runtime.HandleError(err)
	glog.Infof("Dropping %v out of the queue: %v", key, err)
}
