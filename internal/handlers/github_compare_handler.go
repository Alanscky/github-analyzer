package handlers

import (
	conectionapigithub "github-analyzer/conectionAPIgithub"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GithubCompareResponse struct {
	User1 GithubAnalysisResponse `json:"user1"`
	User2 GithubAnalysisResponse `json:"user2"`
}

func CompareGithubUsers(c *gin.Context) {

	username1 := c.Param("user1")
	username2 := c.Param("user2")

	user1 := buildGithubAnalysis(username1)
	user2 := buildGithubAnalysis(username2)

	response := GithubCompareResponse{
		User1: user1,
		User2: user2,
	}

	c.JSON(http.StatusOK, response)
}

func buildGithubAnalysis(username string) GithubAnalysisResponse {

	user :=
		conectionapigithub.SearchUser(username)

	repos :=
		conectionapigithub.RepositoryListDeclaration(
			user.ReposUrl,
		)

	languageCounts :=
		conectionapigithub.CountLanguages(repos)

	topRepo :=
		conectionapigithub.TopRepository(repos)

	var topLanguage string
	var maxCount int

	for language, count := range languageCounts {

		if count > maxCount {

			maxCount = count
			topLanguage = language
		}
	}

	return GithubAnalysisResponse{

		Name: user.Name,

		Login: user.Login,

		Repos: user.PublicRepos,

		Followers: user.Followers,

		Language: topLanguage,

		PopularRepo: topRepo.Name,

		Languages: languageCounts,

		RepositoryList: repos,
	}
}
