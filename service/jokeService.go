package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parseRandomJSON-test/dto"
)

type JokeService interface {
	ParseJokeNameStruct() string
	ParseJoke() []byte
	MakeUrl() string
	NameAndSurname() []byte
}

func ParseJokeNameStruct() string {
	var result dto.NameJoke
	var lastJoke string
	json.Unmarshal(ParseJoke(), &result)

	lastJoke = result.Value.Joke

	response := fmt.Sprintf(lastJoke)

	return response
}

func ParseJoke() []byte {
	url := MakeUrl()
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func MakeUrl() string {
	var result dto.User
	var FirstName, LastName string
	json.Unmarshal(NameAndSurname(), &result)

	FirstName = result.FirstName
	LastName = result.LastName

	Url := "http://api.icndb.com/jokes/random?firstName=" + FirstName + "&lastName=" + LastName + "&limitTo=nerdy"
	return Url
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
