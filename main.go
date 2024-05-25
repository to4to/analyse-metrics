package main

import (
	"fmt"
	"io"
	"net/http"
)



func main(){



	router:=http.NewServeMux()

router.HandleFunc("POST /data",handleData)
	http.ListenAndServe(":8080",router)
}

func handleData(w http.ResponseWriter, r *http.Request){


	b,err:=io.ReadAll(r.Body)


	if err!=nil{
		panic(err)
	}

	defer r.Body.Close()
	fmt.Println(string(b))

}