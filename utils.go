package main

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func prettyPrint(message string, obj interface{}) {
	data, err := json.MarshalIndent(obj, "\t", "")
	check(err)
	logrus.Debugln(message, string(data))
}
