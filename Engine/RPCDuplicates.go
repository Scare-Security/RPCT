package RPC

func LocateUni(s []string) []string {
	res := make(map[string]bool)
	var result []string
	for _, s := range s {
		if _, ok := res[s]; !ok {
			res[s] = true
			result = append(result, s)
		}
	}
	return result
}
