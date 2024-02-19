package leet

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"time"
)

func CreateToday(context *cli.Context) error {
	if _, err := os.Stat("leetcode"); os.IsNotExist(err) {
		fmt.Println("Algorithm folder does not exist.")
		err := os.Mkdir("leetcode", 0755)
		if err != nil {
			fmt.Println("Error creating algorithm folder:", err)
			return err
		}
		fmt.Println("Algorithm folder created successfully.")
	}
	today := time.Now().Format("2006-01-02")

	fileName := filepath.Join("leetcode", today+".go")
	fmt.Printf("Creating %s file...\n", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating  file:", err)
		return err
	}
	defer file.Close()
	// 写入文件内容
	_, err = file.WriteString(`package leetcode
`)
	if err != nil {
		fmt.Println("Error writing to  file:", err)
		return err
	}
	fmt.Printf("%s file created successfully.\n", fileName)
	return nil
}
