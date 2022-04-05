package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type VeleroVolumeCfg struct {
	IncludeNamespaces     string `yaml:"includeNamespaces,omitempty"`
	ExcludeNamespaces     string `yaml:"excludeNamespaces,omitempty"`
	IncludeVolumeTypes    string `yaml:"includeVolumeTypes,omitempty"`
	ExcludeVolumeTypes    string `yaml:"excludeVolumeTypes,omitempty"`
	IncludeStorageClasses string `yaml:"includeStorageClasses,omitempty"`
	ExcludeStorageClasses string `yaml:"excludeStorageClasses,omitempty"`
	ExcludeJobs           string `yaml:"excludeJobs,omitempty"`
}

type ClusterServerCfg struct {
	// The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.
	MasterURL string `yaml:"masterURL,omitempty"`
	// Path to a kubeconfig. Only required if out-of-cluster.
	KubeConfig string `yaml:"kubeConfig,omitempty"`
	// LeaseLock namespace
	LeaseLockNamespace string `yaml:"leaseLockNamespace,omitempty"`
}

type Config struct {
	ClusterServerCfg *ClusterServerCfg `yaml:"clusterServerCfg,omitempty"`
	VeleroVolumeCfg  *VeleroVolumeCfg  `yaml:"veleroVolumeCfg,omitempty"`
}

// LoadConfig parses configuration file and returns
// an initialized Settings object and an error object if any. For instance if it
// cannot find the configuration file it will set the returned error appropriately.
func LoadConfig(path string) (*Config, error) {
	c := &Config{}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read configuration file: %s, error: %s", path, err)
	}
	if err = yaml.Unmarshal(contents, c); err != nil {
		return nil, fmt.Errorf("Failed to parse configuration, error: %s", err)
	}
	c.initializeDefaults()
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize default value of configuration : %s", err)
	}
	if err = c.validate(); err != nil {
		return nil, fmt.Errorf("Invalid configuration, error: %s", err)
	}
	return c, nil
}

// initializeDefaults initializes empty settings with default values.
func (c *Config) initializeDefaults() error {
	if c.ClusterServerCfg.LeaseLockNamespace == "" {
		namespace, err := getOperatorNamespace()
		if err != nil {
			return err
		}
		c.ClusterServerCfg.LeaseLockNamespace = namespace
	}
	return nil
}

// validate the configuration
func (c *Config) validate() error {
	if c.VeleroVolumeCfg.IncludeNamespaces != "" && c.VeleroVolumeCfg.ExcludeNamespaces != "" {
		return fmt.Errorf("Invalid configuration : specify only one of IncludeNamespaces or ExcludeNamespaces")
	}
	if c.VeleroVolumeCfg.IncludeVolumeTypes != "" && c.VeleroVolumeCfg.ExcludeVolumeTypes != "" {
		return fmt.Errorf("Invalid configuration : specify only one of IncludeVolumeTypes or ExcludeVolumeTypes")
	}
	if c.VeleroVolumeCfg.IncludeStorageClasses != "" && c.VeleroVolumeCfg.ExcludeStorageClasses != "" {
		return fmt.Errorf("Invalid configuration : specify only one of IncludeStorageClasses or ExcludeStorageClasses")
	}
	return nil
}

var errNoNamespace = fmt.Errorf("namespace not found for current environment")

func getOperatorNamespace() (string, error) {
	nsBytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		if os.IsNotExist(err) {
			return "", errNoNamespace
		}
		return "", err
	}
	ns := strings.TrimSpace(string(nsBytes))
	return ns, nil
}
