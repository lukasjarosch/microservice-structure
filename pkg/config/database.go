package config

type MongoDBConfiguration struct {
	Address string `envconfig:"MONGODB_ADDRESS" default:"mongodb://root:root@localhost:27017/"`
}

type MySQLConfiguration struct {
	Address string `envconfig:"MYSQL_ADDRESS" default:"mysql://root:root@localhost:3006/"`
}
