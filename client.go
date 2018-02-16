package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var connect int

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	fmt.Print("Enter Username: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	username := sc.Text()
	fmt.Fprintln(conn, username)
	println("1. Join 2. Create")
	sc.Scan()
	flag := sc.Text()
	fmt.Fprintln(conn, flag)
	for {
		if flag != "1" {
			fmt.Println("Enter Username To Chat With: ")
			sc = bufio.NewScanner(os.Stdin)
			sc.Scan()
			username = sc.Text()
			fmt.Fprintln(conn, username)
			sc = bufio.NewScanner(conn)
			sc.Scan()
			if sc.Text() == "1" {
				connect = 1
				println("Connection With " + username + " Established")
				break
			}
		} else {
			fmt.Println("Waiting For A Connection To Be Established")
			sc := bufio.NewScanner(conn)
			sc.Scan()
			user := sc.Text()
			connect = 1
			println("Connection With " + user + " Established")
			if flag == "1" {
				fmt.Fprintln(conn, user)
			}
			break
		}
	}
	defer conn.Close()
	for {
		go read(conn)
		if connect == 0 {
			break
		}
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		fmt.Fprintln(conn, text)
	}
	defer conn.Close()
}

func read(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)
		scanner.Scan()
		text := scanner.Text()
		if text == "" {
			fmt.Println("Disconnected From The Server.")
			conn.Close()
			connect = 0
			break
		}
		fmt.Println(text)
	}
}
