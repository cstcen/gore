package gonfig

func SetSdev0() {
	g := GetInstance().Gore
	g.Cache.Hosts = []string{"10.251.43.21:6379"}
	g.Cache.Password = "sdev0#redis#3aPVNy"

	g.Elasticsearch.URL = "http://10.251.104.6:9200"
	g.Elasticsearch.Password = "allDev#es#3aPVNy"

	g.Mongo.Hosts = []string{"10.251.104.15:27017"}

	g.Mysql.Dsn.Address = "10.251.43.32:3306"

	g.Kafka.Consumer.Brokers = []string{"10.251.104.6:9200"}

	g.Redis.Hosts = []string{"10.251.43.21:6379"}
	g.Redis.Password = "sdev0#redis#3aPVNy"
}
