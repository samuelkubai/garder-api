package api

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/models"
    "app/internal"
)

type PullRequestsController struct {
    DB *gorm.DB
}

func (ctrl PullRequestsController) GetPullRequests(w http.ResponseWriter, r *http.Request) {
   enableCors(&w)

   var respond internal.Response
   var pullRequests []models.PullRequest
   var pullRequest models.PullRequest

   params := r.URL.Query()
   if storyID, ok := params["storyID"]; ok {
      ctrl.DB.Preload("Activities", func(db *gorm.DB) *gorm.DB {
        return db.Order("activities.updated_at DESC")
      }).Where("story_id = ?", storyID).Find(&pullRequests)
      respond.AsJson(w, pullRequests)
   } else if pullRequestID, ok := params["id"]; ok {
      ctrl.DB.Preload("Activities", func(db *gorm.DB) *gorm.DB {
        return db.Order("activities.updated_at DESC")
      }).Where("id = ?", pullRequestID).First(&pullRequest)
      respond.AsJson(w, pullRequest)
   } else {
      ctrl.DB.Preload("Activities", func(db *gorm.DB) *gorm.DB {
        return db.Order("activities.updated_at DESC")
      }).Find(&pullRequests)
      respond.AsJson(w, pullRequests)
   }
}
