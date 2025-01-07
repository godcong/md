package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/godcong/md/internal/markdown"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Generate index of markdown files",
	Run:   runIndex,
}

func init() {
	rootCmd.AddCommand(indexCmd)
	indexCmd.Flags().StringP("output", "o", "tree.md", "Output file name")
}

func runIndex(cmd *cobra.Command, args []string) {
	output, _ := cmd.Flags().GetString("output")

	var mds []markdown.Markdown
	for i := range args {
		err := filepath.Walk(args[i], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".md") {
				md := markdown.New(path)
				md.Rel(args[i])
				md.Level()
				mds = append(mds, md)
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error walking directory: %v\n", err)
			return
		}
	}

	// create an index md
	f, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error creating index md: %v\n", err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, _ = w.WriteString("# Markdown Files Index\n\n")

	parent := ""
	for _, md := range mds {
		fmt.Println("Read file: ", md.Path())
		title := md.Title()
		relPath := md.RelPath()
		dir := filepath.Dir(relPath)
		if parent != dir {
			parent = dir
			_, _ = w.WriteString(fmt.Sprintf("## %s\n", dir))
		}
		_, _ = w.WriteString(fmt.Sprintf("- [%s](%s)\n", title, relPath))
	}
	_ = w.Flush()
	fmt.Printf("Index generated: %s\n", output)
}
