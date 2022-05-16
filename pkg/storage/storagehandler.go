package storage

import (
	"encoding/json"
	"os"
)

const (
	filename = "storage.json"
)

var (
	storage = map[string]int{}
)

func Load() error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			if _, err := os.Create(filename); err != nil {
				return err
			}
			if err := os.WriteFile(filename, []byte("{}"), 0644); err != nil {
				return err
			}
			data, err = os.ReadFile(filename)
			if err != nil {
				return err
			}

		}
	}
	err = json.Unmarshal(data, &storage)
	return err
}

func GetViewingCount(videourl string) int {
	if val, ok := storage[videourl]; ok {
		return val
	} else {
		return 0
	}
}

func IncreaseViewingCount(videourl string) {
	if val, ok := storage[videourl]; ok {
		storage[videourl] = val + 1
	} else {
		storage[videourl] = 1
	}
}

func Save() error {
	data, err := json.Marshal(storage)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
