package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gadhittana01/todolist/config"
)

func startHTTPServer(handler http.Handler, c *config.GlobalConfig) error {
	log.Println("Serving HTTP on ports :" + strconv.Itoa(c.HTTP.Port))
	port := fmt.Sprintf(":%d", c.HTTP.Port)
	return http.ListenAndServe(port, handler)
}
