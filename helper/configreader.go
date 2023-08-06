package helper

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ObsCreds ObsCreds `toml:"Obs_Creds"`
}

type ObsCreds struct {
	Username string `toml:"obs_username"`
	Password string `toml:"obs_password"`
}

func ReadConfig(filePath string) (*Config, error) {
	var config Config

	configFileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if _, err := toml.Decode(string(configFileContent), &config); err != nil {
		return nil, err
	}

	// Print the extracted values
	fmt.Println("ReadConfig - Username:", config.ObsCreds.Username)
	fmt.Println("ReadConfig - Password:", config.ObsCreds.Password)

	return &config, nil
}
