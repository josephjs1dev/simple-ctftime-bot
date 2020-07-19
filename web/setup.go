package web

import (
	applog "github.com/josephsalimin/simple-ctftime-bot/internal/pkg/log"
)

// InitAppLog initializes app logger
func InitAppLog() error {
	return applog.Init()
}

// SyncAppLog sync app log
func SyncAppLog() {
	defer applog.Sync()
}
