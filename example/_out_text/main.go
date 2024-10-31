package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 命令行参数
	sourceDir := flag.String("src", ".", "源目录路径")
	outputFile := flag.String("out", "output.txt", "输出文件路径")
	exclude := flag.String("exclude", ".git,.idea,node_modules,vendor", "要排除的目录，用逗号分隔")
	fileTypes := flag.String("types", ".go,.java,.py,.js,.cpp,.h,.c,.txt", "要包含的文件类型，用逗号分隔")
	flag.Parse()

	// 转换排除目录和文件类型为map，便于快速查找
	excludeDirs := make(map[string]bool)
	for _, dir := range strings.Split(*exclude, ",") {
		excludeDirs[dir] = true
	}

	fileExtensions := make(map[string]bool)
	for _, ext := range strings.Split(*fileTypes, ",") {
		fileExtensions[ext] = true
	}

	// 创建输出文件
	out, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("创建输出文件失败: %v\n", err)
		return
	}
	defer out.Close()

	// 遍历目录
	err = filepath.Walk(*sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查是否是要排除的目录
		if info.IsDir() {
			if excludeDirs[info.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// 检查文件扩展名
		ext := strings.ToLower(filepath.Ext(path))
		if !fileExtensions[ext] {
			return nil
		}

		// 写入文件路径作为分隔符
		fmt.Fprintf(out, "\n\n=== %s ===\n\n", path)

		// 打开并读取文件
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("打开文件 %s 失败: %v", path, err)
		}
		defer file.Close()

		// 复制文件内容到输出文件
		_, err = io.Copy(out, file)
		if err != nil {
			return fmt.Errorf("复制文件 %s 内容失败: %v", path, err)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("处理过程中发生错误: %v\n", err)
		return
	}

	fmt.Println("文件合并完成!")
}

// go run . -src=../../../tp-protocol-sdk-go -out=result.txt -types=.go,.java -exclude=.git,vendor
