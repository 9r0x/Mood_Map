// test project main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/gaocegege/hackys-backend-writer/api"
	"github.com/gaocegege/hackys-backend-writer/pkg/log"
)

const (
	port = 8088
)

func main() {
	api.Init()
	server := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: restful.DefaultContainer}
	log.Infof("circle server listening on %d", port)
	log.Fatal(server.ListenAndServe())
}
