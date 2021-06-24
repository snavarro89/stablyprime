package main

import (
	api "github.com/snavarro89/stablyprime/api/numbers"
	app "github.com/snavarro89/stablyprime/app"
	primeget "github.com/snavarro89/stablyprime/functions/prime/get"
)

func main() {
	application := app.App{
		Data: app.Data{
			Numbers: api.NumbersModel{DB: nil},
		},
	}
	primeget.Data(application)
	primeget.Aws()
}
