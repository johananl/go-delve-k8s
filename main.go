package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", HelloHandler)

	rand.Seed(time.Now().UnixNano())

	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	log.Println("Listening on :8080")
	select {}
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	res, err := data()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Fprintf(w, res)
}

type Data struct {
	ID        uuid.UUID `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Foo       string    `json:"foo"`
	Bar       int       `json:"bar"`
}

func data() (string, error) {
	in := rand.Intn(5)

	foo, err := genRandomString(in)
	if err != nil {
		return "", fmt.Errorf("calculating foo: %w", err)
	}

	d := Data{
		ID:        uuid.New(),
		Timestamp: time.Now().UTC(),
		Foo:       foo,
		Bar:       3,
	}

	j, err := json.Marshal(d)
	if err != nil {
		return "", fmt.Errorf("marshalling response: %w", err)
	}

	return string(j), nil
}

func genRandomString(in int) (string, error) {
	if in == 0 {
		return "", fmt.Errorf("oops")
	}

	var out string

	for i := 0; i < 5; i++ {
		// Generate a random rune in the 97-122 ASCII range (a-z).
		offset := rand.Intn(26)
		c := rune(97 + offset)
		out += string(c)
	}

	return out, nil
}
