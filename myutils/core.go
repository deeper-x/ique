// Package myutils has some helpers and utils
package myutils

import "log"

// FailsOnError debug utils
func FailsOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%v %v", msg, err)
	}
}
