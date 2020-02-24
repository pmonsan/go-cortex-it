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

  redis<br> 
  &emsp;&emsp;redisManager.go<br>
  services<br> 
  &emsp;&emsp;common.go<br>
  &emsp;&emsp;request.go<br>
  &emsp;&emsp;request_test.go<br>
  &emsp;&emsp;statistics.go<br>
  &emsp;&emsp;statistics_test.go<br>
  main<br>
  &emsp;&emsp;main.go<br>

- redisManager.go :<br>
  1. func SaveData(key string, value int ) error : store requests hits in redis server<br>
  2. func GetData(key string) (int, error) : Gets request hits from redis server<br>
  3. func GetMostFrequentRequest() (string, int) : get the most frequent request : the most used request (having the maximum of hits in redis server)<br>

  &emsp;&emsp;the format of redis keys is : "FIZZBUZZ|int1-int2-limit-str1-str2"<br>

- common.go :<br>
func CommonFunc(int1, int2, limit int, str1, str2 string) []string :<br>
  1. Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.<br>
  2. Returns an array of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1,
all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.<br>

- request.go :<br> 
func Request(w http.ResponseWriter, r *http.Request) :
Exposes a REST API endpoint that:<br>
  1. Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.<br>
  2. Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1,
all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.<br>

- request_test.go :<br>
implementation of an unit test for request.go<br>

- statistics.go :<br> 
func Statistics(w http.ResponseWriter, r *http.Request) :<br>
  1. Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:<br>
  2. Accept no parameter<br>
  3. Return the parameters corresponding to the most used request, as well as the number of hits for this request<br>

  &emsp;&emsp;the format of return is : {'int1': '{int1}', 'int2': '{int2}', 'limit': '{limit}', 'str1': '{str1}', 'str2': '{str2}', 'hits': '{hits}'}<br><br>

- statistics_test.go :<br>
implementation of an unit test for statistics.go  


USAGE
---------------------------

install a redis server with default port 6379<br>

to install dependencies :<br>
go get github.com/go-redis/redis<br>
go get github.com/gorilla/mux<br>

to run : <br>
go run main.go<br>

to execute the tests :<br>
go test fizzbuzz/services

  
URL
---------------------------
http://localhost:8080/?int1=3&int2=5&limit=100&str1=fizz&str2=buzz

http://localhost:8080/statistics
	