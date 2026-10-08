package service

import "htmx-server/shared/types"

var DataMap = &map[string][]types.Book{}

func GetDataMap() map[string][]types.Book {
	return *DataMap
}
