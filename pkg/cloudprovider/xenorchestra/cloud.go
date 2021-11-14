package xenorchestra

import (
	"io"
	"io/ioutil"

	cloudprovider "k8s.io/cloud-provider"
	klog "k8s.io/klog/v2"
)

const (
	// RegisteredProviderName is the name of the cloud provider registered with
	// Kubernetes.
	RegisteredProviderName string = "xenorchestra"

	// ProviderName is the name used for constructing Provider ID
	ProviderName string = "xenorchestra"

	// ClientName is the user agent passed into the controller client builder.
	ClientName string = "xo-cloud-controller-manager"

	// dualStackFeatureGateEnv is a required environment variable when enabling dual-stack nodes
	// https://kubernetes.io/docs/concepts/services-networking/dual-stack/
	dualStackFeatureGateEnv string = "ENABLE_ALPHA_DUAL_STACK"
)

func init() {
	cloudprovider.RegisterCloudProvider(RegisteredProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		cfg, err := ioutil.ReadAll(config)
		if err != nil {
			klog.Errorf("ReadAll failed: %s", err)
			return nil, err
		}

		return newXenOrchestra(cfg, true)
	})
}

// Creates new Controller node interface and returns
func newXenOrchestra(cfg []byte, finalize ...bool) (*XenOrchestra, error) {
	return nil, nil
}

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping or run custom controllers specific to the cloud provider.
// Any tasks started here should be cleaned up when the stop channel closes.
func (xo *XenOrchestra) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {

}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (xo *XenOrchestra) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return nil, false
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (xo *XenOrchestra) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

// InstancesV2 is an implementation for instances and should only be implemented by external cloud providers.
// Implementing InstancesV2 is behaviorally identical to Instances but is optimized to significantly reduce
// API calls to the cloud provider when registering and syncing nodes. Implementation of this interface will
// disable calls to the Zones interface. Also returns true if the interface is supported, false otherwise.
func (xo *XenOrchestra) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
// DEPRECATED: Zones is deprecated in favor of retrieving zone/region information from InstancesV2.
// This interface will not be called if InstancesV2 is enabled.
func (xo *XenOrchestra) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (xo *XenOrchestra) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (xo *XenOrchestra) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (xo *XenOrchestra) ProviderName() string {
	return "xenorchestra"
}

// HasClusterID returns true if a ClusterID is required and set
func (xo *XenOrchestra) HasClusterID() bool {
	return false
}
