package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"z0nix007/service"

	"github.com/boltdb/bolt"
)

const channelSecret = ""
const channelAccessToken = ""

type LineWebhook struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Message    struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

func main() {
	http.HandleFunc("/webhook", handleWebhook)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	db, err := service.OpenDB()
	if err != nil {
		log.Fatal("Error opening BoltDB:", err)
	}
	defer service.CloseDB()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body := &LineWebhook{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Println("Error decoding request:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	commandHandlers := map[string]func(*bolt.DB, []string, string,string){
		"/quecar":    handleQuecar,
		"/getquecar": handleGetQuecar,
		"/getfrom":   handleGetFrom,
		"/getTo":     handleGetTo,
		"/getcarType":handleGetCarType,
		"/from":      handleFrom,
		"/to":        handleTo,
		"/carType":   handleCarType,
	}

	for _, event := range body.Events {
		fmt.Println(event.Message.Text)

		msg := strings.Split(event.Message.Text, " ")
		text := event.Message.Text
		if handler, exists := commandHandlers[msg[0]]; exists {
			handler(db, msg[1:], event.ReplyToken,text)
		} else {
			replyMessage(event.ReplyToken, "Unknown command: "+msg[0])
		}
	}
}

func handleQuecar(db *bolt.DB, args []string, replyToken string, text string) {
	if len(args) != 1 {
		replyMessage(replyToken, "Invalid command. Usage: /quecar <value>")
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		replyMessage(replyToken, "Invalid input. Please provide a valid number.")
		return
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("myBucket"))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}
		valueBytes := make([]byte, 8) // 8 bytes to store an int64
		binary.BigEndian.PutUint64(valueBytes, uint64(num))
		err = bucket.Put([]byte("quecarCout"), valueBytes)
		return err
	})

	if err != nil {
		log.Fatal("Error updating value in BoltDB:", err)
	}
	replyMessage(replyToken, "Success to "+args[0])
}

func handleCarType(db *bolt.DB, args []string, replyToken string, text string) {
	if len(args) != 1 {
		replyMessage(replyToken, "Invalid command. Usage: /carType <value>")
		return
	}
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("myBucket"))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}

		err := bucket.Put([]byte("carType"), []byte(args[0]))
		return err
	})

	if err != nil {
		log.Fatal("Error updating value in BoltDB:", err)
	}
	replyMessage(replyToken, "Success to "+args[0])
}


func handleGetCarType(db *bolt.DB, args []string, replyToken string, text string) {
	carType, err := service.GetValueFromDB(db, "carType")
	if err != nil {
		log.Fatal("Error updating value in BoltDB:", err)
	}
	replyMessage(replyToken, "Success to "+string(carType))
}


func handleGetQuecar(db *bolt.DB, args []string, replyToken string, text string) {
	QuecarCout, err := service.GetValueFromDB(db, "quecarCout")
	if err != nil {
		replyMessage(replyToken, "Failed to retrieve value from DB.")
		return
	}
	quecarCout := int(binary.BigEndian.Uint64(QuecarCout))
	str := strconv.Itoa(quecarCout)
	fmt.Println(str)
	replyMessage(replyToken, str)
}

func handleFrom(db *bolt.DB, args []string, replyToken string, text string) {

	if len(args) == 0 {
		replyMessage(replyToken, "Invalid command. Usage: /from <value>")
		return
	}
	result := strings.Replace(text, "/from ", "", 1)
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("myBucket"))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}

		err := bucket.Put([]byte("from"), []byte(result))
		return err
	})

	if err != nil {
		log.Fatal("Error updating value in BoltDB:", err)
	}

	replyMessage(replyToken, "Success to "+result)
}

func handleGetFrom(db *bolt.DB, args []string, replyToken string, text string) {
	From, err := service.GetValueFromDB(db, "from")
	if err != nil {
		log.Fatal("Error updating value in BoltDB:", err)
	}
	replyMessage(replyToken, "Success to "+ string(From))
}


func handleTo(db *bolt.DB, args []string, replyToken string,text string) {
	if len(args) == 0 {
		replyMessage(replyToken, "Invalid command. Usage: /to <value>")
		return
	}
	result := strings.Replace(text, "/to ", "", 1)
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("myBucket"))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}

		err := bucket.Put([]byte("to"), []byte(result))
		return err
	})

	if err != nil {
		log.Fatal("Error updating value in BoltDB:", err)
	}
	replyMessage(replyToken, "Success to "+result)
}


func handleGetTo(db *bolt.DB, args []string, replyToken string,text string) {
	From, err := service.GetValueFromDB(db, "to")
	if err != nil {
		log.Fatal("Error updating value in BoltDB:", err)
	}
	replyMessage(replyToken, "Success to "+string(From))
}


func replyMessage(replyToken, text string) {
	url := "https://api.line.me/v2/bot/message/reply"
	data := map[string]interface{}{
		"replyToken": replyToken,
		"messages": []map[string]interface{}{
			{
				"type": "text",
				"text": text,
			},
		},
	}

	client := &http.Client{}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error encoding reply:", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+channelAccessToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending reply:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK status code received:", resp.StatusCode)
	}
}
