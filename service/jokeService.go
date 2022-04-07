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
}

func ParseJokeNameStruct() string {
	var result dto.NameJoke
	var lastJoke string
	json.Unmarshal(parseJoke(), &result)

	lastJoke = result.Value.Joke

	response := fmt.Sprintf(lastJoke)

	return response
}

func parseJoke() []byte {
	url := makeUrl()
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("Second %d API broken", url)
	}

	return body
}

func makeUrl() string {
	var result dto.User
	var FirstName, LastName string
	json.Unmarshal(nameAndSurname(), &result)

	FirstName = result.FirstName
	LastName = result.LastName

	Url := "http://api.icndb.com/jokes/random?firstName=" + FirstName + "&lastName=" + LastName + "&limitTo=nerdy"
	return Url
}
func nameAndSurname() []byte {
	url := "https://names.mcquay.me/api/v0/"

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("problems with %d API", url)
	}
	return body
}
