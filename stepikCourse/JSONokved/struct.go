package main

type okved []struct {
	GlobalID       int    `json:"global_id"`
	SystemObjectID string `json:"system_object_id"`
	SignatureDate  string `json:"signature_date"`
	Nomdescr       string `json:"Nomdescr"`
	Razdel         string `json:"Razdel"`
	Kod            string `json:"Kod"`
	Name           string `json:"Name"`
	Idx            string `json:"Idx"`
}
