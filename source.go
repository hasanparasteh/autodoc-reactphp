package main

import "os"

type Project struct {
	rootPath    string
	totalRoutes uint32
}

func (project *Project) Default() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	project.rootPath = dir
	project.totalRoutes = 0
	return nil
}

func (project Project) HasSrcDir() (bool, error) {
	isPathExists, pathError := project.IsPathExists()
	if !isPathExists {
		return false, pathError
	}

	entries, err := os.ReadDir(project.rootPath)
	if err != nil {
		return false, err
	}

	var founded bool = false
	for _, entry := range entries {
		if entry.Name() == "src" {
			founded = true
			break
		}

	}

	return founded, nil
}

func (project *Project) IsPathExists() (bool, error) {
	stat, err := os.Stat(project.rootPath)
	if err != nil {
		return false, err
	}

	if stat.IsDir() {
		return true, nil
	}

	return false, nil
}
