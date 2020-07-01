package repository

import (
	"github.com/ofonimefrancis/google-analytics/model"
	"github.com/ofonimefrancis/google-analytics/counter"
)
type Stats string

const (
	StatsLocationCountry Stats = "country"
	StatsLocationCity          = "city"
	StatsDeviceType            = "deviceType"
	StatsDevicePlatform        = "devicePlatform"
	StatsDeviceOS              = "os"
	StatsDeviceBrowser         = "browser"
	StatsDeviceLanguage        = "language"
	StatsReferral              = "referral"
)

// EventRepository is storage repository for Events
type EventRepository struct {
	locationCountry *counter.Counter
	locationCity    *counter.Counter
	deviceType      *counter.Counter
	devicePlatform  *counter.Counter
	deviceOS        *counter.Counter
	deviceBrowser   *counter.Counter
	deviceLanguage  *counter.Counter
	referral        *counter.Counter
}
