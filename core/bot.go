package core

import (
	"fmt"
	"net/http"
	"time"

	"agent301/helper"
)

func processResponse(res map[string]interface{}) map[string]interface{} {
	var result map[string]interface{}
	// Mengakses data dari map
	if _, exists := res["msg"].(string); exists && (res["msg"].(string) == "success") {
		// Akses data dari map
		if data, exists := res["data"].(map[string]interface{}); exists {
			result = data
		}
	} else {
		helper.PrettyLog("error", fmt.Sprintf("Request failed: %v ", res["msg"].(string)))
	}

	return result
}

// TODO
func launchBot(username string, query string, apiUrl string, referUrl string, isAutoBindWallet bool) {
	client := &Client{
		apiURL:     apiUrl,
		referURL:   referUrl,
		httpClient: &http.Client{},
	}

	var tasks = map[string]string{
		"tgStatus":    "claimTg",
		"xStatus":     "claimX",
		"shareStatus": "claimShare",
	}

	// Login Account
	req, err := client.login(query)
	if err != nil {
		helper.PrettyLog("error", fmt.Sprintf("Failed to login: %v", err))
		return
	}

	res, err := handleResponse(req)
	if err != nil {
		fmt.Println("Error handling response:", err)
		return
	}

	authToken := processResponse(res)

	if token, exists := authToken["token"].(string); exists {
		client.authToken = token
	} else {
		helper.PrettyLog("error", fmt.Sprintf("%s | Failed Get Access Token...", username))
		return
	}

	if len(client.authToken) > 0 {
		req, err = client.detailAccount()
		if err != nil {
			helper.PrettyLog("error", fmt.Sprintf("Failed to get detail account: %v", err))
			return
		}

		res, err = handleResponse(req)
		if err != nil {
			fmt.Println("Error handling response:", err)
			return
		}

		userData := processResponse(res)

		helper.PrettyLog("success", fmt.Sprintf("%s | Balance: %.0f | Wallet Address: %s", username, userData["totalBonus"].(float64), userData["tonWallet"].(string)))

		// Daily CheckIn
		if fmt.Sprintf("%.0f", userData["checkInDate"].(float64)) != time.Now().Format("20060102") {
			req, err = client.dailyCheckIn(query)
			if err != nil {
				helper.PrettyLog("error", fmt.Sprintf("Failed to daily check in: %v", err))
			}

			res, err = handleResponse(req)
			if err != nil {
				fmt.Println("Error handling response:", err)
			}

			status := processResponse(res)

			if status["result"].(bool) {
				helper.PrettyLog("success", fmt.Sprintf("%s | Daily Check In Successfully...", username))
			} else {
				helper.PrettyLog("error", fmt.Sprintf("%s | Daily Check In Failed | Try Again Letter...", username))
			}
		}

		// Completing Task
		for key, task := range tasks {
			if int(userData[key].(float64)) != 1 {
				req, err = client.claimTask(task)
				if err != nil {
					helper.PrettyLog("error", fmt.Sprintf("Failed to completing %s: %v", task, err))
					continue
				}

				res, err = handleResponse(req)
				if err != nil {
					fmt.Println("Error handling response:", err)
					continue
				}

				status := processResponse(res)

				if task == "claimTg" {
					if int(status["status"].(float64)) != 1 {
						helper.PrettyLog("error", fmt.Sprintf("%s | Failed Completing Task %s | Try Again Letter...", username, task))
					} else {
						helper.PrettyLog("success", fmt.Sprintf("%s | Completing Task %s Successfully...", username, task))
					}
				} else {
					if status["data"] == nil {
						helper.PrettyLog("success", fmt.Sprintf("%s | Completing Task %s Successfully...", username, task))
					} else {
						helper.PrettyLog("error", fmt.Sprintf("%s | Failed Completing Task %s | Try Again Letter...", username, task))
					}
				}
			}
		}

		// Bind Wallet
		if isAutoBindWallet && (userData["tonWallet"].(string) == "") {
			walletFile := helper.ReadFileTxt("wallet.txt")
			if walletFile == nil {
				helper.PrettyLog("error", "Wallet data not found")
				return
			}

			wallets := helper.SplitTextByColon(walletFile)

			isWalletFound := false
			for key, wallet := range wallets {
				if key == username {
					req, err = client.bindWallet(wallet)
					if err != nil {
						helper.PrettyLog("error", fmt.Sprintf("Failed to bind wallet: %v", err))
					}

					res, err = handleResponse(req)
					if err != nil {
						fmt.Println("Error handling response:", err)
					}

					status := processResponse(res)

					if status["result"].(bool) {
						helper.PrettyLog("success", fmt.Sprintf("%s | Bind Wallet Successfully | Wallet Address: %s", username, wallet))
						isWalletFound = true
					} else {
						helper.PrettyLog("error", fmt.Sprintf("%s | Bind Wallet Failed : %v", username, status["result"].(bool)))
						isWalletFound = true
					}
				}

			}

			if !isWalletFound {
				helper.PrettyLog("error", fmt.Sprintf("%s | Bind Wallet Failed | Cannot Find This Account In Wallet File...", username))
			}

		}
	}
}
