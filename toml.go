package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

// Conf ...
type Conf struct {
	Name      string
	CountI    int32
	CountJ    int32
	IsKilled  bool
	KillTime  string
	StartTime string
	EndTime   string
}
type CfConf struct {
	KillProcessName    string
	AllowMinutesToPlay int32
}

func ReadConfigs(path string) (conf CfConf, err error) {
	var config CfConf
	if _, err := toml.DecodeFile(path, &config); err != nil {
		fmt.Println(err)
	}
	return config, err
}
func WriteToml(path string, i int32, j int32, isKilled bool, killTime string, startTime string, endTime string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	encode := toml.NewEncoder(file)
	if err = encode.Encode(&Conf{Name: "gameCount", CountI: i, CountJ: j, IsKilled: isKilled, KillTime: killTime, StartTime: startTime, EndTime: endTime}); err != nil {
		return
	}
}

func ReadToml(path string) (conf Conf, err error) {
	var config Conf
	if _, err := toml.DecodeFile(path, &config); err != nil {
		fmt.Println(err)
	}
	return config, err
}
func ExistsFile(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
