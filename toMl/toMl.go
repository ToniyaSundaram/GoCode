package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

var tomlData = `[validator]
address = "vaUjHHvCLmUzbzrz7xgF266xZEDvRa6RSDC"

[tendermint]
seeds = ""
persistentPeers = ""
listenAddress = "tcp://127.0.0.1:45566"
moniker = "moniker"
tendermintRoot = "data"

[rpc]

[logging]
ExcludeTrace = false
NonBlocking = false
[logging.RootSink]
  [logging.RootSink.Output]
    OutputType = "stderr"
    Format = "json"`

type feature1 struct {
	Enable  bool
	Userids []string
}

type feature2 struct {
	Enable bool
}

type tomlConfig struct {
	Title string
	F1    feature1 `toml:"feature1"`
	F2    feature2 `toml:"feature2"`
}

type Config struct {
	Validator  ValidatorConfig  `toml:"validator"`
	Tendermint TendermintConfig `toml:"tendermint"`
}

type RPCConfig struct {
}

type TendermintConfig struct {
	// Initial peers we connect to for peer exchange
	Seeds string `toml:"seeds"`
	// Peers to which we automatically connect
	PersistentPeers string `toml:"persistentPeers"`
	ListenAddress   string `toml:"listenAddress"`
	Moniker         string `toml:"moniker"`
	TendermintRoot  string `toml:"tendermintRoot"`
}

type ValidatorConfig struct {
	Address string `toml:"address"`
}

func main() {
	var conf tomlConfig
	if _, err := toml.Decode(tomlData, &conf); err != nil {
		log.Fatal(err)
	}
	log.Printf("title: %s", conf.Title)
	log.Printf("Feature 1: %#v", conf.F1)
	log.Printf("Feature 2: %#v", conf.F2)
}
