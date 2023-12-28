package entities

type Repo struct {
	Repos []RepoItem `yaml:"repos" json:"repos"`
}
type RepoItem struct {
	Source   string   `yaml:"source" json:"source"`
	Platform string   `yaml:"platform" json:"platform"`
	Branch   []string `yaml:"branch" json:"branch"`
	Target   string   `yaml:"target" json:"target"`
}
