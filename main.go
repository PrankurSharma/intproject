package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
        "time"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
collection := helper.ConnectDB()
func schedule(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")

	var sched models.Meeting
	_ = json.NewDecoder(r.Body).Decode(&sched)

	result, err := collection.InsertOne(context.TODO(), sched)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}
func meet(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")

	var sched models.Meeting
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&sched)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(sched)
}
func allmeet(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")

	var scheds []models.Meeting
        starttime, _ := mux.Vars(r)["starttime"]
        endtime, _ := mux.Vars(r)["endtime"]
	cur, err := collection.Find(context.TODO(), bson.M{{"_starttime": starttime},{"_endtime":endtime}})

	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		
		var sched models.Meeting
		
		err := cur.Decode(&sched)
		if err != nil {
			log.Fatal(err)
		}

		
		scheds = append(scheds, sched)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(scheds) 
}
func pmeet(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
	var scheds []models.Participants
        email, _ := mux.Vars(r)["email"]

	cur, err := collection.Find(context.TODO(), bson.M{"_email": email})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var sched models.Participants
		err := cur.Decode(&sched) 
		if err != nil {
			log.Fatal(err)
		}

		
		scheds = append(scheds, sched)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(scheds) 
}
func main() {
	
	r := mux.NewRouter()

  	
	r.HandleFunc("/meetings", schedule).Methods("POST")
	r.HandleFunc("/meetings/{id}", meet).Methods("GET")
	r.HandleFunc("/meetings?start={starttime}&end={endtime}", allmeet).Methods("GET")
	r.HandleFunc("/meetings?participant={email}", pmeet).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}
