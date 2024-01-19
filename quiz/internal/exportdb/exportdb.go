package exportdb

import (
	"bytes"
	"log"
	"os/exec"
	"quiz/internal/common"
)

func RunExportDb() {
	exeFile := common.GetDotEnvVariable("MYSQLDUMP_BIN")
	fileName := common.GetDumpFilePath()
	log.Printf("Process create file dump: %s\n", fileName)

	cmd := exec.Command(
		exeFile,
		"-u", common.GetDotEnvVariable("DBUSER"),
		"--password="+common.GetDotEnvVariable("DBPASS"),
		common.GetDotEnvVariable("DBNAME"),
		"--result-file="+fileName,
	)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Exportdb ERROR: %v\nStdErr: %s\nOutput: %s", err, stderr.String(), out.String())
	}
	log.Printf("Success. Command output:\n%s\n", out.String())
}
