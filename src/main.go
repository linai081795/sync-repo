package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"src/src/entities"
	"src/src/platform"
	"src/src/utils"
)

func main() {
	utils.InitRequest()

	repoFile, err := os.ReadFile("repo.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}
	var repos entities.Repo

	err = yaml.Unmarshal(repoFile, &repos)

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, i2 := range repos.Repos {
		commit := utils.GetCommit(i2)
		platform.GetSourceCommit(commit, i2)
	}
}
