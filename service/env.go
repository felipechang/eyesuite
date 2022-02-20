package service

import (
	"log"
	"os"
)

type Constants struct {
	AccessSecret  string
	RefreshSecret string
}

func GetEnv() *Constants {
	cts := &Constants{
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}
	if cts.AccessSecret == "" || cts.RefreshSecret == "" {
		log.Panicln("jwt env missing")
	}
	return cts
}
