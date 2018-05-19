package main

type Config struct {
	URL         string
	API_VERSION string
}

func initConfig() *Config {
	conf := &Config{
		URL:         "https://api.coinmarketcap.com/",
		API_VERSION: "v2/",
	}

	return conf
}
