package config

import (
	"encoding/json"
	"flag"
	"github.com/owulveryck/gorchestrator/logger"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Config struct {
	Executor struct {
		URL string `json:"url"`
	} `json:"executor,omitempty"`
	HTTP struct {
		BindAddress string `json:"bind_address"`
		BindPort    int    `json:"bind_port"`
		Certificate string `json:"certificate,omitempty"`
		Key         string `json:"key,omitempty"`
		Scheme      string `json:"scheme,omitempty"`
	} `json:"http"`
	Log logger.Log `json:"log"`
}

var (
	config     *Config
	configLock = new(sync.RWMutex)
	configFile string
)

func loadConfig(fail bool) error {
	flag.Parse()
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	temp := new(Config)
	if err = json.Unmarshal(file, temp); err != nil {
		return err
	}
	configLock.Lock()
	config = temp
	configLock.Unlock()
	return nil
}

func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

// go calls init on start
func init() {
	flag.StringVar(&configFile, "config", "config.json", "Config file")
	loadConfig(true)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)
	go func() {
		for {
			<-s
			loadConfig(false)
		}
	}()
}
