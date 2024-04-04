package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/subosito/gotenv"
)

func request(method string, url string, payload interface{}) ([]byte, error) {

	var dataBody bytes.Buffer

	if payload != nil {
		dataByte, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		dataBody = *bytes.NewBuffer(dataByte)
	}

	req, err := http.NewRequest(method, url, &dataBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}

func sendCall(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("http://%v", os.Getenv("PLUTO_HOST"))
	fmt.Println("url", url)
	_, err := request("GET", url, nil)
	if err != nil {
		b, _ := json.Marshal(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(b)
		return
	}

	b, _ := json.Marshal("success")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(b)
}

func main() {

	gotenv.Load()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/send", sendCall)

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}
