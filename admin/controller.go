package admin

import (
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"
  "net/http"

  "github.com/alexeyismirnov/go-vue-test/models"
)

type ApiController struct {
  DB *gorm.DB
}

// curl -H "Content-type: text" localhost:8080/api/v1/getMessage
func (api ApiController) GetMessage(c *gin.Context) {
  var message models.Message
  api.DB.First(&message)

  c.String(http.StatusOK, message.Content)
}

// curl -H "Content-type: text"  -X POST -d 'new message' localhost:8080/api/v1/setMessage
func (api ApiController) SetMessage(c *gin.Context) {
  str, _ := c.GetRawData()
  message := models.Message{ID: 1, Content: string(str)}
  api.DB.Save(&message)

  c.String(http.StatusOK, message.Content)
}
