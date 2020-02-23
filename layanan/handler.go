package layanan

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Handler for layanan handler
type Handler struct {
	Handler Object
}

// ProvideLayananAPI for layanan handler
func ProvideLayananAPI(db *gorm.DB) Handler {
	model := ProvideObject(db)
	return Handler{Handler: model}
}

// Add layanan
func (h Handler) Add() gin.HandlerFunc {
	var layanan LayananMapper
	return func(c *gin.Context) {
		err := c.BindJSON(&layanan)
		if err != nil {
			log.Panicln("layanan handler.go: ", err)
			c.Status(http.StatusBadRequest)
			return
		}

		newLayanan := h.Handler.Save(toTable(layanan))
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "status": "success", "data": toResponse(newLayanan)})
	}
}

// Show all layanans
func (h Handler) Show() gin.HandlerFunc {

	return func(c *gin.Context) {
		layanans := h.Handler.Show()
		if len(layanans) <= 0 {
			log.Println("layanan handler.go: data is empty")
			c.Status(http.StatusOK)
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "status": "success", "data": toMultipleResponse(layanans)})
	}
}
