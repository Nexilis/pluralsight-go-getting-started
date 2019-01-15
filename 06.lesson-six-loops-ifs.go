package main

func lessonSixLoopsIfs() {
	if foo := 2; foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	switch bear := 1; bear {
	case 1:
		println("one")
	case 2:
		println("two")
	default:
		println("three or more")
	}

	for i := 0; i < 5; i++ {
		println(i)
	}

	j := 0
	for {
		j++
		print(j)

		if j > 5 {
			break
		}
	}

	s := []string{"foo", "bar", "buz"}

	for idx, v := range s {
		println(idx, v)
	}

	myMap := make(map[string]string)
	myMap["first"] = "foo"
	myMap["second"] = "bar"
	myMap["third"] = "buz"

	for k, v := range myMap {
		println(k, v)
	}
}
