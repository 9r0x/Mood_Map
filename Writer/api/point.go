package api

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/gaocegege/hackys-backend-writer/cognitive"
	"github.com/gaocegege/hackys-backend-writer/pkg/log"
	"gopkg.in/pg.v4"
)

var db = pg.Connect(&pg.Options{
	Addr:     "postgres:5432",
	User:     "postgres",
	Password: "password",
	Database: "writer",
})

func createPoint(request *restful.Request, response *restful.Response) {
	point := &Point{}
	err := request.ReadEntity(point)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusBadRequest, "Writer Service unable to parse request body")
		return
	}

	log.Infof("Point: %s\n", point)
	go PushPointToDB(point)
	var createPointResponse CreatePointResponse
	createPointResponse.ErrorMessage = ""
	response.WriteHeaderAndEntity(http.StatusAccepted, createPointResponse)
}

func PushPointToDB(point *Point) {
	text := point.Content
	result, err := cognitive.RecognizeText(text)
	if err != nil {
		log.Warnf("Err when push point to DB: %v", err)
		return
	}
	log.Infof("Result: %v", result)

	pointInDB := &PointInDB{}
	pointInDB.TimeStamp = point.Date
	pointInDB.X = point.Location.Lat
	pointInDB.Y = point.Location.Lng
	if result.Documents[0].Score > 0.5 {
		// Happy
		pointInDB.TagID = 5
	} else {
		// Anger
		pointInDB.TagID = 1
	}
	err = db.Create(pointInDB)
	if err != nil {
		log.Warnf("DB error: %v", err)
	}
}
