package bridge

import (
	"fmt"
)

// Database 数据库接口
type Database interface {
	Query(sql string) (result string, err error)
}

// MySql 数据库
type MySql struct {
}

func (m *MySql) Query(sql string) (result string, err error) {
	return fmt.Sprintf("MySql result for sql: %s", sql), nil
}

// Oracle 数据库
type Oracle struct {
}

func (m *Oracle) Query(sql string) (result string, err error) {
	return fmt.Sprintf("Oracle result for sql: %s", sql), nil
}
