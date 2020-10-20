package http

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/domain"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/usecase"
)

func NewRouter(alarmController usecase.AlarmController, port int) {

	ADDR := ":"+strconv.Itoa(port)

	router := gin.Default()
	router.POST("/blinkt", func(c *gin.Context) {

		alarm := new(domain.Alarm)
		if err := c.BindJSON(alarm); err != nil {
			log.Warnf("could not bind JSON", err, c)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err := alarmController.TriggerAlarm(*alarm)
		if err != nil {
			log.Warnf("could not trigger Alarm",err)
			c.AbortWithStatus(http.StatusBadRequest)
		}
		c.String(200, "Trigger Alarm")
	})

	router.Run(ADDR)
}
