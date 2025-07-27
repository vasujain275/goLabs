package store

import (
	"errors"
	"strconv"
)


type DataStore interface{
	CreateShortCode(url string) string
	GetShortCode(code string) (string,error)
}

type MemoryStore struct{
	Urls map[string]string
	Id int
}

func (m *MemoryStore) CreateShortCode(url string) string {
	m.Id++
	code := "url-" + strconv.Itoa(m.Id)
	m.Urls[code] = url

	return code
}

func (m *MemoryStore) GetShortCode(code string) (string,error){

	url,ok := m.Urls[code]

	if !ok {
		return "", errors.New("Short Code Not Found")
	}

	return url,nil
}
