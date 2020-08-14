package main

import "fmt"

type rectangle struct {
	X      int
	Y      int
	width  int
	height int
}

func (r rectangle) left() int {
	return r.X
}

func (r rectangle) right() int {
	return r.X + r.width
}

func (r rectangle) top() int {
	return r.Y + r.height
}

func (r rectangle) bottom() int {
	return r.Y
}

func (r rectangle) print() {
	fmt.Println("X: ", r.X)
	fmt.Println("Y: ", r.Y)
	fmt.Println("width: ", r.width)
	fmt.Println("height: ", r.height)
}

func (r rectangle) withinX(item int) bool {
	if item > r.left() && item < r.right() {
		return true
	}

	return false
}

func (r rectangle) _withinX(item int) bool {
	if item >= r.left() && item <= r.right() {
		return true
	}

	return false
}

func (r rectangle) withinY(item int) bool {
	if item > r.bottom() && item < r.top() {
		return true
	}

	return false
}

func (r rectangle) _withinY(item int) bool {
	if item >= r.bottom() && item <= r.top() {
		return true
	}

	return false
}

func checkLR(guy rectangle, girl rectangle) rectangle {
	XrightOverlap := guy.withinX(girl.left()) && !guy.withinX(girl.right())
	XleftOverlap := !guy.withinX(girl.left()) && guy.withinX(girl.right())
	XcenterOverlap := guy._withinX(girl.left()) && guy._withinX(girl.right())
	XouterOverlap := girl._withinX(guy.left()) && girl._withinX(guy.right())

	width := -1
	X := -1

	if XcenterOverlap {
		width = girl.width
		X = girl.X
	} else if XouterOverlap {
		width = guy.width
		X = guy.X
	} else if XrightOverlap {
		width = guy.right() - girl.left()
		X = girl.left()
	} else if XleftOverlap {
		width = girl.right() - guy.left()
		X = guy.left()
	}

	YupperOverlap := guy.withinY(girl.bottom()) && !guy.withinY(girl.top())
	YlowerOverlap := !guy.withinY(girl.bottom()) && guy.withinY(girl.top())
	YcenterOverlap := guy._withinY(girl.top()) && guy._withinY(girl.bottom())
	YouterOverlap := girl._withinY(guy.top()) && girl._withinY(guy.bottom())

	height := -1
	Y := -1

	if YcenterOverlap {
		height = girl.height
		Y = girl.Y
	} else if YouterOverlap {
		height = guy.height
		Y = guy.Y
	} else if YupperOverlap {
		height = guy.top() - girl.bottom()
		Y = girl.bottom()
	} else if YlowerOverlap {
		height = girl.top() - guy.bottom()
		Y = guy.bottom()
	}

	return rectangle{X: X, Y: Y, width: width, height: height}
}

func main() {
	guy := rectangle{X: 1, Y: 1, width: 6, height: 3}
	girl := rectangle{X: 2, Y: 1, width: 3, height: 2}
	res := checkLR(guy, girl)
	res.print()
}
