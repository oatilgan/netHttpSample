package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"netHttpSample/model"
)

func main() {

	requestBody, _ := json.Marshal(map[string]string{
		"user":     "userName",
		"password": "password",
	})
	resp, err := http.Post("tokenUrl", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var prohandelToken model.ProhandelToken

	json.Unmarshal([]byte(body), &prohandelToken)

	client := &http.Client{}

	req, _ := http.NewRequest("GET", "getArticleSizeNumberUrl", nil)

	req.Header.Add("Authorization", "Bearer "+prohandelToken.Access_Token)

	res, errClinet := client.Do(req)

	if errClinet != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, errr := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println(string(body))

	// fmt.Printf("Access Token: %s\nRefresh Token: %s\nExpiresIn: %d", prohandelToken.Access_Token, prohandelToken.Refresh_Token, prohandelToken.Expires_In)

	// log.Println(string(body))
}
