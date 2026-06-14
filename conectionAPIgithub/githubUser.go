package conectionapigithub

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type githubUser struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Location    string `json:"location"`
	ReposUrl    string `json:"repos_url"`
}

func SearchUser(name string) githubUser {
	url := "https://api.github.com/users/" + name

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while doing petition: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var userGithub githubUser

	err = json.NewDecoder(resp.Body).Decode(&userGithub)
	if err != nil {
		log.Fatal("Error while decoding json: ", err)
	}

	return userGithub
}

func DescriptionUser(user githubUser) {
	fmt.Printf("\nname: %s\nlogin: %s\npublic repos: %d\nfollowers: %d\nlocation: %s\nrepos-url: %s\n",
		user.Name, user.Login, user.PublicRepos, user.Followers, user.Location, user.ReposUrl)
}
