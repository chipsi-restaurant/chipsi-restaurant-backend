package main

import "chipsiBackend/bootstrap"

func main() {
	app := bootstrap.App()
	defer app.CloseDbConnection()
}
