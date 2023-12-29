package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

const HOST_DEFAULT = "http://127.0.0.1"
const PAGE_SIZE_DEFAULT = 20

func GetPort() int {
	res, _ := strconv.Atoi(GetDotEnvVariable("PORT"))
	return res
}

func GetServerInfo(req *http.Request) string {
	clientIp := getClientIpAddr(req)
	if clientIp == "" || clientIp == "127.0.0.1" {
		return fmt.Sprintf("%s:%d", HOST_DEFAULT, GetPort())
	}
	return fmt.Sprintf("%s:%d", GetDotEnvVariable("HOST_ROUTER"), GetPort())
}

func getClientIpAddr(req *http.Request) string {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	return host
}

func GetDotEnvVariable(key string) string {
	err := godotenv.Load("quiz/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
