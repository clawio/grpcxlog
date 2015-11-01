package grpcxlog

import (
	"github.com/rs/xlog"
)

type Log struct {
	xlog.Logger
}

func (l Log) Fatal(args ...interface{}) {
	l.Error(args)
}

func (l Log) Fatalf(format string, args ...interface{}) {
	l.Errorf(format, args)
}

func (l Log) Fatalln(args ...interface{}) {
	l.Fatal(args)
}

func (l Log) Print(args ...interface{}) {
	l.Error(args)
}

func (l Log) Printf(format string, args ...interface{}) {
	l.Errorf(format, args)
}

func (l Log) Println(args ...interface{}) {
	l.Print(args)
}
