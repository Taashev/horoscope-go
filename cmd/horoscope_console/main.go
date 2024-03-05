package run

import (
	getName "horoscope_console/internal/get_name"
	greeting "horoscope_console/internal/greeting"
)

func Run() {
	name := getName.GetName()

	greeting.PrintGreeting(name)
}
