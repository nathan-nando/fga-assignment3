package functions

import (
	"assignment-3/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func JsonReload() {
	for {
		rand.Seed(time.Now().UnixNano())
		water := rand.Intn(100-1) + 1
		wind := rand.Intn(100-1) + 1

		data := model.StatusData{}
		data.Wind = wind
		data.Water = water

		dataJSON, err := json.Marshal(data)

		if err != nil {
			log.Fatal("error occured while marshalling data to json", err.Error())
		}

		err = ioutil.WriteFile("data.json", dataJSON, 0644)

		if err != nil {
			log.Fatal("error occured while writing data to data.json file", err.Error())
		}

		time.Sleep(15 * time.Second)
	}
}

func WebReload(w http.ResponseWriter, r *http.Request) {
	fileData, err := ioutil.ReadFile("data.json")

	if err != nil {
		log.Fatal("error occured while reading data from data.json file", err.Error())
	}

	var statusData model.StatusData

	err = json.Unmarshal(fileData, &statusData)
	if err != nil {
		log.Fatal("error occured while unMarshalling from data.json file", err.Error())
	}

	waterVal := statusData.Water
	windVal := statusData.Wind
	var (
		waterStatus string
		windStatus  string
	)
	waterValue := strconv.Itoa(waterVal)
	windValue := strconv.Itoa(windVal)

	if waterVal <= 5 {
		waterStatus = "Aman"
	} else if waterVal > 5 && waterVal <= 8 {
		waterStatus = "Siaga"
	} else if waterVal > 8 {
		waterStatus = "Bahaya"
	} else {
		waterStatus = "Ketinggian Air Tidak Diketahui"
	}

	if windVal <= 6 {
		windStatus = "Aman"
	} else if windVal > 6 && windVal <= 15 {
		windStatus = "Siaga"
	} else if windVal > 15 {
		windStatus = "Bahaya"
	} else {
		windStatus = "Kecepatan Angin Tidak Diketahui"
	}

	data := map[string]string{
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
		"waterValue":  waterValue,
		"windValue":   windValue,
	}

	tpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal("error occured while parsing index.html file", err.Error())
	}
	tpl.Execute(w, data)
}
