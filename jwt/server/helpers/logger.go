package helpers

import "log"

func ErrorLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorMsgLog(msg string, err error) {
	if err != nil {
		log.Fatalf("message: %v\nerr:%v", msg, err)
	}
}
