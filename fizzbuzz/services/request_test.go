package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequest(t *testing.T) {
	Populate()

	var int1 = 2
	var int2 = 3
	var limit = 30
	var str1 = "Cortex"
	var str2 = "IT"
	var url = fmt.Sprintf("/?int1=%d&int2=%d&limit=%d&str1=%s&str2=%s", int1, int2, limit, str1, str2)

	var contentRes = ""
	var fres = CommonFunc(int1, int2, limit, str1, str2)
	for i := 1; i <= len(fres); i++ {
		contentRes = fmt.Sprintf("%s%s%s", contentRes, fres[i-1], "\n")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Request)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned %v", status)
	}

	if rr.Body.String() != contentRes {
		t.Errorf("Wrong server response!")
	}

	//fmt.Println("Request Test OK")
}