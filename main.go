package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var urls = []string{
	"https://bsc-dataseed1.binance.org",
	"https://localhost:8545",
}

func main() {

	for {
		ticker := time.NewTicker(2 * time.Second)
		now := time.Now()
		fmt.Println(now.String(), "\n\n")

		for _, url := range urls {
			pengNum, queued, err := getTxpoolStatus(url)
			if err != nil {
				fmt.Println("have err")
				return
			}

			blocknumber, txCount, e := getTxCountNum(url)
			if e != nil {
				fmt.Println("have err")
				return
			}

			fmt.Println("url   "+url+"\tblockNumber:", blocknumber, "txCount:", txCount, "\n", "\t\tpending :", pengNum, "\n", "\t\tqueued :", queued)

		}

		<-ticker.C
	}

}

func getTxpoolStatus(url string) (pengNum, queued uint64, err error) {

	s := `{"jsonrpc":"2.0","method":"` + "txpool_status" + `","id":1}`

	//s:=`{"jsonrpc":"2.0","method":"txpool_status","id":6}`
	jsonStr := []byte(s)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return 0, 0, err
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return 0, 0, e
	}

	//处理body数据

	var res map[string]interface{}
	unmarshal := json.Unmarshal(body, &res)

	if unmarshal != nil {
		return 0, 0, unmarshal
	}

	if i, ok := res["result"]; ok {

		i2 := i.(map[string]interface{})

		sp := i2["pending"].(string)
		sq := i2["queued"].(string)

		u, i3 := strconv.ParseUint(sp, 0, 64)
		if i3 != nil {
			fmt.Println(i3)
		}
		parseUint, i4 := strconv.ParseUint(sq, 0, 64)
		if i4 != nil {
			fmt.Println(i4)
		}

		pengNum, queued = u, parseUint

		return
	} else {
		fmt.Println(res)
		return
	}

	return
}

func getTxCountNum(url string) (blocknumber, txCount uint64, err error) {

	s := `{"jsonrpc":"2.0","method":"` + "eth_blockNumber" + `","id":1}`

	//s:=`{"jsonrpc":"2.0","method":"txpool_status","id":6}`
	jsonStr := []byte(s)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return 0, 0, err
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return 0, 0, e
	}

	//处理body数据

	var res map[string]interface{}
	unmarshal := json.Unmarshal(body, &res)

	if unmarshal != nil {
		return 0, 0, unmarshal
	}

	if i, ok := res["result"]; ok {
		sp := i.(string)
		u, i3 := strconv.ParseUint(sp, 0, 64)
		if i3 != nil {
			fmt.Println(i3)
		}

		s := `{"jsonrpc":"2.0","method":"eth_getBlockTransactionCountByNumber","params":["` + sp + `"],"id":1}`

		jsonStr := []byte(s)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		resp, err := client.Do(req)

		if err != nil {
			return 0, 0, err
		}

		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			return 0, 0, e
		}

		//处理body数据

		var res map[string]interface{}
		unmarshal := json.Unmarshal(body, &res)

		if unmarshal != nil {
			return 0, 0, unmarshal
		}

		if y, ok := res["result"]; ok {

			sp1 := y.(string)
			count, i4 := strconv.ParseUint(sp1, 0, 64)
			if i4 != nil {
				fmt.Println(i4)
			}

			return u, count, nil
		} else {
			fmt.Println(res)
			return 0, 0, nil
		}

	} else {
		fmt.Println(res)
		return 0, 0, nil
	}

}
