package main

import "time"

type Ticket struct {
	Id           string    `bson:"_id"`
	Type         string    `bson:"type"`
	StartTime    time.Time `bson:"start_time"`
	EndTime      time.Time `bson:"end_time"`
	ServiceAgent string    `bson:"service_agent_id"`
	Slot         string    `bson:"slot_id"`
}

type ServiceAgent struct {
	Id     string `bson:"_id"`
	Name   string `bson:"name"`
	Status string `bson:"status"`
}

type Slot struct {
	Id           string    `bson:"_id"`
	Time         time.Time `bson:"time"`
	Status       string    `bson:"status"`
	ServiceAgent string    `bson:"service_agent_id"`
}
