package conectionapigithub

import "fmt"

func TopRepository(repositoryList []GithubRepository) GithubRepository {
	var topRepository = repositoryList[0]
	for counter := 1; counter < len(repositoryList); counter++ {
		if repositoryList[counter].Stargazers > topRepository.Stargazers {
			topRepository = repositoryList[counter]
		}
	}
	return topRepository
}

func CountLanguages(repos []GithubRepository) map[string]int {

	counts := make(map[string]int)

	for _, repo := range repos {
		if repo.Language == "" {
			continue
		}
		counts[repo.Language]++
	}

	return counts
}

func RepositoryDescriptionList(repo []GithubRepository) {

	for c := 0; c < len(repo); c++ {
		fmt.Printf("\nRepo number: %d\nDescription: %s\nSize: %d\nStargazers: %d\nWatchers: %d\nLanguage: %s\nForks: %d\nCreated at: %s\nUpdated at: %s\n",
			c+1, repo[c].Description, repo[c].Size, repo[c].Stargazers, repo[c].Watchers, repo[c].Language, repo[c].Forks, repo[c].Creatated_at, repo[c].Updated_at)
	}
}
