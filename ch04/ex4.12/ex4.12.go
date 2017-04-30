package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const indexFile = "index.json"
const urlFormat = "https://xkcd.com/%d/info.0.json"

type Index struct {
	Infos    map[int]*Info
	FilePath string
}

type Info struct {
	Title      string
	Transcript string
	ImgURL     string `json:"img"`
}

func newInfoFromURL(comicURL string) (*Info, error) {
	resp, err := http.Get(comicURL)

	if resp.StatusCode != http.StatusOK || err != nil {
		resp.Body.Close()
		return nil, err
	}

	var info Info
	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &info, nil
}

func newIndex(filePath string) *Index {
	var infos map[int]*Info

	out, err := ioutil.ReadFile(filePath)
	if err != nil {
		infos = make(map[int]*Info)
	} else {
		json.Unmarshal(out, &infos)
	}

	return &Index{
		Infos:    infos,
		FilePath: filePath,
	}
}

func (i *Index) build(fromID, toID int) {
	newInfos := make(map[int]*Info)
	for j := fromID; j < toID; j++ {
		if _, exists := i.Infos[j]; exists {
			continue
		}
		newInfo, err := newInfoFromURL(fmt.Sprintf(urlFormat, j))
		if err != nil {
			continue
		}
		newInfos[j] = newInfo
	}
	i.addInfos(newInfos)
}

func (i *Index) addInfos(infos map[int]*Info) {
	for id := range infos {
		i.Infos[id] = infos[id]
	}
}

func (i *Index) save() {
	out, err := json.MarshalIndent(i.Infos, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile(i.FilePath, out, 0644)
}

func (i *Index) search(query string) []*Info {
	var foundInfos []*Info
	for _, info := range i.Infos {
		if strings.Contains(info.Title, query) || strings.Contains(info.Transcript, query) {
			foundInfos = append(foundInfos, info)
		}
	}
	return foundInfos
}

var fromID = flag.Int("from", 0, "id from which to build index")
var toID = flag.Int("to", 01, "id to which to build index")
var search = flag.String("search", "", "search for transcript and title")

func main() {
	flag.Parse()
	index := newIndex(indexFile)

	if *fromID != 0 && *toID != 0 {
		index.build(*fromID, *toID)
		index.save()
	}

	if *search != "" {
		foundInfos := index.search(*search)
		for _, foundInfo := range foundInfos {
			fmt.Printf("URL: %v\nTranscript: %v\n\n", foundInfo.ImgURL, foundInfo.Transcript)
		}
	}
}
