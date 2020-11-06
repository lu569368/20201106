package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type XTFSSearchAPIRequest struct {
	DirPath                string         `json:"dirPath"`
	FsName                 string         `json:"fsName"`
	Constraint             bool           `json:"constraint"`
	Recursive              bool           `json:"recursive"`
	ShowMeta               bool           `json:"showMeta"`
	ShowOwn                bool           `json:"showOwn"`
	RawMode                bool           `json:"rawMode"`
	FetchId                bool           `json:"fetchId"`
	TimeStamp              bool           `json:"timeStamp"`
	Rank                   string         `json:"rank"`
	Ascending              bool           `json:"ascending"`
	ExportDataset          string         `json:"exportDataset"`
	ExportDatasetType      string         `json:"exportDatasetType"`
	ExportDatasetAuthRead  string         `json:"exportDatasetAuthRead"`
	ExportDatasetAuthWrite string         `json:"exportDatasetAuthWrite"`
	Silent                 bool           `json:"silent"`
	FsExpress              XTFSExpress    `json:"fsExpress"`
	MetasExpress           XTMetasExpress `json:"metasExpress"`
	Cursor                 int            `json:"cursor"`
	Limit                  int            `json:"limit"`
	Handler                string         `json:"handler"`
	AlgoType               string
}

type XTFSExpress struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	User        string `json:"user"`
	Group       string `json:"group"`
	Mode        string `json:"mode"`
	SizeGreater string `json:"sizeGreater"`
	SizeLess    string `json:"sizeLess"`
	CTimeStart  string `json:"cTimeStart"`
	CTimeEnd    string `json:"cTimeEnd"`
	MTimeStart  string `json:"mTimeStart"`
	MTimeEnd    string `json:"mTimeEnd"`
}

type XTMetasExpress struct {
	AndOr          bool            `json:"andOr"`
	Metas          []XTFSMetaEntry `json:"metas"`
	ConstraintMeta bool            `json:"constraintMeta"`
	TTimeStart     string          `json:"tTimeStart"`
	TTimeEnd       string          `json:"tTimeEnd"`
}

type XTFSMetaEntry struct {
	Key           string `json:"key"`
	Value         string `json:"value"`
	MatchAllValue bool   `json:"matchAllValue"`
}

var jsonString1 string = `{
 "dirPath": "/dir1-1",
 "fsName": "xtao",
 "constraint": false,
 "recursive": true,
 "showMeta": false,
 "showOwn": false,
 "rawMode": false,
 "fetchId": false,
 "timeStamp": false,
 "rank": "",
 "ascending": true,
 "exportDataset": "",
 "exportDatasetType": "",
 "exportDatasetAuthRead": "",
 "exportDatasetAuthWrite": "",
 "silent": false,
 "fsExpress": {
  "type": "",
  "name": "",
  "user": "",
  "group": "",
  "mode": "",
  "sizeGreater": "0",
  "sizeLess": "3",
  "cTimeStart": "",
  "cTimeEnd": "",
  "mTimeStart": "",
  "mTimeEnd": ""
 },
 "metasExpress": {
  "andOr": false,
  "metas": null,
  "constraintMeta": true,
  "tTimeStart": "",
  "tTimeEnd": ""
 },
 "cursor": 0,
 "limit": 1000,
 "handler": "",
 "AlgoType": ""
}`
var jsonString2 string = `{
 "ascending": true,
 "constraint": false,
 "cursor": 0,
 "dirPath": "/dir1-1",
 "exportDataset": "",
 "exportDatasetAuthRead": "",
 "exportDatasetAuthWrite": "",
 "exportDatasetType": "",
 "fetchId": true,
 "fsExpress": {
  "cTimeEnd": "",
  "cTimeStart": "",
  "group": "",
  "mTimeEnd": "",
  "mTimeStart": "",
  "mode": "",
  "name": "",
  "sizeGreater": 0,
  "sizeLess": -1,
  "type": "",
  "user": ""
 },
 "fsName": "xtao",
 "handler": "ceshi",
 "limit": 59,
 "metasExpress": {
  "andOr": false,
  "constraintMeta": false,
  "metas": {},
  "tTimeEnd": "",
  "tTimeStart": ""
 },
 "path": "/",
 "rank": "",
 "rawMode": true,
 "recursive": false,
 "showMeta": true,
 "showOwn": true,
 "silent ": true,
 "timeStamp": ""
}`

func Decode(r io.Reader) (u *XTFSSearchAPIRequest, err error) {
	u = new(XTFSSearchAPIRequest)
	err = json.NewDecoder(r).Decode(u)
	if err != nil {
		return
	}
	return
}

func main() {
	user, err := Decode(strings.NewReader(jsonString1))
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatalln(err)
	} else {
		fmt.Printf("%#v\n", user)
	}

	fmt.Println("----------------------------")
	user2, err := Decode(strings.NewReader(jsonString2))
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatalln(err)
	} else {
		fmt.Printf("%#v\n", user2)
	}

}
