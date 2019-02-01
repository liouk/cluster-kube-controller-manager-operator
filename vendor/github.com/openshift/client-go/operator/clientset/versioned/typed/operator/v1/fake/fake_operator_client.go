// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOperatorV1 struct {
	*testing.Fake
}

func (c *FakeOperatorV1) Consoles() v1.ConsoleInterface {
	return &FakeConsoles{c}
}

func (c *FakeOperatorV1) KubeAPIServers() v1.KubeAPIServerInterface {
	return &FakeKubeAPIServers{c}
}

func (c *FakeOperatorV1) KubeControllerManagers() v1.KubeControllerManagerInterface {
	return &FakeKubeControllerManagers{c}
}

func (c *FakeOperatorV1) OpenShiftAPIServers() v1.OpenShiftAPIServerInterface {
	return &FakeOpenShiftAPIServers{c}
}

func (c *FakeOperatorV1) OpenShiftControllerManagers() v1.OpenShiftControllerManagerInterface {
	return &FakeOpenShiftControllerManagers{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOperatorV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
