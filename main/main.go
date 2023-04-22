package main

import (
	"context"
	"fmt"
	"github.com/kondroid00/sample-server-2022/main/environment"
	"github.com/kondroid00/sample-server-2022/main/infra"
	"github.com/kondroid00/sample-server-2022/main/infra/di"
	infraService "github.com/kondroid00/sample-server-2022/main/infra/service"
	"github.com/kondroid00/sample-server-2022/main/server/middleware"
	"github.com/kondroid00/sample-server-2022/package/env"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/kondroid00/sample-server-2022/main/interface/openapi"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	sentryecho "github.com/getsentry/sentry-go/echo"
)

const DEFAULT_PORT = "8080"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	environment, err := environment.New(env.EnvFuncImpl(os.Getenv))
	if err != nil {
		log.Fatalln(err)
	}
	if err := environment.SetConfig(); err != nil {
		log.Fatalln(err)
	}

	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatalln(err)
	}
	swagger.Servers = nil

	i, err := getInfra(environment.CurrentEnv())
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()
	h, err := di.NewHandler(i)
	if err != nil {
		log.Fatalln(err)
	}

	openapi.RegisterHandlers(e, h)

	e.Use(
		echomiddleware.Logger(),
		oapimiddleware.OapiRequestValidator(swagger),
		sentryecho.New(sentryecho.Options{}),
		middleware.NewHttpStatusInterceptor(),
	)

	go func() {
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	switch <-stop {
	// kill -SIGHUP XXXX
	case syscall.SIGHUP:
		log.Println("hungup")

		// kill -SIGINT XXXX or Ctrl+c
	case syscall.SIGINT:
		log.Println("Warikomi")

		// kill -SIGTERM XXXX
	case syscall.SIGTERM:
		log.Println("force stop")

		// kill -SIGQUIT XXXX
	case syscall.SIGQUIT:
		log.Println("stop and core dump")
	default:
		log.Println("Unknown signal.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	log.Println("Start GracefulStop")
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Finish graceful stop")
}

func getInfra(env environment.EnvType) (infra.Service, error) {
	switch env {
	case environment.Debug:
		fallthrough
	case environment.Local:
		return infra.SetConfig(&infraService.Local{})
	default:
		return infra.SetConfig(&infraService.Local{})
	}
}
