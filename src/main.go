package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/ini.v1"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln("Fail happen,", err)
		os.Exit(1)
	}
}

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
	CheckErr(err)
	address := cfg.Section("address")
	ip, err := address.GetKey("ip")
	CheckErr(err)
	port, err := address.GetKey("port")
	CheckErr(err)

	log.Println("config of ip:", ip.String(), ", port:", port.String())
	connaddr := ip.String() + ":" + port.String()
	log.Println("http address:", connaddr)

	//http
	http.HandleFunc("/", ServeHTTP)
	log.Fatal(http.ListenAndServe(connaddr, nil))
}
