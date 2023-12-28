package utils

import (
	"encoding/json"
	"log"
	"os"
	"src/src/entities"
)

func GetCommit(repoInfo entities.RepoItem) entities.Commit {
	checkAndCreateDir(repoInfo.Platform, repoInfo.Source)
	file := getJsonFile(repoInfo.Platform, repoInfo.Source)

	var commit entities.Commit
	commit.RepoName = repoInfo.Source

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&commit)
	if err != nil {
		commit.Branchs
		return commit
	}

	return commit
}

func checkAndCreateDir(platform string, repoName string) {
	dirPath := "repo/" + platform + "/" + repoName

	_, err := os.Stat(dirPath)

	if err == nil {
		return
	}

	err = os.MkdirAll(dirPath, os.ModeDir)

	if err != nil {
		log.Fatal(err)
		return
	}

}

func getJsonFile(platform string, repoName string) *os.File {
	filePath := "repo/" + platform + "/" + repoName + "/commit.json"

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0660)

	if err == nil {
		return nil
	}

	return file
}
