package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat:",version(),"started at",config.Address)

	//多路复用器
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))

    //处理器
    mux.Handle("/static/",http.StripPrefix("/static/",files))

    //处理器函数  index
	mux.HandleFunc("/",index)

	mux.HandleFunc("/err",err)

	mux.HandleFunc("/login",login)

	mux.HandleFunc("/logout",logout)

	mux.HandleFunc("/signup",signup)

	mux.HandleFunc("/signup_account", signupAccount)

	mux.HandleFunc("/authenticate", authenticate)


	mux.HandleFunc("/thread/new", newThread)

	mux.HandleFunc("/thread/create", createThread)

	mux.HandleFunc("/thread/post", postThread)

	mux.HandleFunc("/thread/read", readThread)


    //绑定端口
	server := &http.Server{
      Addr:config.Address,
      Handler:mux,
      ReadTimeout : time.Duration(config.ReadTimeout  * int64(time.Second)),
	  WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
	  MaxHeaderBytes: 1 << 20,
	}

	//监听
	server.ListenAndServe()
}


