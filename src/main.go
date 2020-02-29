package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/ini.v1"
)

func PrintRequest(req *http.Request) {
	log.Println(req.RemoteAddr, "connected")
	log.Println(req.Method, req.RequestURI)
	log.Println(req.Header)
	req.ParseForm()
	log.Println(req.Form)
}

//http://127.0.0.1:8000/test?name=gongluck&age=18
func ServeHTTP(respon http.ResponseWriter, req *http.Request) {
	PrintRequest(req)
	responstr := "You are requesting " + req.RequestURI
	respon.Write([]byte(responstr))
}

func main() {
	// config
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Fatalln("Fail to read config file, ", err)
		os.Exit(1)
	}
	address := cfg.Section("address")
	if address == nil {
		log.Fatalln("Fail to read address section")
		os.Exit(2)
	}
	ip, err := address.GetKey("ip")
	if err != nil {
		log.Fatalln("Fail to get ip config, ", err)
		os.Exit(3)
	}
	port, err := address.GetKey("port")
	if err != nil {
		log.Fatalln("Fail to get port config, ", err)
		os.Exit(4)
	}

	log.Println("config of ip:", ip.String(), ", port:", port.String())
	connaddr := ip.String() + ":" + port.String()
	log.Println("http address:", connaddr)

	//http
	http.HandleFunc("/", ServeHTTP)
	log.Fatal(http.ListenAndServe(connaddr, nil))
}
