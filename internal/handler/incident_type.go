package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariaNomidi/eco-app/internal/service"
)

type IncidentTypeHandler struct {
	service *service.IncidentTypeService
}

func NewIncidentTypeHandler(s *service.IncidentTypeService) *IncidentTypeHandler {
	return &IncidentTypeHandler{service: s}
}

func (h *IncidentTypeHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.service.GetAll(ctx)
	if err != nil {
		log.Printf("GetAll IncidentType failed: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch incident types",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
