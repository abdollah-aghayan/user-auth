package httprouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HealthRequest Health request interface
type HealthRequest interface {
	getHealth(c *gin.Context)
}

//HealthRoute struct
type healthRoute struct {
	HealthRequest
}

//HealthRoutes health route variable
var healthRoutes = newHealthRoute()

func newHealthRoute() HealthRequest {
	return &healthRoute{}
}

func (h *healthRoute) getHealth(c *gin.Context) {
	res := map[string]string{
		"status": "pass",
	}
	c.JSON(http.StatusOK, res)
}
