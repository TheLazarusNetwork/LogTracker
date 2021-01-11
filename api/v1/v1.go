package apiv1

import (
	"github.com/TheLazarusNetwork/LogTracker/api/v1/log"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup, private bool) {
	v1 := r.Group("/v1")
	{
		if private {
			// Privately accessible APIs
		} else {
			// Publicly accessible APIs
			log.ApplyRoutes(v1)
		}
	}
}
