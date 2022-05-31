package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

func SetQueryParams(params any, q *url.Values) error {
	var paramsMap map[string]interface{}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to Marshal params: %v", err)
	}

	err = json.Unmarshal(jsonParams, &paramsMap)
	if err != nil {
		return fmt.Errorf("failed Unmarshal: %v", err)
	}

	setQueryParamsFromMap(paramsMap, q)
	return nil
}

func setQueryParamsFromMap(paramsMap map[string]interface{}, params *url.Values) {
	for key, value := range paramsMap {
		if value == nil {
			continue
		}

		switch v := value.(type) {
		case string:
			params.Set(key, v)
		case int:
			params.Set(key, strconv.Itoa(v))
		case float64:
			params.Set(key, strconv.FormatFloat(v, 'f', 0, 64))
		case map[string]interface{}:
			setQueryParamsFromMap(v, params)
		default:
			log.Printf("%s Type is %T. continued", key, v)
			continue
		}
	}
}

func firstLetterToLower(value string) string {
	v := []byte(value)

	v[0] = v[0] | ('a' - 'A')

	return string(v)
}
