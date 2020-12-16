package xormplus

//数据库连接配置
type Config struct {
	ShowSql        bool     `hcl:"show_sql"`
	MaxIdle        int      `hcl:"max_idle"`
	MaxConn        int      `hcl:"max_conn"`
	Master         string   `hcl:"master"`
	Slaves         []string `hcl:"slaves"`
	UseMasterSlave bool     `hcl:"use_master_slave"`
}
