package main

import (
	"github.com/codeinuit/semantics-files-checker/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	handler.NewRouter(&handler.Handler{
		Log: logrus.New(),
	}).Run()
}
