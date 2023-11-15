package main
import (
"fmt"  
"net"
"log"
)
func getPortAndIP() {
var serverIP string
var port int
fmt.Println("Input server IP")
  fmt.Scan(serverIP)
  fmt.Println("Input Port")
    fmt.Scan(port)
return port
return serverIP
}
func main() {
  getPortAndIP()
  fmt.Println(serverIP + ":" + port)
}
