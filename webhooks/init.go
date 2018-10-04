package webhooks

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/models"
)

func RegisterWebhooks(mux *http.ServeMux, db *gorm.DB) {
    pt := PivotalTracker{ DB: db }
    gh := Github{ DB: db }
    // Register all the various webhooks here.
    mux.HandleFunc("/webhooks/pt/activity", pt.HandleActivity);
    mux.HandleFunc("/webhooks/gh/pull-requests", gh.HandlePullRequestsActivity)
}

func LogStoryActivity(story *models.Story, activity_type string, db *gorm.DB) {
  db.Model(story).Association("Activity").Append(&models.Activity{Type: activity_type, StoryID: int(story.ID)})
}

func LogCommentActivity(comment *models.Comment, activity_type string, db *gorm.DB) {
  db.Model(comment).Association("Activity").Append(&models.Activity{Type: activity_type, StoryID: comment.StoryID, PullRequestID: comment.PullRequestID})
}

func LogPullRequestActivity(pullRequest *models.PullRequest, activity_type string, db *gorm.DB) {
  db.Model(pullRequest).Association("Activity").Append(&models.Activity{Type: activity_type, StoryID: pullRequest.StoryID, PullRequestID: int(pullRequest.ID)})
}

func LogLabelsActivity(pullRequest *models.PullRequest, activity_type string, db *gorm.DB) {
  for _, label := range pullRequest.PullRequestLabels {
    db.Model(label).Association("Activity").Append(&models.Activity{Type: activity_type, StoryID: pullRequest.StoryID, PullRequestID: int(pullRequest.ID)})
  }
}
