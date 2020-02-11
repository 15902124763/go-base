package db

import (
	"encoding/json"
	"fmt"
)

// 表t_test映射
type Ttest struct {
	Id       int64
	Name     string
	Password string
}

func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}

// 增
func Insert() {
	conn, err := Conn()
	if err == nil {
		exec, err := conn.Exec("INSERT INTO t_test(name,password ) VALUES(?,?)", "小红", "123456")
		if err == nil {
			affected, _ := exec.RowsAffected()
			fmt.Print(affected)
		}

		if err != nil {
			err_handler(err)
		}
	}
}

// 删
func Delete() {
	conn, err := Conn()
	if err != nil {
		err_handler(err)
		return
	}
	exec, err := conn.Exec("delete from t_test where name = ?", "小红")
	if err != nil {
		err_handler(err)
		return
	}
	affected, err := exec.RowsAffected()
	if err == nil {
		fmt.Print("删除行数：", affected)
	}

}

// 改
func Update() {
	conn, err := Conn()
	if err != nil {
		err_handler(err)
		return
	}

	exec, err := conn.Exec("update t_test set password = ? where name = ?", "11111", "小红")
	if err == nil {
		affected, _ := exec.RowsAffected()
		fmt.Print("更新:", affected)
	}
}

// 查
func Query() {
	conn, err := Conn()
	if err == nil {
		query, err := conn.Query("select id Id, name Name,password Password from t_test")
		if err != nil {
			err_handler(err)
			return
		}

		for query.Next() {
			var test Ttest
			err := query.Scan(&test.Id, &test.Name, &test.Password)
			if err != nil {
				err_handler(err)
			}

			test_json, err := json.Marshal(test)
			if err == nil {
				fmt.Print("\n", string(test_json))
			}
		}
	}
}
