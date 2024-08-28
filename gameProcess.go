package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func isLinuxProcessRunning(processName string) bool {
	cmd := exec.Command("ps", "-C", processName)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error while checking process status")
		return false
	}
	if len(output) > 0 {
		return true
	}
	return false
}

func isProcessExist(appName string) (bool, string, int) {
	appary := make(map[string]int)
	cmd := exec.Command("cmd", "/C", "tasklist")
	output, _ := cmd.Output() //fmt.Printf("fields: %v\n", output)
	n := strings.Index(string(output), "System")
	if n == -1 {
		fmt.Println("no System process find")
		os.Exit(1)
	}
	data := string(output)[n:]
	fields := strings.Fields(data)
	for k, v := range fields {
		if v == appName {
			appary[appName], _ = strconv.Atoi(fields[k+1])

			return true, appName, appary[appName]
		}
	}

	return false, appName, -1
}

func killProcessByName(procname string) int {
	kill := exec.Command("taskkill", "/IM", procname, "/T", "/F")
	err := kill.Run()
	if err != nil {
		return -1
	}
	return 0
}
