package main

import (
    "fmt"
    "net"
	"os"
	"encoding/csv"
	"io"
	"log"
	"encoding/json"
	"strings"

)

type alldata struct
{
	Date string `json:"date"`
	Positive string `json:"positive"`
	Tests string `json:"tests"`
	Expired string `json:"expired"`
	Admitted string `json:"admitted"`
	Discharged string `json:"discharged"`
	Region string `json:"region"`
}

// type Qry struct
// {
// 	Q string `json:"query"`
// 	D string `json:"date"`
// 	R string `json:"region"`
// }

const (
    CONN_HOST = "localhost"
    CONN_PORT = "4040"
    CONN_TYPE = "tcp"
)

func main(){
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
 
for{  

  conn.Write([]byte("\nPlease enter your choice (date or region)\n"))
  buf1 := make([]byte, 1024)

  reqLen, err := conn.Read(buf1)
  if err != nil {
    fmt.Println("Error reading:", err.Error(), reqLen)
  }

  conn.Write([]byte("\nPlease enter your value for (date or region)\n"))
  buf2 := make([]byte, 1024)
  
  reqLen2, err2 := conn.Read(buf2)
  if err2 != nil {
    fmt.Println("Error reading:", err2.Error(), reqLen2)
  }

//   str := string(buf)
//   fmt.Println(str)
// 	 obj := Qry{}
//   json.Unmarshal([]byte(str), &obj)
//   fmt.Printf("Query: %s", obj.R)

// var ch string = string(buf1)
// var val string = string(buf2)

	// a := query(ch, val)
a := query("region", "sindh")
  conn.Write([]byte("\n" + a))
  
}
}


func query(Choice string, Value string) string{
			// Open the file
	csvfile, err := os.Open("covid_final_data.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	var data_arr []alldata
	
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data := alldata{
			Date: record[4],
			Positive: record[2],
			Tests: record[3],
			Expired: record[6],
			Admitted: record[10],
			Discharged: record[5],
			Region: record[9],
		}
		
		
		// if (strings.ToLower(Choice) == "region"){
			// fmt.Println(Choice, Value)
			if strings.ToLower(Value)==strings.ToLower(data.Region){
				data_arr = append(data_arr,data)
			
			}
		// }

		// if (strings.ToLower(Choice) == "date"){
			if strings.ToLower(Value) == strings.ToLower(data.Date) {
				data_arr = append(data_arr,data)
			}
		// }
	
	}
	jsonData, _ := json.Marshal(data_arr)
	return string(jsonData)
	// return Choice+Value
}
