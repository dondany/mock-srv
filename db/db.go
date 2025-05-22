package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

type Database struct {
	Filename string
	Data     map[string][]map[string]interface{}
}

func LoadDatabase(filename string) (*Database, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error("Couldn't read database file", err)
		return nil, err
	}

	var dataMap map[string][]map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		slog.Error("Couldn't parse database file", err)
		return nil, err
	}
	return &Database{filename, dataMap}, nil
}

func (db Database) Save() error {
	data, err := json.MarshalIndent(db.Data, "", " ")
	if err != nil {
		slog.Error("Couldn't marshal database file", err)
		return err
	}
	err = os.WriteFile(db.Filename, data, 0644)
	return nil
}

func (db Database) FindById(root string, id float64) (map[string]interface{}, error) {
	rootElement, ok := db.Data[root]
	if !ok {
		slog.Error("Couldn't find root element", root)
		return nil, errors.New("not found")
	}

	for _, element := range rootElement {
		if element["id"] == id {
			return element, nil
		}
	}

	return nil, errors.New("not found")
}

func (db Database) FindAll(root string, filters map[string][]string) ([]map[string]interface{}, error) {
	rootElement, ok := db.Data[root]
	if !ok {
		slog.Error("Couldn't find root element", root)
		return nil, errors.New("not found")
	}

	var filteredElements []map[string]interface{}
	for _, item := range rootElement {
		matches := true
		for key, values := range filters {
			itemVal, ok := item[key]
			if !ok || fmt.Sprint(itemVal) != values[0] {
				matches = false
				break
			}
		}
		if matches {
			filteredElements = append(filteredElements, item)
		}
	}

	return filteredElements, nil
}

func (db Database) Insert(root string, element map[string]interface{}) error {
	rootElement, ok := db.Data[root]
	if !ok {
		slog.Error("Couldn't find root element", root)
		return errors.New("not found")
	}

	id := generateNextId(rootElement)
	element["id"] = id

	rootElement = append(rootElement, element)
	db.Data[root] = rootElement
	db.Save()
	return nil
}

func (db Database) Update(root string, id float64, updates map[string]interface{}) (map[string]interface{}, error) {
	rootElement, ok := db.Data[root]
	if !ok {
		slog.Error("Couldn't find root element", root)
		return nil, errors.New("not found")
	}

	index := -1
	for i, element := range rootElement {
		if element["id"] == id {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, errors.New("not found")
	}

	updates["id"] = id
	rootElement[index] = updates
	db.Save()
	return updates, nil
}

func (db Database) Delete(root string, id float64) error {
	rootElement, ok := db.Data[root]
	if !ok {
		slog.Error("Couldn't find root element", root)
		return errors.New("not found")
	}

	var filteredElement []map[string]interface{}
	for _, element := range rootElement {
		if element["id"] != id {
			filteredElement = append(filteredElement, element)
		}
	}
	if len(filteredElement) == len(rootElement) {
		return errors.New("not found")
	}

	db.Data[root] = filteredElement
	db.Save()
	return nil
}

func generateNextId(collection []map[string]interface{}) float64 {
	var max float64 = 0
	for _, item := range collection {
		if idVal, ok := item["id"].(float64); ok && idVal > max {
			max = idVal
		}
	}
	return max + 1
}
