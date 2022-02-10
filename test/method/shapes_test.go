package method

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	// got 변수를 Perimeter 함수의 결과값으로 생성
	got := Perimeter(rectangle)
	want := 40.0

	// 변수 got 과 변수 want의 값이 다르면
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	// got 변수를 Area 함수의 결과값으로 생성
	got := Area(rectangle)
	want := 72.0

	// 변수 got 과 변수 want의 값이 다르면
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Perimeter 함수
// Rectangle 타입의 변수(rectangle)를 받으며 리턴하는 값의 타입은 float64
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area 함수
// Rectangle 타입의 변수(rectangle)를 받으며 리턴하는 값의 타입은 float64
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}