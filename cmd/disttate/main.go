package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/glog"

	"github.com/ckatsak/crdgentest/controller"
	dsv1a1 "github.com/ckatsak/crdgentest/pkg/apis/tests.ckatsak/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	podName, podNs, podIP string
)

func init() {
	_ = flag.Set("logtostderr", "true")
}

func main() {
	flag.Parse()

	// Parse env vars.
	podName, podNs, podIP = os.Getenv("POD_NAME"), os.Getenv("POD_NS"), os.Getenv("POD_IP")
	glog.Infof("pod name: %q", podName)
	glog.Infof("pod namespace: %q", podNs)
	glog.Infof("pod IP: %q", podIP)

	// Create a DisttateController and warm up its cache.
	dsc := controller.NewDisttateController(nil, podNs)
	if err := dsc.Indexer.Add(
		&dsv1a1.Disttate{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "yolodisttate",
				Namespace: podNs,
			},
			//Spec: &dsv1a1.DisttateSpec{
			//	Bitset: dsv1a1.NewBitSet(5),
			//},
		}); err != nil {
		glog.Errorf("error warming up the cache: %v", err)
	}

	// Run the controller until the stop channel is closed.
	stopChan := make(chan struct{})
	go dsc.Run(stopChan)

	// Wait for a SIGTERM.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	<-signalChan

	// Mock graceful shut down.
	glog.Infof("Notifying everyone to shut down and waiting 2 sec...")
	close(stopChan)
	time.Sleep(2 * time.Second)
	glog.Infof("Shutting down.")
}
