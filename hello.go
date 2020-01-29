package main

func ifTest () {
	var i = 30
	var max = 100
	if val := i * 2; val < max {
		println("true")
	} else {
		println("false")
	}
}

func forTest () {
	names := []string{"a", "b", "c"}

	for index, name := range names {
		println(index, name)
	}
}

func funTest (msg string) {
	msg += "1"
	println(msg)
}

func funTest2 (msg *string) {
	*msg = "change message"
	println(*msg)
}

func goRoutineTest () {
	ch := make(chan int)

	go func () {
		ch <- 123
	}()

	var i int
	i = <- ch
	println(i)
}

func main () {
	goRoutineTest()
	//forTest()
	//ifTest()
	var a int
	a = 10
	println(a)
	b := `aaaaa
bbbbb
ccccc
ddddd`
	println("hello wolrd")
	println(b)

	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)

	println(u, f)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

	var k2 int = 10
	var p = &k2
	println(p, *p)
}
