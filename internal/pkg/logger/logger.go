package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

var plain *zap.Logger
var sugared *zap.SugaredLogger

// Init initialize log and sugaredLog
// For now, let's use zap Development mode
func Init() (err error) {
	option := zap.AddCallerSkip(1)
	if plain, err = zap.NewDevelopment(option); err != nil {
		return err
	}

	sugared = plain.Sugar()

	return nil
}

// Sync calls plain.Sync()
func Sync() error {
	if plain != nil {
		return plain.Sync()
	}

	return nil
}

// PlainInfo calls plain.Info if present, else calls log.Println
func PlainInfo(msg string, fields ...zap.Field) {
	if plain != nil {
		plain.Info(msg, fields...)
	} else {
		log.Println(msg)
	}
}

// Info calls sugared.Info if present, else calls log.Println
func Info(msg string, v ...interface{}) {
	if sugared != nil {
		sugared.Info(msg, v)
	} else {
		log.Println(fmt.Sprintf(msg, v...))
	}
}

// Infof calls sugared.Infof if present, else calls log.Println
func Infof(msg string, v ...interface{}) {
	if sugared != nil {
		sugared.Infof(msg, v)
	} else {
		log.Println(fmt.Sprintf(msg, v...))
	}
}

// Warnf calls sugared.Infof if present, else calls log.Println
func Warnf(msg string, v ...interface{}) {
	if sugared != nil {
		sugared.Warnf(msg, v)
	} else {
		log.Println(fmt.Sprintf(msg, v...))
	}
}

// Errorf calls sugared.Errorf if present, else calls log.Println
func Errorf(msg string, v ...interface{}) {
	if sugared != nil {
		sugared.Errorf(msg, v)
	} else {
		log.Println(fmt.Sprintf(msg, v...))
	}
}
