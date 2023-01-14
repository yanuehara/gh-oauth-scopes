// Copyright (c) 2023 @yanuehara (https://github.com/yanuehara)
package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"

	"github.com/cli/go-gh"
)

func main() {
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Request("GET", "https://api.github.com/user", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        fmt.Println(err)
		return
    }

	type User struct{
		Login string `json:login`
	}
	var data User

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Acessing github as @%s with oauth scopes %s", data.Login, resp.Header.Get("X-Oauth-Scopes"))
}
