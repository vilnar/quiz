package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"quiz/internal/common"
	"runtime"
)

func main() {
	exeFile := common.GetDotEnvVariable("MYSQL_BIN")
	fmt.Printf("exe file: %s\n", exeFile)
	fileName := common.GetDumpFilePath()
	fmt.Printf("Get data from file dump: %s\n", fileName)

	arg := fmt.Sprintf(
		`%s --user %s --password=%s --database %s < %s`,
		common.GetDotEnvVariable("MYSQL_BIN"),
		common.GetDotEnvVariable("DBUSER"), common.GetDotEnvVariable("DBPASS"),
		common.GetDotEnvVariable("DBNAME"),
		fileName,
	)
	cmd := exec.Command("bash", "-c", arg)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", arg)
	}
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Exportdb ERROR: %v\nStdErr: %s\nOutput: %s", err, stderr.String(), out.String())
	}
	fmt.Printf("Success. Command output:\n%s\n", out.String())
}
