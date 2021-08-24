// Copyright 2021 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package docker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/docker/distribution/reference"
	"github.com/juju/errors"
	"github.com/juju/featureflag"
	"gopkg.in/yaml.v2"

	"github.com/juju/juju/feature"
)

// The default server address - https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#-em-secret-docker-registry-em-
const defaultServerAddress = "https://index.docker.io/v1/"

// TokenAuthConfig contains authorization information for token auth.
// Juju does not support the docker credential helper because k8s does not support it either.
// https://kubernetes.io/docs/concepts/containers/images/#configuring-nodes-to-authenticate-to-a-private-registry
type TokenAuthConfig struct {
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// IdentityToken is used to authenticate the user and get
	// an access token for the registry.
	IdentityToken string `json:"identitytoken,omitempty" yaml:"identitytoken,omitempty"`

	// RegistryToken is a bearer token to be sent to a registry
	RegistryToken string `json:"registrytoken,omitempty" yaml:"registrytoken,omitempty"`
}

// Empty checks if the auth information is empty.
func (ac TokenAuthConfig) Empty() bool {
	return ac.IdentityToken == "" && ac.RegistryToken == ""
}

// Validate validates the spec.
func (ac *TokenAuthConfig) Validate() error {
	return nil
}

// BasicAuthConfig contains authorization information for basic auth.
type BasicAuthConfig struct {
	// Auth is the base64 encoded "username:password" string.
	Auth string `json:"auth,omitempty" yaml:"auth,omitempty"`

	// Username holds the username used to gain access to a non-public image.
	Username string `json:"username" yaml:"username"`

	// Password holds the password used to gain access to a non-public image.
	Password string `json:"password" yaml:"password"`
}

// Empty checks if the auth information is empty.
func (ba BasicAuthConfig) Empty() bool {
	return ba.Auth == "" && ba.Username == "" && ba.Password == ""
}

// Validate validates the spec.
func (ba BasicAuthConfig) Validate() error {
	return nil
}

// ImageRepoDetails contains authorization information for connecting to a Registry.
type ImageRepoDetails struct {
	BasicAuthConfig `json:",inline" yaml:",inline"`
	TokenAuthConfig `json:",inline" yaml:",inline"`

	// Repository is the namespace of the image repo.
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// ServerAddress is the auth server address.
	ServerAddress string `json:"serveraddress,omitempty" yaml:"serveraddress,omitempty"`
}

func (rid ImageRepoDetails) AuthEqual(r ImageRepoDetails) bool {
	return reflect.DeepEqual(rid.BasicAuthConfig, r.BasicAuthConfig) &&
		reflect.DeepEqual(rid.TokenAuthConfig, r.TokenAuthConfig)
}

func (rid ImageRepoDetails) IsPrivate() bool {
	return !rid.BasicAuthConfig.Empty() || !rid.TokenAuthConfig.Empty()
}

type dockerConfigData struct {
	Auths map[string]ImageRepoDetails `json:"auths"`
}

// SecretData returns secret data format.
func (rid ImageRepoDetails) SecretData() ([]byte, error) {
	if rid.BasicAuthConfig.Empty() && rid.TokenAuthConfig.Empty() {
		// No auth information is required for a public repository.
		return nil, nil
	}
	rid.Repository = ""
	o := dockerConfigData{
		Auths: map[string]ImageRepoDetails{
			rid.ServerAddress: rid,
		},
	}
	return json.Marshal(o)
}

// String returns yaml format.
func (rid *ImageRepoDetails) String() string {
	d, _ := yaml.Marshal(rid)
	return string(d)
}

// Validate validates the spec.
func (rid ImageRepoDetails) Validate() error {
	if rid.Repository == "" {
		return errors.NotValidf("empty repository")
	}
	_, err := reference.ParseNormalizedNamed(rid.Repository)
	if err != nil {
		return errors.NewNotValid(err, fmt.Sprintf("docker image path %q", rid.Repository))
	}
	if err := rid.BasicAuthConfig.Validate(); err != nil {
		return errors.Annotatef(err, "validating basic auth config for repository %q", rid.Repository)
	}
	if err := rid.TokenAuthConfig.Validate(); err != nil {
		return errors.Annotatef(err, "validating token auth config for repository %q", rid.Repository)
	}
	return nil
}

func fileExists(p string) (bool, error) {
	info, err := os.Stat(p)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, errors.Trace(err)
	}
	return !info.IsDir(), nil
}

// NewImageRepoDetails tries to parse a file path or file content and returns an instance of ImageRepoDetails.
func NewImageRepoDetails(contentOrPath string) (*ImageRepoDetails, error) {
	data := []byte(contentOrPath)
	isPath, err := fileExists(contentOrPath)
	if err == nil && isPath {
		data, err = ioutil.ReadFile(contentOrPath)
		if err != nil {
			return nil, errors.Trace(err)
		}
	}
	o := &ImageRepoDetails{}
	err = yaml.Unmarshal(data, o)
	if err != nil {
		return &ImageRepoDetails{Repository: contentOrPath}, nil
	}
	if o.ServerAddress == "" {
		o.ServerAddress = defaultServerAddress
	}
	if o.IsPrivate() && !featureflag.Enabled(feature.PrivateRegistry) {
		return nil, errors.New(
			fmt.Sprintf("private registry support is under development, enable feature flag %q to test it out", feature.PrivateRegistry),
		)
	}
	return o, nil
}
