package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const lockFileName = ".internet-outages-monitor.lock"

var lockFilePath string

// Init ...
func Init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Failed to initialize storage due to error getting user home dir: %s\n", err)
		panic(err)
	}
	lockFilePath = filepath.Join(homeDir, lockFileName)
	_, err = os.Stat(lockFilePath)
	if err != nil {
		createLockFile()
		removeLockFile()
	}
}

// AquireLock ...
func AquireLock() {
	_, err := os.Stat(lockFilePath)
	if err != nil {
		createLockFile()
		return
	}
}

// ReleaseLock ...
func ReleaseLock() time.Time {
	f, err := os.Stat(lockFilePath)
	if err != nil {
		panic(fmt.Errorf("Failed to release lock due to stat lock file failed: %s", err))
	}
	modTime := f.ModTime()
	err = os.Remove(lockFilePath)
	if err != nil {
		panic(fmt.Errorf("Failed to release lock due to error in removing lock file: %s", err))
	}
	return modTime
}

func createLockFile() {
	_, err := os.Create(lockFilePath)
	if err != nil {
		log.Println(fmt.Errorf("Failed to create lock file due to %s", err))
		panic(err)
	}
}

func removeLockFile() {
	err := os.Remove(lockFilePath)
	if err != nil {
		log.Println(fmt.Errorf("Failed to remove lock file due to %s", err))
		panic(err)
	}
}
