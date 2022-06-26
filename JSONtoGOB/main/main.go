package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type Container struct {
	Template *Metric
	Metrics  []*Metric
}

type Metric struct {
	Time     *time.Time
	Metadata map[string]interface{}
	Data     map[string]interface{}
}

func main() {

	Timestamp_data := time.Now()
	metadata := map[string]interface{}{
		"ifname":    "asf",
		"lane":      5,
		"something": "random"}
	data := map[string]interface{}{
		"ifHCInOctets":  125125,
		"ifHCOutOctets": 125125,
	}
	Metric_data1 := Metric{
		&Timestamp_data,
		metadata,
		data}

	res := make([]*Metric, 1)
	for i := 0; i < 1; i++ {
		res[i] = &Metric_data1
	}
	skogulGOB_TestData := Container{Metrics: res}
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	err := encoder.Encode(skogulGOB_TestData)
	if err != nil {
		fmt.Println("Error occurred in Encoding", err)
	}
	b := buf.Bytes()
	f, e := os.Create("testdata.gob")
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	n, e1 := f.Write(b)
	if e1 != nil {
		fmt.Println(e)
	}
	fmt.Println(n, b)

	var GOB_decode_buffer *Container
	f, err_f := os.Open("testdata.gob")
	if err_f != nil {
		fmt.Println("fileopen error", err_f)
	}
	defer f.Close()
	decoder := gob.NewDecoder(f)
	err_d := decoder.Decode(&GOB_decode_buffer)
	if err_d != nil {
		fmt.Println("Decode error", err_d)
	}
	fmt.Println(*GOB_decode_buffer.Metrics[0])

}
