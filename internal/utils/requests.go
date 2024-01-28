package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type agifyResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type genderizeResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type country struct {
	CountryId   string  `json:"country_Id"`
	Probability float64 `json:"probability"`
}

type nationalizeResponse struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []country `json:"country"`
}

var client = http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	//req, err := http.NewRequest(http.MethodGet, url, nil)
	//if err != nil {
	//	return err
	//}
	//req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

func GetAge(name string) (int, error) {
	answer := agifyResponse{}
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	if err := getJson(url, &answer); err != nil {
		return 0, err
	}
	return answer.Age, nil
}

func GetGender(name string) (string, error) {
	answer := genderizeResponse{}
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	if err := getJson(url, &answer); err != nil {
		return "", err
	}
	return answer.Gender, nil
}

func GetNationality(name string) (string, error) {
	answer := nationalizeResponse{}
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	if err := getJson(url, &answer); err != nil {
		return "", err
	}
	if len(answer.Country) == 0 {
		return "", errors.New("not nationality in response")
	}
	return answer.Country[0].CountryId, nil
}
