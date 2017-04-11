package main

func stackMap() bool {
	stackMap := make(map[string]string)
	for _, v := range []string{"just", "for", "fun"} {
		stackMap[v] = v
	}
	if len(stackMap) > 2 {
		return true
	}
	return false
}

func heapMap() map[string]string {
	heapMap := make(map[string]string)
	for _, v := range []string{"just", "for", "fun"} {
		heapMap[v] = v
	}
	return heapMap
}

func main() {

}
