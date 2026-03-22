package main

import (
	"fmt"
	"log"
	"os"
)

var DefaultRules = map[string]string{
	".jpg":  "Images",
	".jpeg": "Images",
	".png":  "Images",
	".pdf":  "Documents",
	".doc":  "Documents",
	".docx": "Documents",
	".txt":  "Documents",
	".mp3":  "Music",
	".wav":  "Music",
	".mp4":  "Video",
	".avi":  "Video",
	".zip":  "Archives",
	".rar":  "Archives",
}

type FileOrganizer struct {
	sourceDir      string
	rulesMap       map[string]string
	processedFiles int
	logFile        *os.File
}

func NewFileOrganizer(sourceDir string) (*FileOrganizer, error) {
	if sourceDir == "" {
		return nil, fmt.Errorf("путь к директории не может быть пустым")
	}

	info, err := os.Stat(sourceDir)
	if err != nil {
		return nil, fmt.Errorf("директория не найдена: %w", err)
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("указанный путь не является директорией: %s", sourceDir)
	}

	return &FileOrganizer{
		sourceDir: sourceDir,
		rulesMap:  DefaultRules,
	}, nil
}

func (fo *FileOrganizer) initLog() error {
	file, err := os.OpenFile("organizer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("не удалось открыть лог-файл: %w", err)
	}

	fo.logFile = file
	log.SetOutput(file)
	return nil
}

func (fo *FileOrganizer) logSuccess(message string) {
	log.Printf("[SUCCESS] %s", message)
}

func (fo *FileOrganizer) logError(message string) {
	log.Printf("[ERROR] %s", message)
}

func (fo *FileOrganizer) Close() error {
	if fo.logFile != nil {
		return fo.logFile.Close()
	}
	return nil
}

func main() {
	organizer, err := NewFileOrganizer(".")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	defer organizer.Close()

	if err := organizer.initLog(); err != nil {
		fmt.Printf("Ошибка логирования: %v\n", err)
		return
	}

	organizer.logSuccess("Тест логирования")
	fmt.Println("Логирование работает")
}
