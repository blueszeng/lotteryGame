package config

const (
	RedisAddr               string = "127.0.0.1:6379"
	RedisPassword           string = ""
	RedisDB                 int    = 0
	Dbname                  string = "lottery"                         // 数据库名称
	Mysqlconnstring         string = "root:123456@tcp(127.0.0.1:3306)" // mysql连接字符串
	Mysqlconncap            int    = 2048                              // mysql连接池容量
	Mysqlmaxallowedpacketmb int    = 1                                 // mysql通信缓冲区的最大长度，单位MB，默认1MB

)
