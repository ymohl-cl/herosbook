package main

import (
	"encoding/json"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/ymohl-cl/herosbook/cmd/api/constant"
	"github.com/ymohl-cl/herosbook/pkg/cassandra"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
)

type ssl struct {
	Cert   string `json:"certificat" valid:"required"`
	Key    string `json:"key" valid:"required"`
	Domain string `json:"domain" valid:"required"`
}

type config struct {
	SSL       ssl               `json:"api" valid:"required"`
	SQL       postgres.Conf     `json:"psql" valid:"required"`
	CQL       cassandra.Conf    `json:"cassandra" valid:"required"`
	Constants constant.Constant `json:"constant_params" valid:"required"`
}

func configure() (*config, error) {
	var c config

	// open config file
	file, err := os.Open("/bin/config_example.json")
	if err != nil {
		return nil, err
	}

	// get infos file
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// create buffer file
	data := make([]byte, stat.Size())

	// read file
	if _, err := file.Read(data); err != nil {
		return nil, err
	}

	// unmarshall JSON
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	// check ValidateStruct
	if ok, err := govalidator.ValidateStruct(&c); !ok {
		return nil, err
	}

	return &c, nil
}
