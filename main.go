package main

import (
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/alamre/godemo/libs"
)

//Port : smtp port
const Port = "587"

//Host : smtp host
const Host = "smtp.office365.com"

//Username : smtp username
const Username = "test@test.ly"

//Password : smtp password
const Password = "password"

func main() {
	diskStatus := libs.DiskUsage("/Volumes")

	free := diskStatus.Free / 1024 / 1024
	used := diskStatus.Used / 1024 / 1024
	all := diskStatus.All / 1024 / 1024

	allSpace := strconv.FormatUint(all, 10)
	freeSpace := strconv.FormatUint(free, 10)
	usedSpace := strconv.FormatUint(used, 10)

	// put list of email
	to := []string{
		"test@test.ly",
	}

	// pass the body of email
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: Zabbix Alert\n"
	body := []byte(subject + mime + "<html><body><h2 style='color: white; background-color: red;'>The capacity of storage the Zabbix server has reached a risk limit</h2><h3> All space : " + allSpace + " - MB</h3> <h3> Used space : " + usedSpace + " - MB</h3> <h3> Free space : " + freeSpace + " - MB</h3></body></html>")

	auth := libs.LoginAuth(Username, Password)
	err := smtp.SendMail(Host+":"+Port, auth, Username, to, body)

	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
}
