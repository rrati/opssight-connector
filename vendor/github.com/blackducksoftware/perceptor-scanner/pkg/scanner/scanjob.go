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

package scanner

import (
	"fmt"

	"github.com/blackducksoftware/perceptor-scanner/pkg/common"
)

type ScanJob struct {
	PullSpec              string
	Sha                   string
	HubProjectName        string
	HubProjectVersionName string
	HubScanName           string
}

func NewScanJob(repository string, sha string, hubProjectName string, hubProjectVersionName string, hubScanName string) *ScanJob {
	pullSpec := fmt.Sprintf("%s@sha256:%s", repository, sha)
	return &ScanJob{
		PullSpec:              pullSpec,
		Sha:                   sha,
		HubProjectName:        hubProjectName,
		HubProjectVersionName: hubProjectVersionName,
		HubScanName:           hubScanName}
}

func (sj *ScanJob) image() *common.Image {
	return &common.Image{PullSpec: sj.PullSpec}
}
