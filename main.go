package main

import (
	"encodin/json"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
)

//ping
func Chek(w http.ResponseWriter,r *http.Request){
	name:=r.Form.Get("name")
	if name==""{
		w.Write([]byte("hey!! call me by my name"))

	}else{
	
		w.Write([]byte("thanks for calling by my name"))
	}
	fmt.Fprintf(w,"working")
	
}


//error 
type resp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}


//pokemon api
func Search(w http.ResponseWriter,r *http.Request){
	var jinput struct{
		Name string `json:"name"`
	}
	//input:=r.Form.Get("name")


	if err:=
	json.NewDecoder(r.Body).Decode(&jinput);err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp{
			Status: "error",
			Msg:	"cant leave empty",
		})
		return
	}

	resp,err:=http.Get("https://pokeapi.co/api/v2/pokemon/Name")
	if err !=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp{
			Status: "error",
			Msg:	"API error!!",
		})
		return

	}
	body,err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal(err)
	}

	
	sb:=string(body)
	fmt.Fprintf(w,"%s",sb)





}
func main() {
	router:=http.NewServeMux()
	router.HandleFunc("/find",Search)
	router.HandleFunc("/ping",Chek)


	http.ListenAndServe("0.0.0.0.8080",router)

	
}