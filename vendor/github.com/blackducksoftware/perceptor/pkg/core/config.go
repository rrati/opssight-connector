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

package core

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// Config contains all configuration for Perceptor
type Config struct {
	HubHost                         string
	HubUser                         string
	HubUserPasswordEnvVar           string
	HubClientTimeoutMilliseconds    int
	HubPort                         int
	PruneOrphanedImagesPauseMinutes int
	ConcurrentScanLimit             int
	UseMockMode                     bool
	Port                            int
	LogLevel                        string
}

// HubClientTimeout converts the milliseconds to a duration
func (config *Config) HubClientTimeout() time.Duration {
	return time.Duration(config.HubClientTimeoutMilliseconds) * time.Millisecond
}

// GetLogLevel .....
func (config *Config) GetLogLevel() (log.Level, error) {
	return log.ParseLevel(config.LogLevel)
}
