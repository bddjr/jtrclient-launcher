package main

import "time"

func myInit() {
	parseArgs()
	delay := *argDelay
	if delay > 0 {
		println("Waiting", delay, "millisecond...")
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	setEnv()
	completeUpgrade()
}
