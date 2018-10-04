package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Label struct {
    gorm.Model
    Name string `json:"name"`
    Color string `json:"color"`
    PullRequestsLabels []*PullRequestLabel `json:"pullRequests,omitempty"`
}
