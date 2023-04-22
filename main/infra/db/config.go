package db

import "github.com/kondroid00/sample-server-2022/main/environment"

type (
	DBHost string
)

const (
	// DBHost
	MainDB DBHost = environment.DB_MAIN_NAME
	ReadDB DBHost = environment.DB_READ_NAME
)
