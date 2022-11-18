package server

import (
	"encoding/json"
	"net/http"
)

type Weather struct {
	City     string `json:"city"`
	Forecast string `json:"forecast"`
}

func GetWeather(url string) (*Weather, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//Defer statement is used to execute a function call just before the surrounding function
	//where the defer statement is present returns.
	defer resp.Body.Close()
	var weather Weather
	//func NewDecoder(r io.Reader) *Decoder is a function defined in the encoding/json package
	//which reads a JSON input stream, stores it in the buffer of a Decoder object,
	//and returns the Decoder object.
	//	the Decoder object can then be used to perform methods such as Decode(), More(), InputOutputOffset(), Buffered().
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, err
	}
	return &weather, nil

}
