package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/boltdb/bolt"
)

type CheckResponse struct {
	Success bool `json:"success"`
}

type UserInfo struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Timestamp string `json:"timestamp"`
}

func TestApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	test := make(map[string]bool)
	test["success"] = true

	j, _ := json.Marshal(test)
	w.Write(j)
}

func CheckApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var userInfo *UserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		fmt.Println(err)
	}

	successChance := rand.Intn(101)
	var checkResponse bool

	if successChance <= 80 {
		checkResponse = true
	} else {
		checkResponse = false
	}

	response := &CheckResponse{
		Success: checkResponse,
	}

	j, _ := json.Marshal(response)
	w.Write(j)
	postData(userInfo)
}

func postData(userInfo *UserInfo) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, openDbErr := bolt.Open("users.db", 0600, nil)
	if openDbErr != nil {
		fmt.Println("ERR: ", openDbErr)
	}
	defer db.Close()

	// Store the user model in the user bucket using the username as the key.
	updateErr := db.Update(func(tx *bolt.Tx) error {
		b, insertErr := tx.CreateBucketIfNotExists([]byte("Users"))
		if insertErr != nil {
			return insertErr
		}

		encoded, encodeErr := json.Marshal(userInfo)
		if encodeErr != nil {
			return encodeErr
		}

		return b.Put([]byte(userInfo.Username), encoded)
	})

	if updateErr != nil {
		fmt.Println("ERR:", updateErr)
	}
}
