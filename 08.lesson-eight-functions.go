package main

func lessonEightFunctions() {
	res := add(1, 2, 3)
	println(res)

	numOfTerms, sum := numOfTermsAndSum(4, 53, 7, 7, 45)
	println(numOfTerms, sum)

	anonymousFunction()
}

func add(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}

	return result
}

func numOfTermsAndSum(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}

func anonymousFunction() {
	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)
		return
	}

	numTerms, sum := addFunc(1, 3)
	println(numTerms, sum)
}
