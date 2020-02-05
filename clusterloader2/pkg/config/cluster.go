/*
Copyright 2018 The Kubernetes Authors.

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

package config

// ClusterLoaderConfig represents all flags used by CLusterLoader
type ClusterLoaderConfig struct {
	ClusterConfig            ClusterConfig `json: clusterConfig`
	ReportDir                string        `json: reportDir`
	EnablePrometheusServer   bool          `json: enablePrometheusServer`
	TearDownPrometheusServer bool          `json: tearDownPrometheusServer`
	TestConfigPath           string        `json: testConfigPath`
	TestOverridesPath        []string      `json: testOverrides`
	PrometheusConfig         PrometheusConfig
}

// ClusterConfig is a structure that represents cluster description.
type ClusterConfig struct {
	KubeConfigPath             string   `json: kubeConfigPath`
	Nodes                      int      `json: nodes`
	Provider                   string   `json: provider`
	MasterIPs                  []string `json: masterIPs`
	MasterInternalIPs          []string `json: masterInternalIPs`
	MasterName                 string   `json: masterName`
	KubemarkRootKubeConfigPath string   `json: kubemarkRootKubeConfigPath`
}

// PrometheusConfig represents all flags used by prometheus.
type PrometheusConfig struct {
	EnableServer       bool
	TearDownServer     bool
	ScrapeEtcd         bool
	ScrapeNodeExporter bool
	ScrapeKubelets     bool
	ScrapeKubeProxy    bool
}

// GetMasterIp returns the first master ip, added for backward compatibility.
// TODO(mmatt): Remove this method once all the codebase is migrated to support multiple masters.
func (c *ClusterConfig) GetMasterIp() string {
	if len(c.MasterIPs) > 0 {
		return c.MasterIPs[0]
	}
	return ""
}

// GetMasterInternalIp returns the first internal master ip, added for backward compatibility.
// TODO(mmatt): Remove this method once all the codebase is migrated to support multiple masters.
func (c *ClusterConfig) GetMasterInternalIp() string {
	if len(c.MasterInternalIPs) > 0 {
		return c.MasterInternalIPs[0]
	}
	return ""
}
