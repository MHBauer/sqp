package main

import (
	"encoding/json"
	"testing"
)

func TestJson(t *testing.T) {

	rawjson0 := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000,"thing":{"more":"stuff"}}`)
	rawjson1 := []byte(`{"Name":"Bob","Body":"Hello","Time":1294706395881547000,"thing":{"more":"stuff"}}`)
	rawjson2 := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000,"thing":{"more":"stuff"}}`)

	var parsed0 map[string]interface{}
	var parsed1 map[string]interface{}
	var parsed2 map[string]interface{}

	json.Unmarshal(rawjson0, &parsed0)
	json.Unmarshal(rawjson1, &parsed1)
	json.Unmarshal(rawjson2, &parsed2)

	if isDifferent(parsed0, parsed0) {
		t.Fatal("same should not be different")
	}
	if isDifferent(parsed1, parsed1) {
		t.Fatal("same should not be different")
	}
	if isDifferent(parsed2, parsed2) {
		t.Fatal("same should not be different")
	}
	// different args
	if !isDifferent(parsed0, parsed1) {
		t.Fatal("different should not be different")
	}
	if !isDifferent(parsed1, parsed2) {
		t.Fatal("different should not be different")
	}
	// these are actually the same
	if isDifferent(parsed2, parsed0) {
		t.Fatal("different should not be different")
	}
}
