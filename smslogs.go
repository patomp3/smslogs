package smslog

import (
	"encoding/json"
	"fmt"
	"time"
)

//LogInfo for ..
type LogInfo struct {
	Timestamp string   `json:"@timestamp"`
	Tags      []string `json:"tags"`
	OrderNo   string   `json:"orderno"`
	//TrackingNo      string   `json:"trackingno"`
	CorrelationID   string `json:"correlationid"`
	ApplicationName string `json:"applicationname"`
	FunctionName    string `json:"functionname"`
	OrderDate       string `json:"orderdate"`
	OrderType       string `json:"ordertype"`
	TVSNo           string `json:"tvsno"`
	MobileNo        string `json:"mobileno"`
	SerialNo        string `json:"serialno"`
	Reference1      string `json:"reference1"`
	Reference2      string `json:"reference2"`
	Reference3      string `json:"reference3"`
	Reference4      string `json:"reference4"`
	Reference5      string `json:"reference5"`
	Request         string `json:"request"`
	Response        string `json:"response"`
	Start           string `json:"start"`
	End             string `json:"end"`
	Duration        string `json:"duration"`
	Level           string `json:"level"`
}

var debugLevel int

// Level for debug
type Level int

const (
	ALL Level = iota + 1
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
)

// New for create logInfo
func New(appName string, funcName string, level Level) *LogInfo {
	//debugLevel := l
	//fmt.Printf("debug level = %d\n", debugLevel)
	log := &LogInfo{Timestamp: time.Now().Format(time.RFC3339Nano), Level: getName(level), ApplicationName: appName, FunctionName: funcName}

	return log
}

// PrintLog for send log to stdoutput
func (l *LogInfo) PrintLog() {
	jsonStr, _ := json.Marshal(l)
	fmt.Printf("%s\n", string(jsonStr))
}

/*func (l *LogInfo) SubmitLog(params ...Level) {
	//
	total := 0
	for _, num := range params {
		total = total + 1
		_ = num
	}
	if total > 0 {
		l.Level = getName(params[0])
	}
	jsonStr, _ := json.Marshal(l)

	fmt.Printf("%s\n", string(jsonStr))
}*/

/*func main() {

	log := New(INFO)
	log.PrintLog()

	log.TVSNo = "6666666"
	log.PrintLog()

	tag := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	log.Tags = tag
	log.PrintLog()
	//displayLog(ALL, "Hello")
}*/

func getName(level Level) string {
	var myReturn string

	switch level {
	case ALL:
		myReturn = "ALL"
	case TRACE:
		myReturn = "TRACT"
	case DEBUG:
		myReturn = "DEBUG"
	case INFO:
		myReturn = "INFO"
	case WARN:
		myReturn = "WARN"
	case ERROR:
		myReturn = "ERROR"
	}
	return myReturn
}
