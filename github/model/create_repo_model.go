package model

type CreateRepoModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProject  bool   `json:"has_project"`
	HasWiki     bool   `json:"has_wiki"`
}
