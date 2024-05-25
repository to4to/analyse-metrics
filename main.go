package main

import "net/http"



func main(){



	router:=http.NewServeMux()

router.HandleFunc("POST /data",handleData)
	http.ListenAndServe(":8080",router)
}

func handleData(w http.ResponseWriter, r *http.Request){


	

}