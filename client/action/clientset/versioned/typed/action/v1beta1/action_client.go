// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	http "net/http"

	actionv1beta1 "github.com/stolostron/cluster-lifecycle-api/action/v1beta1"
	scheme "github.com/stolostron/cluster-lifecycle-api/client/action/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ActionV1beta1Interface interface {
	RESTClient() rest.Interface
	ManagedClusterActionsGetter
}

// ActionV1beta1Client is used to interact with features provided by the action.open-cluster-management.io group.
type ActionV1beta1Client struct {
	restClient rest.Interface
}

func (c *ActionV1beta1Client) ManagedClusterActions(namespace string) ManagedClusterActionInterface {
	return newManagedClusterActions(c, namespace)
}

// NewForConfig creates a new ActionV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ActionV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ActionV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ActionV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ActionV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new ActionV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ActionV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ActionV1beta1Client for the given RESTClient.
func New(c rest.Interface) *ActionV1beta1Client {
	return &ActionV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := actionv1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ActionV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
