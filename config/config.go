package config

import "time"

const (
	//gin监听的地址和端口
	ListenAddr = "0.0.0.0:9090"
	KubeConfig = "/Users/adoo/.kube/config"
	//KubeConfig = "C:\\custom\\project\\go\\config"
	//查看日志的行数
	PodLogTailLine = 2000
	//管理员账号密码
	AdminUser = "admin"
	AdminPwd = "123456"
	//数据库配置
	DbType="mysql"
	DbHost="localhost"
	DbPort=3306
	DbName="k8s_platform"
	DbUser="root"
	DbPassword="123456"
	//打印mysql debug日志开关
	LogMode= false
	//连接池配置
	MaxIdleConns = 10 //最大空闲连接
	MaxOpenConns = 100 //最大连接数
	MaxLifeTime = 30 * time.Second //最大生存时间
)