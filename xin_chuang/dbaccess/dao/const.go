package dao

// 数据库类型
const (
	// DBTypeMysql mysql数据库
	DBTypeMysql = "mysql"
	// DBTypeKingBaseMysql kingbase数据库，支持mysql兼容模式
	DBTypeKingBaseMysql = "kingbase_mysql"
	// DBTypeKingBasePgsql kingbase数据库，支持pgsql兼容模式
	DBTypeKingBasePgsql = "kingbase_pgsql"
	// DBTypePostgres postgres数据库
	DBTypePostgres = "postgres"
	// DBTypeSqlite sqlite3数据库
	DBTypeSqlite = "sqlite3"
	// DBTypeMemory memory数据库
	DBTypeMemory = "memory"
)
