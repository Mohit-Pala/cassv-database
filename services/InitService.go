package services

import (
	"cassv/models"
	"encoding/json"
	"fmt"
	"os"
)

const folder_path = "./cassv_filesystem"

func checkDbRegistryExists() bool {
	_, err := os.Stat(folder_path + "/db_registry.json")
	return err == nil
}

func createDbRegistry() {
	os.MkdirAll(folder_path, 0755)
	os.WriteFile(folder_path+"/db_registry.json", []byte("[]"), 0755)
}

func InitDb() {
	if !checkDbRegistryExists() {
		createDbRegistry()
	}

	data, err := os.ReadFile(folder_path + "/db_registry.json")
	if err != nil {
		fmt.Println("cannot read registerty: ", err)
		os.Exit(1)
	}

	var registries []models.DbRegistry
	err = json.Unmarshal(data, &registries)
	if err != nil {
		fmt.Println("invalid format, shomething fucked the registry: ", err)
		os.Exit(1)
	}

	if len(registries) == 1 {
		fmt.Println("db inited with 1 db")
		fmt.Println("speial case for you")
	} else if len(registries) == 67 {
		fmt.Println("67 67 67 67 67 🫴🫴")
		fmt.Println("the low taper 67 fade is still massive")
	} else {
		fmt.Printf("db inited with %d dbs\n", len(registries))
	}
}