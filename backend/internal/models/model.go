package models

type Item struct {
    ID   string `json:"id" dynamodbav:"ID"`
    Name string `json:"name" dynamodbav:"Name"`
}