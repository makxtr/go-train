package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

/**
In the Go file, write a program to perform a GET request on the route htttp://coderbyte.com/api/challenges/json/wizard-list and sort the map keys alphabetically.
The sorting should be case-insensitive, and the original data structure should be preserved (e.g., slices should remain slices, maps should remain maps).
Keep in mind that while the JSON format uses null, in Go, this is represented as nil.

Next, remove any duplicate maps from slices.
Two maps are considered duplicates if they have the same keys and values in the same order.
Only the first occurrence should be preserved when a slice contains duplicate maps.

Finally, remove any map properties with all values set to an empty string or nil (equivalent to null in JSON).
If all values in a map are removed, retain the key with an empty map {} instead of removing the property entirely.
Then print a list of modified maps. Be sure to use json.Marshal on the final list to convert it back to JSON format.
*/

func main() {
	resp, err := http.Get("https://coderbyte.com/api/challenges/json/wizard-list")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error:", err)
		return
	}

	processed := Process(data)

	result, err := json.Marshal(processed)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(result))
}

type CustomMap map[string]interface{}

func (m CustomMap) MarshalJSON() ([]byte, error) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		ki, kj := keys[i], keys[j]
		li, lj := strings.ToLower(ki), strings.ToLower(kj)
		if li != lj {
			return li < lj
		}
		return ki < kj
	})

	var buf strings.Builder
	buf.WriteByte('{')
	for i, k := range keys {
		if i > 0 {
			buf.WriteByte(',')
		}
		keyBytes, _ := json.Marshal(k)
		buf.Write(keyBytes)
		buf.WriteByte(':')
		valBytes, err := json.Marshal(m[k])
		if err != nil {
			return nil, err
		}
		buf.Write(valBytes)
	}
	buf.WriteByte('}')
	return []byte(buf.String()), nil
}

func Process(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		newMap := make(CustomMap)
		for k, val := range v {
			processedVal := Process(val)
			if processedVal == nil {
				continue
			}
			if str, ok := processedVal.(string); ok && str == "" {
				continue
			}
			newMap[k] = processedVal
		}
		return newMap
	case []interface{}:
		newSlice := make([]interface{}, 0, len(v))
		seen := make(map[string]bool)
		for _, val := range v {
			processedVal := Process(val)

			if cm, ok := processedVal.(CustomMap); ok {
				bytes, _ := json.Marshal(cm)
				str := string(bytes)
				if seen[str] {
					continue
				}
				seen[str] = true
			}

			newSlice = append(newSlice, processedVal)
		}
		return newSlice
	default:
		return v
	}
}
