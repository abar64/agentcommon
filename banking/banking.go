
ackage banking

import (
	"fmt"
	"strconv"
  "encoding/json"
	"fmt"
)

// TO DO _ READ THESE VALUES FROM SOMEWHERE
var fileIdentifier = "DartC4DFile.exe"
var routingGateway = "656565"

var cd4_record_count = 1

/*
	type  c4dTransactionRecord struct{
		getMsgType,
		reconTransRef,
		branch,
		pannumber,
		routingGateway,
		used_settlement_date,
		recTransDate,
		recTransTime,
		tranType,
		//1.35, //
		transAmount,
	}
*/
func FileC4DHello() string {
	return "FileC4DHello Hello, world."
}

/*
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}*/

func C4D_Header(record []string) {

	fmt.Printf("\"SH\",\"02\",\"%s\",\"%s\",\"%07d\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n",
		"time_of_generation",
		"pbsFileIdentifier",
		69,
		"R_COUNT",
		"DART",
		"APP_NAME",
		"PRODUCT_VERSION",
		"GetFileInFullPath(DART_file)")

}

func C4D_Transaction_Accept(record []string) string {
	//fmt.Printf("%#v\n", record)
	var getMsgType = "C4" //or D
	var branch = record[3-1]
	var tranType = record[7-1]
	var pannumber = record[11-1]
	//	var expDate = record[12-1]
	//var transAmount  = record[13-1]
	var transAmount float64

	transAmount, err := strconv.ParseFloat(record[13-1], 64)
	if err == nil {
		fmt.Printf("") //%.2f\n", transAmount)
	}
	//		fmt.Printf("%.2f\n", transAmount)*/
	var used_settlement_date = "12/34/56"
	var recTransDate = record[21-1][:8]
	var recTransTime = record[21-1][9:]
	var reconTransRef = record[23-1]

	line := fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"20%s\",\"%s\",\"%s\",\"%.2f\",\"%s\",\"%s\",\"%d\"\n",
		getMsgType,
		reconTransRef,
		branch,
		pannumber,
		routingGateway,
		used_settlement_date,
		recTransDate,
		recTransTime,
		tranType,
		//1.35, //
		transAmount,
		"discrepency",
		fileIdentifier,
		cd4_record_count,
	)
	cd4_record_count += 1

	return line
}

func C4D_Transaction_Pending(record []string) {
	//fmt.Printf("%#v\n", record)
	var getMsgType = "C4" //or D
	var branch = record[3-1]
	var tranType = record[7-1]
	var pannumber = record[11-1]
	//	var expDate = record[12-1]
	var transAmount = record[13-1]
	var used_settlement_date = "12/34/56"
	var recTransDate = record[18-1][:8]
	var recTransTime = record[18-1][9:]
	var reconTransRef = record[20-1]
	//	var authMethod = record[23-1]

	fmt.Printf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"20%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%d\"\n",
		getMsgType,
		reconTransRef,
		branch,
		pannumber,
		routingGateway,
		used_settlement_date,
		recTransDate,
		recTransTime,
		tranType,
		transAmount,
		"",
		fileIdentifier,
		cd4_record_count,
	)
	cd4_record_count += 1
}

func C4D_Transaction_Reject(record []string) {
	//	fmt.Printf("%#v\n", record)
	var getMsgType = "C4" //or D
	var branch = record[3-1]
	var tranType = record[7-1]
	var pannumber = record[10-1]
	//	var expDate = record[11-1]
	var transAmount = record[12-1]
	var used_settlement_date = "12/34/56"
	var recTransDate = record[17-1][:8]
	var recTransTime = record[17-1][9:]
	var reconTransRef = record[19-1]
	//	var authMethod = record[22-1]

	fmt.Printf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"20%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%d\"\n",
		getMsgType,
		reconTransRef,
		branch,
		pannumber,
		routingGateway,
		used_settlement_date,
		recTransDate,
		recTransTime,
		tranType,
		transAmount,
		"",
		fileIdentifier,
		cd4_record_count,
	)
	cd4_record_count += 1
}



//var fileIdentifier = "DartC4DFile.exe"
//var routingGateway = "656565"

const (
	Initial_State   int = -1
	Header_Record   int = 00
	Accepted_Record int = 01
	Pending_Record  int = 02
	Rejected_Record int = 03
	Trailer_Record  int = 99
	End_Of_Run      int = 9000
	sequence_Error  int = 9999
)

const (
	Header_Record_Type    = 0
	Header_Data_File_Date = 1
)

const (
	Accepted_Header_Record_Type    = 0
	Accepted_Merchant_ID           = 3
	Accepted_Transaction_Type      = 7
	Accepted_PAN                   = 11
	Accepted_Card_Exp_Date         = 12
	Accepted_Transaction_Amount    = 13
	Accepted_Transaction_Date_time = 21
	Accepted_Payee_Ref_OTR         = 23
	Accepted_Auth_Method           = 28
)

const (
	Pending_Header_Record_Type    = 0
	Pending_Merchant_ID           = 3
	Pending_Transaction_Type      = 7
	Pending_PAN                   = 11
	Pending_Card_Exp_Date         = 12
	Pending_Transaction_Amount    = 13
	Pending_Transaction_Date_time = 18
	Pending_Payee_Ref_OTR         = 20
	Pending_Auth_Method           = 23
)

const (
	Rejected_Header_Record_Type    = 0
	Rejected_Merchant_ID           = 3
	Rejected_Transaction_Type      = 7
	Rejected_PAN                   = 10
	Rejected_Card_Exp_Date         = 11
	Rejected_Transaction_Amount    = 12
	Rejected_Transaction_Date_time = 17
	Rejected_Payee_Ref_OTR         = 19
	Rejected_Auth_Method           = 22
)

//type CsvLine struct {}

// type dataSlice []*DartTransactionLine
type DartTransactionInterface interface {
	MsgType() string
	ReconTransRef() string
	Branch() string
	Pannumber() string
	RoutingGateway() string
	Used_settlement_date() string
	RecTransDate() string
	RecTransTime() string
	TranType() string
	TransAmount() string
	TransDiscrepancy() string
}

type DartTransactionLine struct {
	MsgType              string `json:"msgType"`
	ReconTransRef        string `json:"reconTransRef"`
	Branch               string `json:"branch"`
	Pannumber            string `json:"pannumber"`
	RoutingGateway       string `json:"routingGateway"`
	Used_settlement_date string `json:"used_settlement_date"`
	RecTransDate         string `json:"recTransDate"`
	RecTransTime         string `json:"recTransTime"`
	TranType             string `json:"tranType"`
	TransAmount          string `json:"transAmount"`
	TransDiscrepancy     string `json:"transDiscrepancy"`
}

func StringArraytoDartTransactionJson(data []string) string {

	dartTransaction := DartTransactionLine{}

	//	fmt.Println(data)
	/*
		//	v := reflect.ValueOf(dartTransaction)
		v := reflect.ValueOf(dartTransaction)

		values := make([]interface{}, v.NumField())

		for i := 0; i < v.NumField(); i++ {
			//		values[i] = v.Field(i).Interface()
			values[i] = data[i]
		}

		fmt.Println(values)
		dartTransaction = values*/
	for itemNo, item := range data {
		switch itemNo {
		case 0:
			dartTransaction.MsgType = item
		case 1:
			dartTransaction.ReconTransRef = item
		case 2:
			{
				dartTransaction.Branch = item
			}
		case 5:
			{
				dartTransaction.RoutingGateway = item
			}
		case 6: //7
			dartTransaction.TranType = item
		case 10: //11
			dartTransaction.Pannumber = item
		case 12: //13
			dartTransaction.TransAmount = item
		case 20:
			{ //21 Need to spit item
				newFunction(dartTransaction, item)
				dartTransaction.RecTransTime = item
			}
		case 29: //30
			dartTransaction.Used_settlement_date = item
		case 99:
			dartTransaction.TransDiscrepancy = item
		default:
		}

	}

	a, _ := json.Marshal(dartTransaction)
	//a, _ := json.Marshal(values)
	//fmt.Println(string(a)) // ["foo","bar","baz"]

	return string(a)
}

func newFunction(dartTransaction DartTransactionLine, item string) {
	dartTransaction.RecTransDate = item
}

/*
// Len is part of sort.Interface.

	func (d dataSlice) Len() int {
		println(len(d))
		return len(d)
	}

// Swap is part of sort.Interface.

	func (d dataSlice) Swap(i, j int) {
		println(i, j)
		d[i], d[j] = d[j], d[i]
	}

// Less is part of sort.Interface. We use count as the value to sort by

	func (d dataSlice) Less(i, j int) bool {
		return d[i].branch < d[j].branch
		//return true
	}
*/
var currentState int = Initial_State

func CheckState(recordType int) int {
	currentState = checkState(recordType)
	return currentState
}

func checkState(recordType int) int {
	// Current State
	dartRecordStates := [5][5]int{
		{Accepted_Record, Pending_Record, Rejected_Record, Trailer_Record},
		{Accepted_Record, Pending_Record, Rejected_Record, Trailer_Record},
		{Accepted_Record, Pending_Record, Rejected_Record, Trailer_Record},
		{Trailer_Record, End_Of_Run}}

	switch currentState {
	case Initial_State:
		if recordType == Header_Record {
			return recordType
		}
		return sequence_Error
	case Header_Record:
		for _, num := range dartRecordStates[Header_Record] {
			if num == recordType {
				return recordType
			}
		}
		return sequence_Error
	case Accepted_Record:
		for _, num := range dartRecordStates[Header_Record] {
			if num == recordType {
				return recordType
			}
		}
		return sequence_Error
	case Pending_Record:
		for _, num := range dartRecordStates[Header_Record] {
			if num == recordType {
				return recordType
			}
		}
		return sequence_Error
	case Rejected_Record:
		for _, num := range dartRecordStates[Header_Record] {
			if num == recordType {
				return recordType
			}
		}
		return sequence_Error
	case Trailer_Record:
		for _, num := range dartRecordStates[Header_Record] {
			if num == recordType {
				return recordType
			}
		}
		return sequence_Error
	default:
		fmt.Printf("Don't know type %T\n", currentState)
	}
	return sequence_Error
}
