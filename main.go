package main

import (
	"github.com/codeinuit/semantics-files-checker/internal/handler"
)

func main() {
	handler.NewRouter().Run()
}
