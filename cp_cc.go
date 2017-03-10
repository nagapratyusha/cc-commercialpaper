/*
Copyright 2016 IBM

Licensed under the Apache License, Version 2.0 (the "License")
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Licensed Materials - Property of IBM
© Copyright IBM Corp. 2016
*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var quotePrefix = "qt:"
var propertyPrefix = "pt:"
var proposalPrefix = "pr:"
var agreementPrefix = "ag:"
var deedPrefix = "de:"
var notificationPrefix = "de:"


var cpPrefix = "cp:"
var accountPrefix = "acct:"
var accountsKey = "accounts"

var recentLeapYear = 2016

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func generateCUSIPSuffix(issueDate string, days int) (string, error) {

	t, err := msToTime(issueDate)
	if err != nil {
		return "", err
	}

	maturityDate := t.AddDate(0, 0, days)
	month := int(maturityDate.Month())
	day := maturityDate.Day()

	suffix := seventhDigit[month] + eigthDigit[day]
	return suffix, nil

}

const (
	millisPerSecond     = int64(time.Second / time.Millisecond)
	nanosPerMillisecond = int64(time.Millisecond / time.Nanosecond)
)

func msToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(msInt/millisPerSecond,
		(msInt%millisPerSecond)*nanosPerMillisecond), nil
}

type Owner struct {
	Company  string `json:"company"`
	Quantity int    `json:"quantity"`
}


type Notification struct {
	NotificationId string `json:"notificationId"`
	Parameter1  string  `json:"parameter1"`
	Parameter2  string  `json:"parameter2"`
	Parameter3  string  `json:"parameter3"`
	Parameter4  string  `json:"parameter4"`
	Parameter5  string  `json:"parameter5"`
	Parameter6  string  `json:"parameter6"`
	Parameter7  string  `json:"parameter7"`
	Parameter8  string  `json:"parameter8"`
	Parameter9  string  `json:"parameter9"`
	Parameter10  string  `json:"parameter10"`
	
}

type Quote struct {
	QuoteNo     string  `json:"quoteno"`
	Item    string  `json:"item"`
	Qty    string  `json:"qty"`
	ShipTerm    string  `json:"shipterm"`
	ShipDate    string  `json:"shipdate"`
	ItemDetails    string  `json:"itemdetails"`
	Status    string  `json:"status"`
	Issuer    string  `json:"issuer"`
	IssueDate string  `json:"issueDate"`
	ModifiedOn    string  `json:"modifiedon"`
	RequesterOrg    string  `json:"requesterorg"`
	Country    string  `json:"country"`
	Parameter1 	string  `json:"parameter1"`
	Parameter2  string   `json:"parameter2"`
	Parameter3  string   `json:"parameter3"`
	Parameter4  string   `json:"parameter4"`
	Parameter5  string   `json:"parameter5"`
}


type Property struct {
	PropId     string  `json:"propid"`
	PropOwner    string  `json:"owner"`
	Tax    string  `json:"tax"`
	PropType    string  `json:"proptype"`
	Measure    string  `json:"mesaure"`
	MesaureDisp    string  `json:"mesauredisp"`
	Address    string  `json:"address"`
	Location    string  `json:"location"`
	Latitude string  `json:"latitude"`
	Longitude    string  `json:"longitude"`
	Histories    []History `json:"history"`  
	Litigations  []Litigation  `json:"litigations"`
	Parameter1 string  `json:"parameter1"`
	Parameter2 string  `json:"parameter2"`
	Parameter3 string  `json:"parameter3"`
	Parameter4 string  `json:"parameter4"`
	Parameter5 string  `json:"parameter5"`
	Parameter6 string  `json:"parameter6"`
}

type Proposal struct {
	ProposalNo string  `json:"proposalNo"`
	PropId     string  `json:"propid"`
	ProposedBy    string  `json:"proposedby"`
	ProposedPrice    string  `json:"proposedprice"`
	ProposedDate    string  `json:"proposeddate"`
	Parameter1 string  `json:"parameter1"`
	Parameter2 string  `json:"parameter2"`
	Parameter3 string  `json:"parameter3"`
	Parameter4 string  `json:"parameter4"`
	Parameter5 string  `json:"parameter5"`
	Parameter6 string  `json:"parameter6"`
}


type SaleAgreement struct {
	AgreementNo string  `json:"agreementno"`
	PropId     string  `json:"propid"`
	Parties []Party `json:"buyer"`
	Loan Loan `json:"loan"`
	Parameter1 string  `json:"parameter1"`
	Parameter2 string  `json:"parameter2"`
	Parameter3 string  `json:"parameter3"`
	Parameter4 string  `json:"parameter4"`
	Parameter5 string  `json:"parameter5"`
	Parameter6 string  `json:"parameter6"`
	SignedOn string  `json:"signedon"`
}

type SaleDeed struct {
	DeedNo string  `json:"deedno"`
	AgreementNo string  `json:"agreementno"`
	Registrar Registrar `json:"registrar"`
	Settlement []Settlement `json:"settlement"`
	SignedOn string  `json:"signedon"`
}

type History struct {
	HistoryOwner     string  `json:"owner"`
	Location    string  `json:"location"`
	From       string `json:"from"`
	To       string     `json:"to"`
}

type Litigation struct {
	Data     string  `json:"data"`
}

type Party struct {
	PartyName     string  `json:"name"`
	PartyDOB    string  `json:"dob"`
	PartyIDType    string  `json:"idtype"`
	PartyIDNumber    string  `json:"idnumber"`
	PartyAddress    string  `json:"address"`
	PartyType    string `json:"type"`
}

type Loan struct {
	Bank     string  `json:"bank"`
	Branch    string  `json:"branch"`
	LoanAmount    string  `json:"amount"`
	LoanType    string  `json:"type"`
	LoanPercentage    string  `json:"percentage"`
	ROI    string  `json:"roi"`
	Tenure string  `json:"tenure"`
	ApprovedOn string  `json:"aprovedon"`
	AppliedOn string  `json:"appliedon"`
}

type Registrar struct {
	Name     string  `json:"name"`
	Location    string  `json:"location"`
}

type Settlement struct {
	SettlementType     string  `json:"type"`
	SettlementAmount    string  `json:"amount"`
	SettlementRef		string `json:"ref"`
	SettlementDate		string `json:"date"`
}


type CP struct {
	CUSIP     string  `json:"cusip"`
	Ticker    string  `json:"ticker"`
	Par       float64 `json:"par"`
	Qty       int     `json:"qty"`
	Discount  float64 `json:"discount"`
	Maturity  int     `json:"maturity"`
	Owners    []Owner `json:"owner"`
	Issuer    string  `json:"issuer"`
	IssueDate string  `json:"issueDate"`
}

type Account struct {
	ID          string   `json:"id"`
	Prefix      string   `json:"prefix"`
	CashBalance float64  `json:"cashBalance"`
	AssetsIds   []string `json:"assetIds"`
}

type Transaction struct {
	CUSIP       string  `json:"cusip"`
	FromCompany string  `json:"fromCompany"`
	ToCompany   string  `json:"toCompany"`
	Quantity    int     `json:"quantity"`
	Discount    float64 `json:"discount"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Initialize the collection of commercial paper keys
	fmt.Println("Initializing paper keys collection")
	var blank []string
	var blank1 []string
	var blank2 []string
	var blank3 []string
	var blank4 []string
	var blank5 []string
	var blank6 []string

	blankBytes, _ := json.Marshal(&blank)
	err := stub.PutState("PaperKeys", blankBytes)
	if err != nil {
		fmt.Println("Failed to initialize paper key collection")
	}

	blankBytes1, _ := json.Marshal(&blank1)
	err1 := stub.PutState("PropertyKeys", blankBytes1)
	if err1 != nil {
		fmt.Println("Failed to initialize paper key collection")
	}

	blankBytes2, _ := json.Marshal(&blank2)
	err2 := stub.PutState("ProposalKeys", blankBytes2)
	if err2 != nil {
		fmt.Println("Failed to initialize paper key collection")
	}

	blankBytes3, _ := json.Marshal(&blank3)
	err3 := stub.PutState("AgreementKeys", blankBytes3)
	if err3 != nil {
		fmt.Println("Failed to initialize paper key collection")
	}

	blankBytes4, _ := json.Marshal(&blank4)
	err4 := stub.PutState("DeedKeys", blankBytes4)
	if err4 != nil {
		fmt.Println("Failed to initialize paper key collection")
	}

	blankBytes5, _ := json.Marshal(&blank5)
	err5 := stub.PutState("NotificationKeys", blankBytes5)
	if err5 != nil {
		fmt.Println("Failed to initialize paper key collection")
	}
	blankBytes6, _ := json.Marshal(&blank6)
	err6 := stub.PutState("QuoteKeys", blankBytes6)
	if err6 != nil {
		fmt.Println("Failed to initialize paper key collection")
	}


	fmt.Println("Initialization complete")
	return nil, nil
}

func (t *SimpleChaincode) createAccounts(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	//  				0
	// "number of accounts to create"
	var err error
	numAccounts, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("error creating accounts with input")
		return nil, errors.New("createAccounts accepts a single integer argument")
	}
	//create a bunch of accounts
	var account Account
	counter := 1
	for counter <= numAccounts {
		var prefix string
		suffix := "000A"
		if counter < 10 {
			prefix = strconv.Itoa(counter) + "0" + suffix
		} else {
			prefix = strconv.Itoa(counter) + suffix
		}
		var assetIds []string
		account = Account{ID: "company" + strconv.Itoa(counter), Prefix: prefix, CashBalance: 10000000.0, AssetsIds: assetIds}
		accountBytes, err := json.Marshal(&account)
		if err != nil {
			fmt.Println("error creating account" + account.ID)
			return nil, errors.New("Error creating account " + account.ID)
		}
		err = stub.PutState(accountPrefix+account.ID, accountBytes)
		counter++
		fmt.Println("created account" + accountPrefix + account.ID)
	}

	fmt.Println("Accounts created")
	return nil, nil

}

func (t *SimpleChaincode) createAccount(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// Obtain the username to associate with the account
	if len(args) != 1 {
		fmt.Println("Error obtaining username")
		return nil, errors.New("createAccount accepts a single username argument")
	}
	username := args[0]

	// Build an account object for the user
	var assetIds []string
	suffix := "000A"
	prefix := username + suffix
	var account = Account{ID: username, Prefix: prefix, CashBalance: 10000000.0, AssetsIds: assetIds}
	accountBytes, err := json.Marshal(&account)
	if err != nil {
		fmt.Println("error creating account" + account.ID)
		return nil, errors.New("Error creating account " + account.ID)
	}

	fmt.Println("Attempting to get state of any existing account for " + account.ID)
	existingBytes, err := stub.GetState(accountPrefix + account.ID)
	if err == nil {

		var company Account
		err = json.Unmarshal(existingBytes, &company)
		if err != nil {
			fmt.Println("Error unmarshalling account " + account.ID + "\n--->: " + err.Error())

			if strings.Contains(err.Error(), "unexpected end") {
				fmt.Println("No data means existing account found for " + account.ID + ", initializing account.")
				err = stub.PutState(accountPrefix+account.ID, accountBytes)

				if err == nil {
					fmt.Println("created account" + accountPrefix + account.ID)
					return nil, nil
				} else {
					fmt.Println("failed to create initialize account for " + account.ID)
					return nil, errors.New("failed to initialize an account for " + account.ID + " => " + err.Error())
				}
			} else {
				return nil, errors.New("Error unmarshalling existing account " + account.ID)
			}
		} else {
			fmt.Println("Account already exists for " + account.ID + " " + company.ID)
			return nil, errors.New("Can't reinitialize existing user " + account.ID)
		}
	} else {

		fmt.Println("No existing account found for " + account.ID + ", initializing account.")
		err = stub.PutState(accountPrefix+account.ID, accountBytes)

		if err == nil {
			fmt.Println("created account" + accountPrefix + account.ID)
			return nil, nil
		} else {
			fmt.Println("failed to create initialize account for " + account.ID)
			return nil, errors.New("failed to initialize an account for " + account.ID + " => " + err.Error())
		}

	}

}


/* Added by Narayanan L for Trade Finance */
//Quote

func (t *SimpleChaincode) issueQuote(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	
	//need one arg
	if len(args) != 1 {
		fmt.Println("error invalid arguments")
		return nil, errors.New("Incorrect number of arguments. Expecting Quotation record")
	}

	var quote Quote
	var err error
	var account Account

	fmt.Println("Unmarshalling Quote")
	err = json.Unmarshal([]byte(args[0]), &quote)
	if err != nil {
		fmt.Println("error invalid Quote issue")
		return nil, errors.New("Invalid Quote issue")
	}

	

	fmt.Println("Marshalling Quote bytes")
	quote.QuoteNo = account.Prefix + quote.QuoteNo

	fmt.Println("Getting State on CP " + quote.QuoteNo)
	cpRxBytes, err := stub.GetState(quotePrefix + quote.QuoteNo)
	if cpRxBytes == nil {
		fmt.Println("QuoteNo does not exist, creating it")
		cpBytes, err := json.Marshal(&quote)
		if err != nil {
			fmt.Println("Error marshalling quote")
			return nil, errors.New("Error issuing quote")
		}
		err = stub.PutState(quotePrefix+quote.QuoteNo, cpBytes)
		if err != nil {
			fmt.Println("Error issuing paper")
			return nil, errors.New("Error issuing quote")
		}

		

		// Update the paper keys by adding the new key
		fmt.Println("Getting Paper Keys")
		keysBytes, err := stub.GetState("QuoteKeys")
		if err != nil {
			fmt.Println("Error retrieving paper keys")
			return nil, errors.New("Error retrieving paper keys")
		}
		var keys []string
		err = json.Unmarshal(keysBytes, &keys)
		if err != nil {
			fmt.Println("Error unmarshel keys")
			return nil, errors.New("Error unmarshalling paper keys ")
		}

		fmt.Println("Appending the new key to Paper Keys")
		foundKey := false
		for _, key := range keys {
			if key == quotePrefix+quote.QuoteNo {
				foundKey = true
			}
		}
		if foundKey == false {
			keys = append(keys, quotePrefix+quote.QuoteNo)
			keysBytesToWrite, err := json.Marshal(&keys)
			if err != nil {
				fmt.Println("Error marshalling keys")
				return nil, errors.New("Error marshalling the keys")
			}
			fmt.Println("Put state on QuoteKeys")
			err = stub.PutState("QuoteKeys", keysBytesToWrite)
			if err != nil {
				fmt.Println("Error writting keys back")
				return nil, errors.New("Error writing the keys back")
			}
		}

		fmt.Println("Issue commercial paper %+v\n", quote)
		return nil, nil
	} else {
		fmt.Println("QuoteNo exists")

		var quoterx Quote
		fmt.Println("Unmarshalling CP " + quote.QuoteNo)
		err = json.Unmarshal(cpRxBytes, &quoterx)
		if err != nil {
			fmt.Println("Error unmarshalling cp " + quote.QuoteNo)
			return nil, errors.New("Error unmarshalling cp " + quote.QuoteNo)
		}

		quoterx.Qty = quoterx.Qty + quote.Qty

		


		cpWriteBytes, err := json.Marshal(&quoterx)
		if err != nil {
			fmt.Println("Error marshalling cp")
			return nil, errors.New("Error issuing commercial paper")
		}
		err = stub.PutState(quotePrefix+quote.QuoteNo, cpWriteBytes)
		if err != nil {
			fmt.Println("Error issuing paper")
			return nil, errors.New("Error issuing commercial paper")
		}

		fmt.Println("Updated commercial paper %+v\n", quoterx)
		return nil, nil
		
	}
}
func GetAllQuotes(stub shim.ChaincodeStubInterface) ([]Quote, error) {

	var allquote []Quote

	fmt.Println("retrieving quote Keys ")

	// Get list of all the keys
	keysBytes, err := stub.GetState("QuoteKeys")
	if err != nil {
		fmt.Println("Error retrieving quote Keys ")
		return nil, errors.New("Error retrieving quote Keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling quote keys")
		return nil, errors.New("Error unmarshalling quote keys")
	}

	// Get all the cps
	for _, value := range keys {
		cpBytes, err := stub.GetState(value)

		var quote Quote
		err = json.Unmarshal(cpBytes, &quote)
		if err != nil {
			fmt.Println("Error retrieving quote " + value)
			return nil, errors.New("Error retrieving quote " + value)
		}

		fmt.Println("Appending quote" + value)
		allquote = append(allquote, quote)
	}

	return allquote, nil
}



/* Added by Narayanan L for Land Record Management*/

//Notification

func (t *SimpleChaincode) addNotification(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	
	//need one arg
	if len(args) != 1 {
		fmt.Println("error invalid arguments")
		return nil, errors.New("Incorrect number of arguments. Expecting Quotation record")
	}

	var notification Notification
	var err error
	//var account Account

	fmt.Println("Unmarshalling notification")
	err = json.Unmarshal([]byte(args[0]), &notification)
	if err != nil {
		fmt.Println("error invalid notification issue" + args[0])
		return nil, errors.New("Invalid notification issue" + args[0])
	}

	

	fmt.Println("Marshalling notification bytes")
	//notification.PropId = notificationPrefix + notification.NotificationId

	fmt.Println("Getting State on notification " + notification.NotificationId)
	cpRxBytes, err := stub.GetState(notificationPrefix + notification.NotificationId)
	if cpRxBytes == nil {
		fmt.Println("PropId does not exist, creating it")
		cpBytes, err := json.Marshal(&notification)
		if err != nil {
			fmt.Println("Error marshalling notification")
			return nil, errors.New("Error issuing notification")
		}
		err = stub.PutState(notificationPrefix+notification.NotificationId, cpBytes)
		if err != nil {
			fmt.Println("Error issuing paper")
			return nil, errors.New("Error issuing notification")
		}

		

		// Update the paper keys by adding the new key
		fmt.Println("Getting Paper Keys")
		keysBytes, err := stub.GetState("NotificationKeys")
		if err != nil {
			fmt.Println("Error retrieving paper keys")
			return nil, errors.New("Error retrieving paper keys")
		}
		var keys []string
		err = json.Unmarshal(keysBytes, &keys)
		if err != nil {
			fmt.Println("Error unmarshel keys")
			return nil, errors.New("Error unmarshalling paper keys ")
		}

		fmt.Println("Appending the new key to Paper Keys")
		foundKey := false
		for _, key := range keys {
			if key == notificationPrefix+notification.NotificationId {
				foundKey = true
			}
		}
		if foundKey == false {
			keys = append(keys, notificationPrefix+notification.NotificationId)
			keysBytesToWrite, err := json.Marshal(&keys)
			if err != nil {
				fmt.Println("Error marshalling keys")
				return nil, errors.New("Error marshalling the keys")
			}
			fmt.Println("Put state on NotificationKeys")
			err = stub.PutState("NotificationKeys", keysBytesToWrite)
			if err != nil {
				fmt.Println("Error writting keys back")
				return nil, errors.New("Error writing the keys back")
			}
		}

		fmt.Println("Issue commercial paper %+v\n", notification)
		return nil, nil
	} else {
		fmt.Println("QuoteNo exists")

		var notificationrx Notification
		fmt.Println("Unmarshalling CP " + notification.NotificationId)
		err = json.Unmarshal(cpRxBytes, &notificationrx)
		if err != nil {
			fmt.Println("Error unmarshalling cp " + notification.NotificationId)
			return nil, errors.New("Error unmarshalling cp " + notification.NotificationId)
		}

		//quoterx.Qty = quoterx.Qty + quote.Qty

		notificationrx = notification

		cpWriteBytes, err := json.Marshal(&notificationrx)
		if err != nil {
			fmt.Println("Error marshalling cp")
			return nil, errors.New("Error issuing commercial paper")
		}
		err = stub.PutState(notificationPrefix+notification.NotificationId, cpWriteBytes)
		if err != nil {
			fmt.Println("Error issuing paper")
			return nil, errors.New("Error issuing commercial paper")
		}

		fmt.Println("Updated commercial paper %+v\n", notificationrx)
		return nil, nil
	}
}

func GetAllNotifications(stub shim.ChaincodeStubInterface) ([]Notification, error) {

	var allNotification []Notification

	// Get list of all the keys
	keysBytes, err := stub.GetState("NotificationKeys")
	if err != nil {
		fmt.Println("Error retrieving Notification keys")
		return nil, errors.New("Error retrieving Notification keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling Notification keys")
		return nil, errors.New("Error unmarshalling Notification keys")
	}

	// Get all the cps
	for _, value := range keys {
		cpBytes, err := stub.GetState(value)

		var notification Notification
		err = json.Unmarshal(cpBytes, &notification)
		if err != nil {
			fmt.Println("Error retrieving cp " + value)
			return nil, errors.New("Error retrieving cp " + value)
		}

		fmt.Println("Appending CP" + value)
		allNotification = append(allNotification, notification)
	}

	return allNotification, nil
}

//property

func (t *SimpleChaincode) addProperty(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	
	//need one arg
	if len(args) != 1 {
		fmt.Println("error invalid arguments")
		return nil, errors.New("Incorrect number of arguments. Expecting Quotation record")
	}

	var property Property
	var err error
	//var account Account

	fmt.Println("Unmarshalling Property")
	err = json.Unmarshal([]byte(args[0]), &property)
	if err != nil {
		fmt.Println("error invalid Property issue" + args[0])
		return nil, errors.New("Invalid Property issue" + args[0])
	}

	

	fmt.Println("Marshalling Property bytes")
	//property.PropId = propertyPrefix + property.propid

	fmt.Println("Getting State on Property " + property.PropId)
	cpRxBytes, err := stub.GetState(propertyPrefix + property.PropId)
	if cpRxBytes == nil {
		fmt.Println("PropId does not exist, creating it")
		cpBytes, err := json.Marshal(&property)
		if err != nil {
			fmt.Println("Error marshalling property")
			return nil, errors.New("Error issuing property")
		}
		err = stub.PutState(propertyPrefix+property.PropId, cpBytes)
		if err != nil {
			fmt.Println("Error issuing paper")
			return nil, errors.New("Error issuing property")
		}

		

		// Update the paper keys by adding the new key
		fmt.Println("Getting Paper Keys")
		keysBytes, err := stub.GetState("PropertyKeys")
		if err != nil {
			fmt.Println("Error retrieving paper keys")
			return nil, errors.New("Error retrieving paper keys")
		}
		var keys []string
		err = json.Unmarshal(keysBytes, &keys)
		if err != nil {
			fmt.Println("Error unmarshel keys")
			return nil, errors.New("Error unmarshalling paper keys ")
		}

		fmt.Println("Appending the new key to Paper Keys")
		foundKey := false
		for _, key := range keys {
			if key == propertyPrefix+property.PropId {
				foundKey = true
			}
		}
		if foundKey == false {
			keys = append(keys, propertyPrefix+property.PropId)
			keysBytesToWrite, err := json.Marshal(&keys)
			if err != nil {
				fmt.Println("Error marshalling keys")
				return nil, errors.New("Error marshalling the keys")
			}
			fmt.Println("Put state on PropertyKeys")
			err = stub.PutState("PropertyKeys", keysBytesToWrite)
			if err != nil {
				fmt.Println("Error writting keys back")
				return nil, errors.New("Error writing the keys back")
			}
		}

		fmt.Println("Issue commercial paper %+v\n", property)
		return nil, nil
	} else {
		fmt.Println("QuoteNo exists")

		var propertyrx Property
		fmt.Println("Unmarshalling CP " + property.PropId)
		err = json.Unmarshal(cpRxBytes, &propertyrx)
		if err != nil {
			fmt.Println("Error unmarshalling cp " + property.PropId)
			return nil, errors.New("Error unmarshalling cp " + property.PropId)
		}

		//quoterx.Qty = quoterx.Qty + quote.Qty

		propertyrx = property

		cpWriteBytes, err := json.Marshal(&propertyrx)
		if err != nil {
			fmt.Println("Error marshalling cp")
			return nil, errors.New("Error issuing commercial paper")
		}
		err = stub.PutState(propertyPrefix+property.PropId, cpWriteBytes)
		if err != nil {
			fmt.Println("Error issuing paper")
			return nil, errors.New("Error issuing commercial paper")
		}

		fmt.Println("Updated commercial paper %+v\n", propertyrx)
		return nil, nil
	}
}

func GetAllProperties(stub shim.ChaincodeStubInterface) ([]Property, error) {

	var allProperties []Property

	// Get list of all the keys
	keysBytes, err := stub.GetState("PropertyKeys")
	if err != nil {
		fmt.Println("Error retrieving paper keys")
		return nil, errors.New("Error retrieving paper keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling paper keys")
		return nil, errors.New("Error unmarshalling paper keys")
	}

	// Get all the cps
	for _, value := range keys {
		cpBytes, err := stub.GetState(value)

		var property Property
		err = json.Unmarshal(cpBytes, &property)
		if err != nil {
			fmt.Println("Error retrieving cp " + value)
			return nil, errors.New("Error retrieving cp " + value)
		}

		fmt.Println("Appending CP" + value)
		allProperties = append(allProperties, property)
	}

	return allProperties, nil
}


//Proposal


func (t *SimpleChaincode) issueProposal(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	
	//need one arg
	if len(args) != 1 {
		fmt.Println("error invalid arguments")
		return nil, errors.New("Incorrect number of arguments. Expecting Quotation record")
	}

	var proposal Proposal
	var err error
	//var account Account

	fmt.Println("Unmarshalling Proposal")
	err = json.Unmarshal([]byte(args[0]), &proposal)
	if err != nil {
		fmt.Println("error invalid proposal issue" + args[0])
		return nil, errors.New("Invalid proposal issue" + args[0])
	}

	

	fmt.Println("Marshalling proposal bytes")
	//property.PropId = propertyPrefix + property.propid

	fmt.Println("Getting State on proposal " + proposal.ProposalNo)
	cpRxBytes, err := stub.GetState(proposalPrefix + proposal.ProposalNo)
	if cpRxBytes == nil {
		fmt.Println("proposalNo does not exist, creating it")
		cpBytes, err := json.Marshal(&proposal)
		if err != nil {
			fmt.Println("Error marshalling proposal")
			return nil, errors.New("Error issuing proposal")
		}
		err = stub.PutState(proposalPrefix + proposal.ProposalNo, cpBytes)
		if err != nil {
			fmt.Println("Error issuing proposal")
			return nil, errors.New("Error issuing proposal")
		}

		

		// Update the paper keys by adding the new key
		fmt.Println("Getting proposal Keys")
		keysBytes, err := stub.GetState("ProposalKeys")
		if err != nil {
			fmt.Println("Error retrieving proposalNo")
			return nil, errors.New("Error retrieving proposalNo")
		}
		var keys []string
		err = json.Unmarshal(keysBytes, &keys)
		if err != nil {
			fmt.Println("Error unmarshel proposalNo")
			return nil, errors.New("Error unmarshalling proposalNo ")
		}

		fmt.Println("Appending the new key to proposalNo Keys")
		foundKey := false
		for _, key := range keys {
			if key == proposalPrefix+proposal.ProposalNo {
				foundKey = true
			}
		}
		if foundKey == false {
			keys = append(keys, proposalPrefix+proposal.ProposalNo)
			keysBytesToWrite, err := json.Marshal(&keys)
			if err != nil {
				fmt.Println("Error marshalling proposalNo")
				return nil, errors.New("Error marshalling the proposalNo")
			}
			fmt.Println("Put state on proposalNo")
			err = stub.PutState("ProposalKeys", keysBytesToWrite)
			if err != nil {
				fmt.Println("Error writting proposalNo back")
				return nil, errors.New("Error writing the proposalNo back")
			}
		}

		fmt.Println("Issue commercial paper %+v\n", proposal)
		return nil, nil
	} else {
		fmt.Println("proposalNo exists")

		var proposalrx Proposal
		fmt.Println("Unmarshalling proposal " + proposal.ProposalNo)
		err = json.Unmarshal(cpRxBytes, &proposalrx)
		if err != nil {
			fmt.Println("Error unmarshalling proposal " + proposal.ProposalNo)
			return nil, errors.New("Error unmarshalling proposal " + proposal.ProposalNo)
		}

		//quoterx.Qty = quoterx.Qty + quote.Qty

		proposalrx = proposal


		cpWriteBytes, err := json.Marshal(&proposalrx)
		if err != nil {
			fmt.Println("Error marshalling proposal")
			return nil, errors.New("Error issuing proposal")
		}
		err = stub.PutState(proposalPrefix+proposal.ProposalNo, cpWriteBytes)
		if err != nil {
			fmt.Println("Error proposal")
			return nil, errors.New("Error issuing proposal")
		}

		fmt.Println("Updated commercial paper %+v\n", proposalrx)
		return nil, nil
	}
}

func GetAllproposal(stub shim.ChaincodeStubInterface) ([]Proposal, error) {

	var allproposal []Proposal

	// Get list of all the keys
	keysBytes, err := stub.GetState("ProposalKeys")
	if err != nil {
		fmt.Println("Error retrieving proposal Keys ")
		return nil, errors.New("Error retrieving proposal Keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling proposal keys")
		return nil, errors.New("Error unmarshalling proposal keys")
	}

	// Get all the cps
	for _, value := range keys {
		cpBytes, err := stub.GetState(value)

		var proposal Proposal
		err = json.Unmarshal(cpBytes, &proposal)
		if err != nil {
			fmt.Println("Error retrieving proposal " + value)
			return nil, errors.New("Error retrieving proposal " + value)
		}

		fmt.Println("Appending proposal" + value)
		allproposal = append(allproposal, proposal)
	}

	return allproposal, nil
}


//Agreement


func (t *SimpleChaincode) issueSaleAgreement(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	
	//need one arg
	if len(args) != 1 {
		fmt.Println("error invalid arguments")
		return nil, errors.New("Incorrect number of arguments. Expecting Quotation record")
	}

	var saleAgreement SaleAgreement
	var err error
	//var account Account

	fmt.Println("Unmarshalling Property")
	err = json.Unmarshal([]byte(args[0]), &saleAgreement)
	if err != nil {
		fmt.Println("error invalid saleAgreement issue" + args[0])
		return nil, errors.New("Invalid saleAgreement issue" + args[0])
	}

	

	fmt.Println("Marshalling saleAgreement bytes")
	//property.PropId = propertyPrefix + property.propid

	fmt.Println("Getting State on saleAgreement " + saleAgreement.AgreementNo)
	cpRxBytes, err := stub.GetState(agreementPrefix + saleAgreement.AgreementNo)
	if cpRxBytes == nil {
		fmt.Println("AgreementNo does not exist, creating it")
		cpBytes, err := json.Marshal(&saleAgreement)
		if err != nil {
			fmt.Println("Error marshalling saleAgreement")
			return nil, errors.New("Error issuing saleAgreement")
		}
		err = stub.PutState(agreementPrefix + saleAgreement.AgreementNo, cpBytes)
		if err != nil {
			fmt.Println("Error issuing saleAgreement")
			return nil, errors.New("Error issuing saleAgreement")
		}

		

		// Update the paper keys by adding the new key
		fmt.Println("Getting saleAgreement Keys")
		keysBytes, err := stub.GetState("AgreementKeys")
		if err != nil {
			fmt.Println("Error retrieving AgreementKeys")
			return nil, errors.New("Error retrieving AgreementKeys")
		}
		var keys []string
		err = json.Unmarshal(keysBytes, &keys)
		if err != nil {
			fmt.Println("Error unmarshel AgreementKeys")
			return nil, errors.New("Error unmarshalling AgreementKeys ")
		}

		fmt.Println("Appending the new key to AgreementKeys Keys")
		foundKey := false
		for _, key := range keys {
			if key == agreementPrefix+saleAgreement.AgreementNo {
				foundKey = true
			}
		}
		if foundKey == false {
			keys = append(keys, agreementPrefix+saleAgreement.AgreementNo)
			keysBytesToWrite, err := json.Marshal(&keys)
			if err != nil {
				fmt.Println("Error marshalling AgreementKeys")
				return nil, errors.New("Error marshalling the AgreementKeys")
			}
			fmt.Println("Put state on AgreementKeys")
			err = stub.PutState("AgreementKeys", keysBytesToWrite)
			if err != nil {
				fmt.Println("Error writting AgreementKeys back")
				return nil, errors.New("Error writing the AgreementKeys back")
			}
		}

		fmt.Println("Issue commercial paper %+v\n", saleAgreement)
		return nil, nil
	} else {
		fmt.Println("AgreementNo exists")

		var saleAgreementrx SaleAgreement
		fmt.Println("Unmarshalling saleAgreement " + saleAgreement.AgreementNo)
		err = json.Unmarshal(cpRxBytes, &saleAgreementrx)
		if err != nil {
			fmt.Println("Error unmarshalling saleAgreement " + saleAgreement.AgreementNo)
			return nil, errors.New("Error unmarshalling saleAgreement " + saleAgreement.AgreementNo)
		}

		//quoterx.Qty = quoterx.Qty + quote.Qty

		saleAgreementrx = saleAgreement
		


		cpWriteBytes, err := json.Marshal(&saleAgreementrx)
		if err != nil {
			fmt.Println("Error marshalling saleAgreement")
			return nil, errors.New("Error issuing saleAgreement")
		}
		err = stub.PutState(agreementPrefix+saleAgreement.AgreementNo, cpWriteBytes)
		if err != nil {
			fmt.Println("Error saleAgreement")
			return nil, errors.New("Error issuing saleAgreement")
		}

		fmt.Println("Updated commercial paper %+v\n", saleAgreementrx)
		return nil, nil
	}
}

func GetAllAgreement(stub shim.ChaincodeStubInterface) ([]SaleAgreement, error) {

	var allSaleAgreement []SaleAgreement

	// Get list of all the keys
	keysBytes, err := stub.GetState("AgreementKeys")
	if err != nil {
		fmt.Println("Error retrieving SaleAgreement Keys ")
		return nil, errors.New("Error retrieving SaleAgreement Keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling SaleAgreement keys")
		return nil, errors.New("Error unmarshalling SaleAgreement keys")
	}

	// Get all the cps
	for _, value := range keys {
		cpBytes, err := stub.GetState(value)

		var saleAgreement SaleAgreement
		err = json.Unmarshal(cpBytes, &saleAgreement)
		if err != nil {
			fmt.Println("Error retrieving saleAgreement " + value)
			return nil, errors.New("Error retrieving saleAgreement " + value)
		}

		fmt.Println("Appending saleAgreement" + value)
		allSaleAgreement = append(allSaleAgreement, saleAgreement)
	}

	return allSaleAgreement, nil
}


//Deeds


func (t *SimpleChaincode) issueSaleDeeds(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	
	//need one arg
	if len(args) != 1 {
		fmt.Println("error invalid arguments")
		return nil, errors.New("Incorrect number of arguments. Expecting Quotation record")
	}

	var saleDeed SaleDeed
	var err error
	//var account Account

	fmt.Println("Unmarshalling Property")
	err = json.Unmarshal([]byte(args[0]), &saleDeed)
	if err != nil {
		fmt.Println("error invalid saleDeed issue" + args[0])
		return nil, errors.New("Invalid saleDeed issue" + args[0])
	}

	

	fmt.Println("Marshalling saleDeed bytes")
	//property.PropId = propertyPrefix + property.propid

	fmt.Println("Getting State on saleDeed " + saleDeed.DeedNo)
	cpRxBytes, err := stub.GetState(deedPrefix + saleDeed.DeedNo)
	if cpRxBytes == nil {
		fmt.Println("DeedNo does not exist, creating it")
		cpBytes, err := json.Marshal(&saleDeed)
		if err != nil {
			fmt.Println("Error marshalling saleDeed")
			return nil, errors.New("Error issuing saleDeed")
		}
		err = stub.PutState(deedPrefix + saleDeed.DeedNo, cpBytes)
		if err != nil {
			fmt.Println("Error issuing saleDeed")
			return nil, errors.New("Error issuing saleDeed")
		}

		

		// Update the paper keys by adding the new key
		fmt.Println("Getting saleDeed Keys")
		keysBytes, err := stub.GetState("DeedKeys")
		if err != nil {
			fmt.Println("Error retrieving DeedKeys")
			return nil, errors.New("Error retrieving DeedKeys")
		}
		var keys []string
		err = json.Unmarshal(keysBytes, &keys)
		if err != nil {
			fmt.Println("Error unmarshel DeedKeys")
			return nil, errors.New("Error unmarshalling DeedKeys ")
		}

		fmt.Println("Appending the new key to DeedKeys Keys")
		foundKey := false
		for _, key := range keys {
			if key == deedPrefix+saleDeed.DeedNo {
				foundKey = true
			}
		}
		if foundKey == false {
			keys = append(keys, deedPrefix+saleDeed.DeedNo)
			keysBytesToWrite, err := json.Marshal(&keys)
			if err != nil {
				fmt.Println("Error marshalling DeedKeys")
				return nil, errors.New("Error marshalling the DeedKeys")
			}
			fmt.Println("Put state on DeedKeys")
			err = stub.PutState("DeedKeys", keysBytesToWrite)
			if err != nil {
				fmt.Println("Error writting DeedKeys back")
				return nil, errors.New("Error writing the DeedKeys back")
			}
		}

		fmt.Println("Issue commercial paper %+v\n", saleDeed)
		return nil, nil
	} else {
		fmt.Println("DeedNo exists")

		var saleDeedrx SaleDeed
		fmt.Println("Unmarshalling saleDeed " + saleDeed.DeedNo)
		err = json.Unmarshal(cpRxBytes, &saleDeedrx)
		if err != nil {
			fmt.Println("Error unmarshalling saleDeed " + saleDeed.DeedNo)
			return nil, errors.New("Error unmarshalling saleDeed " + saleDeed.DeedNo)
		}

		//quoterx.Qty = quoterx.Qty + quote.Qty

		saleDeedrx = saleDeed


		cpWriteBytes, err := json.Marshal(&saleDeedrx)
		if err != nil {
			fmt.Println("Error marshalling saleDeed")
			return nil, errors.New("Error issuing saleDeed")
		}
		err = stub.PutState(deedPrefix+saleDeed.DeedNo, cpWriteBytes)
		if err != nil {
			fmt.Println("Error saleDeed")
			return nil, errors.New("Error issuing saleDeed")
		}

		fmt.Println("Updated commercial paper %+v\n", saleDeedrx)
		return nil, nil
	}
}

func GetAllDeed(stub shim.ChaincodeStubInterface) ([]SaleDeed, error) {

	var allsaleDeed []SaleDeed

	// Get list of all the keys
	keysBytes, err := stub.GetState("DeedKeys")
	if err != nil {
		fmt.Println("Error retrieving saleDeed Keys ")
		return nil, errors.New("Error retrieving saleDeed Keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling saleDeed keys")
		return nil, errors.New("Error unmarshalling saleDeed keys")
	}

	// Get all the cps
	for _, value := range keys {
		cpBytes, err := stub.GetState(value)

		var saleDeed SaleDeed
		err = json.Unmarshal(cpBytes, &saleDeed)
		if err != nil {
			fmt.Println("Error retrieving saleDeed " + value)
			return nil, errors.New("Error retrieving saleDeed " + value)
		}

		fmt.Println("Appending saleDeed" + value)
		allsaleDeed = append(allsaleDeed, saleDeed)
	}

	return allsaleDeed, nil
}









/* Added by Narayanan L for Land Record Management*/




func (t *SimpleChaincode) issueCommercialPaper(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	/*		0
		json
	  	{
			"ticker":  "string",
			"par": 0.00,
			"qty": 10,
			"discount": 7.5,
			"maturity": 30,
			"owners": [ // This one is not required
				{
					"company": "company1",
					"quantity": 5
				},
				{
					"company": "company3",
					"quantity": 3
				},
				{
					"company": "company4",
					"quantity": 2
				}
			],
			"issuer":"company2",
			"issueDate":"1456161763790"  (current time in milliseconds as a string)

		}
	*/
	//need one arg
	if len(args) != 1 {
		fmt.Println("error invalid arguments")
		return nil, errors.New("Incorrect number of arguments. Expecting commercial paper record")
	}

	var cp CP
	var err error
	var account Account

	fmt.Println("Unmarshalling CP")
	err = json.Unmarshal([]byte(args[0]), &cp)
	if err != nil {
		fmt.Println("error invalid paper issue")
		return nil, errors.New("Invalid commercial paper issue")
	}

	//generate the CUSIP
	//get account prefix
	fmt.Println("Getting state of - " + accountPrefix + cp.Issuer)
	accountBytes, err := stub.GetState(accountPrefix + cp.Issuer)
	if err != nil {
		fmt.Println("Error Getting state of - " + accountPrefix + cp.Issuer)
		return nil, errors.New("Error retrieving account " + cp.Issuer)
	}
	err = json.Unmarshal(accountBytes, &account)
	if err != nil {
		fmt.Println("Error Unmarshalling accountBytes")
		return nil, errors.New("Error retrieving account " + cp.Issuer)
	}

	account.AssetsIds = append(account.AssetsIds, cp.CUSIP)

	// Set the issuer to be the owner of all quantity
	var owner Owner
	owner.Company = cp.Issuer
	owner.Quantity = cp.Qty

	cp.Owners = append(cp.Owners, owner)

	suffix, err := generateCUSIPSuffix(cp.IssueDate, cp.Maturity)
	if err != nil {
		fmt.Println("Error generating cusip")
		return nil, errors.New("Error generating CUSIP")
	}

	fmt.Println("Marshalling CP bytes")
	cp.CUSIP = account.Prefix + suffix

	fmt.Println("Getting State on CP " + cp.CUSIP)
	cpRxBytes, err := stub.GetState(cpPrefix + cp.CUSIP)
	if cpRxBytes == nil {
		fmt.Println("CUSIP does not exist, creating it")
		cpBytes, err := json.Marshal(&cp)
		if err != nil {
			fmt.Println("Error marshalling cp")
			return nil, errors.New("Error issuing commercial paper")
		}
		err = stub.PutState(cpPrefix+cp.CUSIP, cpBytes)
		if err != nil {
			fmt.Println("Error issuing paper")
			return nil, errors.New("Error issuing commercial paper")
		}

		fmt.Println("Marshalling account bytes to write")
		accountBytesToWrite, err := json.Marshal(&account)
		if err != nil {
			fmt.Println("Error marshalling account")
			return nil, errors.New("Error issuing commercial paper")
		}
		err = stub.PutState(accountPrefix+cp.Issuer, accountBytesToWrite)
		if err != nil {
			fmt.Println("Error putting state on accountBytesToWrite")
			return nil, errors.New("Error issuing commercial paper")
		}

		// Update the paper keys by adding the new key
		fmt.Println("Getting Paper Keys")
		keysBytes, err := stub.GetState("PaperKeys")
		if err != nil {
			fmt.Println("Error retrieving paper keys")
			return nil, errors.New("Error retrieving paper keys")
		}
		var keys []string
		err = json.Unmarshal(keysBytes, &keys)
		if err != nil {
			fmt.Println("Error unmarshel keys")
			return nil, errors.New("Error unmarshalling paper keys ")
		}

		fmt.Println("Appending the new key to Paper Keys")
		foundKey := false
		for _, key := range keys {
			if key == cpPrefix+cp.CUSIP {
				foundKey = true
			}
		}
		if foundKey == false {
			keys = append(keys, cpPrefix+cp.CUSIP)
			keysBytesToWrite, err := json.Marshal(&keys)
			if err != nil {
				fmt.Println("Error marshalling keys")
				return nil, errors.New("Error marshalling the keys")
			}
			fmt.Println("Put state on PaperKeys")
			err = stub.PutState("PaperKeys", keysBytesToWrite)
			if err != nil {
				fmt.Println("Error writting keys back")
				return nil, errors.New("Error writing the keys back")
			}
		}

		fmt.Println("Issue commercial paper %+v\n", cp)
		return nil, nil
	} else {
		fmt.Println("CUSIP exists")

		var cprx CP
		fmt.Println("Unmarshalling CP " + cp.CUSIP)
		err = json.Unmarshal(cpRxBytes, &cprx)
		if err != nil {
			fmt.Println("Error unmarshalling cp " + cp.CUSIP)
			return nil, errors.New("Error unmarshalling cp " + cp.CUSIP)
		}

		cprx.Qty = cprx.Qty + cp.Qty

		for key, val := range cprx.Owners {
			if val.Company == cp.Issuer {
				cprx.Owners[key].Quantity += cp.Qty
				break
			}
		}

		cpWriteBytes, err := json.Marshal(&cprx)
		if err != nil {
			fmt.Println("Error marshalling cp")
			return nil, errors.New("Error issuing commercial paper")
		}
		err = stub.PutState(cpPrefix+cp.CUSIP, cpWriteBytes)
		if err != nil {
			fmt.Println("Error issuing paper")
			return nil, errors.New("Error issuing commercial paper")
		}

		fmt.Println("Updated commercial paper %+v\n", cprx)
		return nil, nil
	}
}

func GetAllCPs(stub shim.ChaincodeStubInterface) ([]CP, error) {

	var allCPs []CP

	// Get list of all the keys
	keysBytes, err := stub.GetState("PaperKeys")
	if err != nil {
		fmt.Println("Error retrieving paper keys")
		return nil, errors.New("Error retrieving paper keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling paper keys")
		return nil, errors.New("Error unmarshalling paper keys")
	}

	// Get all the cps
	for _, value := range keys {
		cpBytes, err := stub.GetState(value)

		var cp CP
		err = json.Unmarshal(cpBytes, &cp)
		if err != nil {
			fmt.Println("Error retrieving cp " + value)
			return nil, errors.New("Error retrieving cp " + value)
		}

		fmt.Println("Appending CP" + value)
		allCPs = append(allCPs, cp)
	}

	return allCPs, nil
}

func GetCP(cpid string, stub shim.ChaincodeStubInterface) (CP, error) {
	var cp CP

	cpBytes, err := stub.GetState(cpid)
	if err != nil {
		fmt.Println("Error retrieving cp " + cpid)
		return cp, errors.New("Error retrieving cp " + cpid)
	}

	err = json.Unmarshal(cpBytes, &cp)
	if err != nil {
		fmt.Println("Error unmarshalling cp " + cpid)
		return cp, errors.New("Error unmarshalling cp " + cpid)
	}

	return cp, nil
}

func GetCompany(companyID string, stub shim.ChaincodeStubInterface) (Account, error) {
	var company Account
	companyBytes, err := stub.GetState(accountPrefix + companyID)
	if err != nil {
		fmt.Println("Account not found " + companyID)
		return company, errors.New("Account not found " + companyID)
	}

	err = json.Unmarshal(companyBytes, &company)
	if err != nil {
		fmt.Println("Error unmarshalling account " + companyID + "\n err:" + err.Error())
		return company, errors.New("Error unmarshalling account " + companyID)
	}

	return company, nil
}

// Still working on this one
func (t *SimpleChaincode) transferPaper(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	/*		0
		json
	  	{
			  "CUSIP": "",
			  "fromCompany":"",
			  "toCompany":"",
			  "quantity": 1
		}
	*/
	//need one arg
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting commercial paper record")
	}

	var tr Transaction

	fmt.Println("Unmarshalling Transaction")
	err := json.Unmarshal([]byte(args[0]), &tr)
	if err != nil {
		fmt.Println("Error Unmarshalling Transaction")
		return nil, errors.New("Invalid commercial paper issue")
	}

	fmt.Println("Getting State on CP " + tr.CUSIP)
	cpBytes, err := stub.GetState(cpPrefix + tr.CUSIP)
	if err != nil {
		fmt.Println("CUSIP not found")
		return nil, errors.New("CUSIP not found " + tr.CUSIP)
	}

	var cp CP
	fmt.Println("Unmarshalling CP " + tr.CUSIP)
	err = json.Unmarshal(cpBytes, &cp)
	if err != nil {
		fmt.Println("Error unmarshalling cp " + tr.CUSIP)
		return nil, errors.New("Error unmarshalling cp " + tr.CUSIP)
	}

	var fromCompany Account
	fmt.Println("Getting State on fromCompany " + tr.FromCompany)
	fromCompanyBytes, err := stub.GetState(accountPrefix + tr.FromCompany)
	if err != nil {
		fmt.Println("Account not found " + tr.FromCompany)
		return nil, errors.New("Account not found " + tr.FromCompany)
	}

	fmt.Println("Unmarshalling FromCompany ")
	err = json.Unmarshal(fromCompanyBytes, &fromCompany)
	if err != nil {
		fmt.Println("Error unmarshalling account " + tr.FromCompany)
		return nil, errors.New("Error unmarshalling account " + tr.FromCompany)
	}

	var toCompany Account
	fmt.Println("Getting State on ToCompany " + tr.ToCompany)
	toCompanyBytes, err := stub.GetState(accountPrefix + tr.ToCompany)
	if err != nil {
		fmt.Println("Account not found " + tr.ToCompany)
		return nil, errors.New("Account not found " + tr.ToCompany)
	}

	fmt.Println("Unmarshalling tocompany")
	err = json.Unmarshal(toCompanyBytes, &toCompany)
	if err != nil {
		fmt.Println("Error unmarshalling account " + tr.ToCompany)
		return nil, errors.New("Error unmarshalling account " + tr.ToCompany)
	}

	// Check for all the possible errors
	ownerFound := false
	quantity := 0
	for _, owner := range cp.Owners {
		if owner.Company == tr.FromCompany {
			ownerFound = true
			quantity = owner.Quantity
		}
	}

	// If fromCompany doesn't own this paper
	if ownerFound == false {
		fmt.Println("The company " + tr.FromCompany + "doesn't own any of this paper")
		return nil, errors.New("The company " + tr.FromCompany + "doesn't own any of this paper")
	} else {
		fmt.Println("The FromCompany does own this paper")
	}

	// If fromCompany doesn't own enough quantity of this paper
	if quantity < tr.Quantity {
		fmt.Println("The company " + tr.FromCompany + "doesn't own enough of this paper")
		return nil, errors.New("The company " + tr.FromCompany + "doesn't own enough of this paper")
	} else {
		fmt.Println("The FromCompany owns enough of this paper")
	}

	amountToBeTransferred := float64(tr.Quantity) * cp.Par
	amountToBeTransferred -= (amountToBeTransferred) * (cp.Discount / 100.0) * (float64(cp.Maturity) / 360.0)

	// If toCompany doesn't have enough cash to buy the papers
	if toCompany.CashBalance < amountToBeTransferred {
		fmt.Println("The company " + tr.ToCompany + "doesn't have enough cash to purchase the papers")
		return nil, errors.New("The company " + tr.ToCompany + "doesn't have enough cash to purchase the papers")
	} else {
		fmt.Println("The ToCompany has enough money to be transferred for this paper")
	}

	toCompany.CashBalance -= amountToBeTransferred
	fromCompany.CashBalance += amountToBeTransferred

	toOwnerFound := false
	for key, owner := range cp.Owners {
		if owner.Company == tr.FromCompany {
			fmt.Println("Reducing Quantity from the FromCompany")
			cp.Owners[key].Quantity -= tr.Quantity
			//			owner.Quantity -= tr.Quantity
		}
		if owner.Company == tr.ToCompany {
			fmt.Println("Increasing Quantity from the ToCompany")
			toOwnerFound = true
			cp.Owners[key].Quantity += tr.Quantity
			//			owner.Quantity += tr.Quantity
		}
	}

	if toOwnerFound == false {
		var newOwner Owner
		fmt.Println("As ToOwner was not found, appending the owner to the CP")
		newOwner.Quantity = tr.Quantity
		newOwner.Company = tr.ToCompany
		cp.Owners = append(cp.Owners, newOwner)
	}

	fromCompany.AssetsIds = append(fromCompany.AssetsIds, tr.CUSIP)

	// Write everything back
	// To Company
	toCompanyBytesToWrite, err := json.Marshal(&toCompany)
	if err != nil {
		fmt.Println("Error marshalling the toCompany")
		return nil, errors.New("Error marshalling the toCompany")
	}
	fmt.Println("Put state on toCompany")
	err = stub.PutState(accountPrefix+tr.ToCompany, toCompanyBytesToWrite)
	if err != nil {
		fmt.Println("Error writing the toCompany back")
		return nil, errors.New("Error writing the toCompany back")
	}

	// From company
	fromCompanyBytesToWrite, err := json.Marshal(&fromCompany)
	if err != nil {
		fmt.Println("Error marshalling the fromCompany")
		return nil, errors.New("Error marshalling the fromCompany")
	}
	fmt.Println("Put state on fromCompany")
	err = stub.PutState(accountPrefix+tr.FromCompany, fromCompanyBytesToWrite)
	if err != nil {
		fmt.Println("Error writing the fromCompany back")
		return nil, errors.New("Error writing the fromCompany back")
	}

	// cp
	cpBytesToWrite, err := json.Marshal(&cp)
	if err != nil {
		fmt.Println("Error marshalling the cp")
		return nil, errors.New("Error marshalling the cp")
	}
	fmt.Println("Put state on CP")
	err = stub.PutState(cpPrefix+tr.CUSIP, cpBytesToWrite)
	if err != nil {
		fmt.Println("Error writing the cp back")
		return nil, errors.New("Error writing the cp back")
	}

	fmt.Println("Successfully completed Invoke")
	return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//need one arg
	if len(args) < 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ......")
	}

	if args[0] == "GetAllCPs" {
		fmt.Println("Getting all CPs")
		allCPs, err := GetAllCPs(stub)
		if err != nil {
			fmt.Println("Error from getallcps")
			return nil, err
		} else {
			allCPsBytes, err1 := json.Marshal(&allCPs)
			if err1 != nil {
				fmt.Println("Error marshalling allcps")
				return nil, err1
			}
			fmt.Println("All success, returning allcps")
			return allCPsBytes, nil
		}
	} else if args[0] == "GetCP" {
		fmt.Println("Getting particular cp")
		cp, err := GetCP(args[1], stub)
		if err != nil {
			fmt.Println("Error Getting particular cp")
			return nil, err
		} else {
			cpBytes, err1 := json.Marshal(&cp)
			if err1 != nil {
				fmt.Println("Error marshalling the cp")
				return nil, err1
			}
			fmt.Println("All success, returning the cp")
			return cpBytes, nil
		}
	} else if args[0] == "GetCompany" {
		fmt.Println("Getting the company")
		company, err := GetCompany(args[1], stub)
		if err != nil {
			fmt.Println("Error from getCompany")
			return nil, err
		} else {
			companyBytes, err1 := json.Marshal(&company)
			if err1 != nil {
				fmt.Println("Error marshalling the company")
				return nil, err1
			}
			fmt.Println("All success, returning the company")
			return companyBytes, nil
		}
	} else if args[0] == "GetAllProperties" {
		fmt.Println("Getting all CPs")
		allProperties, err := GetAllProperties(stub)
		if err != nil {
			fmt.Println("Error from GetAllProperties")
			return nil, err
		} else {
			allPropertiesBytes, err1 := json.Marshal(&allProperties)
			if err1 != nil {
				fmt.Println("Error marshalling allcps")
				return nil, err1
			}
			fmt.Println("All success, returning allcps")
			return allPropertiesBytes, nil
		}
	} else if args[0] == "GetAllproposal" {
		fmt.Println("Getting all Proposal")
		allProposal, err := GetAllproposal(stub)
		if err != nil {
			fmt.Println("Error from GetAllproposal")
			return nil, err
		} else {
			allProposalBytes, err1 := json.Marshal(&allProposal)
			if err1 != nil {
				fmt.Println("Error marshalling allProposal")
				return nil, err1
			}
			fmt.Println("All success, returning allProposal")
			return allProposalBytes, nil
		}
	} else if args[0] == "GetAllQuotes" {
		fmt.Println("Getting all Quote")
		allquote, err := GetAllQuotes(stub)
		if err != nil {
			fmt.Println("Error from GetAllQuotes")
			return nil, err
		} else {
			allquoteBytes, err1 := json.Marshal(&allquote)
			if err1 != nil {
				fmt.Println("Error marshalling allquote")
				return nil, err1
			}
			fmt.Println("All success, returning allquote")
			return allquoteBytes, nil
		}
	}else if args[0] == "GetAllAgreement" {
		fmt.Println("Getting all Agreement")
		allSaleAgreement, err := GetAllAgreement(stub)
		if err != nil {
			fmt.Println("Error from GetAllAgreement")
			return nil, err
		} else {
			allSaleAgreementBytes, err1 := json.Marshal(&allSaleAgreement)
			if err1 != nil {
				fmt.Println("Error marshalling allSaleAgreement")
				return nil, err1
			}
			fmt.Println("All success, returning allSaleAgreement")
			return allSaleAgreementBytes, nil
		}
	} else if args[0] == "GetAllDeed" {
		fmt.Println("Getting all Deed")
		allSaleDeed, err := GetAllDeed(stub)
		if err != nil {
			fmt.Println("Error from GetAllDeed")
			return nil, err
		} else {
			allSaleDeedBytes, err1 := json.Marshal(&allSaleDeed)
			if err1 != nil {
				fmt.Println("Error marshalling allSaleDeed")
				return nil, err1
			}
			fmt.Println("All success, returning allSaleDeed")
			return allSaleDeedBytes, nil
		}
	} else if args[0] == "GetAllNotifications" {
		fmt.Println("Getting all Notification")
		allNotification, err := GetAllNotifications(stub)
		if err != nil {
			fmt.Println("Error from GetAllNotifications")
			return nil, err
		} else {
			allNotificationBytes, err1 := json.Marshal(&allNotification)
			if err1 != nil {
				fmt.Println("Error marshalling allNotification")
				return nil, err1
			}
			fmt.Println("All success, returning allNotification")
			return allNotificationBytes, nil
		}
	} else {
		fmt.Println("Generic Query call")
		bytes, err := stub.GetState(args[0])

		if err != nil {
			fmt.Println("Some error happenend")
			return nil, errors.New("Some Error happened")
		}

		fmt.Println("All success, returning from generic")
		return bytes, nil
	}
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	if function == "issueCommercialPaper" {
		fmt.Println("Firing issueCommercialPaper")
		//Create an asset with some value
		return t.issueCommercialPaper(stub, args)
	} else if function == "transferPaper" {
		fmt.Println("Firing cretransferPaperateAccounts")
		return t.transferPaper(stub, args)
	} else if function == "createAccounts" {
		fmt.Println("Firing createAccounts")
		return t.createAccounts(stub, args)
	} else if function == "createAccount" {
		fmt.Println("Firing createAccount")
		return t.createAccount(stub, args)
	} else if function == "init" {
		fmt.Println("Firing init")
		return t.Init(stub, "init", args)
	} else if function == "issueQuote" { //Added for Trade finance 
		fmt.Println("Firing issueQuote")
		return t.issueQuote(stub, args)
	} else if function == "addProperty" { //Added for Trade finance 
		fmt.Println("Firing addProperty")
		return t.addProperty(stub, args)
	} else if function == "issueProposal" { //Added for Trade finance 
		fmt.Println("Firing issueProposal")
		return t.issueProposal(stub, args)
	} else if function == "issueSaleAgreement" { //Added for Trade finance 
		fmt.Println("Firing issueSaleAgreement")
		return t.issueSaleAgreement(stub, args)
	} else if function == "issueSaleDeeds" { //Added for Trade finance 
		fmt.Println("Firing issueSaleDeeds")
		return t.issueSaleDeeds(stub, args)
	} else if function == "addNotification" { //Added for Trade finance 
		fmt.Println("Firing addNotification")
		return t.addNotification(stub, args)
	}


	return nil, errors.New("Received unknown function invocation")
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("Error starting Simple chaincode: %s", err)
	}
}

//lookup tables for last two digits of CUSIP
var seventhDigit = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "P",
	15: "Q",
	16: "R",
	17: "S",
	18: "T",
	19: "U",
	20: "V",
	21: "W",
	22: "X",
	23: "Y",
	24: "Z",
}

var eigthDigit = map[int]string{
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
	16: "G",
	17: "H",
	18: "J",
	19: "K",
	20: "L",
	21: "M",
	22: "N",
	23: "P",
	24: "Q",
	25: "R",
	26: "S",
	27: "T",
	28: "U",
	29: "V",
	30: "W",
	31: "X",
}
