package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var m int = 5 //package level variable. Only visible to the package

//M is ...
var M int = 6 //package level variable. Globally visible

const i int = 9

func typeConversion() {
	var j int = 5
	var i float32
	fmt.Printf("%v, %T\n", i, i)
	i = float32(j)
	fmt.Printf("%v, %T\n", i, i)

	var str string
	str = strconv.Itoa(j)
	fmt.Printf("%v, %T\n", str, str)
}

func booleanType() {
	i := 2 == 1
	var k bool
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", k, k)
}

func realNumber() {
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))

	var m complex64 = complex(5, 12)
	fmt.Printf("%v, %T\n", m, m)

	s := "Hello World!"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}

func characters() {
	var character rune = 't'
	fmt.Printf("%v, %T\n", strconv.QuoteRune(character), strconv.QuoteRune(character))
}

func constants() {
	const i int = 2
	const j = 3
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", i+j, i+j)

	const (
		_ = iota
		a = 2 * iota
		b
		c
	)

	// const (
	// 	a1 = iota
	// 	b2
	// 	c2
	// )

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
}

func arrays() {
	grades := [3]int{1, 2, 3}
	fmt.Printf("Grades: %v\n", grades)

	student := [...]string{"jcbkj", "jcnj"}
	fmt.Printf("Students: %v\n", student)
	student[0] = "Hello"
	fmt.Printf("Students: %v\n", student)
	fmt.Printf("Students: %v\n", len(student))

	var matrix [3][3]int = [3][3]int{[3]int{1, 2, 3}, [3]int{1, 2, 3}}

	var matrix1 [3][3]int
	matrix1[0] = [3]int{1, 2, 3}
	matrix1[1] = [3]int{0, 9, 0}

	fmt.Println(matrix)
	fmt.Println(matrix1)

	b := student  //copying student array to b
	c := &student //pointing to sudent array
	c[0] = "Benura"
	b[0] = "Wenura"

	fmt.Printf("Students: %v\n", student)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

}

func slices() {
	//slices are reference type
	a := []int{1, 2, 3, 4, 5, 6}
	b := a
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("cap(a): %v\n", cap(a))
	fmt.Printf("len(a): %v\n", len(a))

	/*
	  Supports for both arrays and slices  -> a := [...]int{1,2,3,4,5,6}
	*/
	c := a[:]   //slice of all elements
	d := a[3:]  //slice from 4th element
	e := a[:6]  //slice first 6 elements
	f := a[3:6] //slice 4th 5th 6th elements

	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("f: %v\n", f)

	//Creating slices
	sl := make([]int, 3, 5)
	fmt.Printf("%v, %T\n", sl, sl)
	sl = append(sl, 1, 2, 3)
	fmt.Println(sl)
	fmt.Println("Cap:", cap(sl))

	//concatanating slices
	sl = append(sl, []int{10, 20, 30}...)
	fmt.Println("Contacatanated:", sl)
	fmt.Println("Cap:", cap(sl))

}

// Doctor is ...
type Doctor struct {
	name       string
	age        int
	companions []string
}

func structMap() {
	structMap := map[string]string{
		"name":     "Benura",
		"age":      "24",
		"location": "Colombo",
	}
	structMap["dob"] = "05/05/1997"
	delete(structMap, "age")
	fmt.Println(structMap)

	aDoctor := Doctor{
		name:       "Benura",
		age:        21,
		companions: []string{"leo", "geo"},
	}

	fmt.Println("Dcotor: ", aDoctor)

	aStudent := struct{ name string }{name: "Benura"}
	bStudent := &aStudent
	bStudent.name = "Bena"
	fmt.Println("Student a: ", aStudent)
	fmt.Println("Student b: ", bStudent)
}

// Animal is ...
type Animal struct {
	Name   string
	Origin string
}

// Bird is ...
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// Composistion relationship
func inheritance() {
	b := Bird{}
	b.Name = "Crow"
	b.Origin = "Sri Lanka"
	b.SpeedKPH = 45
	b.CanFly = true
	fmt.Println(b)
}

// AnimalTag is ...
type AnimalTag struct {
	Name   string `required max:"100"`
	Origin string
}

//tag
func tag() {
	t := reflect.TypeOf(AnimalTag{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

func conditionIf() {
	population := map[string]int{
		"Califonia": 10000,
		"Texas":     19990,
		"Ohio":      98000,
	}

	if pop, ok := population["Texas"]; ok {
		fmt.Println(ok, " ", pop)
	}

	switch 7 {
	case 1, 4, 5:
		fmt.Println(1)
	case 2, 7, 8:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("Else")
	}

	switch m := 2 + 5; m {
	case 1, 4, 5:
		fmt.Println(1)
	case 2, 7, 8:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("Else")
	}

	m := 3

	// Implicit break between cases. No need to define.
	switch {
	case m <= 2:
		fmt.Println("true")
	case m > 2:
		fmt.Println("false")
	}

	// To execute the next case (Condition does not matter)
	switch {
	case m <= 4:
		fmt.Println("less than 4")
		fallthrough
	case m >= 5:
		fmt.Println("less than 5")
	}

	var k interface{} = 1
	fmt.Println("K : ", k)
	switch k.(type) {
	case int:
		fmt.Printf("%T\n", k)
		break //breaking out early
		fmt.Println("K type:", reflect.TypeOf(k).Kind())
	case float32:
		fmt.Printf("%T", k)
	default:
		fmt.Println("Wrong")
	}

}

func loops() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
	// 	fmt.Println(i, j)
	// }

	// i := 0
	// for ; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Loop:
	// 	for i := 0; i < 5; i++ {
	// 		for j := 0; j < 5; j++ {
	// 			fmt.Println(i)
	// 			if i == 2 {
	// 				break Loop
	// 			}
	// 		}
	// 	}

	// 	s := []int{1, 2, 3}
	// 	for k, v := range s {
	// 		fmt.Println(k, "=>", v)
	// 	}

	// str := "Hello Go!"
	// for k, v := range str {
	// 	fmt.Println(k, "=>", string(v))
	// }

	mapTest := map[string]int{
		"Cal":     1000,
		"Texa":    2000,
		"Virg":    3000,
		"Florida": 9000,
	}

	for _, v := range mapTest {
		fmt.Println(v)
	}

}

func deferFunc() {

	// fmt.Println("Start")        //1
	// defer fmt.Println("Middle") //after the main function executed
	// fmt.Println("End1")         //2
	// fmt.Println("End2")         //3

	// executes according to LIFO (Last In First Out)
	defer fmt.Println("Start")  //4
	defer fmt.Println("Middle") //3
	defer fmt.Println("End1")   //2
	defer fmt.Println("End2")   //1

	// res, err := http.Get("https://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(robots)

	a := "start"         //  initial value
	defer fmt.Println(a) //assign "start" to a
	a = "end"            // not applicable to the defer
}

func exceptionPanic() {
	// a, b := 1, 0
	// fmt.Println(a / b)
	// panic("Division by 0 not allowed")
	// fmt.Println("End")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// panic will execute after defer
	// fmt.Println("start")            //1
	// defer fmt.Println("defer func") //2
	// panic("Wrong")                  //3
	// fmt.Println("end")

	fmt.Println("================================")
	fmt.Println("start")
	fmt.Println("about to panic")
	defer func() {
		//To hanle the panic
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	panic("Something wrong panic")
	fmt.Println("Done")

}

type myStruct struct {
	name string
	age  int
}

func pointers() {
	var a int = 2
	var b *int = &a
	fmt.Println(a, *b)
	*b = 13
	fmt.Println(a, *b)
	a = 15
	fmt.Println(a, *b)

	var ms *myStruct
	ms = new(myStruct)
	(*ms).name = "Benura" // == ms.name = "Benura"
	fmt.Println(ms)
	fmt.Println((*ms))
}

func varArgs(values ...int) {
	result := 0
	for _, v := range values {
		result += v
	}

	fmt.Println(result)
}

func returFunc(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func returnRef(values ...int) *int {
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func returnRef1(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot Divide by Zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// Writer is ...
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter is ...
type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func goRoutine() {
	fmt.Println("Hello Go!")
}

var mutex = sync.RWMutex{}
var counter int = 0
var wg = sync.WaitGroup{}

func main() {
	var (
		m = 4
		o = 5
	)
	k := 3.7
	fmt.Printf("%v, %T\n", k, k)
	fmt.Println(m, o)
	typeConversion()
	booleanType()
	realNumber()
	characters()
	constants()
	arrays()
	slices()
	structMap()
	inheritance()
	tag()
	conditionIf()
	loops()
	deferFunc()
	exceptionPanic()
	pointers()
	varArgs(1, 2, 3, 4)
	fmt.Println("Sum is: ", returFunc(1, 2, 3, 4))
	s := returnRef(1, 2, 3, 4)
	fmt.Println("Ref sum: ", *s)
	fmt.Println("Ref sum: ", *s)
	b, err := divide(2, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	/**
	  annonymous functions
	*/
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	f := func(i int) {
		fmt.Println("Hello Go ", i)
	}

	var divide func(float64, float64) (float64, error) = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot Divide by Zero")
		}
		return a / b, nil
	}
	f(2)
	d, err := divide(2, 0)
	if err != nil {
		fmt.Println(err)
		fmt.Println(d)
	}

	/**
	  Methods
	*/
	g := greeter{
		greeting: "Hello",
		name:     "Benura",
	}
	g.greet()

	// interfaces
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go"))

	// goRoutine
	go goRoutine()
	time.Sleep(100 * time.Millisecond)

	var msg string = "Hello Go!"
	wg.Add(2)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	go func() {
		fmt.Println("Hello World")
		wg.Done()
	}()
	msg = "Good Bye"
	wg.Wait()

	// Mutex and goRuotine
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock()
		go printHello()
		mutex.Lock()
		go writeCount()
	}
	wg.Wait()

	/**
	  Channel
	*/
	ch := make(chan int)
	wg.Add(2)
	//read only channel
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	// write only channel
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

func printHello() {
	fmt.Printf("Hello # %v\n", counter)
	mutex.RUnlock()
	wg.Done()
}

func writeCount() {
	counter++
	mutex.Unlock()
	wg.Done()
}
