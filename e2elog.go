package smslog

import (
	"encoding/json"
	"fmt"
	"time"
)

//LogInfoE2E for ..
type LogInfoE2E struct {
	Prefix    string `json:"@prefix"`
	Suffix    string `json:"@suffix"`
	Timestamp string `json:"@timestamp"`

	Level           string  `json:"level"`
	LogCat          string  `json:"log_cate"`
	TxId            string  `json:"tx_id"`
	StepTxId        string  `json:"step_txid"`
	StepName        string  `json:"step_name"`
	EndpointName    string  `json:"endpoint_name"`
	Endpoint        string  `json:"endpoint"`
	ResultCode      string  `json:"result_code"`
	ResultDesc      string  `json:"result_desc"`
	StartDate       string  `json:"start_date"`
	EndDate         string  `json:"end_date"`
	ElapsedTime     float64 `json:"elapsed_time"`
	ResultIndicator string  `json:"result_indicator"`

	Request  string `json:"request"`
	Response string `json:"response"`
}

// NewE2E for create LogInfoE2E
func NewE2E(prefix string, suffix string, txid string) *LogInfoE2E {
	//debugLevel := l
	//fmt.Printf("debug level = %d\n", debugLevel)
	log := &LogInfoE2E{Prefix: prefix, Suffix: suffix, TxId: txid}

	return log
}

// CloseE2E dbinfo
func (l *LogInfoE2E) CloseE2E() {
	l = nil
}

// PrintLogE2E for send log to stdoutput
func (l *LogInfoE2E) PrintLogE2E(req string, res string) {
	if l.Timestamp == "" {
		l.Timestamp = time.Now().Format(time.RFC3339Nano)
	}
	l.Request = req
	l.Response = res

	// write log E2E
	jsonStr, _ := json.Marshal(l)
	fmt.Printf("%s\n", string(jsonStr))
}
