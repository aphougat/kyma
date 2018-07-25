// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kyma-project/kyma/components/api-controller/pkg/clients/authentication.istio.io/clientset/versioned/typed/authentication.istio.io/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAuthenticationV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAuthenticationV1alpha1) Policies(namespace string) v1alpha1.PolicyInterface {
	return &FakePolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAuthenticationV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}