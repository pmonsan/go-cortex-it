The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

Bonus question :
- Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request


Implementation:
-----------------
- three packages containing :<br>

  redis 
    redisManager.go
  services 
    common.go
	request.go
	request_test.go
	statistics.go
	statistics_test.go
  main
    main.go

- redisManager.go :<br>
func SaveData(key string, value int ) error : store requests hits in redis server
func GetData(key string) (int, error) : Gets request hits from redis server
func GetMostFrequentRequest() (string, int) : get the most frequent request : the most used request (having the maximum of hits in redis server)

the format of redis keys is : "FIZZBUZZ|int1-int2-limit-str1-str2"

- common.go :<br>
func CommonFunc(int1, int2, limit int, str1, str2 string) []string :<br>
  1.Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
  2.Returns an array of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1,
all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

- request.go :<br> 
func Request(w http.ResponseWriter, r *http.Request) :
Exposes a REST API endpoint that:
- Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1,
all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

- request_test.go :
implementation of an unit test for request.go



- statistics.go :<br> 
func Statistics(w http.ResponseWriter, r *http.Request) :
- Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request

the format of return is : {'int1': '{int1}', 'int2': '{int2}', 'limit': '{limit}', 'str1': '{str1}', 'str2': '{str2}', 'hits': '{hits}'}

- statistics_test.go :<br>
implementation of an unit test for statistics.go  


USAGE
---------------------------

to install dependencies :
go get github.com/go-redis/redis
go get github.com/gorilla/mux

to run : 
go run main.go

to execute the tests :
go test fizzbuzz/services

  
URL
---------------------------
http://localhost:8080/?int1=3&int2=5&limit=100&str1=fizz&str2=buzz

http://localhost:8080/statistics
	