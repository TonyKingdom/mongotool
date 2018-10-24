// Copyright © 2018 TonyKingdom
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/globalsign/mgo"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Map map[string][]string

var dbfile = "database.log"

var missdb = "report.log"

var diffdb = "diff.log"

func Run() {
	mongoUrl := "mongodb://" + Conf.user + ":" + Conf.pass + "@" + Conf.host + ":" + strconv.Itoa(Conf.port)
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	dbs, err := session.DatabaseNames()
	if err != nil {
		panic(err)
	}
	dbmap := make(Map)
	for _, db := range dbs {
		tables, _ := session.DB(db).CollectionNames()
		dbmap[db] = tables
	}

	if _, err := os.Stat(Conf.path); os.IsNotExist(err) {
		os.MkdirAll(Conf.path, 755)
	}

	filename := Conf.path + "/" + dbfile
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println(filename + " not exist, creating with current dbs...")
		ioutil.WriteFile(filename, encodeMap(dbmap).Bytes(), 644)
		return
	}
	tmptbs, _ := ioutil.ReadFile(filename)
	decodedMap := decodeMap(bytes.NewBuffer(tmptbs))
	missmap := diffMap(dbmap, decodedMap)

	ts := time.Now().Format("20060102150405")
	missing := Conf.path + "/" + missdb + "-" + ts
	ioutil.WriteFile(filename+"-"+ts, encodeMap(dbmap).Bytes(), 644)

	ioutil.WriteFile(missing, []byte(missmap.String()), 644)
	files, err := filepath.Glob(Conf.path + "/" + dbfile + "-*")
	if err != nil {
		panic(err)
	}
	rotateFile(files, filecount)
	missfiles, err := filepath.Glob(Conf.path + "/" + missdb + "-*")
	if err != nil {
		panic(err)
	}
	rotateFile(missfiles, filecount)

	ioutil.WriteFile(filename, encodeMap(dbmap).Bytes(), 644)
}

func Diff(file1, file2 string) {
	if len(file1) == 0 || len(file2) == 0 {
		panic("--file1 or --file2 not specified")
	}
	dbfile1, _ := ioutil.ReadFile(file1)
	decodedMap1 := decodeMap(bytes.NewBuffer(dbfile1))
	dbfile2, _ := ioutil.ReadFile(file2)
	decodedMap2 := decodeMap(bytes.NewBuffer(dbfile2))
	missMap := diffMap(decodedMap1, decodedMap2)
	missing := Conf.path + "/" + diffdb
	ioutil.WriteFile(missing, []byte(missMap.String()), 644)
	fmt.Println(missMap)
	fmt.Println("The above diff content also recorded in file: " + missing)
}

func diffMap(mp1, mp2 Map) Map {
	missmap := make(Map)
	for key, value := range mp1 {
		if _, ok := mp2[key]; !ok {
			missmap[key] = nil
		}
		var misstbl []string
		for _, tbl := range value {
			var exist bool
			for _, tb := range mp2[key] {
				if tb == tbl {
					exist = true
					break
				}
			}
			if !exist {
				misstbl = append(misstbl, tbl)
			}
		}
		if len(misstbl) > 0 {
			missmap[key] = misstbl
		}
	}
	return missmap
}

func (ma Map) String() string {
	var buffer bytes.Buffer
	for key, value := range ma {
		buffer.WriteString("Database: ")
		buffer.WriteString(key)
		buffer.WriteString(", Tables: ")
		for _, val := range value {
			buffer.WriteString(val)
			buffer.WriteString(" ")
		}
		buffer.WriteString("\r\n")
	}
	return buffer.String()
}

func encodeMap(mp Map) *bytes.Buffer {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	e.Encode(mp)
	return b
}

func decodeMap(buf *bytes.Buffer) Map {
	var decodedMap Map
	d := gob.NewDecoder(buf)
	err := d.Decode(&decodedMap)
	if err != nil {
		panic(err)
	}
	return decodedMap
}

func rotateFile(files []string, filecount int) []string {
	sort.Slice(files, func(a, b int) bool {
		atime := strings.Split(files[a], "-")
		btime := strings.Split(files[b], "-")
		if len(atime) == 2 && len(btime) == 2 {
			t1, _ := time.Parse("20060102150405", atime[1])
			t2, _ := time.Parse("20060102150405", btime[1])
			return t1.Before(t2)
		}
		return false
	})

	lenth := len(files)
	if lenth > filecount {
		for _, file := range files[:lenth-filecount] {
			os.Remove(file)
		}
		files = files[lenth-filecount:]
	}
	return files
}
