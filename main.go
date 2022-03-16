package reverseimageapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type ReverseImageApiAnswerResult struct {
	Url    string `json:"url"`
	Image  string `json:"image"`
	Height string `json:"height"`
	Width  string `json:"width"`
	Title  string `json:"title"`
}
type ReverseImageApiAnswer struct {
	CreditsRemaining int                           `json:"credits_remaining"`
	Count            int                           `json:"count"`
	Results          []ReverseImageApiAnswerResult `json:"results"`
}

func Search(apiKey string, imageUrl string) (ReverseImageApiAnswer, error) {
	urlForSearch := "https://reverseimageapi.com/api/search?key="
	urlForSearch += apiKey
	urlForSearch += "&url="
	urlForSearch += url.QueryEscape(imageUrl)
	
	response, err := http.Get(urlForSearch)
	if err != nil {
		return ReverseImageApiAnswer{},err
	}
	
	answer := ReverseImageApiAnswer{}
	err = json.NewDecoder(response.Body).Decode(&answer)
	if err != nil {
		return ReverseImageApiAnswer{},err
	}
	return answer,nil
}
