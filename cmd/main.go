package main

import (
	"chipsiBackend/api/route"
	"chipsiBackend/bootstrap"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDbConnection()
	router := route.Setup(app.Db)

	app.Log.Info(fmt.Sprintf("Server is listening on PORT: %d", app.Cfg.Server.Port))
	addr := ":" + strconv.Itoa(app.Cfg.Server.Port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		app.Log.Error("Error starting Server: ", "err", err)
		os.Exit(1)
	}
}
