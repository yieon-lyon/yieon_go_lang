package method

import (
	"testing"
)

// 같은 폴더 안에 TestArea 함수가 존재하여 TestArea2 함수로 지정
func TestArea2(t *testing.T) {

	// 배열 구현체 정의
	areaTests := []struct {
		// 구현체에 들어갈 파라미터와 타입을 지정
		// Shape 인터페이스 타입을 사용하는 shape 변수 지정
		shape Shape
		// float64 타입을 사용하는 want 변수 지정
		want  float64
	}{
		// Shape: XXX, want: XXX
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	// areaTests 반복문
	for _, tt := range areaTests {
		// index 0 = {Rectangle{12, 6}, 72.0}, index 1 = {Circle{10}, 314.1592653589793}
		// index 0 tt.shape = Rectangle{12, 6}, index 1 tt.shape = Circle{10}
		got := tt.shape.Area2()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}