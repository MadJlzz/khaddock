package engine

import (
	"os"

	"gopkg.in/yaml.v3"
)

type KhaddockBase struct {
	ApiVersion string              `yaml:"apiVersion"`
	Kind       string              `yaml:"kind"`
	Metadata   KhaddockMetadata    `yaml:"metadata"`
	Labels     LabelsConfiguration `yaml:"labels"`
}

type KhaddockMetadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type LabelsConfiguration map[string]string

type SailorConfiguration struct {
	KhaddockBase `yaml:",inline"`
	Specs        struct {
		SysctlConfiguration map[string]string `yaml:"sysctl"`
	}
}

func Unmarshal(filepath string, out interface{}) error {
	d, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(d, out)
	if err != nil {
		return err
	}

	return nil
}
