package main

import (
	"fmt"
	"time"
)

// TCP Connection

// func main() {
// 	fmt.Println("Hello World")
// 	listener, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer listener.Close()

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		go handleConnection(conn)
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()
// }

// DNS lookup

// func main() {
// 	fmt.Println("Hello World")
// 	ips, err := net.LookupHost("wisc.edu")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ips_split := strings.Split(ips[0], ".")
// 	fmt.Println(ips_split)
// 	// ipv4 verification, each segement should be between 0 and 255 and the total length should be 4
// 	if len(ips_split) != 4 {
// 		fmt.Println("Not a valid address")
// 	}
// 	for _, v := range ips_split {
// 		int_v, err := strconv.Atoi(v)
// 		if err != nil {
// 			fmt.Println("Not valid")
// 			break
// 		}
// 		if !(int_v <= 255) && !(int_v >= 0) {
// 			fmt.Println("Not valid")
// 			break
// 		}
// 	}
// 	fmt.Println("True")
// }

// API Call
// type jokes struct {
// 	Type      string `json:"type"`
// 	Setup     string `json:"setup"`
// 	Punchline string `json:"punchline"`
// 	Id        int    `json:"id"`
// }

// func main() {
// 	fmt.Println("Hello World!")
// 	baseUrl := "https://official-joke-api.appspot.com/jokes/ten"
// 	response, err := http.Get(baseUrl)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(response.Status)
// 	jokeObj := []jokes{}
// 	if response.StatusCode == 200 {
// 		contents, err := ioutil.ReadAll(response.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		err = json.Unmarshal(contents, &jokeObj)
// 		if err != nil {
// 			fmt.Println("Error unmarshalling", err)
// 		}
// 	}
// 	sort.Slice(jokeObj, func(i, j int) bool {
// 		return jokeObj[i].Id < jokeObj[j].Id
// 	})
// 	for _, value := range jokeObj {
// 		fmt.Println(value)
// 	}
// }

// Socket connection

// defining node structure
// type Node struct {
// 	name     string
// 	incoming chan string
// 	outgoing []net.Conn
// }

// //constructing new node
// func NewNode(name string) *Node {
// 	return &Node{
// 		name:     name,
// 		incoming: make(chan string),
// 		outgoing: make([]net.Conn, 0),
// 	}
// }

// // adding new outgoing channel
// func (node *Node) addNewConnection(conn net.Conn) {
// 	node.outgoing = append(node.outgoing, conn)
// }

// // listening to any data
// func (node *Node) startListening() {
// 	for {
// 		select {
// 		case data := <-node.incoming:
// 			fmt.Println(data)
// 		}
// 	}
// }

// // fowarding data to all outgoing node
// func (node *Node) fowardData(data string) {
// 	for _, v := range node.outgoing {
// 		_, err := v.Write([]byte(data))
// 		if err != nil {
// 			fmt.Println("error forwarding", err)
// 		}
// 	}
// }

// func main() {
// 	fmt.Println("Starting...")

// 	nodeA := NewNode("A")
// 	nodeB := NewNode("B")
// 	nodeC := NewNode("C")

// 	go nodeA.startListening()
// 	go nodeB.startListening()
// 	go nodeC.startListening()

// 	connectionAB, err := net.Dial("tcp", "localhost:8080")
// 	if err != nil {
// 		fmt.Println("Error making connection to A->B", err)
// 	}
// 	connectionAC, err := net.Dial("tcp", "localhost:8081")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	nodeA.addNewConnection(connectionAB)
// 	nodeA.addNewConnection(connectionAC)

// 	nodeA.incoming <- "Hello World"

// 	select {}
// }

// LRU Cache
// type Node struct {
// 	key  int
// 	val  int
// 	next *Node
// 	prev *Node
// }

// type LRUCache struct {
// 	capacity int
// 	lru_map  map[int]*Node
// 	left     *Node
// 	right    *Node
// }

// func Constructor(capacity int) LRUCache {
// 	left_fake := &Node{val: -1, next: nil, prev: nil}
// 	right_fake := &Node{val: -1, next: nil, prev: left_fake}
// 	left_fake.next = right_fake
// 	return LRUCache{
// 		capacity: capacity,
// 		lru_map:  make(map[int]*Node, capacity),
// 		left:     left_fake,
// 		right:    right_fake,
// 	}

// }

// func (this *LRUCache) remove(node *Node) {
// 	prev := node.prev
// 	nxt := node.next
// 	prev.next = nxt
// 	nxt.prev = prev
// }

// func (this *LRUCache) insert(node *Node) {
// 	new_prev := this.right.prev
// 	new_nxt := this.right
// 	new_prev.next = node
// 	new_nxt.prev = node
// 	node.next = new_nxt
// 	node.prev = new_prev
// }

// func (this *LRUCache) Get(key int) int {
// 	if _, ok := this.lru_map[key]; !ok {
// 		return -1
// 	}

// 	this.remove(this.lru_map[key])
// 	this.insert(this.lru_map[key])
// 	return this.lru_map[key].val
// }

// func (this *LRUCache) Put(key int, value int) {
// 	if _, ok := this.lru_map[key]; ok {
// 		this.remove(this.lru_map[key])
// 	}
// 	temp := &Node{
// 		key: key,
// 		val: value,
// 	}
// 	this.lru_map[key] = temp
// 	this.insert(this.lru_map[key])
// 	length := len(this.lru_map)
// 	if length > this.capacity {
// 		lru := this.left.next
// 		this.remove(lru) // remove the left most node
// 		delete(this.lru_map, lru.key)
// 	}
// }

// func main() {
// 	capacity := 2
// 	obj := Constructor(capacity)

// 	// Test case operations
// 	obj.Put(1, 1)
// 	obj.Put(2, 2)
// 	fmt.Println(obj.Get(1)) // Output: 1
// 	obj.Put(3, 3)
// 	fmt.Println(obj.Get(2)) // Output: -1
// 	obj.Put(4, 4)
// 	fmt.Println(obj.Get(1)) // Output: -1
// 	fmt.Println(obj.Get(3)) // Output: 3
// 	fmt.Println(obj.Get(4)) // Output: 4

// }

// Time keeping LRU leetcode

// type rateLimiter struct {
// 	packetCount       int
// 	lastSecondStart   int64
// 	lastFiveSecond    [5]int
// 	lastFiveSecondPtr int
// }

// func Constructor() *rateLimiter {
// 	return &rateLimiter{}
// }

// func (r1 *rateLimiter) isAllowed(timeStamp int64) bool {
// 	now := time.Now().Unix()

// 	if now != r1.lastSecondStart {
// 		r1.lastSecondStart = now
// 		r1.packetCount = 1
// 	} else {
// 		r1.packetCount++
// 	}

// 	r1.lastFiveSecondPtr = int(now % 5)
// 	r1.lastFiveSecond[r1.lastFiveSecondPtr] = r1.packetCount

// 	if r1.packetCount <= 3 {
// 		return true
// 	}

// 	sum := 0
// 	for _, count := range r1.lastFiveSecond {
// 		sum += count
// 	}

// 	return sum <= 10
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	r1 := Constructor()

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		timestamp, err := strconv.ParseInt(line, 10, 64)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}

// 		if r1.isAllowed(timestamp) {
// 			fmt.Println("a")
// 		} else {
// 			fmt.Println("d")
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading input:", err)
// 		os.Exit(1)
// 	}
// }

// Time based key-value store

// type Pair struct {
// 	val       string
// 	timeStamp int
// }

// type TimeMap struct {
// 	timeMap map[string][]Pair
// }

// func Constructor() TimeMap {
// 	return TimeMap{
// 		timeMap: make(map[string][]Pair, 0),
// 	}
// }

// func (this *TimeMap) set(key string, value string, timestamp int) string {
// 	pair := Pair{
// 		val:       value,
// 		timeStamp: timestamp,
// 	}
// 	if _, ok := this.timeMap[key]; !ok {
// 		this.timeMap[key] = []Pair{}
// 	}
// 	this.timeMap[key] = append(this.timeMap[key], pair)
// 	return ""
// }

// func (this *TimeMap) get(key string, timestamp int) string {
// 	if _, ok := this.timeMap[key]; !ok {
// 		return ""
// 	}
// 	arr := this.timeMap[key]
// 	if arr[0].timeStamp > timestamp {
// 		return ""
// 	}
// 	left, right := 0, len(arr)-1
// 	for left < right {
// 		mid := int((left + right) / 2)
// 		if arr[mid].timeStamp == timestamp {
// 			return arr[mid].val
// 		} else if arr[mid].timeStamp < timestamp {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	if arr[left].timeStamp > timestamp {
// 		return arr[left-1].val
// 	}
// 	return arr[left].val
// }

// func main() {
// 	fmt.Println("Hello World")
// 	obj := Constructor()
// 	fmt.Println(obj.set("foo", "bar", 1))  // store the key "foo" and value "bar" along with timestamp = 1.
// 	fmt.Println(obj.get("foo", 1))         // return "bar"
// 	fmt.Println(obj.get("foo", 3))         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
// 	fmt.Println(obj.set("foo", "bar2", 4)) // store the key "foo" and value "bar2" along with timestamp = 4.
// 	fmt.Println(obj.get("foo", 4))         // return "bar2"
// 	fmt.Println(obj.get("foo", 5))         // return "bar2"
// }

// implementing a stack

// func main() {

// 	requests := make(chan int, 5)
// 	for i := 1; i <= 5; i++ {
// 		requests <- i
// 	}
// 	close(requests)

// 	limiter := time.Tick(200 * time.Millisecond)

// 	for req := range requests {
// 		<-limiter
// 		fmt.Println("request", req, time.Now())
// 	}

// 	burstyLimiter := make(chan time.Time, 3)

// 	for i := 0; i < 3; i++ {
// 		burstyLimiter <- time.Now()
// 	}

// 	go func() {
// 		for t := range time.Tick(200 * time.Millisecond) {
// 			burstyLimiter <- t
// 		}
// 	}()

// 	burstyRequests := make(chan int, 5)
// 	for i := 1; i <= 5; i++ {
// 		burstyRequests <- i
// 	}
// 	close(burstyRequests)
// 	for req := range burstyRequests {
// 		<-burstyLimiter
// 		fmt.Println("request", req, time.Now())
// 	}
// }

// Rate Limiter

// type messageBody struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// }

// func requestHandle(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("content-type", "application/json")
// 	writer.WriteHeader(http.StatusOK)
// 	message := messageBody{
// 		Status:  "Successful",
// 		Message: "Hello World!",
// 	}

// 	err := json.NewEncoder(writer).Encode(&message)
// 	if err != nil {
// 		return
// 	}

// }

// func main() {

// }

// package main

// import (
//     "fmt"
//     "time"
// )

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick((1000 / 3) * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	select {}

}
