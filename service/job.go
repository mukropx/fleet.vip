package service

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/valyala/fasthttp"
	"golang.org/x/exp/slices"
)

type APIResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	TID     string       `json:"tid"`
	Data    ResponseData `json:"data"`
}

type ResponseData struct {
	Items      []OrderItem `json:"items"`
	Pagination Pagination  `json:"pagination"`
}

type OrderItem struct {
	BaseInfo BaseInfo   `json:"base_info"`
	SiteInfo []SiteInfo `json:"site_info"`
}

type BaseInfo struct {
	SerialNo        string `json:"serial_no"`
	Track           string `json:"track"`
	Area            string `json:"area"`
	CarType         int    `json:"car_type"`
	CarTypeText     string `json:"car_type_text"`
	RunningMileage  string `json:"running_mileage"`
	StopTime        int64  `json:"stop_time"`
	FinishForTime   string `json:"finish_for_time"`
	PlanArrivedTime int64  `json:"plan_arrived_time"`
	Price           int    `json:"price"`
}

type SiteInfo struct {
	OrderNo           int     `json:"order_no"`
	StoreName         string  `json:"store_name"`
	Address           string  `json:"address"`
	RunningMileage    *string `json:"running_mileage"`
	PlanArrivedTime   int64   `json:"plan_arrived_time"`
	PlanDepartureTime *int64  `json:"plan_departure_time"`
	Lat               float64 `json:"lat"`
	Lng               float64 `json:"lng"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	TotalCount  int `json:"totalCount"`
}

type GrabJonRespo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Tid     string `json:"tid"`
	Data    string `json:"data"`
}



func GetJob() {
	db, err := OpenDB()
	if err != nil {
		log.Fatal("Error opening BoltDB:", err)
	}
	defer db.Close()
	From, err := GetValueFromDB(db, "from")
	To, err := GetValueFromDB(db, "to")
	CarType, err := GetValueFromDB(db, "carType")
	dataArrayFrom := strings.Split(string(From), ",")
	dataArrayTo := strings.Split(string(To), ",")
	dataCarType := strings.Split(string(CarType), ",")
        fmt.Println(dataArrayFrom,dataArrayTo,dataCarType)
	url := "https://gwapi.fleet.vip/gw/fms/grab_order/list"
	method := "POST"
	pageSize := 100
	pageNum := 1
	var allResponses []APIResponse

	for {
		payload := fmt.Sprintf(`{
			"area": "",
			"car_type": "",
			"page_size": %d,
			"page_num": %d
		}`, pageSize, pageNum)

		req := fasthttp.AcquireRequest()
		req.SetRequestURI(url)
		req.Header.SetMethod(method)
		req.Header.SetUserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82")
		req.Header.Add("Fms-Session-Id", Session)
		req.Header.Add("Content-Type", "application/json")
		req.SetBodyString(payload)

		resp := fasthttp.AcquireResponse()
		client := &fasthttp.Client{}
		err := client.Do(req, resp)
		if err != nil {
			fmt.Println(err)
			return
		}

		body := resp.Body()
		var response APIResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		allResponses = append(allResponses, response)

		// ตรวจสอบว่ายังมีหน้าถัดไปหรือไม่
		if len(response.Data.Items) < pageSize {
			break
		} else {
			pageNum++
		}

		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}

	for _, response := range allResponses {
		fmt.Println("Status Code:", response.Code)
		fmt.Println("Message:", response.Message)
		fmt.Println("Current Page:", response.Data.Pagination.CurrentPage)
		fmt.Println("Per Page:", response.Data.Pagination.PerPage)
		fmt.Println("Total Count:", response.Data.Pagination.TotalCount)
		for _, item := range response.Data.Items {
			fmt.Println(item)
			fmt.Println(slices.Contains(dataArrayFrom, item.SiteInfo[0].StoreName) && slices.Contains(dataArrayTo, item.SiteInfo[1].StoreName))
			if slices.Contains(dataCarType, item.BaseInfo.CarTypeText) {
				if slices.Contains(dataArrayFrom, item.SiteInfo[0].StoreName) && slices.Contains(dataArrayTo, item.SiteInfo[1].StoreName) {
					if item.BaseInfo.CarType == 100 {
						GrabJob(100, 21521, 90, item.BaseInfo.SerialNo, "กรุงเทพ")
					}
					if item.BaseInfo.CarType == 101 {
						GrabJob(101, 3089, 90, item.BaseInfo.SerialNo, "กรุงเทพ")
					}
					if item.BaseInfo.CarType == 200 {
						GrabJob(200, 132454, 90, item.BaseInfo.SerialNo, "กรุงเทพ")
					}
					if item.BaseInfo.CarType == 203 {
						GrabJob(203, 139605, 90, item.BaseInfo.SerialNo, "กรุงเทพ")
					}
				}
			}
		}
	}
}

func GrabJob(carType int, carId int, driverId int, serialNo string, province string) {
	QuecarCout, err := GetValueFromDB(db, "quecarCout")
	if err != nil {
		return
	}
	quecarCout := int(binary.BigEndian.Uint64(QuecarCout))
	if quecarCout == 0 {
		return
	}
	url := "https://gwapi.fleet.vip/gw/fms/grab_order/grab"
	lineType := 1
	payload := fmt.Sprintf(`{"car_type": %d, "car_id": %d, "driver_id": %d, "serial_no": "%s", "line_type": %d, "province": "%s"}`,
		carType, carId, driverId, serialNo, lineType, province)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json;charset=UTF-8")
	req.Header.Set("Host", "gwapi.fleet.vip")
	req.Header.Set("Fms-Session-Id", Session)
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.5790.110 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Origin", "https://supplier-th.fleet.vip")
	req.Header.Set("Referer", "https://supplier-th.fleet.vip/")

	req.SetRequestURI(url)
	req.SetBodyString(payload)

	client := &fasthttp.Client{}
	err = client.Do(req, resp)
	if err != nil {
		log.Fatal("Error sending HTTP request:", err)
	}
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	body := resp.Body()

	var response GrabJonRespo
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	fmt.Println("Code:", response.Code)
	fmt.Println("Message:", response.Message)
	fmt.Println("Tid:", response.Tid)
	fmt.Println("Data:", response.Data)
	fmt.Println("Response Body:", string(body))
	SendLineNotifyMessage(string(body))
	if response.Message == "success" {
		err = db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("myBucket"))
			if bucket == nil {
				return fmt.Errorf("Bucket not found")
			}
			valueBytes := make([]byte, 8) // 8 bytes to store an int64
			binary.BigEndian.PutUint64(valueBytes, uint64(quecarCout-1))
			err = bucket.Put([]byte("quecarCout"), valueBytes)
			return err
		})

		if err != nil {
			log.Fatal("Error updating value in BoltDB:", err)
		}
	}
}
