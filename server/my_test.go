package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Weather
	expectedError error
}

func TestGetWeather(t *testing.T) {
	var tests = []Tests{
		{
			name: "basic-request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"city": "Denve, Co", "forecast": "sunny"}`))
			})),
			response: &Weather{
				City:     "Denve, Co",
				Forecast: "sunny",
			},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()
			resp, err := GetWeather(test.server.URL) //server jis ip p start hua h vhi pass kr rha h
			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("failed:expected %v, got %v\n", test.response, resp)
			}
			if !errors.Is(err, test.expectedError) {
				t.Errorf("expected error failed. expected %v, got %v\n", err, test.expectedError)
			}

		})
	}
}
