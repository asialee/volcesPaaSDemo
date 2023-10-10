package main

import (
    "fmt"
    "net/http"
//    "strings"
    "log"
    "io/ioutil"
    "net"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)


func readFile() (string){
    f, err := ioutil.ReadFile("output")
    if err != nil {
        return "read fail"
    }
    return string(f)
}

func RemoteIp(req *http.Request) string {
    remoteAddr := req.RemoteAddr
    if ip := req.Header.Get(XRealIP); ip != "" {
        remoteAddr = ip
    } else if ip = req.Header.Get(XForwardedFor); ip != "" {
        remoteAddr = ip
    } else {
        remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
    }

    if remoteAddr == "::1" {
        remoteAddr = "127.0.0.1"
    }

    return remoteAddr
}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
    // r.ParseForm()  // 解析参数，默认是不会解析的
    // fmt.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
    // fmt.Println("path", r.URL.Path)
    // fmt.Println("scheme", r.URL.Scheme)
    // fmt.Println(r.Form["url_long"])
    // for k, v := range r.Form {
    //     fmt.Println("key:", k)
    //     fmt.Println("val:", strings.Join(v, ""))
    // }
    ip := RemoteIp(r)
    //fmt.Fprintf(w, ip)
    fmt.Fprintf(w, "client ip:",ip) // 这个写入到 w 的是输出到客户端的
}

func main() {
    http.HandleFunc("/", sayhelloName) // 设置访问的路由
    err := http.ListenAndServe(":9090", nil) // 设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}