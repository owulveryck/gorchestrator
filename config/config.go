package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

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
