package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpResponse struct {
	Fact map[string]int
}

var apiKey string = "" //X-Yandex-API-Key

func main() {

	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://api.weather.yandex.ru/v2/forecast?lat=55.75396&lon=37.620393&extra=false&limit=1&hours=false", nil,
	)
	req.Header.Add("X-Yandex-API-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var r httpResponse
	json.Unmarshal(body, &r)

	if r.Fact != nil {
		fmt.Printf("Температура в Москве %v ℃", r.Fact["temp"])
	} else {
		fmt.Print("Не удалось получить ответ от api, скорее всего не указан api key")
	}

}
