package main

import (
	"net/http"

	"github.com/anatoliin/metrics-collector/internal/config"
	"github.com/anatoliin/metrics-collector/internal/router"
)

func main() {
	options := config.InitOptions()
	r := router.NewRouter()

	err := http.ListenAndServe(options.RunAddress, r)
	if err != nil {
		panic(err)
	}
}
