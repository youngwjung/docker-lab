package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("/app/password.txt")

	if err != nil {
		if os.Getenv("DEBUG") == "TRUE" {
			log.Fatal(err)
		} else {
			log.Fatal("애플리케이션에 치명적인 오류가 있습니다. 상세한 에러 로그를 보고 싶으면 환경변수 DEBUG의 값을 TRUE 설정하고 애플리케이션을 구동하세요.")
		}
	}

	password := strings.TrimSpace(string(data))

	if password == "8lwRp92WTJ" {
		fmt.Println("암호는 \"love dive\" 입니다.")
	} else {
		if os.Getenv("DEBUG") == "TRUE" {
			log.Fatal("패스워드가 틀립니다.")
		} else {
			log.Fatal("애플리케이션에 치명적인 오류가 있습니다. 상세한 에러 로그를 보고 싶으면 환경변수 DEBUG의 값을 TRUE 설정하고 애플리케이션을 구동하세요.")
		}
	}
}
