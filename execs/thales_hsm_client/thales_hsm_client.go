package main

import (
	"encoding/hex"
	"fmt"
	"github.com/rkbalgi/go/net"
	"os"
	"time"
)

// this program keeps writing a NC command to the hsm at the given ip and port
//every 5 seconds
func main() {

	client := net.NewNetCatClient("127.0.0.1:1500", net.MLI_2E)
	err := client.OpenConnection()
	if err != nil {
		fmt.Println("error", err.Error())
		os.Exit(1)
	}
	i := 0

	for {
		hdr := fmt.Sprintf("%012d", i)
		i++
		fmt.Printf("%s\n", hdr)
		data := make([]byte, 0)
		data = append(data, []byte(hdr)...)
		
		//NC command data
		cmd_data, _ := hex.DecodeString("4e43")
		data = append(data, cmd_data...)

		err := client.Write(data)
		if err != nil {
			net.HandleError(err)
		}
		time.Sleep(time.Millisecond * 5000)
	}

}
