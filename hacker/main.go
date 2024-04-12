package main

import (
	"fmt"
	"log"
	"net/http"
	"net/netip"
	"strconv"
)

// prefill data
// type ip_data struct {
// 	User_id  int
// 	IpAdrres string
// }

var userData map[int]string

var listenAddr = "localhost:8000"

func handleRequests(writer http.ResponseWriter, request *http.Request) {
	user_id := request.URL.Query().Get("user_id")
	ip_addres := request.URL.Query().Get("edge_ip")
	fmt.Println(user_id, ip_addres)

	user_id_num, err := strconv.Atoi(user_id)
	if err != nil {
		fmt.Printf(err.Error()) // json encoder bad request 400 request invalid
	}

	if _, ok := userData[user_id_num]; !ok {
		fmt.Printf("User ID invalid") // json encode not found 404 or 400
	}

	// arr := strings.Split(ip_addres, ".")

	ip_address_from_data := userData[user_id_num] // 192.168.0.0 - 192.168.255.255
	ip_addre_d_cidr, err := netip.ParsePrefix(ip_address_from_data)
	if err != nil {
		fmt.Println(err) // 500 internal server bug
	}
	ip_addres_byte, errParse := netip.ParseAddr(ip_addres)
	if errParse != nil {
		fmt.Println(err) // 400 bad request
	}

	if ip_addre_d_cidr.Contains(ip_addres_byte) {
		fmt.Println("Can transmit") // 200 OK message : can transmit
	} else {
		fmt.Println("Cannot Transmit") // 403 Forbidden : cannot transmit
	}

}

func main() {
	log.Printf("listening on http://%s\n", listenAddr)
	userData = make(map[int]string)
	userData[123] = "192.168.0.0/24"
	userData[456] = "1.0.0.0/24"
	userData[666] = "fe80:c001::/48"

	http.HandleFunc("/v1/edge-ip-acl", handleRequests)

	log.Fatal(http.ListenAndServe(listenAddr, nil))

}
