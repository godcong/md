package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Generate index of markdown files",
	Run:   runIndex,
}

func init() {
	rootCmd.AddCommand(indexCmd)
	indexCmd.Flags().StringP("dir", "d", ".", "Directory to scan")
	indexCmd.Flags().StringP("output", "o", "index.md", "Output file name")
}

func runIndex(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("dir")
	output, _ := cmd.Flags().GetString("output")

	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".md") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return
	}

	// 创建索引文件
	f, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error creating index file: %v\n", err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("# Markdown Files Index\n\n")

	for _, file := range files {
		title := extractTitle(file)
		relPath, _ := filepath.Rel(dir, file)
		w.WriteString(fmt.Sprintf("- [%s](%s)\n", title, relPath))
	}

	w.Flush()
	fmt.Printf("Index generated: %s\n", output)
}

func extractTitle(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return filepath.Base(filename)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}

	return filepath.Base(filename)
}
