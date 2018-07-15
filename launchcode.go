package main

import "os/exec"
import "fmt"

/*
#!/bin/bash
echo -n "Current compositing state: "
qdbus org.kde.kwin /KWin compositingActive
echo -n "Changing compositing state..."
qdbus org.kde.kwin /KWin toggleCompositing
echo -n "Current compositing state: "
qdbus org.kde.kwin /KWin compositingActive
echo -n "All done. Press ENTER to close window: "
read ENTRY
*/
func main() {
	/*binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}*/
	binary, lookErr := exec.LookPath("qdbus")
	if lookErr != nil {
		panic(lookErr)
	}

	fmt.Println(binary)
}
