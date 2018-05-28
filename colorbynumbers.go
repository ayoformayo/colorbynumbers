package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tr "github.com/ayoformayo/colorbynumbers/proto"
	jsonpb "github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func fetchGeoJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "combined2.geojson")
}
func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/public/index.html")
}

func fetchGTrain(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://datamine.mta.info/mta_esi.php?key=c1fe4f67509bc093ccd9b8f9b73857f6&feed_id=31")
	fmt.Println(resp.Status)
	in, _ := ioutil.ReadAll(resp.Body)
	message := tr.FeedMessage{}
	err := proto.Unmarshal(in, &message)
	if err != nil {
		log.Fatalln("<<<<FAILED TO UNMARSHAL>>>>", err)
	}

	marshaller := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	mess := marshaller.Marshal(w, &message)
	fmt.Println(mess)
	// if er != nil {
	// 	log.Fatalln("<<<<FAILED TO UNMARSHAL>>>>", err)
	// }
	// json.NewEncoder(w).Encode(mess)
	// j, _ := json.Marshal(mess)
	// w.Write(j)
	// fmt.Println("mess=", mess)
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("public/build/index.html")
// 	t.Execute(w)
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
// 	// http.Handle("/", http.FileServer(http.Dir("./public/build")))
// 	http.Handle("/gtrain", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

func main() {
	http.HandleFunc("/", static)
	http.HandleFunc("/gtrain", fetchGTrain)
	http.HandleFunc("/geojson", fetchGeoJSON)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
