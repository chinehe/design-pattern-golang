package bridge

// DataAccess 数据库访问
type DataAccess struct {
	db Database
}

// SetDatabase 设置数据库
func (access *DataAccess) SetDatabase(db Database) {
	access.db = db
}

// Query 查询
func (access *DataAccess) Query(sql string) (string, error) {
	return access.db.Query(sql)
}
