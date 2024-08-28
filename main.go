// Copyright (C) 2018 Betalo AB - All Rights Reserved

package main

import (
	"fmt"
	"forward/WindowsUI"
	"strings"
	"sync"
	"time"
)

var i int32
var j int32
var AllowMinutesToPlay int32
var isKilled bool
var killTime string
var startTime string
var endTime string
var StrProcess string

func main() {
	//time.Minute to time.Second

	var GameTicker = time.NewTicker(time.Minute * 1) // 每分钟执行一次
	i = 0
	j = 0
	AllowMinutesToPlay = 34
	isKilled = false
	killTime = ""
	startTime = "" //no use it
	endTime = ""
	cf := "./configs.toml"
	StrProcess = "Palgame.exe,YURI.exe,gamemd.exe,msedge.exe"
	if ExistsFile(cf) {
		cfConf, err := ReadConfigs(cf)
		if err != nil {
			fmt.Println("config file not exist.")
		}
		StrProcess = cfConf.KillProcessName
		AllowMinutesToPlay = cfConf.AllowMinutesToPlay
	}
	fmt.Println(StrProcess)
	fp := "./gamekill.toml"
	if ExistsFile(fp) {
		conf, err := ReadToml(fp)
		if err != nil {
			i = 0
			j = 0
			isKilled = false
			killTime = ""
			startTime = ""
			endTime = ""
		} else {
			i = conf.CountI
			j = conf.CountJ
			isKilled = conf.IsKilled
			killTime = conf.KillTime
			startTime = conf.StartTime
			endTime = conf.EndTime
			if killTime != "" {
				stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", killTime, time.Local)
				t := time.Now()
				diff := t.Sub(stamp)
				m := int32(diff.Minutes())
				if m > 175 {
					i = 0
					j = 0
					isKilled = false
					killTime = ""
					startTime = ""
					endTime = ""
				} else {
					j = m
				}
			}
			if endTime != "" && !isKilled {
				stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
				t := time.Now()
				diff := t.Sub(stamp)
				m := int32(diff.Minutes())
				if m > 160 {
					i = 0
					j = 0
					isKilled = false
					killTime = ""
					startTime = ""
					endTime = ""
				}
			}
		}
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for range GameTicker.C {

			if isGameProcessExist() && !isKilled {
				if i == 0 {
					startTime = time.Now().Format("2006-01-02 15:04:05")
				}
				endTime = time.Now().Format("2006-01-02 15:04:05")
				i++
				if i > AllowMinutesToPlay-2 && i < AllowMinutesToPlay+1 { //32
					WindowsUI.WTSSendMessage(1, "Hello", "The Game time is coming to end, Please save !", WindowsUI.MB_YESNO, 30)
				}
				if i > AllowMinutesToPlay {
					KillAllGameProcess()
					isKilled = true
					killTime = time.Now().Format("2006-01-02 15:04:05")
					startTime = ""
					endTime = ""
				}
			}
			if isKilled {
				j++
				KillAllGameProcess()
				if j > 180 {
					i = 0
					j = 0
					isKilled = false
					killTime = ""
					startTime = ""
					endTime = ""
				}
			}
			fmt.Println(i, j, isKilled, killTime)
			WriteToml(fp, i, j, isKilled, killTime, startTime, endTime)
		}
	}()
	wg.Wait()
	fmt.Println("bye")
}
func KillAllGameProcess() {
	tmp := strings.Split(StrProcess, ",")
	for i := 0; i < len(tmp); i++ {
		killProcessByName(strings.TrimSpace(tmp[i]))
	}
}

func isGameProcessExist() bool {

	/*
		a, _, _ := isProcessExist("Palgame.exe")
		b, _, _ := isProcessExist("YURI.exe")
		c, _, _ := isProcessExist("gamemd.exe")
		d, _, _ := isProcessExist("msedge.exe")
			if a || b || c || d {
			return true
		} else {
			return false
		}*/
	boolTmp := false
	tmp := strings.Split(StrProcess, ",")
	for i := 0; i < len(tmp); i++ {
		a, _, _ := isProcessExist(strings.TrimSpace(tmp[i]))
		if a {
			boolTmp = true
		}
	}
	return boolTmp
}
