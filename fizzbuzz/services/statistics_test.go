package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatistics(t *testing.T) {
	Populate()

	req, err := http.NewRequest("GET", "/statistics", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Statistics)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned %v", status)
	}

	fmt.Println("Response Body : ", rr.Body.String())

	if rr.Body.String() != "{'int1': '3', 'int2': '5', 'limit': '100', 'str1': 'fizz', 'str2': 'buzz', 'hits': '10'}" {
		t.Errorf("Wrong server response!")
	}

	//fmt.Println("Statistics Test OK")
}