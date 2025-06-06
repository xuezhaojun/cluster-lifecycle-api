// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	http "net/http"

	scheme "github.com/stolostron/cluster-lifecycle-api/client/imageregistry/clientset/versioned/scheme"
	imageregistryv1alpha1 "github.com/stolostron/cluster-lifecycle-api/imageregistry/v1alpha1"
	rest "k8s.io/client-go/rest"
)

type ImageregistryV1alpha1Interface interface {
	RESTClient() rest.Interface
	ManagedClusterImageRegistriesGetter
}

// ImageregistryV1alpha1Client is used to interact with features provided by the imageregistry.open-cluster-management.io group.
type ImageregistryV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ImageregistryV1alpha1Client) ManagedClusterImageRegistries(namespace string) ManagedClusterImageRegistryInterface {
	return newManagedClusterImageRegistries(c, namespace)
}

// NewForConfig creates a new ImageregistryV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ImageregistryV1alpha1Client, error) {
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

// NewForConfigAndClient creates a new ImageregistryV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ImageregistryV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ImageregistryV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ImageregistryV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ImageregistryV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ImageregistryV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ImageregistryV1alpha1Client {
	return &ImageregistryV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := imageregistryv1alpha1.SchemeGroupVersion
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
func (c *ImageregistryV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
