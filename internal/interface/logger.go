package _interface

import "github.com/adolsalamanca/go-clean-boilerplate/pkg/logger"

type Logger interface {
	Debug(msg string, fields ...logger.Field)
	Info(msg string, fields ...logger.Field)
	Warn(msg string, fields ...logger.Field)
	Error(msg string, fields ...logger.Field)
}
