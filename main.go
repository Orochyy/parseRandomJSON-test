package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type NameJoke struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

func main() {
	fmt.Println(parseJokeNameStruct())
}

func NameAndSurname() []byte {
	url := "https://names.mcquay.me/api/v0/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func MakeUrl() string {
	var result User
	var FirstName, LastName string
	json.Unmarshal(NameAndSurname(), &result)

	FirstName = result.FirstName
	LastName = result.LastName

	nullUrl := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=" + FirstName + "&lastName=" + LastName + "&limitTo=nerdy")
	return nullUrl
}

func parseJoke() []byte {
	url := MakeUrl()
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func parseJokeNameStruct() string {
	var result NameJoke
	var lastJoke string
	json.Unmarshal(parseJoke(), &result)

	lastJoke = result.Value.Joke

	response := fmt.Sprintf(lastJoke)

	return response
}
