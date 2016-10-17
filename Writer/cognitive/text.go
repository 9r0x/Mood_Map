package cognitive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gaocegege/hackys-backend-writer/pkg/log"
)

const (
	URL = "https://westus.api.cognitive.microsoft.com/text/analytics/v2.0/sentiment"
)

func RecognizeText(text string) (*TextResult, error) {
	textA, err := New("ea376e2dcce14fd9aa8f0b8161f97a7e")
	if err != nil {
		panic(err)
	}

	return textA.Recognize(text)
}

func New(key string) (*TextAnalytics, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid Key")
	}
	return &TextAnalytics{
		BingKey: key,
	}, nil
}

func (emotion *TextAnalytics) Recognize(text string) (*TextResult, error) {
	apiURL := URL
	textRequest := TextRequest{
		Documents: make([]TextReq, 0),
	}
	textRequest.Documents = append(textRequest.Documents, TextReq{
		Language: "en",
		ID:       "1",
		Text:     text,
	})
	buf, err := json.Marshal(textRequest)
	if err != nil {
		log.Warnf("err: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(buf))
	if err != nil {
		log.Warnf("err: %v", err)
		return nil, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", emotion.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Warnf("err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		result := TextResult{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Warnf("err: %v", err)
			return nil, err
		}

		return &result, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 429 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Warnf("err: %v", err)
			return nil, err
		}

		return nil, fmt.Errorf(result.Code)
	}

	return nil, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}
