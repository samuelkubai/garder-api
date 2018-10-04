package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Activity struct {
    gorm.Model
    Type string `json:"type"`
    PullRequest PullRequest
    PullRequestID int
    Story Story
    StoryID int
    ActorID int
    ActorType string
}
