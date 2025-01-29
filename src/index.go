package core

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const APP_NAME = "simple-pre-commit"

func CheckConfig() error {
	file, err := os.Open("Makefile")
	if err != nil {
		return err
	}
	defer file.Close()
	preCommitReg := regexp.MustCompile(`^pre-commit\s*:`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if preCommitReg.MatchString(line) {
			return nil
		}
	}
	msg := fmt.Sprintf("%s: pre-commit config not found in Makefile", APP_NAME)
	return errors.New(msg)
}

func GetGitProjectRoot() (string, error) {
	startDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	parentDir := ""
	for {
		gitDir := filepath.Join(startDir, ".git")
		if info, err := os.Stat(gitDir); err == nil && info.IsDir() {
			return startDir, nil
		}
		parentDir = filepath.Dir(startDir)
		if startDir == parentDir {
			msg := fmt.Sprintf("%s: .git folder not found", APP_NAME)
			return "", errors.New(msg)
		}
		startDir = parentDir
	}
}

func SetHook(gitRoot string) error {
	hookFile := filepath.Join(gitRoot, ".git", "hooks", "pre-commit")
	command := `make pre-commit`
	file, err := os.Create(hookFile)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(command)
	if err != nil {
		return err
	}
	err = os.Chmod(hookFile, 0755)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("%s: Successfully set git hooks", APP_NAME)
	fmt.Println(msg)
	return nil
}
