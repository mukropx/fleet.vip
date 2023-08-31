package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("./db/my.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("myBucket"))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}

		bucket.ForEach(func(key, value []byte) error {
			fmt.Printf("Key: %s, Value: %s\n", key, value)
			return nil
		})
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}


// package main

// import (
// 	"encoding/binary"
// 	"fmt"
// 	"log"

// 	"github.com/boltdb/bolt"
// )

// func main() {
// 	db, err := bolt.Open("./db/my.db", 0666, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Start a read-write transaction
// 	err = db.Update(func(tx *bolt.Tx) error {
// 		// Create or get the existing bucket named "myBucket"
// 		bucket, err := tx.CreateBucketIfNotExists([]byte("myBucket"))
// 		if err != nil {
// 			return fmt.Errorf("Error creating bucket: %s", err)
// 		}
// 		fmt.Println("Bucket 'myBucket' created or found successfully.")

// 		// Update the value associated with the key "quecarCout" to be an int
// 		myInt := 1
// 		valueBytes := make([]byte, 8) // 8 bytes to store an int64
// 		binary.BigEndian.PutUint64(valueBytes, uint64(myInt))
// 		err = bucket.Put([]byte("quecarCout"), valueBytes)
// 		if err != nil {
// 			return err
// 		}

// 		// Update the values associated with the keys "from" and "to"
// 		err = bucket.Put([]byte("from"), []byte("2SRG_BDC-สำโรง,5SRG_PDC-สำโรง,NMD_SP-หนามแดง,SHOP IMPRELIUM SUMLHONG,SHOP THE MASTER UDOMSUK,SHOP YAKE SI UDOM,SRG_SP-สำโรง,2PKS_BDC-แพรกษา,MPK_SP-เมืองแพรกษา,PKM_SP-แพรกษาใหม่,PKS_SP-แพรกษา,SHOP NAM DEANG,SHOP NIKOM BANGPU PRAKSA,2BMN_BDC-บางเมืองใหม่,2BPU_BDC-บางปู,BMG_SP-บางเมือง,BMN_SP-บางเมืองใหม่,BPU_SP-บางปู,PNM_SP-ปากน้ำ,SHOP PTT THEPHARAK,TBM_SP-ท้ายบ้านใหม่,TIB_SP-ท้ายบ้าน,TPR_SP-เทพารักษ์"))
// 		if err != nil {
// 			return err
// 		}
// 		err = bucket.Put([]byte("to"), []byte("05 Las_hub-ลาซาล,05 LAS_HUB-ลาซาล,21 BPL_BHUB-บางพลี,77 SCB_HUB-บางพลี"))
// 		if err != nil {
// 			return err
// 		}
// 		err = bucket.Put([]byte("carType"), []byte("4W"))
// 		if err != nil {
// 			return err
// 		}

// 		quecarCoutBytes := bucket.Get([]byte("quecarCout"))
// 		if quecarCoutBytes == nil {
// 			// Handle the case when the key does not exist
// 			return fmt.Errorf("Key not found")
// 		}

// 		// Convert the byte array to an int64
// 		quecarCout := int(binary.BigEndian.Uint64(quecarCoutBytes))

// 		// Print the result
// 		fmt.Printf("quecarCout: %d\n", quecarCout)

// 		// Print all keys and values in the bucket
// 		err = bucket.ForEach(func(key, value []byte) error {
// 			fmt.Printf("Key: %s, Value: %q\n", key, value)
// 			return nil
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }


// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"github.com/gorilla/websocket"
// )

// func main() {
// 	// URL ของ WebSocket server
// 	url := "wss://msg-api-th.fleet.vip/connection/websocket"

// 	// สร้างการเชื่อมต่อ WebSocket
// 	c, _, err := websocket.DefaultDialer.Dial(url, nil)
// 	if err != nil {
// 		log.Fatal("Error connecting to WebSocket:", err)
// 	}
// 	defer c.Close()

// 	// ส่งข้อความครั้งแรกทันทีเมื่อเปิดการเชื่อมต่อ
// 	// connectMessage := []byte(`{"connect":{"token":"eyJUeXBlIjoiSnd0IiwidHlwIjoiSldUIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMjYxNCIsImV4cCI6MTY5MTMwNzM1OH0.5YO7RYRV1nkN17NkUtdySwXHVtg2T19MTKS2qVXOR7k","name":"js"},"id":1} {"subscribe":{"channel":"fms_grab_order"},"id":2}
// 	// `)
// 	// err = c.WriteMessage(websocket.TextMessage, connectMessage)
// 	// if err != nil {
// 	// 	log.Println("Error sending 'connect' message:", err)
// 	// 	return
// 	// }
// 	connectData := map[string]interface{}{
// 		"connect": map[string]string{
// 			"token": "eyJUeXBlIjoiSnd0IiwidHlwIjoiSldUIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMjYxNCIsImV4cCI6MTY5MTMwODg5Mn0.yA0wwO80m4EnU2jq6XxQv-iqj-g60Z2zxgs3GtZzKTU",
// 			"name":  "js",
// 		},
// 		"id": 1,
// 	}

// 	subscribeData := map[string]interface{}{
// 		"subscribe": map[string]string{
// 			"channel": "fms_grab_order",
// 		},
// 		"id": 2,
// 	}

// 	// แปลงข้อมูล JSON object เป็น JSON string
// 	connectMessage, err := json.Marshal(connectData)
// 	if err != nil {
// 		log.Println("Error marshaling 'connect' message:", err)
// 		return
// 	}

// 	subscribeMessage, err := json.Marshal(subscribeData)
// 	if err != nil {
// 		log.Println("Error marshaling 'subscribe' message:", err)
// 		return
// 	}

// 	// ส่งข้อความที่คุณสร้างไปยังเซิร์ฟเวอร์ WebSocket
// 	err = c.WriteMessage(websocket.TextMessage, connectMessage)
// 	if err != nil {
// 		log.Println("Error sending 'connect' message:", err)
// 		return
// 	}

// 	err = c.WriteMessage(websocket.TextMessage, subscribeMessage)
// 	if err != nil {
// 		log.Println("Error sending 'subscribe' message:", err)
// 		return
// 	}

// 	// อ่านข้อความที่ได้รับจาก WebSocket server
// 	for {
// 		_, message, err := c.ReadMessage()
// 		if err != nil {
// 			log.Println("Error reading message:", err)

// 			// สร้างการเชื่อมต่อใหม่ทันทีเมื่อเกิด error
// 			// reconnect(c)
// 			return
// 		}

// 		// แสดงข้อความที่ได้รับจาก WebSocket server
// 		fmt.Printf("Received message: %s\n", message)

// 		// ตรวจสอบว่าข้อความที่ได้รับเป็น {} ให้ทำการส่ง {} กลับไป
// 		if string(message) == "{}" {
// 			responseMessage := []byte(`{}`)
// 			err := c.WriteMessage(websocket.TextMessage, responseMessage)
// 			if err != nil {
// 				log.Println("Error sending response message:", err)
// 				return
// 			}
// 		}
// 	}
// }

// func readMessages(c *websocket.Conn) {
// 	for {
// 		_, message, err := c.ReadMessage()
// 		if err != nil {
// 			log.Println("Error reading message:", err)

// 			// สร้างการเชื่อมต่อใหม่ทันทีเมื่อเกิด error
// 			// reconnect(c)
// 			return
// 		}

// 		// แสดงข้อความที่ได้รับจาก WebSocket server
// 		fmt.Printf("Received message: %s\n", message)

// 		// ตรวจสอบว่าข้อความที่ได้รับเป็น {} ให้ทำการส่ง {} กลับไป
// 		if string(message) == "{}" {
// 			responseMessage := []byte(`{}`)
// 			err := c.WriteMessage(websocket.TextMessage, responseMessage)
// 			if err != nil {
// 				log.Println("Error sending response message:", err)
// 				return
// 			}
// 		}
// 	}
// }

// func reconnect(c *websocket.Conn) {
// 	for {
// 		log.Println("WebSocket connection closed. Reconnecting...")

// 		// รอเป็นเวลา 1 นาทีก่อนที่จะเชื่อมต่อใหม่
// 		time.Sleep(1 * time.Minute)

// 		// ทำการเชื่อมต่อใหม่
// 		newConn, _, err := websocket.DefaultDialer.Dial("wss://msg-api-th.fleet.vip/connection/websocket", nil)
// 		if err == nil {
// 			log.Println("Reconnected to WebSocket server.")
// 			// ปิดการเชื่อมต่อเก่า
// 			c.Close()
// 			// หยุดลูปการเชื่อมต่อใหม่เมื่อเชื่อมต่อสำเร็จ
// 			break
// 		}
// 		log.Println("Error reconnecting to WebSocket:", err)
// 	}

// 	// เริ่มต้นอ่านข้อความใหม่จากการเชื่อมต่อใหม่
// 	readMessages(c)
// }

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	url := "https://gwapi.fleet.vip/gw/fms/grab_order/grab"
// 	jsonData := `{"car_type":100,"car_id":135866,"driver_id":20601,"serial_no":"VA2023072900140000003","line_type":1,"province":"กรุงเทพ"}`

// 	reqBody := bytes.NewBuffer([]byte(jsonData))
// 	req, err := http.NewRequest("POST", url, reqBody)
// 	if err != nil {
// 		fmt.Println("Error creating HTTP request:", err)
// 		return
// 	}

// 	req.Header.Set("Host", "gwapi.fleet.vip")
// 	req.Header.Set("Content-Length", "135")
// 	req.Header.Set("Sec-Ch-Ua", "")
// 	req.Header.Set("Fms-Session-Id", "1691493598_3dbee3c9c0d8c36d7718c0386639bfe07c924f6264929f6d49de8b613cd44ae2_12616")
// 	req.Header.Set("Accept-Language", "en-US")
// 	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.5790.110 Safari/537.36")
// 	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
// 	req.Header.Set("Accept", "application/json, text/plain, */*")
// 	req.Header.Set("Cache-Control", "no-cache")
// 	req.Header.Set("Sec-Ch-Ua-Platform", "\"\"")
// 	req.Header.Set("Origin", "https://supplier-th.fleet.vip")
// 	req.Header.Set("Sec-Fetch-Site", "same-site")
// 	req.Header.Set("Sec-Fetch-Mode", "cors")
// 	req.Header.Set("Sec-Fetch-Dest", "empty")
// 	req.Header.Set("Referer", "https://supplier-th.fleet.vip/")
// 	req.Header.Set("Accept-Encoding", "gzip, deflate")
// 	req.Header.Set("Connection", "close")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending HTTP request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	var body bytes.Buffer
// 	_, err = body.ReadFrom(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return
// 	}

// 	fmt.Println("Response Status Code:", resp.Status)
// 	fmt.Println("Response Body:", body.String())
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// type APIResponse struct {
// 	Code    int          `json:"code"`
// 	Message string       `json:"message"`
// 	TID     string       `json:"tid"`
// 	Data    ResponseData `json:"data"`
// }

// type ResponseData struct {
// 	Items      []OrderItem `json:"items"`
// 	Pagination Pagination  `json:"pagination"`
// }

// type OrderItem struct {
// 	BaseInfo BaseInfo   `json:"base_info"`
// 	SiteInfo []SiteInfo `json:"site_info"`
// }

// type BaseInfo struct {
// 	SerialNo        string `json:"serial_no"`
// 	Track           string `json:"track"`
// 	Area            string `json:"area"`
// 	CarType         int    `json:"car_type"`
// 	CarTypeText     string `json:"car_type_text"`
// 	RunningMileage  string `json:"running_mileage"`
// 	StopTime        int64  `json:"stop_time"`
// 	FinishForTime   string `json:"finish_for_time"`
// 	PlanArrivedTime int64  `json:"plan_arrived_time"`
// 	Price           int    `json:"price"`
// }

// type SiteInfo struct {
// 	OrderNo           int     `json:"order_no"`
// 	StoreName         string  `json:"store_name"`
// 	Address           string  `json:"address"`
// 	RunningMileage    *string `json:"running_mileage"`
// 	PlanArrivedTime   int64   `json:"plan_arrived_time"`
// 	PlanDepartureTime *int64  `json:"plan_departure_time"`
// 	Lat               float64 `json:"lat"`
// 	Lng               float64 `json:"lng"`
// }

// type Pagination struct {
// 	CurrentPage int `json:"currentPage"`
// 	PerPage     int `json:"perPage"`
// 	TotalCount  int `json:"totalCount"`
// }

// func main() {
// 	method := "POST"
// 	url := "https://gwapi.fleet.vip/gw/fms/grab_order/list"
// 	seesion := "1691576357_21868d7852e4009b5439e2107e4236cd4ca274caae3117dd7e5cbc10b9fd3d99_12614"

// 	pageSize := 100
// 	pageNum := 1
// 	var allResponses []APIResponse

// 	client := &http.Client{}

// 	for {
// 		payload := strings.NewReader(fmt.Sprintf(`{
// 			"area": "",
// 			"car_type": "",
// 			"page_size": %d,
// 			"page_num": %d
// 		}`, pageSize, pageNum))

// 		req, err := http.NewRequest(method, url, payload)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82")
// 		req.Header.Add("Fms-Session-Id", seesion)
// 		req.Header.Add("Content-Type", "application/json")

// 		res, err := client.Do(req)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		defer res.Body.Close()

// 		body, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		var response APIResponse
// 		err = json.Unmarshal(body, &response)
// 		if err != nil {
// 			fmt.Println("Error decoding JSON:", err)
// 			return
// 		}

// 		allResponses = append(allResponses, response)

// 		// ตรวจสอบว่ายังมีหน้าถัดไปหรือไม่
// 		if response.Data.Pagination.PerPage < pageSize {
// 			break
// 		} else {
// 			pageNum++
// 		}
// 	}
// 	fmt.Println(allResponses)
// 	// ทำอะไรกับข้อมูลที่รวมมาทั้งหมด
// 	// ...
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// type APIResponse struct {
// 	Code    int          `json:"code"`
// 	Message string       `json:"message"`
// 	TID     string       `json:"tid"`
// 	Data    ResponseData `json:"data"`
// }

// type ResponseData struct {
// 	Items      []OrderItem `json:"items"`
// 	Pagination Pagination  `json:"pagination"`
// }

// type OrderItem struct {
// 	BaseInfo BaseInfo   `json:"base_info"`
// 	SiteInfo []SiteInfo `json:"site_info"`
// }

// type BaseInfo struct {
// 	SerialNo        string `json:"serial_no"`
// 	Track           string `json:"track"`
// 	Area            string `json:"area"`
// 	CarType         int    `json:"car_type"`
// 	CarTypeText     string `json:"car_type_text"`
// 	RunningMileage  string `json:"running_mileage"`
// 	StopTime        int64  `json:"stop_time"`
// 	FinishForTime   string `json:"finish_for_time"`
// 	PlanArrivedTime int64  `json:"plan_arrived_time"`
// 	Price           int    `json:"price"`
// }

// type SiteInfo struct {
// 	OrderNo           int     `json:"order_no"`
// 	StoreName         string  `json:"store_name"`
// 	Address           string  `json:"address"`
// 	RunningMileage    *string `json:"running_mileage"`
// 	PlanArrivedTime   int64   `json:"plan_arrived_time"`
// 	PlanDepartureTime *int64  `json:"plan_departure_time"`
// 	Lat               float64 `json:"lat"`
// 	Lng               float64 `json:"lng"`
// }

// type Pagination struct {
// 	CurrentPage int `json:"currentPage"`
// 	PerPage     int `json:"perPage"`
// 	TotalCount  int `json:"totalCount"`
// }
// func isValidPagination(pagination Pagination) bool {
// 	if pagination.CurrentPage <= 0 || pagination.PerPage <= 0 || pagination.TotalCount < 0 {
// 		return false
// 	}

// 	// ตรวจสอบว่าหน้าปัจจุบันคือหน้าที่มีข้อมูลอยู่ โดยไม่เกิน totalCount
// 	return (pagination.CurrentPage-1)*pagination.PerPage < pagination.TotalCount
// }

// func main() {
// 	url := "https://gwapi.fleet.vip/gw/fms/grab_order/list"
// 	method := "POST"
// 	session := "1691578412_911473d92be17493b9f0cac3093c4a3535402167aefd46aaec7a2259511090c2_12614" // แทนที่ด้วย Session ID จริง

// 	pageSize := 100
// 	pageNum := 1
// 	var allResponses []APIResponse

// 	client := &http.Client{}

// 	for {
// 		payload := strings.NewReader(fmt.Sprintf(`{
// 			"area": "",
// 			"car_type": "",
// 			"page_size": %d,
// 			"page_num": %d
// 		}`, pageSize, pageNum))

// 		req, err := http.NewRequest(method, url, payload)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82")
// 		req.Header.Add("Fms-Session-Id", session)
// 		req.Header.Add("Content-Type", "application/json")

// 		res, err := client.Do(req)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		defer res.Body.Close()

// 		body, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		var response APIResponse
// 		err = json.Unmarshal(body, &response)
// 		if err != nil {
// 			fmt.Println("Error decoding JSON:", err)
// 			return
// 		}

// 		allResponses = append(allResponses, response)

// 		// ตรวจสอบว่ายังมีหน้าถัดไปหรือไม่
// 		if len(response.Data.Items) < pageSize {
// 			break
// 		} else {
// 			pageNum++
// 		}
// 	}
// 	// แสดงข้อมูลทั้งหมดที่ได้รับมา
// 	for _, response := range allResponses {
// 		fmt.Println("Current Page:", response.Data.Pagination.CurrentPage)
// 		fmt.Println("Per Page:", response.Data.Pagination.PerPage)
// 		fmt.Println("Total Count:", response.Data.Pagination.TotalCount)
// 		fmt.Println("---")
// 	}
// }
