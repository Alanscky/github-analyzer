package handlers

import (
	conectionapigithub "github-analyzer/conectionAPIgithub"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGithubUser(c *gin.Context) {

	username :=
		c.Param("username")

	user := conectionapigithub.SearchUser(username)
	repos := conectionapigithub.RepositoryListDeclaration(user.ReposUrl)
	languageCounts := conectionapigithub.CountLanguages(repos)
	topRepo := conectionapigithub.TopRepository(repos)

	var topLanguage string
	var maxCount int

	for language, count := range languageCounts {

		if count > maxCount {
			maxCount = count
			topLanguage = language
		}
	}

	response := GithubAnalysisResponse{
		Name:           user.Name,
		Login:          user.Login,
		Repos:          user.PublicRepos,
		Followers:      user.Followers,
		Language:       topLanguage,
		PopularRepo:    topRepo.Name,
		Languages:      languageCounts,
		RepositoryList: repos,
	}

	c.JSON(http.StatusOK, response)
}

type GithubAnalysisResponse struct {
	Name           string                                `json:"name"`
	Login          string                                `json:"login"`
	Repos          int                                   `json:"repos"`
	Followers      int                                   `json:"followers"`
	Language       string                                `json:"language"`
	PopularRepo    string                                `json:"popularRepo"`
	Languages      map[string]int                        `json:"languages"`
	RepositoryList []conectionapigithub.GithubRepository `json:"repositoryList"`
}
