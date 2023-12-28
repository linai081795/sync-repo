package platform

import (
	"fmt"
	"os"
	"src/src/entities"
	"src/src/utils"
)

type GiteeCommit struct {
	Sha string `json:"sha"`
}

var request = utils.Request()
var BASE_URL = "https://gitee.com/api/v5/"
var GITEE_ACCOUNT = os.Getenv("secrets.GITEE_ACCOUNT")
var GITEE_PASSWORD = os.Getenv("secrets.GITEE_PASSWORD")
var GITEE_TOKEN = os.Getenv("secrets.GITEE_TOKEN")

func GetSourceCommit(commit entities.Commit, repo entities.RepoItem) {

	var url = BASE_URL + "repos/" + GITEE_ACCOUNT + "/" + repo.Source + "/commits"

	for _, branch := range repo.Branch {
		var params = map[string]string{
			"access_token": GITEE_TOKEN,
			"sha":          branch,
		}
		var result []GiteeCommit

		_, _ = request.SetResult(&result).SetQueryParams(params).Get(url)

		if result != nil {
			fmt.Println("Sha", result[0].Sha)
		}
	}

	return
}
