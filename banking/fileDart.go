package banking

import (
	"encoding/json"
	"fmt"
)

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

func FileDartHello() string {
	return "FileDartHello Hello, world."
}

func StringArraytoDartJson(data []string) string {

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
		case 1:
			dartTransaction.MsgType = item
		case 2:
			dartTransaction.ReconTransRef = item
		case 3:
			dartTransaction.Branch = item
		case 4:
			dartTransaction.Pannumber = item
		case 5:
			dartTransaction.RoutingGateway = item
		case 6:
			dartTransaction.Used_settlement_date = item
		case 7:
			dartTransaction.RecTransDate = item
		case 8:
			dartTransaction.RecTransTime = item
		case 9:
			dartTransaction.TranType = item
		case 10:
			dartTransaction.TransAmount = item
		case 11:
			dartTransaction.TransDiscrepancy = item
		}

	}

	a, _ := json.Marshal(dartTransaction)
	//a, _ := json.Marshal(values)
	//fmt.Println(string(a)) // ["foo","bar","baz"]

	return string(a)
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
func CheckState(currentState int, recordType int) int {
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
