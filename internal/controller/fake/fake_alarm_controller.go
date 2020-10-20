package fake

import (
	"fmt"
	"time"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/domain"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/usecase"
)

type AlarmController struct {
}

func NewAlarmController() usecase.AlarmController{
	return &AlarmController{}
}
//TriggerAlarm will simulate the duration and sleep for the time.
func (ac *AlarmController) TriggerAlarm(alarm domain.Alarm) error{
	fmt.Printf("Trigger the Color Red: %d, Green %d and Blue %d for %d MilliSeconds\n", alarm.Red, alarm.Green, alarm.Blue, alarm.DurationMs)
	time.Sleep(time.Millisecond * time.Duration(alarm.DurationMs))
	return nil
}