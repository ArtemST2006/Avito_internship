package schemas

/*
🧱 1. /team/add — POST
Запрос (requestBody):
json
{
  "team_name": "string",
  "members": [
    {
      "user_id": "string",
      "username": "string",
      "is_active": true
    }
  ]
}
Ответ (201 Created):
json
{
  "team": {
    "team_name": "string",
    "members": [
      {
        "user_id": "string",
        "username": "string",
        "is_active": true
      }
    ]
  }
}
Ошибки:
400: {"error": {"code": "TEAM_EXISTS", "message": "..."}}
🧱 2. /team/get — GET
Параметр:
team_name (в query)
Ответ (200 OK):
json
{
  "team_name": "string",
  "members": [
    {
      "user_id": "string",
      "username": "string",
      "is_active": true
    }
  ]
}
Ошибки:
404: {"error": {"code": "NOT_FOUND", "message": "..."}}
🧱 3. /users/setIsActive — POST
Запрос:
json
{
  "user_id": "string",
  "is_active": true
}
Ответ (200 OK):
json

{
  "user": {
    "user_id": "string",
    "username": "string",
    "team_name": "string",
    "is_active": true
  }
}
Ошибки:
404: {"error": {"code": "NOT_FOUND", "message": "..."}}

🧱 4. /pullRequest/create — POST
Запрос:
json
{
  "pull_request_id": "string",
  "pull_request_name": "string",
  "author_id": "string"
}
Ответ (201 Created):
json
{
  "pr": {
    "pull_request_id": "string",
    "pull_request_name": "string",
    "author_id": "string",
    "status": "OPEN",
    "assigned_reviewers": ["string"],
    "createdAt": "2025-10-24T12:34:56Z"
  }
}
Ошибки:
404: {"error": {"code": "NOT_FOUND", "message": "..."}}
409: {"error": {"code": "PR_EXISTS", "message": "..."}}

🧱 5. /pullRequest/merge — POST
Запрос:
json
{
  "pull_request_id": "string"
}
Ответ (200 OK):
json
{
  "pr": {
    "pull_request_id": "string",
    "pull_request_name": "string",
    "author_id": "string",
    "status": "MERGED",
    "assigned_reviewers": ["string"],
    "mergedAt": "2025-10-24T12:34:56Z"
  }
}
Ошибки:
404: {"error": {"code": "NOT_FOUND", "message": "..."}}

🧱 6. /pullRequest/reassign — POST
Запрос:
json
{
  "pull_request_id": "string",
  "old_user_id": "string"
}
Ответ (200 OK):
json
{
  "pr": {
    "pull_request_id": "string",
    "pull_request_name": "string",
    "author_id": "string",
    "status": "OPEN",
    "assigned_reviewers": ["string"]
  },
  "replaced_by": "string"
}
Ошибки:
404: {"error": {"code": "NOT_FOUND", "message": "..."}}
409: {"error": {"code": "PR_MERGED", "message": "..."}}
409: {"error": {"code": "NOT_ASSIGNED", "message": "..."}}
409: {"error": {"code": "NO_CANDIDATE", "message": "..."}}

🧱 7. /users/getReview — GET
Параметр:
user_id (в query)
Ответ (200 OK):
json
{
  "user_id": "string",
  "pull_requests": [
    {
      "pull_request_id": "string",
      "pull_request_name": "string",
      "author_id": "string",
      "status": "OPEN"
    }
  ]
}
*/

// pullRequest/merge — POST
type PullRqMergeRequest struct {
	PullRequestId string `json:"pull_request_id" binding:"required"`
}

type PullRqResponse struct {
	Pr PullRequest `json:"pr"`
}

// pullRequest/create — POST
type CreatePullRequestRequest struct {
	PullRequestID   string `json:"pull_request_id" binding:"required"`
	PullRequestName string `json:"pull_request_name" binding:"required"`
	AuthorID        string `json:"author_id" binding:"required"`
}

// pullRequest/reassign — POST
type PRChangeAuthorRequest struct {
	PullRequestID string `json:"pull_request_id" binding:"required"`
	OldUserID     string `json:"old_user_id" binding:"required"`
}

type PRChangeAuthorResponse struct {
	Pr         PullRequest `json:"pr"`
	ReplacedBy string      `json:"replaced_by"`
}

// team/get — GET
type GetTeamResponse struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}

// team/add — GET
type CreateTeamRequest struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}

type Member struct {
	UserID   string `json:"user_id"`
	UserName string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type CreateTeamResponse struct {
	Team GetTeamResponse `json:"team"`
}

//users/setIsActive

type ActieveUserRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	IsActive bool   `json:"is_active"`
}

type ActiveUserResponse struct {
	User User `json:"user"`
}

// users/getRew
type GetUserPRResponse struct {
	UserID       string            `json:"user_id"`
	PullRequests []PullRequestInfo `json:"pull_requests"`
}

type PullRequestInfo struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
	Status          string `json:"status"`
}
