package main

import "fmt"

/**

 m := map[string][]string{
	"group1": {"apple", "banana"},
	"group2": {"carrot"},
	}


	newValues := []string{"banana", "cherry"}
	key := "group1"

	m := map[string][]string{
		"group1": {"apple", "banana", "cherry"}
		"group2": {"carrot"}
	}

*/

func MergeToMap(m map[string][]string, s []string, key string) {
	if _, ok := m[key]; ok {
		set := make(map[string]struct{}, len(m[key]))
		for _, item := range m[key] {
			set[item] = struct{}{}
		}
		for _, n := range s {
			if _, ok = set[n]; !ok {
				m[key] = append(m[key], n)
			}
		}
	}
}

func main() {
	m := map[string][]string{
		"group1": {"apple", "banana"},
		"group2": {"carrot"},
	}

	newValues := []string{"banana", "cherry"}
	key := "group21"

	MergeToMap(m, newValues, key)

	fmt.Println(m)
}
