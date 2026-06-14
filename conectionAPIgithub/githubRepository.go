package conectionapigithub

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type GithubRepository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Size        int    `json:"size"`

	Stargazers int `json:"stargazers_count"`
	Watchers   int `json:"watchers"`

	Language string `json:"language"`
	Forks    int    `json:"forks"`

	Creatated_at string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}

func RepositoryListDeclaration(reposUrl string) []GithubRepository {
	resp, err := http.Get(reposUrl)
	if err != nil {
		fmt.Println("Error while doing petition: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var repositoryList []GithubRepository

	err = json.NewDecoder(resp.Body).Decode(&repositoryList)
	if err != nil {
		log.Fatal("Error while decoding json: ", err)
	}

	return repositoryList
}

func RepositoryDescription(repo GithubRepository) {
	fmt.Printf("\nDescription: %s\nSize: %d\nStargazers: %d\nWatchers: %d\nLanguage: %s\nForks: %d\nCreated at: %s\nUpdated at: %s\n",
		repo.Description, repo.Size, repo.Stargazers, repo.Watchers, repo.Language, repo.Forks, repo.Creatated_at, repo.Updated_at)
}
