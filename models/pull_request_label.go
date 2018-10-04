package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type PullRequestLabel struct {
    gorm.Model
    PullRequest *PullRequest
    PullRequestID int
    Label *Label
    LabelID int
    Activity []Activity `gorm:"polymorphic:Actor;"`
}

func (*PullRequestLabel) TableName() string {
  return "pr_label";
}
