package main

import (
	"file-organizer/logger"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/**
* PURPOSE: To Organize immediate files by their type (image, document, video) etc.
* The file names will be read one by one and depending on their type/extension they will be moved to respective directory
 */

var log = logger.New(true)

const (
	Images    = "images"
	Audios    = "audios"
	Documents = "documents"
	Unknown   = "unknown"
)

var dirMap = map[string]bool{
	Images:    true,
	Audios:    true,
	Documents: true,
	Unknown:   true,
}

var extToDir = map[string]string{
	".png":  Images,
	".jpg":  Images,
	".jpeg": Images,

	".mp3": Audios,
	".wav": Audios,

	".txt":  Documents,
	".docx": Documents,
}

func getDirectoryForFile(ext string) string {
	if dir, ok := extToDir[ext]; ok {
		return dir
	}
	return Unknown
}

func ensureDir(name string) error {
	log.Debug("Ensuring directory %s", name)
	return os.MkdirAll(name, 0750)
}

func isCategoryDir(dirPath string) bool {
	log.Debug("Checking if %s is category directory", dirPath)
	dirName := filepath.Base(dirPath)
	_, ok := dirMap[dirName]
	return ok
}

func getDirEntries(dir string) ([]os.DirEntry, error) {
	return os.ReadDir(dir)
}

func createCategoryDirectories(dir string, dryRun bool) {
	log.Debug("Creating category directory")
	for _, knownDir := range []string{Images, Audios, Documents, Unknown} {
		if dryRun {
			log.Info("[DRY RUN] would ensure directory %s exists in %s", knownDir, dir)
			continue
		}
		dirPath := filepath.Join(dir, knownDir)

		if err := ensureDir(dirPath); err != nil {
			log.Error("%s", err)
		}

	}
}

func moveFilesToCategoryDirectories(srcPath string, dstPath string) {
	log.Info("Moving %s -> %s", srcPath, dstPath)
	if err := os.Rename(srcPath, dstPath); err != nil {
		log.Error("%s", err)
	}
}

func FileOrganizer(dirPath string, dryRun bool) {

	log.Info("Entered Directory: %s", dirPath)

	if isCategoryDir(dirPath) {
		log.Info("%s is a category directory", dirPath)
		return
	}

	entries, err := getDirEntries(dirPath)
	if err != nil {
		log.Error("unable to get directories", err)
		return
	}

	log.Info("Directory Entries: %v", entries)

	fmt.Println("===============================================================================================================")

	if dryRun {
		log.Info("[DRY RUN] entered directory named: %s\n", dirPath)
	}

	createCategoryDirectories(dirPath, dryRun)

	for _, entry := range entries {

		if dryRun {
			log.Info("[DRY RUN] %s has entry: %s\n", dirPath, entry.Name())
		}

		if entry.IsDir() {
			if !isCategoryDir(entry.Name()) {
				childPath := filepath.Join(dirPath, entry.Name())
				FileOrganizer(childPath, dryRun)

			}
			continue
		}

		fileName := entry.Name()
		ext := strings.ToLower(filepath.Ext(fileName))

		categoryDir := getDirectoryForFile(ext)

		srcPath := filepath.Join(dirPath, fileName)
		destPath := filepath.Join(dirPath, categoryDir, fileName)

		if dryRun {
			log.Info("[DRY RUN] would move %s -> %s \n", srcPath, destPath)
			continue
		}
		moveFilesToCategoryDirectories(srcPath, destPath)
	}
}

func main() {
	var folder string
	var dryRun bool
	flag.StringVar(&folder, "folder", "folder", "Directory to organize")
	flag.BoolVar(&dryRun, "dry-run", false, "simulate file moves")
	flag.Parse()

	if folder == "folder" {
		log.Info("Folder argument is required: go run ./main.go --folder <folder> name")
		os.Exit(1)
	}

	absPath, err := filepath.Abs(folder)
	if err != nil {
		log.Error("Failed to get absolute path:", err)
		return
	}
	FileOrganizer(absPath, dryRun)
}
