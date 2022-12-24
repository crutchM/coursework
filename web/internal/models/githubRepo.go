package models

type GithubRepository struct {
	Id            int64  `json:"id" db:"id"`
	FullName      string `json:"full_name" db:"full_name"`
	IsPrivate     bool   `json:"private" db:"is_private"`
	RepositoryUrl string `json:"html_url" db:"url"`
	Description   string `json:"description" db:"descr"`
	CanFork       bool   `json:"fork" db:"can_fork"`
	CreatedAt     string `json:"created_at" db:"created"`
	UpdatedAt     string `json:"updated_at" db:"updated"`
	PushedAt      string `json:"pushed_at" db:"pushed"`
	Size          int    `json:"size" db:"size"`
	Language      string `json:"language" db:"language"`
	Forks         int    `json:"forks" db:"forks"`
	Issues        int    `json:"open_issues" db:"issues"`
	Watchers      int    `json:"watchers" db:"watchers"`
	Subscribers   int    `json:"subscribers_count" db:"subscribers"`
}
