package entities

import "time"

type Commit struct {
	LatestSyncTime time.Time `json:"latestSyncTime"`
	RepoName       string    `json:"repoName"`
	Branchs        []struct {
		Name   string `json:"name"`
		Commit string `json:"commit"`
	} `json:"branchs"`
}
