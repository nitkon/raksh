// Copyright 2019 IBM Corp
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

package app

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestInsertVaultSecret(t *testing.T) {
	testVaultSecretName := "test-vault-secret"
	expectSecrets := []corev1.EnvVar{
		{
			Name: "SC_VAULT_ADDR",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: testVaultSecretName,
					},
					Key: "vaultAdd",
				},
			},
		},
		{
			Name: "SC_VAULT_TOKEN",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: testVaultSecretName,
					},
					Key: "vaultToken",
				},
			},
		},
		{
			Name: "SC_VAULT_SECRET",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: testVaultSecretName,
					},
					Key: "secretName",
				},
			},
		},
		{
			Name: "SC_VAULT_SYMM_KEY",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: testVaultSecretName,
					},
					Key: "keyName",
				},
			},
		},
	}
	testpod := &corev1.PodSpec{
		Containers: []corev1.Container{
			{Name: "c1"},
			{Name: "c2"},
		},
	}

	insertVaultSecret(testpod, testVaultSecretName)

	for i := range testpod.Containers {
		if equal := reflect.DeepEqual(expectSecrets, testpod.Containers[i].Env); !equal {
			t.Fatalf("For container %s, actual : %+v is not matching the expected: %+v", testpod.Containers[i].Name, testpod.Containers[0].Env, expectSecrets)
		}
	}
}
