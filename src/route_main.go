package main

import (
	"net/http"
	"data"
)

func index(w http.ResponseWriter,r *http.Request)  {
	threads, err := data.Threads()
	if err != nil{
		error_message(w,r,"Cannot get threads")
	}else{
		_, err := session(w, r)
		if err != nil{
			generateHTML(w, threads, "layout", "public.navbar", "index")
		}else{
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

func err(w http.ResponseWriter, r *http.Request)  {
	vals := r.URL.Query()
	_, e := session(w, r)
	if e != nil{
		generateHTML(w,vals.Get("msg"),"layout", "public.navbar", "error")
	}else{
		generateHTML(w,vals.Get("msg"),"layout", "private.navbar", "error")
	}
}

