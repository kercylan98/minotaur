package activity

type activityInterface interface {
	stopTicker()
	refreshTicker(init bool)
}
