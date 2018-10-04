package api

import (
    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "app/models"
    "app/internal"
)

type LabelsController struct {
    DB *gorm.DB
}

func (ctrl LabelsController) GetLabels(w http.ResponseWriter, r *http.Request) {
   enableCors(&w)

   var respond internal.Response
   var labels []models.PullRequestLabel
   var label models.PullRequestLabel

   params := r.URL.Query()
   if labelID, ok := params["id"]; ok {
      ctrl.DB.Preload("PullRequest").Preload("Label").Preload("Activity").Where("id = ?", labelID).Order("updated_at desc").First(&label)
      respond.AsJson(w, label)
   } else {
      ctrl.DB.Preload("PullRequest").Preload("Label").Preload("Activity").Order("updated_at desc").Find(&labels)
      respond.AsJson(w, labels)
   }
}
