/*
Copyright (C) 2018 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package annotations

import (
	"testing"
)

func TestStringMapContains(t *testing.T) {
	testcases := []struct {
		description string
		bigMap      map[string]string
		subset      map[string]string
		retval      bool
	}{
		{
			description: "subset in bigMap",
			bigMap:      map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"},
			subset:      map[string]string{"key1": "value1", "key2": "value2"},
			retval:      true,
		},
		{
			description: "bigmap missing key",
			bigMap:      map[string]string{"key1": "value1", "key3": "value3"},
			subset:      map[string]string{"key1": "value1", "key2": "value2"},
			retval:      false,
		},
		{
			description: "image annotation prefix with same non-json value",
			bigMap:      map[string]string{BDImageAnnotationPrefix: "value", "key1": "value1"},
			subset:      map[string]string{BDImageAnnotationPrefix: "value"},
			retval:      true,
		},
		{
			description: "image annotation prefix with different non-json value",
			bigMap:      map[string]string{BDImageAnnotationPrefix: "value", "key1": "value1"},
			subset:      map[string]string{BDImageAnnotationPrefix: "newValue"},
			retval:      false,
		},
		{
			description: "image annotation prefix with same json value",
			bigMap:      map[string]string{"otherkey": "othervalue", BDImageAnnotationPrefix: CreateVendorVulnerabilityAnnotation(true, "url", 20).AsString()},
			subset:      map[string]string{BDImageAnnotationPrefix: CreateVendorVulnerabilityAnnotation(true, "url", 20).AsString()},
			retval:      true,
		},
		{
			description: "image annotation prefix with different json value",
			bigMap:      map[string]string{"otherkey": "othervalue", BDImageAnnotationPrefix: CreateVendorVulnerabilityAnnotation(true, "url", 20).AsString()},
			subset:      map[string]string{BDImageAnnotationPrefix: CreateVendorVulnerabilityAnnotation(true, "url", 10).AsString()},
			retval:      false,
		},
		{
			description: "pod annotation prefix with same non-json value",
			bigMap:      map[string]string{BDPodAnnotationPrefix: "value", "key1": "value1"},
			subset:      map[string]string{BDPodAnnotationPrefix: "value"},
			retval:      true,
		},
		{
			description: "pod annotation prefix with different non-json value",
			bigMap:      map[string]string{BDPodAnnotationPrefix: "value", "key1": "value1"},
			subset:      map[string]string{BDPodAnnotationPrefix: "newValue"},
			retval:      false,
		},
		{
			description: "pod annotation prefix with same json value",
			bigMap:      map[string]string{"otherkey": "othervalue", BDPodAnnotationPrefix: CreateVendorVulnerabilityAnnotation(true, "url", 20).AsString()},
			subset:      map[string]string{BDPodAnnotationPrefix: CreateVendorVulnerabilityAnnotation(true, "url", 20).AsString()},
			retval:      true,
		},
		{
			description: "pod annotation prefix with different json value",
			bigMap:      map[string]string{"otherkey": "othervalue", BDPodAnnotationPrefix: CreateVendorVulnerabilityAnnotation(true, "url", 20).AsString()},
			subset:      map[string]string{BDPodAnnotationPrefix: CreateVendorVulnerabilityAnnotation(true, "url", 10).AsString()},
			retval:      false,
		},
		{
			description: "value differs",
			bigMap:      map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"},
			subset:      map[string]string{"key1": "value2", "key2": "value2"},
			retval:      false,
		},
	}

	for _, tc := range testcases {
		result := StringMapContains(tc.bigMap, tc.subset)
		if result != tc.retval {
			t.Errorf("[%s] expected %t got %t: bigMap %v, subset %v", tc.description, tc.retval, result, tc.bigMap, tc.subset)
		}
	}
}
