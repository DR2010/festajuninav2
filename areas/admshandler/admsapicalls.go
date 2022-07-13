// Package btcmarketshandler API calls for dishes web
// --------------------------------------------------------------
// .../src/restauranteweb/areas/btcmarkets/btcmarketscalls.go
// --------------------------------------------------------------
package admshandler

import (
	"encoding/json"
	models "festajuninav2/models"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

// TrainingContractGet works
func TrainingContractGet(httpwriter http.ResponseWriter, redisclient *redis.Client, credentials models.Credentials) TrainingContract {

	// var apiserver string
	var emptydisplay TrainingContract

	urlrequest := "https://proxy.admsapi.australianapprenticeships.gov.au/integration/provider/trainingcontract/v1/training-contracts/18"

	fmt.Println("#1 TrainingContractGet...")

	url := fmt.Sprint(urlrequest)

	fmt.Println("#2 TrainingContractGet...")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		fmt.Println("NewRequest: ", err)
		return emptydisplay
	}

	fmt.Println("#3 TrainingContractGet... before GetToken()")

	token := GetToken()

	// via portal requires header
	// req.Header.Set("Ocp-Apim-Subscription-Key", "eb9f7b1620494fb2bdc7815705fd8c7e")
	req.Header.Set("Ocp-Apim-Subscription-Key", "f2c6767fdc4845e7ab9302d7826a9ad5")
	// req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IjZ6U0NrVDQwcmRIV0lsWkZoRzd4OTVvRGlRTSIsImtpZCI6IjZ6U0NrVDQwcmRIV0lsWkZoRzd4OTVvRGlRTSJ9.eyJhdWQiOiJ1cm46c2tsOnBwcm9kOmFwaTphZG1zOnN0YWdpbmciLCJpc3MiOiJodHRwOi8vc3RzLnhhdXRobnAuZGlzLmdvdi5hdS9hZGZzL3NlcnZpY2VzL3RydXN0IiwiaWF0IjoxNjUzODk2NDM1LCJuYmYiOjE2NTM4OTY0MzUsImV4cCI6MTY1MzkwMDAzNSwiZm9yd2FyZGVkY2xpZW50aXAiOiIyMTEuMzEuMTc2LjE3IiwiaHR0cDovL3d3dy5kZWV3ci5nb3YuYXUvMjAxOC8wMy9lc2FtL3ByaW1hcnlwaG9uZW51bWJlciI6IjA0MDg0MjAyMDEiLCJodHRwOi8vd3d3LmRlZXdyLmdvdi5hdS8yMDE3LzAyL2VzYW0vdXNlcmRlY2xhcmF0aW9uYWNjZXB0ZWQiOiIwIiwiaHR0cDovL3d3dy5kZWV3ci5nb3YuYXUvMjAxNy8wMi9lc2FtL3VzZXJhcHBsaWNhdGlvbmFjY291bnRndWlkIjoiQjRBQzk3NjEtODk4Ny00MzQwLUIzRDAtRTA0MTJENjU5NjVEIiwiaHR0cDovL3d3dy5kZWV3ci5nb3YuYXUvMjAxNy8wMi9lc2FtL2NyZWRlbnRpYWxlbmFibGVkIjoiMSIsImh0dHA6Ly93d3cuZGVld3IuZ292LmF1LzIwMTcvMDIvZXNhbS9jb3JlaWRlbnRpdHllbmFibGVkIjoiMSIsImh0dHA6Ly93d3cuZGVld3IuZ292LmF1LzIwMTcvMDIvZXNhbS9hcHBsaWNhdGlvbmFjY291bnRlbmFibGVkIjoiMSIsImZhbWlseV9uYW1lIjoiQ2xhcmtlIiwidW5pcXVlX25hbWUiOiJSRFZXSE01NCIsImdpdmVuX25hbWUiOiJHbGVubiIsImVtYWlsIjoiZ2xlbm5jbGFya2U2N0BnbWFpbC5jb20iLCJodHRwOi8vZGVzZS5nb3YuYXUvYWRtcy9jbGFpbXMvc2l0ZSI6Ik5PU0MiLCJodHRwOi8vZGVzZS5nb3YuYXUvYWRtcy9jbGFpbXMvb3JnIjoiREVQVCIsImh0dHA6Ly9kZXNlLmdvdi5hdS9hZG1zL2NsYWltcy9iYXNlcm9sZSI6IkRQVlciLCJodHRwOi8vZGVld3IuZ292LmF1L2VzLzIwMTgvMDYvY2xhaW1zL3NlcnZpY2VpZGVudGl0eSI6IlkiLCJodHRwOi8vZGVld3IuZ292LmF1L2VzLzIwMTEvMDMvY2xhaW1zL2xhc3RMb2dvbkRhdGVUaW1lU3RhbXAiOiIyMDIyLTA1LTMwIDE3OjQwOjM1Ljc2MzYxMDMgKzEwOjAwIiwiYXV0aG1ldGhvZCI6WyJodHRwOi8vc2NoZW1hcy5taWNyb3NvZnQuY29tL3dzLzIwMDgvMDYvaWRlbnRpdHkvYXV0aGVudGljYXRpb25tZXRob2QvdGxzY2xpZW50IiwiaHR0cDovL3NjaGVtYXMubWljcm9zb2Z0LmNvbS93cy8yMDA4LzA2L2lkZW50aXR5L2F1dGhlbnRpY2F0aW9ubWV0aG9kL3g1MDkiXSwiYXBwdHlwZSI6IkNvbmZpZGVudGlhbCIsImFwcGlkIjoiZTlhZDY4MmEtODJkMC00Y2EzLTkzNjgtM2NlNWJiYWFjMjY0IiwiYXV0aF90aW1lIjoiMjAyMi0wNS0zMFQwNzo0MDozNS43NTBaIiwidmVyIjoiMS4wIn0.oZfVNLUh3UxQt-SGvAUYof59na0VWMKQP7_f3iDHyPVnzgYhQvGFjTPEf1bIQ4R-kG1ns4thHIbvb2OZ0Ar-tSp3tJ8OkdqCDjym_D_k_g12IdNESxqe9HF5pxKR5HwjlppABugrR9nscJZtY76AmKylYAbpfSFuLChOCjf8PvIJk9BAKFU5eGo4-bhPwy4vFpOZigAgtxZVtEUhkpuO0UAybbF6StMHLqZntSlQectkHS-Zdjd5fH5hLCKBzO2u529HNhJ2qUpmVjTWGR4sgvJeKmd_aFvOvp7kieGnshIwOM6Nis37ZnjC1dKvzR9_r0YoFQVDjkfZzr37T8x3kw")
	// req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IjZ6U0NrVDQwcmRIV0lsWkZoRzd4OTVvRGlRTSIsImtpZCI6IjZ6U0NrVDQwcmRIV0lsWkZoRzd4OTVvRGlRTSJ9.eyJhdWQiOiJ1cm46c2tsOnBwcm9kOmFwaTphZG1zOnN0YWdpbmciLCJpc3MiOiJodHRwOi8vc3RzLnhhdXRobnAuZGlzLmdvdi5hdS9hZGZzL3NlcnZpY2VzL3RydXN0IiwiaWF0IjoxNjUzOTA0MTg1LCJuYmYiOjE2NTM5MDQxODUsImV4cCI6MTY1MzkwNzc4NSwiZm9yd2FyZGVkY2xpZW50aXAiOiIyMTEuMzEuMTc2LjE3IiwiaHR0cDovL3d3dy5kZWV3ci5nb3YuYXUvMjAxOC8wMy9lc2FtL3ByaW1hcnlwaG9uZW51bWJlciI6IjA0MDg0MjAyMDEiLCJodHRwOi8vd3d3LmRlZXdyLmdvdi5hdS8yMDE3LzAyL2VzYW0vdXNlcmRlY2xhcmF0aW9uYWNjZXB0ZWQiOiIwIiwiaHR0cDovL3d3dy5kZWV3ci5nb3YuYXUvMjAxNy8wMi9lc2FtL3VzZXJhcHBsaWNhdGlvbmFjY291bnRndWlkIjoiQjRBQzk3NjEtODk4Ny00MzQwLUIzRDAtRTA0MTJENjU5NjVEIiwiaHR0cDovL3d3dy5kZWV3ci5nb3YuYXUvMjAxNy8wMi9lc2FtL2NyZWRlbnRpYWxlbmFibGVkIjoiMSIsImh0dHA6Ly93d3cuZGVld3IuZ292LmF1LzIwMTcvMDIvZXNhbS9jb3JlaWRlbnRpdHllbmFibGVkIjoiMSIsImh0dHA6Ly93d3cuZGVld3IuZ292LmF1LzIwMTcvMDIvZXNhbS9hcHBsaWNhdGlvbmFjY291bnRlbmFibGVkIjoiMSIsImZhbWlseV9uYW1lIjoiQ2xhcmtlIiwidW5pcXVlX25hbWUiOiJSRFZXSE01NCIsImdpdmVuX25hbWUiOiJHbGVubiIsImVtYWlsIjoiZ2xlbm5jbGFya2U2N0BnbWFpbC5jb20iLCJodHRwOi8vZGVzZS5nb3YuYXUvYWRtcy9jbGFpbXMvc2l0ZSI6Ik5PU0MiLCJodHRwOi8vZGVzZS5nb3YuYXUvYWRtcy9jbGFpbXMvb3JnIjoiREVQVCIsImh0dHA6Ly9kZXNlLmdvdi5hdS9hZG1zL2NsYWltcy9iYXNlcm9sZSI6IkRQVlciLCJodHRwOi8vZGVld3IuZ292LmF1L2VzLzIwMTgvMDYvY2xhaW1zL3NlcnZpY2VpZGVudGl0eSI6IlkiLCJodHRwOi8vZGVld3IuZ292LmF1L2VzLzIwMTEvMDMvY2xhaW1zL2xhc3RMb2dvbkRhdGVUaW1lU3RhbXAiOiIyMDIyLTA1LTMwIDE5OjQ5OjQ1LjYyMjY4MzYgKzEwOjAwIiwiYXV0aG1ldGhvZCI6WyJodHRwOi8vc2NoZW1hcy5taWNyb3NvZnQuY29tL3dzLzIwMDgvMDYvaWRlbnRpdHkvYXV0aGVudGljYXRpb25tZXRob2QvdGxzY2xpZW50IiwiaHR0cDovL3NjaGVtYXMubWljcm9zb2Z0LmNvbS93cy8yMDA4LzA2L2lkZW50aXR5L2F1dGhlbnRpY2F0aW9ubWV0aG9kL3g1MDkiXSwiYXBwdHlwZSI6IkNvbmZpZGVudGlhbCIsImFwcGlkIjoiZTlhZDY4MmEtODJkMC00Y2EzLTkzNjgtM2NlNWJiYWFjMjY0IiwiYXV0aF90aW1lIjoiMjAyMi0wNS0zMFQwOTo0OTo0NS42MjRaIiwidmVyIjoiMS4wIn0.yA5BYsnBSI0qDlsJjY7zFr7niqI5tA2NrBUmCjJ32xfyHh-e3aT3EySUuHLnCkkCPOwef63ZDXicnEQ1UY8qtEGIWEy4amNzOhANTC_HmpuOCXPxLvpXAKkwDTUMO-LI_ut8_Jn4fSeDEt1mLujgCteFfZuW9qLzN9N58WKjXCkdP5hyd1DSC38qoMp1EX2iukeHKV60aDswxzcKReBzNOaAlINNVwIXO12NntVLk7XizeLO3ByfYj8MVRn8DzPicU1xSAXM1IHqOCMIZF3syXTc8d86uDstJrQdGoZbZWOJQPiu8OUISZXnAroxhQ7o7M6oBNSdg3sx-QM1P_39gw")
	req.Header.Set("Authorization", "Bearer "+token)

	fmt.Println("#4 TrainingContractGet... after GetToken()")
	fmt.Println(" token ...")
	fmt.Println(token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		fmt.Println("Do: ", err)
		return emptydisplay
	}

	fmt.Println("#5 TrainingContractGet...")

	defer resp.Body.Close()

	var tcreturned TrainingContract

	if resp.StatusCode == 401 {
		fmt.Println(" resp.StatusCode ", resp.StatusCode)
		fmt.Println(" err = ", err)
		return tcreturned
	}

	fmt.Println("#6 TrainingContractGet...")

	if err := json.NewDecoder(resp.Body).Decode(&tcreturned); err != nil {
		log.Println(err)
		fmt.Println("err json casting: ", err)
	}

	fmt.Println("#7 TrainingContractGet...")

	fmt.Println("Body: ", resp.Body)
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println("Status): ", resp.Status)

	fmt.Println("tc ApprenticeID: ", tcreturned.ApprenticeID)
	fmt.Println("tc FirstName: ", tcreturned.Apprentice.FirstName)
	fmt.Println("tc Surname: ", tcreturned.Apprentice.Surname)

	fmt.Println("tc returned full: ", tcreturned)

	return tcreturned
}

// TrainingContractGet works
func TrainingContractList(httpwriter http.ResponseWriter, redisclient *redis.Client, credentials models.Credentials) []TrainingContractSummary {

	var emptydisplay []TrainingContractSummary

	urlrequest := "https://proxy.admsapi.australianapprenticeships.gov.au/integration/provider/trainingcontract/v1/training-contracts"

	url := fmt.Sprint(urlrequest)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		fmt.Println("NewRequest: ", err)
		return emptydisplay
	}

	token := GetToken()

	req.Header.Set("Ocp-Apim-Subscription-Key", "f2c6767fdc4845e7ab9302d7826a9ad5")
	req.Header.Set("Authorization", "Bearer "+token)

	// fmt.Println("#4 TrainingContractGet... after GetToken()")
	// fmt.Println(" token ...")
	// fmt.Println(token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		fmt.Println("Do: ", err)
		return emptydisplay
	}

	fmt.Println("#5 TrainingContractGet...")

	defer resp.Body.Close()

	var tcreturned []TrainingContractSummary

	if resp.StatusCode == 401 {
		fmt.Println(" resp.StatusCode ", resp.StatusCode)
		fmt.Println(" err = ", err)
		return tcreturned
	}

	if err := json.NewDecoder(resp.Body).Decode(&tcreturned); err != nil {
		log.Println(err)
		fmt.Println("err json casting: ", err)
	}

	fmt.Println("#7 TrainingContractGet...")

	fmt.Println("Body: ", resp.Body)
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println("Status): ", resp.Status)

	fmt.Println("tc ApprenticeID: ", tcreturned[0].ID)
	fmt.Println("tc FirstName: ", tcreturned[0].ApprenticeFirstName)
	fmt.Println("tc Surname: ", tcreturned[0].ApprenticeSurname)

	fmt.Println("tc returned full: ", tcreturned)

	return tcreturned
}

// ---------------------------------------------------------------------------------
//  GetToken()
// ---------------------------------------------------------------------------------
//
func GetToken() string {

	// var apiserver string
	var emptydisplay string

	// urlrequest := "http://localhost:1411/ADMSSTS/gettoken"
	urlrequest := "http://localhost:1411/ADMSSTS/getfulltoken"

	fmt.Println("#1 GetToken...")

	url := fmt.Sprint(urlrequest)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return emptydisplay
	}

	req.Header.Set("Content-Type", "text/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return emptydisplay
	}

	defer resp.Body.Close()

	var token models.TokenStruct

	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		log.Println(err)
		fmt.Println("err json casting: ", err)
	}

	if resp.StatusCode == 401 {
		fmt.Println(" resp.StatusCode ", resp.StatusCode)
		fmt.Println(" err = ", err)
		return emptydisplay
	}

	return token.AccessToken

}
