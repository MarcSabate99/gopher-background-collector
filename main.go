package main

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const URL = "https://php-noise.com/noise.php?"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type backgroundColor struct {
	Uri string `json:"uri"`
}

func main() {
	var params = buildParams()
	var endPoint = URL + params
	var backgroundColor = &backgroundColor{}
	err := getJson(endPoint, backgroundColor)
	if err != nil {
		println("Error occurred on request")
		return
	}
	downloadError := downloadFile(backgroundColor.Uri)
	if downloadError != nil {
		println("Error occurred on download file")
		return
	}
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		println(err)
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func buildParams() string {
	rand.Seed(time.Now().UnixNano())
	var red = strconv.Itoa(rand.Intn(255-0) + 255)
	var green = strconv.Itoa(rand.Intn(255-0) + 255)
	var blue = strconv.Itoa(rand.Intn(255-0) + 255)
	var numberOfTiles = strconv.Itoa(rand.Intn(50-1) + 50)
	var tileSize = strconv.Itoa(rand.Intn(20-1) + 20)
	var borderWidth = strconv.Itoa(rand.Intn(15-0) + 15)

	return "r=" + red + "&g=" + green + "&b=" + blue + "&tiles=" + numberOfTiles + "&tileSize=" + tileSize + "&borderWidth=" + borderWidth + "&json"
}

func downloadFile(URL string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	var name = RandStringBytesRandom(5)
	file, err := os.Create("images/" + name + ".png")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func RandStringBytesRandom(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
