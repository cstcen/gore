package gonfig

func SetOps(env, appName string) {
	g := GetInstance().Gore
	g.Cache.EnableRing = true
	g.Cache.Hosts = []string{"10.252.1.11:6379"}
	g.Cache.Password = "tenvine@2021"

	g.Elasticsearch.URL = "http://10.252.1.13:9200"

	g.Mongo.Hosts = []string{"10.252.1.13:27027"}

	g.Mysql.Dsn.Address = "10.252.1.15:4725"

	g.Kafka.Consumer.Brokers = []string{"10.252.1.13:9092"}

	g.Redis.Hosts = []string{"10.252.1.11:6379"}
}
