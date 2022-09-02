package banking

import (
	"fmt"
	"strconv"
)

var cd4_record_count = 1

func FileC4DVer() string {
	return "FileC4D v0.0.0.1"
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
	}
*/
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

/*		switch rec {
		case Header_Record:
			//			fmt.Println("Header Record")
			C4D_Header(record)
		case Accepted_Record:
			//			fmt.Println("Accepted Record")
			C4D_Transaction_Accept(record)
		case Pending_Record:
			C4D_Transaction_Pending(record)
		case Rejected_Record:
			C4D_Transaction_Reject(record)
		case Trailer_Record:
			//C4D_Transaction_Trailer(record)
		default:
			fmt.Printf("Don't know type %T\n", record[0])
		}
	}
}
*/
