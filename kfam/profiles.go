// Copyright 2019 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kfam

import (
	"context"

	"github.com/pluralsh/kubeflow-profile-controller/apis/kubeflow.org/v2alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type ProfileInterface interface {
	Create(profile *v2alpha1.Profile) (*v2alpha1.Profile, error)
	Delete(name string, opts *metav1.DeleteOptions) error
	Get(name string, opts metav1.GetOptions) (*v2alpha1.Profile, error)
	List(opts metav1.ListOptions) (*v2alpha1.ProfileList, error)
	Update(profile *v2alpha1.Profile) (*v2alpha1.Profile, error)
}

type ProfileClient struct {
	restClient rest.Interface
}

const Profiles = "profiles"

func (c *ProfileClient) Create(profile *v2alpha1.Profile) (*v2alpha1.Profile, error) {
	result := v2alpha1.Profile{}
	err := c.restClient.
		Post().
		Resource(Profiles).
		Body(profile).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *ProfileClient) Delete(name string, opts *metav1.DeleteOptions) error {
	return c.restClient.
		Delete().
		Resource(Profiles).
		Name(name).
		Body(opts).
		Do(context.TODO()).
		Error()
}

func (c *ProfileClient) Get(name string, opts metav1.GetOptions) (*v2alpha1.Profile, error) {
	result := v2alpha1.Profile{}
	err := c.restClient.
		Get().
		Resource(Profiles).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *ProfileClient) List(opts metav1.ListOptions) (*v2alpha1.ProfileList, error) {
	result := v2alpha1.ProfileList{}
	err := c.restClient.
		Get().
		Resource(Profiles).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *ProfileClient) Update(profile *v2alpha1.Profile) (*v2alpha1.Profile, error) {
	result := v2alpha1.Profile{}
	err := c.restClient.
		Put().
		Resource(Profiles).
		Body(profile).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}
