package common

type Args struct {
	Name   string `short:"n" long:"name" description:"Application name"`
	Env    string `short:"e" long:"env" description:"Environment"`
	Consul string `short:"c" long:"consul" description:"Consul [host:port]"`
	Log    string `short:"l" long:"log" description:"Whether to output log files"`
}
