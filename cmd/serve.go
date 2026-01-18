package cmd

import (
	"github.com/Hamiduzzaman96/Test-Project/config"
	"github.com/Hamiduzzaman96/Test-Project/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.GetRoutes(cnf)
}
