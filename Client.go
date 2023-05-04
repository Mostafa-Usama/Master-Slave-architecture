package main

import (
	"fmt"
	"net"
	"os"
)

var IPs [4]string

func main() {

	// Connect to the master server
	conn, err := net.Dial("tcp", "192.168.43.124:9090") // Replace with the master's IP address
	if err != nil {
		fmt.Println(err)
		return
	} else {
		for i := 0; i < 4; i++ {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)

			}
			IPs[i] = string(buf[:n])
			fmt.Println(IPs[i])
		}

	}

	defer conn.Close()
	fmt.Println("Enter file the name you need: ")
		var file_name string
		fmt.Scan(&file_name)
	var result = ""
	for i := 0; i < 4; i++ {
		conn, err := net.Dial("tcp", IPs[i]+":9090") 
		if err != nil {
			fmt.Println(err)
			return
		}
		
		_,err = conn.Write([]byte(file_name))
		if err != nil{
			fmt.Println(err)
			return
		}
	
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		result += (string(buf[:n]))
	}

	f, err := os.Create("data.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(result)

	if err2 != nil {
		fmt.Println(err2)
	}

}
