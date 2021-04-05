package main

func main() {
	//simple if
	a := true
	if a {
		print("hello world")
	}

	//if int
	b := 1
	if b == 1 {
		print("if b not work")
	}

	//if with map
	mm := map[string]string{
		"firstName": "Raimo",
		"lastName":  "Karhu",
	}
	//in this case firstName and ok will exist only within if
	if firstName, ok := mm["firstName"]; ok {
		println("firstName key exists, = ", firstName)
	} else {
		println("no such key found")
	}

	//in this case lastName exist within this func
	lastName, lastNameExists := mm["lastName"]
	if lastNameExists {
		println("lastName is ", lastName)
	} else {
		println("no lastName found")
	}
	println(len(lastName))

	//if else if
	if !lastNameExists {
		println("No lastName")
	} else if lastName == "Karhu" {
		println("lastName is Karhu")
	} else {
		println("lastName is not Karhu")
	}

	//for
	for {
		println("infinite loop aka while(true)")
		break
	}

	//slice ops
	s1 := []int{3, 4, 5, 6, 7, 8}
	value := 0
	idx := 0

	//while-style
	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value = s1[idx]
		idx++
		println("while-style loop, idx:", idx, "value:", value)
	}

	//normal loop
	for i := 0; i < len(s1); i++ {
		println("c/java-style loop", i, s1[i])
	}

	//range aka for-each
	for idx := range s1 {
		println("range slice by index", idx, s1[idx])
	}

	for idx, val := range s1 {
		println("range slice by idx-value", idx, val)
	}

	//map ops
	for key := range mm {
		println("range map by key", key)
	}

	for key, val := range mm {
		println("range map by key-val", key, val)
	}

	for _, val := range mm {
		println("range map by val", val)
	}

	//switch
	mm["firstName"] = "Raimo"
	mm["flag"] = "OK"
	switch mm["firstName"] {
	case "Raimo", "Raimoke":
		println("switch - name is Raimo")
		//no break needed
	case "Alex":
		if mm["flag"] == "OK" {
			break //exiting switch
		}
		println("switch - name is Alex")
		fallthrough //go to next variant
	default:
		println("switch - some other name")
	}

	//no macarons
	switch {
	case mm["firstName"] == "Raimo":
		println("switch2 - Raimo")
	case mm["lastName"] == "Karhu":
		println("switch - Karhu")
	}

	//exiting loop from inner switch
MyLoop:
	for key, val := range mm {
		println("switch in loop", key, val)
		switch {
		case key == "firstName" && val == "Raimo":
			println("switch - break loop here")
			break MyLoop
		}
	}
}
