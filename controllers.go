package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

//SonarrEventsController ...
type SonarrEventsController struct{}

//Register ...
func (ctrl SonarrEventsController) Register(c *gin.Context) {
	var event SonarrEvent

	if err := c.BindJSON(&event); err != nil {
		c.JSON(406, gin.H{"message": "Invalid event", "form": event})
		log.Println("Bind failed SEC.Create", err)
		log.Println(event)
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Event received"})

	RegisterSonarrEvent(event)
}
