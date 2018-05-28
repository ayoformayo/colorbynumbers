package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	tr "github.com/ayoformayo/colorbynumbers/proto"
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
	message := &tr.FeedMessage{}
	err := proto.Unmarshal(in, message)
	if err != nil {
		log.Fatalln("<<<<FAILED TO UNMARSHAL>>>>", err)
	}
	entity := message.GetEntity()
	// fmt.Println(message)
	now := uint64(time.Now().Unix())
	for _, status := range entity {
		status.GetTripUpdate()
		timestamp := status.Vehicle.GetTimestamp()
		update := status.GetTripUpdate()
		trip := update.GetTrip()
		stopTimeUpdate := update.GetStopTimeUpdate()
		// fmt.Println("update=", update)
		fmt.Println("stopTimeUpdate=", stopTimeUpdate)
		fmt.Println("trip=", trip)
		fmt.Printf("timeDiff = %d\n", now-timestamp)
	}
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
