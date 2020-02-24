package services

import (
	"fizzbuzz/redis"
	"net/http"
	"strconv"
	"strings"
)

/*
- Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request
 */
func Statistics(w http.ResponseWriter, r *http.Request) {
	var result = "{'int1': '{int1}', 'int2': '{int2}', 'limit': '{limit}', 'str1': '{str1}', 'str2': '{str2}', 'hits': '{hits}'}"
	var key, value = redis.GetMostFrequentRequest();
	result = strings.ReplaceAll(result, "{hits}", strconv.Itoa(value))

	key = strings.ReplaceAll(key, "FIZZBUZZ|", "")
	var elements = strings.Split(key,"-")
	result = strings.ReplaceAll(result, "{int1}", elements[0])
	result = strings.ReplaceAll(result, "{int2}", elements[1])
	result = strings.ReplaceAll(result, "{limit}", elements[2])
	result = strings.ReplaceAll(result, "{str1}", elements[3])
	result = strings.ReplaceAll(result, "{str2}", elements[4])

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result))
}

