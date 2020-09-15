package model

type CreateRepoResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login     string `json:"login"`
		AvatarUrl string `json:"avatar_url"`
	} `json:"owner"`
}
