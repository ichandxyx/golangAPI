package main

import (
	"encoding/json"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
)

//ping
func Chek(w http.ResponseWriter,r *http.Request){
	var jinput struct{
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&jinput)
	if jinput.Name==""{
		w.Write([]byte("hey!! call me by my name"))
	}else{
		w.Write([]byte("thanks for calling by my name "))
		fmt.Fprintf(w,"%s\n",jinput.Name)
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
	if err:=
	json.NewDecoder(r.Body).Decode(&jinput)
	err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp{
			Status: "error1",
			Msg:	"no body",
		})
		return
	}
	if jinput.Name==""{
			w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp{
			Status: "error2",
			Msg:	"empty string passed",
		})
		return
	}

		
	pokemon:="https://pokeapi.co/api/v2/pokemon/"+jinput.Name//fmt.Fprintln(w,pokemon)
	ans,err:=http.Get(pokemon)
	if err !=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp{
			Status: "error",
			Msg:	"API error!!",
		})
		return
	}

	body,err2 := ioutil.ReadAll(ans.Body)
	if err2!=nil{
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(resp{
			Status: "error2",
			Msg:	"API error2!!",
		})
		return
	}
	sb:=string(body)
	fmt.Fprintf(w,"%s",sb)
}
func main() {

	router:=http.NewServeMux()
	router.HandleFunc("/find",Search)
	
	router.HandleFunc("/ping",Chek)


	err:=http.ListenAndServe("0.0.0.0:8080", router)
	if err!=nil{
		log.Fatal(err)
	}

	//samople for ./find
	// pass in body {"name":"ditto"}
}