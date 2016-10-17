package api

import "github.com/emicklei/go-restful"

func Init() {
	// Register all rest endpoints.
	ws := &restful.WebService{}
	ws.Consumes(restful.MIME_JSON, "text/plain", "text/event-stream").
		Produces(restful.MIME_JSON, "text/plain", "text/event-stream")

	registerAPIs(ws)

	restful.Add(ws)

	// Add container filter to enable CORS and respond to OPTIONS.
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders: []string{"X-My-Header"},
		// The header "token" is not a standard header. It's used in auth server
		// to handle authentication.
		AllowedHeaders: []string{"Content-Type", "Accept", "token"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer,
	}
	restful.Filter(cors.Filter)
	restful.Filter(restful.OPTIONSFilter())
}

func registerAPIs(ws *restful.WebService) {
	ws.Route(ws.POST("/points").
		To(createPoint).
		Doc("create a point").
		Reads(Point{}).
		Writes(CreatePointResponse{}))
}
