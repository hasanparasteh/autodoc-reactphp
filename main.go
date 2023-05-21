package main

import (
	"flag"
	"fmt"
)

func main() {
	var project Project
	var rootPath string

	flag.StringVar(&rootPath, "path", "", "root project path")
	flag.Parse()

	if rootPath == "" {
		fmt.Println("here")
		project = Project{}
		project.Default()
	} else {
		project = Project{}
		project.rootPath = rootPath
		project.totalRoutes = 0
	}

	var isSrcDetected, err = project.HasSrcDir()
	if !isSrcDetected {
		panic(err)
	}

}
