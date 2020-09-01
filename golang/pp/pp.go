package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"time"
)

var (
	isUDP bool

	listenType string
	listenAddr string
	toAddr     string

	listener net.Listener
)

// data
const (
	MaxReqSize int = 1024
)

func init() {
	flag.BoolVar(&isUDP, "u", false, "-u to use upd")
	flag.StringVar(&listenAddr, "l", "0.0.0.0:9090", "-l listen addr")
	flag.StringVar(&toAddr, "t", "0.0.0.0:9091", "-t to addr")
}

func main() {
	flag.Parse()

	listenType = "tcp"
	if isUDP {
		listenType = "udp"
	}

	var err error
	listener, err = net.Listen(listenType, listenAddr)
	if err != nil {
		log.Fatalln("listen err:", err)
		return
	}
	go accept()

	stop := make(chan bool)
	<-stop
}

func accept() {
	defer func() {
		if rev := recover(); rev != nil {
			log.Println("accept recover", rev)
		}
		go accept()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept err:", conn.RemoteAddr().String(), err)
			continue
		}

		toConn, err := net.DialTimeout(listenType, toAddr, time.Second)
		if err != nil {
			log.Println("dial dst err", err)
			conn.Close()
			return
		}

		go cp(conn, toConn)
		go cp(toConn, conn)
	}
}

func cp(dst net.Conn, src net.Conn) {
	lab := src.RemoteAddr().String() + "->" + dst.RemoteAddr().String()
	log.Println(lab, "cnn")

	defer func() {
		if rev := recover(); rev != nil {
			log.Println(lab, "recover", rev)
		}

	}()

	defer src.Close()
	defer dst.Close()

	for {
		buf, err := ioutil.ReadAll(src)
		if err != nil {
			log.Println(err)
			continue
		}

		if len(buf) > 0 {
			n, err := dst.Write(buf)
			log.Println(lab, n, err, string(buf))
		}

	}
}
