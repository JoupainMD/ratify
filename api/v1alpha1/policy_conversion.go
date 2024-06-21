/*
Copyright The Ratify Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	unversioned "github.com/ratify-project/ratify/api/unversioned"
	"github.com/ratify-project/ratify/internal/constants"
	conversion "k8s.io/apimachinery/pkg/conversion"
)

// Convert unversioned PolicySpec to PolicySpec of v1alpha1.
//
//nolint:revive // ignore linter for autogenerated code
func Convert_unversioned_PolicySpec_To_v1alpha1_PolicySpec(in *unversioned.PolicySpec, out *PolicySpec, _ conversion.Scope) error {
	out.Parameters = in.Parameters
	return nil
}

// Convert unversioned PolicyStatus to PolicyStatus of v1alpha1.
//
//nolint:revive // ignore linter for autogenerated code
func Convert_unversioned_PolicyStatus_To_v1alpha1_PolicyStatus(in *unversioned.PolicyStatus, out *PolicyStatus, _ conversion.Scope) error {
	return nil
}

// Convert unversioned Policy to Policy of v1alpha1.
//
//nolint:revive // ignore linter for autogenerated code
func Convert_unversioned_Policy_To_v1alpha1_Policy(in *unversioned.Policy, out *Policy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	// metadata.name in v1alpha1 is same as spec.type in unversioned.
	out.ObjectMeta.Name = in.Spec.Type
	if err := Convert_unversioned_PolicySpec_To_v1alpha1_PolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return Convert_unversioned_PolicyStatus_To_v1alpha1_PolicyStatus(&in.Status, &out.Status, s)
}

// Convert Policy of v1alpha1 to unversioned Policy.
//
//nolint:revive // ignore linter for autogenerated code
func Convert_v1alpha1_Policy_To_unversioned_Policy(in *Policy, out *unversioned.Policy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	// metadata.name MUST be `ratify-policy` in unversioned.
	out.ObjectMeta.Name = constants.RatifyPolicy
	if err := Convert_v1alpha1_PolicySpec_To_unversioned_PolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	// spec.type in unversioned is same as metadata.name in v1alpha1.
	out.Spec.Type = in.ObjectMeta.Name
	return Convert_v1alpha1_PolicyStatus_To_unversioned_PolicyStatus(&in.Status, &out.Status, s)
}
