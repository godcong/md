package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var diaryCmd = &cobra.Command{
	Use:   "diary",
	Short: "Create a new diary entry",
	Run:   runDiary,
}

func init() {
	rootCmd.AddCommand(diaryCmd)
}

func runDiary(cmd *cobra.Command, args []string) {
	now := time.Now()
	baseDir := "diary"

	// 创建年/月/日目录
	dirPath := filepath.Join(baseDir,
		fmt.Sprintf("%d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%02d", now.Day()))

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// 查找当前最大序号
	files, _ := os.ReadDir(dirPath)
	var mdFiles []os.DirEntry
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
			mdFiles = append(mdFiles, file)
		}
	}
	nextNum := len(mdFiles) + 1

	// 创建新的日记文件
	filename := filepath.Join(dirPath, fmt.Sprintf("%03d.md", nextNum))
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer f.Close()

	// 写入日记模板
	template := fmt.Sprintf("# Diary Entry %d - %s\n\n", nextNum, now.Format("2006-01-02"))
	f.WriteString(template)

	fmt.Printf("Created new diary entry: %s\n", filename)
}
