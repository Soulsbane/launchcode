package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	const (
		suspendArgs = "org.kde.KWin /Compositor suspend"
		resumeArgs  = "org.kde.KWin /Compositor resume"
	)

	qdbus, qbusLookPathErr := exec.LookPath("qdbus")

	if qbusLookPathErr != nil {
		panic(qbusLookPathErr)
	}

	fmt.Println("Let's get this started...")

	suspendCommand := exec.Command(qdbus, suspendArgs)
	suspendErr := suspendCommand.Run()

	fmt.Println("Suspending: ", suspendErr)
	time.Sleep(1 * time.Second)

	codeCommand := exec.Command("/opt/visual-studio-code/bin/code")
	codeCommand.Start()

	time.Sleep(1 * time.Second)

	resumeCommand := exec.Command(qdbus, resumeArgs)
	resumeErr := resumeCommand.Run()
	fmt.Println("Resuming: ", resumeErr)
}
