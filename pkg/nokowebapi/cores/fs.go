package cores

import (
	"os"
	"path/filepath"
)

func GetSourceRootDir() (string, error) {
	var err error
	var sourceRootPath string
	KeepVoid(err, sourceRootPath)

	if sourceRootPath, err = filepath.Abs(os.Args[0]); err != nil {
		return "", err
	}
	sourceRootDir := filepath.Dir(sourceRootPath)
	return sourceRootDir, nil
}

func GetCurrentWorkingDir() (string, error) {
	var err error
	var currentWorkingDir string
	if currentWorkingDir, err = os.Getwd(); err != nil {
		return "", err
	}
	return currentWorkingDir, nil
}

type WorkingDirImpl interface {
	GetSourceRootDir() string
	GetCurrentWorkingDir() string
}

type WorkingDir struct {
	sourceRootDir     string
	currentWorkingDir string
}

func NewWorkingDir(scriptRootDir, currentWorkingDir string) WorkingDirImpl {
	return &WorkingDir{scriptRootDir, currentWorkingDir}
}

func (s *WorkingDir) GetSourceRootDir() string {
	return s.sourceRootDir
}

func (s *WorkingDir) GetCurrentWorkingDir() string {
	return s.currentWorkingDir
}

type WorkingFunc func(WorkingDirImpl)

func (w WorkingFunc) Call(workingDir WorkingDirImpl) {
	w(workingDir)
}

func SetWorkingDir(cb WorkingFunc) (WorkingDirImpl, error) {
	var err error
	var sourceRootDir string
	var currWorkingDir string
	KeepVoid(err, sourceRootDir, currWorkingDir)

	if sourceRootDir, err = GetSourceRootDir(); err != nil {
		return nil, err
	}

	if currWorkingDir, err = GetCurrentWorkingDir(); err != nil {
		return nil, err
	}

	workingDir := NewWorkingDir(sourceRootDir, currWorkingDir)
	if err = os.Chdir(sourceRootDir); err != nil {
		return nil, err
	}

	cb.Call(workingDir)

	if err = os.Chdir(currWorkingDir); err != nil {
		return nil, err
	}
	return workingDir, nil
}
