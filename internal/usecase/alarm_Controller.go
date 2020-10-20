package usecase

import "vault.fritz.box/PI-ZERO-BLINKT-SERVER/internal/domain"

type AlarmController interface {
	TriggerAlarm(alarm domain.Alarm) error
}
