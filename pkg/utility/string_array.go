package utility

func CopyStrArr(arr []string) []string {
	rt := make([]string, len(arr), len(arr))
	copy(rt, arr)
	return rt
}

func CopyStringGenericMap(m map[string]interface{}) map[string]interface{} {
	rt := make(map[string]interface{}, len(m))
	for k, v := range m {
		rt[k] = v
	}
	return rt
}

func DelFromArr(target string, arr []string) {
	for i, v := range arr {
		if v == target {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
}

// Return true when the array size is exactly 1 and contains target; otherwise false
func Exactly(arr []string, target string) bool {
	return len(arr) == 1 && target == arr[0]
}

// Returns true when array a contains b; otherwise false.
func StrArrContains(a []string, e string) bool {
	for _, c := range a {
		if c == e {
			return true
		}
	}
	return false
}

// Returns true when array a is a superset of b; otherwise false.
func StrArrContainsAll(a []string, b []string) bool {
	if len(a) < len(b) {
		return false
	}
	for _, e := range b {
		if !StrArrContains(a, e) {
			return false
		}
	}
	return true
}