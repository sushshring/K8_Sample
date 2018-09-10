package model

import (
	`fmt`

	`github.com/go-bongo/bongo`
)

type Request struct {
	bongo.DocumentBase `bson:",inline"`
	UserAgent          string
	Method             string
	Path               string
	Source             string
}

func GetName() string {
	return "request"
}

func (req Request) String() string {
	return fmt.Sprintf("UserAgent: %s, Method: %s, Path: %s, Source: %s\n", req.UserAgent, req.Method, req.Path, req.Source)
}
