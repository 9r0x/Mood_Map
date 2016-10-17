package api

type TypeEnum string

const (
	TextType TypeEnum = "text"
)

type Point struct {
	Content  string   `json:"content"`
	Location Location `json:"location"`
	Type     TypeEnum `json:"type"`
	Date     float32  `json:"date"`
	Source   string   `json:"source"`
}

type Location struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type CreatePointResponse struct {
	// Return the error message. This is used to provide user-facing errors.
	ErrorMessage string `json:"error_msg,omitempty"`
}

type PointInDB struct {
	TableName struct{} `sql:"point,alias:point"`
	ID        int
	X         float32
	Y         float32
	TagID     int32
	TimeStamp float32
}
