package appwifi

import (
	"bytes"
	"log"
	"os/exec"
	"quiz/internal/common"
	"runtime"
)

func RunMobileHotspot() {
	if runtime.GOOS != "windows" {
		log.Fatalf("needs to be implemented for other platforms")
	}
	cmd := exec.Command("cmd", "/C", "start", "MobileHotspot.exe")
	cmd.Dir = common.GetExPath()
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed: %v\nStdErr: %s\nOutput: %s", err, stderr.String(), out.String())
	}
	log.Printf("Subprocess %d, exiting\n", cmd.Process.Pid)
}
