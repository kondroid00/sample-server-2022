package environment

import (
	"cuelang.org/go/cue/cuecontext"
	"fmt"
	"github.com/kondroid00/sample-server-2022/package/env"
	"github.com/kondroid00/sample-server-2022/package/errors"
	"io/ioutil"
	"path/filepath"
)

type EnvType string

const (
	// Debug デバッグ環境
	Debug EnvType = "debug"
	// Local ローカル環境
	Local EnvType = "local"
	// Test テスト環境
	Test EnvType = "test"
	// Development 開発環境
	Development EnvType = "development"
	// Production 本番環境
	Production EnvType = "production"
)

type Environment struct {
	env     EnvType
	envFunc env.EnvFunc
}

func New(envFunc env.EnvFunc) (*Environment, error) {
	env := EnvType(envFunc.Getenv("GO_ENV"))

	switch env {
	case Debug:
		fallthrough
	case Local:
		fallthrough
	case Test:
		fallthrough
	case Development:
		fallthrough
	case Production:
		return &Environment{
			env:     env,
			envFunc: envFunc,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("cannot select %s as env", env))
}

func (e *Environment) CurrentEnv() EnvType {
	return e.env
}

func (e *Environment) SetConfig() error {
	config, err := e.generateConfig()
	if err != nil {
		return err
	}
	configData = config
	return nil
}

func (e *Environment) generateConfig() (Config, error) {
	c := cuecontext.New()
	bytes, err := ioutil.ReadFile(e.getCueFile())
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := c.CompileBytes(bytes).Decode(&config); err != nil {
		return Config{}, err
	}
	switch e.env {
	case Development:
		fallthrough
	case Production:
		e.addSecret(&config)
	}
	return config, nil
}

func (e *Environment) getCueFile() string {
	return filepath.Join("./environment/settings", string(e.env)+".cue")
}

func (e *Environment) addSecret(congfig *Config) {
	for i, v := range congfig.DB.DBConfigs {
		var user string
		var host string
		var password string
		switch v.Name {
		case DB_MAIN_NAME:
			user = e.envFunc.Getenv(ENV_DB_MAIN_USER)
			host = e.envFunc.Getenv(ENV_DB_MAIN_HOST)
			password = e.envFunc.Getenv(ENV_DB_MAIN_PASSWORD)
		case DB_READ_NAME:
			user = e.envFunc.Getenv(ENV_DB_READ_USER)
			host = e.envFunc.Getenv(ENV_DB_READ_HOST)
			password = e.envFunc.Getenv(ENV_DB_READ_PASSWORD)
		}
		congfig.DB.DBConfigs[i].User = user
		congfig.DB.DBConfigs[i].Host = host
		congfig.DB.DBConfigs[i].Password = password
	}
}
