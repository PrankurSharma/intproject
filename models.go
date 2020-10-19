package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import "time"

type Meeting struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
        Participants    *Participants  `json:"Participants,omitempty" bson:"participants,omitempty"`
        Starttime  time.Time  `json:"_starttime" bson:"_starttime"`
        Endtime  time.Time  `json:"_endtime" bson:"_endtime"`
}

type Participants struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"_email,omitempty" bson:"_email,omitempty"`
        RSVP string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}