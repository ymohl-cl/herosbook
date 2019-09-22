package config

import "github.com/kelseyhightower/envconfig"

// ParseEnv get the prefix and the configuration object to search and fill the model structure
func ParseEnv(prefix string, conf interface{}) error {
	var err error

	if err = envconfig.Process(prefix, conf); err != nil {
		return err
	}
	return nil
}
