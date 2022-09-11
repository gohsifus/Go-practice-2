package configs

import (
	"encoding/json"
	"os"
	"task11/errs"
)

type ServerConfig struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	PathToLog string `json:"path_to_log"`
}

func NewConfig() *ServerConfig {
	return &ServerConfig{}
}

func (s *ServerConfig) LoadConfigs(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return errs.Wrap(err)
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		errs.Wrap(err)
	}

	return nil
}
