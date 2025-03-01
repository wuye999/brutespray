package modules

import (
	"fmt"
	"os"
	"strconv"
)

func GetUsersAndPasswords(h *Host, user string, password string, version string) ([]string, []string) {
	userCh := make(chan string)
	passCh := make(chan string)

	go func() {
		defer close(userCh)
		if user != "" {
			if IsFile(user) {
				users, err := ReadUsersFromFile(user)
				if err != nil {
					fmt.Println("Error reading user file:", err)
					os.Exit(1)
				}
				for _, u := range users {
					userCh <- u
				}
			} else {
				userCh <- user
			}
		} else {
			var users []string = GetUsersFromDefaultWordlist(version, h.Service)
			for _, u := range users {
				userCh <- u
			}
		}
	}()

	go func() {
		defer close(passCh)
		if password != "" {
			if IsFile(password) {
				passwords, err := ReadPasswordsFromFile(password)
				if err != nil {
					fmt.Println("Error reading password file:", err)
					os.Exit(1)
				}
				for _, p := range passwords {
					passCh <- p
				}
			} else {
				passCh <- password
			}
		} else {
			var passwords []string = GetPasswordsFromDefaultWordlist(version, h.Service)
			for _, p := range passwords {
				passCh <- p
			}
		}
	}()

	userSlice := []string{}
	for u := range userCh {
		userSlice = append(userSlice, u)
	}

	passwordSlice := []string{}
	for p := range passCh {
		passwordSlice = append(passwordSlice, p)
	}

	return userSlice, passwordSlice
}

func CalcCombinations(userCh []string, passCh []string) int {
	var totalCombinations int
	users := []string{}
	passwords := []string{}

	for u := range userCh {
		users = append(users, strconv.Itoa(u))
	}

	for p := range passCh {
		passwords = append(passwords, strconv.Itoa(p))
	}

	totalCombinations = len(users) * len(passwords)
	return totalCombinations
}

func CalcCombinationsPass(passCh []string) int {
	var totalCombinations int
	passwords := []string{}

	for p := range passCh {
		passwords = append(passwords, strconv.Itoa(p))
	}

	totalCombinations = len(passwords)
	return totalCombinations
}
