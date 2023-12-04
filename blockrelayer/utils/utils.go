package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sap200/env"
	"github.com/syndtr/goleveldb/leveldb"
)

func JSONFileToString(filename string) (string, error) {
	// Read the ABI JSON file
	abiJSON, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(abiJSON), nil
}

func InterfaceToBytes(iface interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(iface)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func HashFunc(a string) string {
	hasher := sha256.New()

	// Write the input string to the hash
	hasher.Write([]byte(a))

	// Get the final hash as a byte slice
	hashBytes := hasher.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

func AddToDB(key, value string) {
	db, err := leveldb.OpenFile(env.LOG_DB_NAME, nil)
	if err != nil {
		fmt.Println("Unable to open db")
		return
	}
	defer db.Close()
	// Write data to the database
	k := []byte(key)
	v := []byte(value)
	err = db.Put(k, v, nil)
	if err != nil {
		fmt.Println("Unable to write to DB")
	}
}

func DoesKeyExists(key string) (bool, error) {
	// Open or create a LevelDB database
	db, err := leveldb.OpenFile(env.LOG_DB_NAME, nil)
	if err != nil {
		fmt.Println("Unable to open db")
		return false, err
	}
	defer db.Close()

	k := []byte(key)
	data, err := db.Get(k, nil)
	if err == nil {
		fmt.Println("data Found:: ", string(data))
		return true, nil
	} else if err == leveldb.ErrNotFound {
		return false, nil
	} else {
		return false, err
	}
}
