package utils

import(
	"net/smtp"
)

func SendMail(receivers []string,message []byte)error{

	sender := "goguruh01@gmail.com"
	password := "Qwertyu!op"


	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("",sender,password,host)
	
	err := smtp.SendMail(host+":587",auth,sender,receivers,message)

	if err != nil{
		panic(err)
	}
	return nil
}

