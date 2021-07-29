package gonfig

import "fmt"

func SetDev(env, appName string) {
	g := GetInstance().Gore
	g.Cache.Hosts = []string{"10.251.43.21:6379"}
	g.Cache.Password = "allDev#redis#3aPVNy"

	g.Elasticsearch.URL = "http://10.251.104.6:9200"
	g.Elasticsearch.Password = "allDev#es#3aPVNy"

	g.Mongo.Hosts = []string{"10.251.104.15:27017"}
	g.Mongo.AppName = fmt.Sprintf("%s_%s", env, appName)
	g.Mongo.Username = fmt.Sprintf("%s_%s_user", env, appName)
	g.Mongo.Password = fmt.Sprintf("%s#%s#5q53AR", env, appName)

	g.Mysql.Dsn.Address = "10.251.43.32:3306"
	g.Mysql.Dsn.Dbname = fmt.Sprintf("%s_%s", env, appName)
	g.Mysql.Dsn.Username = fmt.Sprintf("%s_%s", env, appName)
	g.Mysql.Dsn.Password = fmt.Sprintf("%s_%s#Z3KkhG", env, appName)

	g.Kafka.Consumer.Brokers = []string{"10.251.110.41:9092"}

	g.Redis.Hosts = []string{"10.251.43.21:6379"}
	g.Redis.Password = "allDev#redis#3aPVNy"
}
