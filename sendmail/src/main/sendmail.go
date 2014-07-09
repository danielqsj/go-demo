package main
import (
	"net/smtp"
	"fmt"
	"strings"
)

/*
 *	user : example@example.com login smtp server user
 *	password: xxxxx login smtp server password
 *	host: smtp.example.com:port   smtp.163.com:25
 *	to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */


func SendMail(user, password, host, to, subject, body, mailtype string) error{
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/"+ mailtype + "; charset=UTF-8"
	}else{
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<"+ user +">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	user := "zlyzhq@126.com"
	password := ""
	host := "smtp.126.com:25"
	to := "qsj_daniel@163.com;qsj.daniel@gmail.com"

	subject := "Test send email by golang"
/*
	body := `
	<html>
	<body>
	<h3>
	"Test send email by golang :)"
	</h3>
	</body>
	</html>
	`
*/
	body := "this is test by text/plain golang!!"
	fmt.Println("send email")
//	err := SendMail(user, password, host, to, subject, body, "html")
	err := SendMail(user, password, host, to, subject, body, "plain")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	}else{
		fmt.Println("send mail success!")
	}

}
