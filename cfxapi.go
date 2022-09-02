package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	EndPoint string `json:"EndPoint"`
}

func EndPoint(cfx string) string {
	resp, err := http.Get("https://servers-frontend.fivem.net/api/servers/single/" + cfx)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data Response
	json.Unmarshal(body, &data)
	return data.EndPoint
}
