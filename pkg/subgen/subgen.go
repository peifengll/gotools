package subgen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func ExtractColumnsFromSQL(sqlStatement string) []string {
	// 将换行符替换为空格
	sqlStatement = strings.ReplaceAll(sqlStatement, "\n", " ")
	sqlStatement = regexp.MustCompile(`\s+`).ReplaceAllString(sqlStatement, " ")
	// 匹配 SQL 语句中的列名
	re := regexp.MustCompile(`(?i)select\s+(.*?)\s+from`)
	match := re.FindStringSubmatch(sqlStatement)
	if len(match) > 1 {
		columnsStr := match[1]
		// 去除可能存在的空格，并将列名分割成列表
		re := regexp.MustCompile(`\s*,\s*`)

		columns := re.Split(columnsStr, -1)
		for i := 0; i < len(columns); i++ {
			columns[i] = strings.Replace(columns[i], "`", "", -1)
		}
		return columns
	} else {
		return nil
	}
}

func GetSqlColumns(s string) []string {
	// 有select的情况
	return ExtractColumnsFromSQL(s)
}

func GetAllStruct() (res []string) {
	file, err := os.Open("temp.tmp")
	if err != nil {
		log.Println("读取文件发生错误", err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}
	return res
}

func GetSqlString() string {
	file, err := os.Open("sql.tmp")
	if err != nil {
		return ""
	}
	defer file.Close()

	var sqlStatements strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sqlStatements.WriteString(line)
		sqlStatements.WriteString(" ")
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	return sqlStatements.String()
}

// 获取到json标签后边的值
func extractJSONTag(jsonString string) string {
	// 匹配 JSON 标签部分
	re := regexp.MustCompile(`json:"([^"]+)"`)
	match := re.FindStringSubmatch(jsonString)

	if len(match) > 1 {
		jsonTag := match[1]
		return jsonTag
	} else {
		return ""
	}
}

func ProduceSubStruct(res []string) error {
	file, err := os.Create("resutl.tmp")
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, i := range res {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
	fmt.Println("子结构体成功写入文件")
	return nil
}

func Solve() error {
	sql := GetSqlString()
	dict := make(map[string]bool)
	subcolumns := make([]string, 0)
	columns := GetSqlColumns(sql)
	for _, column := range columns {
		tc := strings.ToUpper(column)
		dict[tc] = true
	}
	res := GetAllStruct()
	for _, line := range res {
		tag := extractJSONTag(line)
		if tag != "" {
			tc := strings.ToUpper(tag)
			if _, ok := dict[tc]; ok {
				subcolumns = append(subcolumns, line)
			}
		}
	}
	err := ProduceSubStruct(subcolumns)
	if err != nil {
		return err
	}
	return nil
}
