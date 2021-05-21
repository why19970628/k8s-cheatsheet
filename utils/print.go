package utils

import "log"

func HandlerCheck(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}
