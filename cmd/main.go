package main

import (
	"context"
	"github.com/uzdemirdl/ready2play/pkg/application"
)

func main() {
	app := application.New()
	app.InitDB(context.Background())
	app.InitHTTP(context.Background())
	err := app.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
