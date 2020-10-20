package rpi

import (
	log "github.com/sirupsen/logrus"
	"time"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/domain"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/usecase"

	"github.com/richrarobi/blinkRpio"
)

type AlarmController struct {
}

func NewAlarmController() usecase.AlarmController {
	return &AlarmController{}
}

func (ac *AlarmController) TriggerAlarm(alarm domain.Alarm) error{
	log.Infof("Trigger Alarm", alarm)

	blinkRpio.Setup()
	blinkRpio.Clear()
	blinkRpio.Show()

	blinkRpio.SetAll(alarm.Red,alarm.Green,alarm.Blue,alarm.Lum)
	blinkRpio.Show()

	time.Sleep(time.Duration(alarm.DurationMs) * time.Millisecond)

	blinkRpio.Clear()
	blinkRpio.Show()

	log.Infof("Stopped ...")

	return nil
}
