package subgen

import (
	"fmt"
	"log"
	"testing"
)

// 使用“，有跳行
func TestExtractColumnsFromSQL(t *testing.T) {
	sql := `
	SELECT
		id,
		cid,
		NAME,
		mobile,
		alert_mobile,
		linkman,
		address,
		email,
		create_time,
		name_show,
		speed_limit,
		wireless_offline_threshold,
		wired_offline_threshold,
		emergency_offline_threshold,
		wired_hbi,
		wireless_hbi 
	FROM
		t_corp 
	WHERE
		cid = ?
	`
	fromSQL := ExtractColumnsFromSQL(sql)
	fmt.Println(fromSQL)
	if len(fromSQL) != 16 {
		log.Fatal("提取SQL中的列失败")
	}
}

// 字符串里还有`
func TestExtractColumnsFromSQL2(t *testing.T) {
	sql2 := " SELECT `mobile`, `type`, `status`, `auth_type`  FROM t_operator WHERE `cid`=? AND status=1"
	fromSQL := ExtractColumnsFromSQL(sql2)
	fmt.Println(fromSQL)
	fmt.Println(len(fromSQL))
}

func TestGetAllStruct(t *testing.T) {
	err := GetAllStruct()
	if err != nil {
		log.Fatal(err)
	}
}

func TestSolve(t *testing.T) {
	err := Solve()
	if err != nil {
		log.Fatal()
	}
}

func TestGetSqlString(t *testing.T) {
	sqlString := GetSqlString()
	fmt.Println(sqlString)
}
