package helper

import "log"

func CheckAndLogError(errMessage string, err string, match bool) {
	if match == false {
		log.Fatalf("ERROR: %s: %s", errMessage, err)
	}
}
