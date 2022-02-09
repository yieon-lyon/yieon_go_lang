package main

// 정수
func Add(x, y int) int {
	return x + y
}

// 반복
// Repeat 함수
// character string -> 함수가 입력받을 타입과 변수명 지정
// func ... ) string -> 함수가 리턴할 타입 지정
func Repeat(character string) string {
	// string 타입의 repeated 변수 생성
	var repeated string
	// 0부터 증감하여 5보다 작은 수 까지의 반복
	for i := 0; i < 5; i++ {
		// 입력 받는 값인 character를 repeated에 추가
		/**
		repeated = "
		character = "ㅎㅇ"
		반복 i=0번째 -> repeated = "ㅎㅇ"
		반복 i=1번째 -> repeated = "ㅎㅇㅎㅇ"
		...
		반복 i=4번째 -> repeated = "ㅎㅇㅎㅇㅎㅇㅎㅇㅎㅇ"
		 */
		repeated += character
	}
	return repeated
}

// 배열과 슬라이스
// Sum 함수
func Sum(numbers []int) int {
	// sum 변수 할당 = 0
	sum := 0
	// 반복문 (numbers 의 배열 길이만큼(range) 반복한다.)
	for _, number := range numbers {
		// ex) numbers = [1,2,3]의 배열에서 각각의 값들을 sum의 값에 더한다.
		/**
		sum = 0
		sum + 1 -> sum = 1
		sum + 2 -> sum = 3
		sum + 3 -> sum = 6
		 */
		sum += number
	}
	// Sum 함수에서 내보낼 값을 설정
	return sum
}
// SumAll 함수
// ...[]int = [[]int, []int, ...]
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	// _ : index를 의미하나 사용하지 않아 _로 표시
	for _, numbers := range numbersToSum {
		// sums에 Sum(numbers) 값을 밀어넣음
		//	(ex1. numbers = {1, 3} -> Sum(numbers) == [4])
		//	(ex2. numbers2 = {5, 6} -> Sum(numbers) == [11])
		// ex1+ex2 -> [4, 11]
		sums = append(sums, Sum(numbers))
	}
	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
		// numbers에 값이 없을 경우 (numbers == {}) sums에 0을 추가함 -> sums == [0]
			sums = append(sums, 0)
		} else {
			// numbers의 index 1 이상부터의 배열을 tail 변수에 할당 (index는 0부터 시작)
			// ex. numbers = {1, 3, 5} -> numbers[1:] == {3, 5}
			tail := numbers[1:]
			// ex. Sum(tail) == 8 -> sums == [8]
			sums = append(sums, Sum(tail))
			// result. numbersToSum의 값이 [{}, {1, 3, 5}]이면 sums은 [0, 8]
		}
	}
	return sums
}