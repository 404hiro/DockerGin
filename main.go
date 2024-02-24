package main

import (
	"DockerGin/src/infrastructure/db"
	"DockerGin/src/infrastructure/logging"
	"DockerGin/src/infrastructure/router"
)

func main() {
	logger := logging.NewLogger("development")
	logger.Debug("debug log")

	sqlhandler := db.New()
	defer sqlhandler.Conn.Close()

	router.Dispatch(sqlhandler)
}
