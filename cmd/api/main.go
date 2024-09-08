package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	"users/config"
	"users/service"

	http_server "users/api/http"

	"github.com/brianvoe/gofakeit"
	"golang.org/x/exp/rand"
)

var configPath = flag.String("config", "", "configuration path")

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone"`
	Password    string    `json:"password"`
	Addresses   []Address `json:"addresses"`
}

func generateUsers(numUsers int) []User {
	gofakeit.Seed(0)
	users := make([]User, numUsers)

	for i := 0; i < numUsers; i++ {
		numAddresses := rand.Intn(5) + 1 // Each user can have between 1 to 5 addresses
		addresses := make([]Address, numAddresses)

		for j := 0; j < numAddresses; j++ {
			addresses[j] = Address{
				Street:  gofakeit.Street(),
				City:    gofakeit.City(),
				State:   gofakeit.State(),
				ZipCode: gofakeit.Zip(),
				Country: gofakeit.Country(),
			}
		}

		users[i] = User{
			ID:          gofakeit.UUID(),
			FirstName:   gofakeit.FirstName(),
			LastName:    gofakeit.LastName(),
			Email:       gofakeit.Email(),
			PhoneNumber: gofakeit.Phone(),
			Password:    gofakeit.Password(true, true, true, true, false, 7),
			Addresses:   addresses,
		}
	}

	return users
}

func saveToJSON(data []User, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
func geturl() string {
	cfg := readConfig()
	host := cfg.Server.Host
	port := cfg.Server.HTTPPort
	url := fmt.Sprintf("%s:%d", host, port)
	return url
}


func createUser(user User) (string, error) {
	
	userJSON, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user data: %v", err)
	}
	createUserEndpoint := fmt.Sprintf("http://%s/api/v1/register", geturl())
	
	resp, err := http.Post(createUserEndpoint, "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		return "", fmt.Errorf("failed to make POST request to register user: %v", err)
	}
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to create user, status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}


	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to extract data from response")
	}
	userID, ok := data["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract user_id from data")
	}

	return userID, nil
}


func createAddress(userID string, address Address) error {
	addressData := map[string]interface{}{
		"user_id":  userID,
		"street":   address.Street,
		"city":     address.City,
		"state":    address.State,
		"zip_code": address.ZipCode,
		"country":  address.Country,
	}


	addressJSON, err := json.Marshal(addressData)
	if err != nil {
		return fmt.Errorf("failed to marshal address data: %v", err)
	}
	createAddressEndpoint := fmt.Sprintf("http://%s/api/v1/addressconc", geturl())
	
	resp, err := http.Post(createAddressEndpoint, "application/json", bytes.NewBuffer(addressJSON))
	if err != nil {
		return fmt.Errorf("failed to make POST request to create address: %v", err)
	}
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create address, status code: %d", resp.StatusCode)
	}

	return nil
}

func ProcessUser(user User) {
	
	userID, err := createUser(user)
	if err != nil {
		log.Printf("Failed to create user: %v\n", err)
		return
	}
	fmt.Printf("User created with ID: %s\n", userID)


	for _, address := range user.Addresses {
		err = createAddress(userID, address)
		if err != nil {
			log.Printf("Failed to create address for user ID %s: %v\n", userID, err)
			return
		}
		fmt.Printf("Address created for user ID: %s\n", userID)
	}
}


func worker(id int, wg *sync.WaitGroup, usersChan <-chan User) {
	defer wg.Done()

	for user := range usersChan {
		fmt.Printf("Worker %d processing user ID: %v\n", id, user.ID)
		ProcessUser(user)
	}
}
func generatefile() {
	const numUsers = 1000 // 1 million users
	const filename = "users_data.json"
	fmt.Printf("Generating %d users with multiple addresses...\n", numUsers)
	usersData := generateUsers(numUsers)
	fmt.Printf("Saving data to %s...\n", filename)
	err := saveToJSON(usersData, filename)
	if err != nil {
		fmt.Printf("Error saving data: %v\n", err)
	} else {
		fmt.Println("Data generation and saving completed.")
	}
}
func addinfo() {
	file, err := os.Open("users_data.json") 
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()


	decoder := json.NewDecoder(file)

	
	if _, err := decoder.Token(); err != nil {
		log.Fatalf("failed to read JSON opening bracket: %v", err)
	}

	
	usersChan := make(chan User, 5)
	var wg sync.WaitGroup

	
	numWorkers := 10
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, usersChan)
	}

	
	for decoder.More() {
		var user User
		
		if err := decoder.Decode(&user); err != nil {
			log.Fatalf("failed to decode user: %v", err)
		}

		
		usersChan <- user
	}

	close(usersChan)

	
	wg.Wait()

	fmt.Println("All users processed successfully!")

}
func main() {
	cfg := readConfig()

	app, err := service.NewAppContainer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		generatefile()
		time.Sleep(5 * time.Second)
		addinfo()
	}()
	http_server.Run(cfg, app)

}

func readConfig() config.Config {
	flag.Parse()

	if cfgPathEnv := os.Getenv("APP_CONFIG_PATH"); len(cfgPathEnv) > 0 {
		*configPath = cfgPathEnv
	}

	if len(*configPath) == 0 {
		log.Fatal("configuration file not found")
	}

	cfg, err := config.ReadStandard(*configPath)

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
