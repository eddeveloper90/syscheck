package main

import (
	"xcheck/config"
)

func Version() string {
	return config.CONFIG.App.Version
}
