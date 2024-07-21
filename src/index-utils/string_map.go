package index_utils

func GetStringMap(strings []string) map[string]int {
	cur_map := make(map[string]int)
	for i, currency := range strings {
		cur_map[currency] = i
	}
	return cur_map
}

func GetInverseStringMap(string_map map[string]int) []string {
	inv := make([]string, len(string_map))
	for k := range string_map {
		inv[string_map[k]] = k
	}
	return inv
}
