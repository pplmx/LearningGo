package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	filePathNames := ListMarkdownFiles("/home/mystic/JetBrains/IdeaProjects/caoyu.info/cc/source")
	ConvertHexoMarkdown2Hugo(filePathNames)
}

func ListMarkdownFiles(path string) []string {
	files, err := WalkMatchedFiles(path, "*.md")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func WalkMatchedFiles(dir string, pattern string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, info.Name()); err != nil {
			return err
		} else if matched {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

//ConvertHexoMarkdown2Hugo Convert all the front matter of hexo markdowns to hugo's
func ConvertHexoMarkdown2Hugo(files []string) {
	newDir := fmt.Sprintf("hugo_posts%d", time.Now().Unix())
	err := os.Mkdir(newDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// traverse all files and read the file line by line
	for idx, file := range files {
		func() {
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(f)

			fmt.Printf("Reading No.%d file: %s\n", idx+1, file)
			scanner := bufio.NewScanner(f)
			counts := 0
			var lines []string
			for scanner.Scan() {
				line := scanner.Text()
				if counts >= 2 {
					lines = append(lines, line)
					continue
				}
				if strings.HasPrefix(line, "---") {
					counts++
				}
				switch {
				// handle the line contains "date:"
				case strings.HasPrefix(line, "date:"):
					line = strings.TrimSpace(line)
					line = HandleDate(line)
					// get the filename to md5
					// append slug and draft is false
					lastIndex := strings.LastIndex(file, "/")
					filename := strings.Split(file[lastIndex+1:], ".md")[0]
					line += fmt.Sprintf("\nslug: %x\ndraft: false", md5.Sum([]byte(filename)))
				// handle the line contains "updated:"
				case strings.HasPrefix(line, "updated:"):
					line = strings.TrimSpace(line)
					line = HandleUpdated(line)
					// handle the line contains "tags:"
					// handle the line contains "categories:"
				}
				lines = append(lines, line)
			}
			fmt.Printf("Updating No.%d file: %s\n", idx+1, file)

			content := strings.Join(lines, "\n")
			if err := HandleContent(newDir, file, content); err != nil {
				log.Fatal("Failed to update file: ", err)
			}
			fmt.Printf("%s Complete to write\n", file)

		}()
	}
}

//HandleContent 将处理完的内容覆盖进去
func HandleContent(newDir string, filePathNames string, result string) error {
	//如果路径分割符为 \ 则替换为 /
	filePathNames = strings.Replace(filePathNames, "\\", "/", -1)

	//获取最后一个分割符后的名称
	index := strings.LastIndex(filePathNames, "/")
	fileName := filePathNames[index+1:]
	//创建文件
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", newDir, fileName), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	_, err = io.WriteString(f, result)
	if err != nil {
		return err
	}
	return nil
}

//HandleDate handle the line contains "date:"
func HandleDate(date string) string {
	// date: 2018-12-12 15:00:00
	// to
	// date: 2018-12-12T15:00:00+08:00
	arr := strings.Split(date, "date:")
	t := strings.TrimSpace(arr[1])
	arr = strings.Split(t, " ")
	date = fmt.Sprintf("date: %sT%s+08:00", arr[0], arr[1])
	return date
}

//HandleUpdated handle the line contains "updated:"
func HandleUpdated(updated string) string {
	// updated: 2018-12-12 15:00:00
	// to
	// lastmod: 2018-12-12T15:00:00+08:00
	arr := strings.Split(updated, "updated:")
	t := strings.TrimSpace(arr[1])
	arr = strings.Split(t, " ")
	lastMod := fmt.Sprintf("lastmod: %sT%s+08:00", arr[0], arr[1])
	return lastMod
}
