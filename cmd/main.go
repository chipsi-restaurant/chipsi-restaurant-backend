package main

import (
	"chipsiBackend/api/route"
	"chipsiBackend/bootstrap"
	"chipsiBackend/internal/scheduler"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDbConnection()

	go scheduler.StartOrderStatusUpdater(app.Db, app.Log)

	router := route.Setup(app)
	app.Log.Info(fmt.Sprintf("Server is listening on PORT: %d", app.Cfg.Server.Port))
	addr := ":" + strconv.Itoa(app.Cfg.Server.Port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		app.Log.Error("Error starting Server: ", "err", err)
		os.Exit(1)
	}
}
