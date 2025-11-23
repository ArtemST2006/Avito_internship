package schemas

// PullRqMergeRequest содержит запрос на объединение pull request
type PullRqMergeRequest struct {
	PullRequestID string `json:"pull_request_id" binding:"required"`
}

// PullRqResponse содержит ответ с данными pull request
type PullRqResponse struct {
	Pr PullRequest `json:"pr"`
}

// CreatePullRequestRequest содержит запрос на создание нового pull request
type CreatePullRequestRequest struct {
	PullRequestID   string `json:"pull_request_id" binding:"required"`
	PullRequestName string `json:"pull_request_name" binding:"required"`
	AuthorID        string `json:"author_id" binding:"required"`
}

// PRChangeAuthorRequest содержит запрос на смену автора pull request
type PRChangeAuthorRequest struct {
	PullRequestID string `json:"pull_request_id" binding:"required"`
	OldUserID     string `json:"old_user_id" binding:"required"`
}

// PRChangeAuthorResponse содержит ответ на запрос смены автора pull request
type PRChangeAuthorResponse struct {
	Pr         PullRequest `json:"pr"`
	ReplacedBy string      `json:"replaced_by"`
}

// GetTeamResponse содержит ответ с информацией о команде
type GetTeamResponse struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}

// CreateTeamRequest содержит запрос на создание новой команды
type CreateTeamRequest struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}

// Member представляет участника команды
type Member struct {
	UserID   string `json:"user_id"`
	UserName string `json:"username"`
	IsActive bool   `json:"is_active"`
}

// CreateTeamResponse содержит ответ на запрос создания команды
type CreateTeamResponse struct {
	Team GetTeamResponse `json:"team"`
}

// ActieveUserRequest содержит запрос на изменение активности пользователя
type ActieveUserRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	IsActive bool   `json:"is_active"`
}

// ActiveUserResponse содержит ответ на запрос изменения активности пользователя
type ActiveUserResponse struct {
	User User `json:"user"`
}

// GetUserPRResponse содержит информацию о pull requests пользователя
type GetUserPRResponse struct {
	UserID       string            `json:"user_id"`
	PullRequests []PullRequestInfo `json:"pull_requests"`
}

// PullRequestInfo содержит основную информацию о pull request
type PullRequestInfo struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
	Status          string `json:"status"`
}

// StatisticResponse содержит статистику по pull requests
type StatisticResponse struct {
	NumberOpen   int                `json:"number_open"`
	NumberMerged int                `json:"number_merged"`
	PRUser       []StatisticForUser `json:"pull_request_statistic"`
}

// StatisticForUser содержит статистику по pull requests для конкретного пользователя
type StatisticForUser struct {
	UserID     string `json:"user_id"`
	NumberOfPR int    `json:"number_pull_request"`
}
