package model

import (
	"fmt"
	"time"
)

var LogRep = Logs{}

type Log struct {
	Time time.Time `json:"time"`
	Body string    `json:"body"`
}
type Logs []Log

func (l *Logs) Set(body string) {
	*l = append(*l, Log{
		Time: time.Now(),
		Body: body,
	})
	fmt.Println(l)
}
func (l *Logs) Get(cnt int) Logs {
	fmt.Println(l)
	fmt.Println(len(*l) - 1 - cnt)
	if len(*l)-1-cnt <= 0 {
		return *l
	}
	return (*l)[len(*l)-1-cnt:]
}

func (l Logs) String() string {
	res := ""
	for _, value := range l {
		res += time.Since(value.Time).String() + " ago >" + value.Body + "\n"
	}
	return res
}
