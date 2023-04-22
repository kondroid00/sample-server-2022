package main

import (
	"context"
	"flag"
	"github.com/kondroid00/sample-server-2022/main/domain/ent/migrate"
	"github.com/kondroid00/sample-server-2022/main/environment"
	"github.com/kondroid00/sample-server-2022/main/infra"
	infraService "github.com/kondroid00/sample-server-2022/main/infra/service"
	"github.com/kondroid00/sample-server-2022/package/env"
	"log"
	"os"
)

func main() {
	envVar := flag.String("env", "none", "environment to choose from debug, local, test, development, production")
	flag.Parse()

	var f env.EnvFunc
	if *envVar != "none" {
		f = env.EnvFuncImpl(func(s string) string {
			return *envVar
		})
	} else {
		f = env.EnvFuncImpl(os.Getenv)
	}

	environment, err := environment.New(f)
	if err != nil {
		log.Fatalln(err)
	}
	if err := environment.SetConfig(); err != nil {
		log.Fatalln(err)
	}

	infra, err := getInfra(environment.CurrentEnv())
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := infra.CloseDB(); err != nil {
			log.Fatalln(err)
		}
	}()
	ctx := context.Background()
	if err := infra.GetDB().Schema.Create(
		ctx,
		migrate.WithDropColumn(false),
		migrate.WithDropIndex(false),
	); err != nil {
		log.Fatalln(err)
	}
}

func getInfra(env environment.EnvType) (infra.Service, error) {
	switch env {
	case environment.Debug:
		fallthrough
	case environment.Local:
		fallthrough
	case environment.Test:
		return infra.SetConfig(&infraService.Local{})
	default:
		return infra.SetConfig(&infraService.Local{})
	}
}
