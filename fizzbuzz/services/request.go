package services

import (
	"fizzbuzz/redis"
	"fmt"
	"net/http"
	"strconv"
)

/*
Populate the redis server for unit tests
 */
func Populate(){
	var  key = "FIZZBUZZ|3-5-100-fizz-buzz"
	var value = 10
	redis.SaveData(key, value)

	key = "FIZZBUZZ|2-3-50-toto-tata"
	value = 8
	redis.SaveData(key, value)

	key = "FIZZBUZZ|5-7-20-cortex-IT"
	value = 4
	redis.SaveData(key, value)

	key = "FIZZBUZZ|6-9-20-IT-Cortex"
	value = 2
	redis.SaveData(key, value)
}

/*
Exposes a REST API endpoint that:
- Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1,
all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
 */
func Request(w http.ResponseWriter, r *http.Request) {

	var content string
	content = ""
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	/*w.Write([]byte(`{"message": "hello world"}`))*/

	int1s, okInt1 := r.URL.Query()["int1"]
	int2s, okInt2 := r.URL.Query()["int2"]
	limits, okLimit := r.URL.Query()["limit"]
	str1s, okStr1 := r.URL.Query()["str1"]
	str2s, okStr2 := r.URL.Query()["str2"]

	if (okInt1 && okInt2 && okLimit && okStr1 && okStr2){

		int1, err1 := strconv.Atoi(int1s[0])
		int2, err2 := strconv.Atoi(int2s[0])
		limit, err3 := strconv.Atoi(limits[0])

		if err1 == nil && err2 == nil && err3 == nil   {
			var str1 = str1s[0]
			var str2 = str2s[0]

			var result = CommonFunc(int1, int2, limit, str1, str2)
			for i := 1; i <= len(result); i++ {
				content = fmt.Sprintf("%s%s%s", content, result[i-1], "\n")
			}

			var key = fmt.Sprintf("FIZZBUZZ|%s-%s-%s-%s-%s", int1s[0], int2s[0], limits[0], str1, str2)
			var value, e = redis.GetData(key)
			if e != nil || value==-1{
				value = 0
			}
			value++
			redis.SaveData(key, value)

			w.Write([]byte(content))
		} else {
			w.Write([]byte("one or more of the parameters int1, int2, limit are not valid!"))
		}

	} else{
		w.Write([]byte("one or more of the parameters int1, int2, limit, str1, str2 are missing!"))
	}
}
