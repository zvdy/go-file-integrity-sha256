package main

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	dirToCheck  = "sample_dir"
	jsonDBPath  = "file_integrity_db.json"
	logFilePath = "integrity_check.log"
)

// FileHash represents the structure to store file hashes
type FileHash struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
}

// calculateFileHash calculates the SHA-256 hash of a file
func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// loadJSONDB loads the JSON database from a file
func loadJSONDB() (map[string]string, error) {
	data, err := os.ReadFile(jsonDBPath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string), nil
		}
		return nil, err
	}

	var db map[string]string
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}

	return db, nil
}

// saveJSONDB saves the JSON database to a file
func saveJSONDB(db map[string]string) error {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(jsonDBPath, data, 0644)
}

// checkFileIntegrity checks the integrity of files in the directory
func checkFileIntegrity() {
	db, err := loadJSONDB()
	if err != nil {
		log.Fatalf("ERROR: Failed to load JSON database: %v", err)
	}

	files, err := os.ReadDir(dirToCheck)
	if err != nil {
		log.Fatalf("ERROR: Failed to read directory: %v", err)
	}

	currentFiles := make(map[string]bool)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(dirToCheck, file.Name())
		currentFiles[filePath] = true

		hash, err := calculateFileHash(filePath)
		if err != nil {
			log.Printf("ERROR: Failed to calculate hash for file %s: %v", filePath, err)
			continue
		}

		if storedHash, exists := db[filePath]; exists {
			if storedHash != hash {
				log.Printf("ALERT: File integrity check failed for %s", filePath)
			}
		} else {
			log.Printf("WARN: New file detected: %s", filePath)
		}

		db[filePath] = hash
	}

	for filePath := range db {
		if !currentFiles[filePath] {
			log.Printf("ALERT: File deleted: %s", filePath)
			delete(db, filePath)
		}
	}

	if err := saveJSONDB(db); err != nil {
		log.Fatalf("ERROR: Failed to save JSON database: %v", err)
	}
}

// logResults logs the results of the integrity check
func logResults(verbose bool) {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("ERROR: Failed to open log file: %v", err)
	}
	defer logFile.Close()

	if verbose {
		log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	} else {
		log.SetOutput(logFile)
	}

	checkFileIntegrity()
}

func main() {
	verbose := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()

	logResults(*verbose)
	log.Println("INFO: File integrity check completed successfully.")
}
