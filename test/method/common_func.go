package method

import "math"

// Area2 함수에 입력할 타입을 지정 (Circle 구현체를 입력받음)
// 반환할 값의 타입을 float64로 지정
func (c Circle) Area2() float64  {
	return math.Pi * c.Radius * c.Radius
}

// Area2 함수에 입력할 타입을 지정 (Rectangle 구현체를 입력받음)
// 반환할 값의 타입을 float64로 지정
func (r Rectangle) Area2() float64  {
	return r.Width * r.Height
}

// Area2 함수에 입력할 타입을 지정 (Triangle 구현체를 입력받음)
// 반환할 값의 타입을 float64로 지정
func (t Triangle) Area2() float64 {
	return (t.Base * t.Height) * 0.5
}