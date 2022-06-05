package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}

		githubUser := os.Getenv("GITHUB_USER")
		githubToken := os.Getenv("GITHUB_TOKEN")
		url := fmt.Sprintf("https://api.github.com/users/%s/repos", githubUser)

		client := &http.Client{}

		res, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("Authorization", "token "+githubToken)

		res, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(res.StatusCode)
		w.Write([]byte(res.Status))
	})

	log.Fatal(http.ListenAndServe(":8080", handler))
}
