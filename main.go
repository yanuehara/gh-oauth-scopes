// Copyright (c) 2023 @yanuehara (https://github.com/yanuehara)
package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"

	"github.com/cli/go-gh"
	"github.com/mgutz/ansi"
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

	oauthScopes := strings.Split(resp.Header.Get("X-Oauth-Scopes"), ",")
	greenText := ansi.ColorFunc("green")
	blueText := ansi.ColorFunc("blue+b")


	fmt.Printf("Acessing github as @%s with oauth scopes: \n", blueText(data.Login))
	for i := 0; i < len(oauthScopes); i++ {
		fmt.Printf("  * %s\n",
			greenText(strings.TrimLeft(oauthScopes[i], " ")),
		)
	}
}
