package app

import "github.com/kelseyhightower/envconfig"

// SSL specification app
type SSL struct {
	Enable bool   `required:"true"`
	Cert   string `required:"false"`
	Key    string `required:"false"`
}

// Authentication specification app
type Authentication struct {
	Enable bool   `required:"true"`
	JwtKey string `required:"false" split_words:"true"`
}

// Config support to the server
type Config struct {
	Ssl  SSL            `required:"true"`
	Port string         `required:"true"`
	Auth Authentication `required:"true"`
}

// NewConfig parse the environment values to return a initialized configuration
func NewConfig(appName string) (Config, error) {
	var err error
	var c Config

	if err = envconfig.Process(appName, &c); err != nil {
		return Config{}, err
	}
	return c, nil
}
