package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const HOST_DEFAULT = "http://127.0.0.1"
const PAGE_SIZE_DEFAULT = 20

func GetPort() int {
	res, _ := strconv.Atoi(GetDotEnvVariable("PORT"))
	return res
}

func GetServerUrlDefault() string {
	return fmt.Sprintf("%s:%d", HOST_DEFAULT, GetPort())
}

func GetServerUrlRouter() string {
	return fmt.Sprintf("%s:%d", GetDotEnvVariable("HOST_ROUTER"), GetPort())
}

func GetServerInfo(req *http.Request) string {
	clientIp := getClientIpAddr(req)
	if clientIp == "" || clientIp == "127.0.0.1" {
		return GetServerUrlDefault()
	}
	return GetServerUrlRouter()
}

func getClientIpAddr(req *http.Request) string {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	return host
}

func GetDotEnvVariable(key string) string {
	pathEnv := filepath.Join(GetProjectRootPath(), "quiz", ".env")
	err := godotenv.Load(pathEnv)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetDbDumpDir() string {
	return filepath.Join(GetProjectRootPath(), "dump")
}

func GetDumpFilePath() string {
	return filepath.Join(GetDbDumpDir(), "quiz.sql")
}

func GetExPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func GetProjectRootPath() string {
	return filepath.Dir(GetExPath())
}

func GetAdminName() string {
	return GetDotEnvVariable("ADMIN_NAME")
}

func GetTemplateFuncMapForAdminHeader() template.FuncMap {
	return template.FuncMap{
		"GetAdminName": GetAdminName,
	}
}

func GetDotEnvVariableForWifi(key string) string {
	pathEnv := filepath.Join(GetExPath(), ".env")
	err := godotenv.Load(pathEnv)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetWifiName() string {
	return GetDotEnvVariableForWifi("SSID")
}

func GetWifiPassword() string {
	return GetDotEnvVariableForWifi("PASSPHRASE")
}
