package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"task11/errs"
)

type PostgresConfig struct {
	host     string `json:"host"`
	port     string `json:"port"`
	user     string `json:"user"`
	password string `json:"password"`
	dbName   string `json:"db_name"`
}

func NewConfig() *PostgresConfig {
	return &PostgresConfig{}
}

//LoadConfigs загрузит логи из json файла
func (p *PostgresConfig) LoadConfigs(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return errs.Wrap(err)
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		errs.Wrap(err)
	}

	return nil
}

func (p *PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.host,
		p.port,
		p.user,
		p.password,
		p.dbName,
	)
}
