package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

var globalMessage string = "Ini pesan global"
var isGlobal bool

type Person struct {
	FirstName string
	LastName  string
	Age       int
	isMarried bool
}

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Contact struct {
	Email       string
	Phone       string
	HomeAddress Address
}

const Pi float64 = 3.14159265359
const AppName = "Aplikasi Keren"
const typedInt int = 100
const untypedInt = 200

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

var packageVar = "Saya di level paket"

type Vertex struct {
	X, Y int
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    float64
	IsActive  bool
}

type Manager struct {
	Person
	Contact
	Department string
	Level      int
}

type Point struct {
	X, Y float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type ConfigError struct {
	FileName string
	Field    string
	Err      error
}

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

type Config struct {
	Timeout *int
	Retries int
}

func helloWorldExample() {
	fmt.Println("Halo, Dunia Go!")
}

func variableExamples() {
	var i int
	var f float64 = 3.14
	var s string
	var b bool = true

	fmt.Println(i, f, s, b)

	var x, y int = 10, 20
	var (
		namaDepan    string = "Budi"
		namaBelakang string = "Santoso"
		umur         int    = 30
	)
	fmt.Println(x, y)
	fmt.Println(namaDepan, namaBelakang, umur)
	fmt.Println(globalMessage)

	message := "Halo dari deklarasi singkat!"
	count := 100
	isValid := true
	pi := 3.14159

	fmt.Println(message, count, isValid, pi)

	host, port := "localhost", 8080
	fmt.Println(host, port)

	count = 101
	fmt.Println(count)

	host, err := "127.0.0.1", error(nil)
	fmt.Println(host, err)

	funcVar := "Saya di level fungsi"
	fmt.Println(packageVar)
	fmt.Println(funcVar)

	if true {
		blockVar := "Saya di level blok if"
		fmt.Println(blockVar)
		fmt.Println(funcVar)

		funcVar := "Saya funcVar di dalam if"
		fmt.Println(funcVar)
	}
	fmt.Println(funcVar)

}

func constantsExample() {
	const localConst = "Ini konstanta lokal"

	fmt.Println(Pi)
	fmt.Println(AppName)
	fmt.Println(localConst)

	fmt.Println("Hari:", Monday, Saturday)

	fmt.Printf("1 GB = %d Bytes\n", GB)

	var myFloat float32 = untypedInt
	fmt.Println(myFloat)
}

func basicTypesExample() {
	var isActive bool = true
	var isEnabled = false
	fmt.Println(isActive, isEnabled)

	var umur int = 25
	var suhu float64 = 26.5
	var nilaiComplex complex128 = 2 + 3i
	var dataByte byte = 'A'
	fmt.Println(umur, suhu, nilaiComplex, dataByte)

	var a int = 10
	var b float64 = 3.5
	var c = float64(a) + b
	fmt.Println(c)

	var greeting string = "Hello"
	var name = "Go Developer"
	var message = greeting + ", " + name + "!"
	fmt.Println(message)

	var multiLine = `Ini adalah string
yang bisa terdiri dari
beberapa baris.
Interpolasi tidak bekerja di sini: ${var}`
	fmt.Println(multiLine)

	fmt.Println(greeting[0])
	fmt.Println(len(greeting))

	for index, runeValue := range "Go语言" {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", index, runeValue, runeValue)
	}
}

func arrayExample() {
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	names := [...]string{"Alice", "Bob", "Charlie"}
	fmt.Println(len(names))
	fmt.Println(names)

	var matrix [2][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	fmt.Println(matrix)
}

func sliceExample() {
	var fruits []string
	fmt.Println(fruits == nil)

	numbers := make([]int, 3, 5)
	fmt.Println(numbers)
	fmt.Println("Len:", len(numbers), "Cap:", cap(numbers))

	colors := []string{"Merah", "Hijau", "Biru"}
	fmt.Println(colors)
	fmt.Println("Len:", len(colors), "Cap:", cap(colors))

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println("Len:", len(s), "Cap:", cap(s))

	s2 := colors[0:2]
	fmt.Println(s2)

	s3 := colors[:2]
	s4 := colors[1:]
	s5 := colors[:]
	fmt.Println(s3, s4, s5)

	numbers = append(numbers, 10)
	fmt.Println(numbers)
	fmt.Println("Len:", len(numbers), "Cap:", cap(numbers))

	numbers = append(numbers, 20, 30)
	fmt.Println(numbers)
	fmt.Println("Len:", len(numbers), "Cap:", cap(numbers))

	moreColors := []string{"Kuning", "Ungu"}
	allColors := append(colors, moreColors...)
	fmt.Println(allColors)

	for index, value := range allColors {
		fmt.Printf("Index: %d, Warna: %s\n", index, value)
	}

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[1][1] = "O"
	fmt.Println(board)
}

func mapExample() {
	var scores map[string]int

	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Println(ages)
	fmt.Println(ages["Alice"])
    fmt.Println(scores) // Demonstrating nil map access is safe for read/print

	populations := map[string]int{
		"Jakarta":  10_000_000,
		"Surabaya": 3_000_000,
		"Bandung":  2_500_000,
	}
	fmt.Println(populations)
	fmt.Println(populations["Jakarta"])
	fmt.Println(populations["Medan"])

	pop, ok := populations["Medan"]
	if ok {
		fmt.Println("Populasi Medan:", pop)
	} else {
		fmt.Println("Data Medan tidak ditemukan.")
	}

	val, exists := ages["Charlie"]
	fmt.Println("Charlie ada?", exists, "Umur:", val)

	ages["Alice"] = 31
	fmt.Println(ages)

	delete(ages, "Bob")
	fmt.Println(ages)

	fmt.Println(len(populations))

	for key, value := range populations {
		fmt.Printf("Kota: %s, Populasi: %d\n", key, value)
	}
}

func structExample() {
	var p1 Person
	p1.FirstName = "Andi"
	p1.LastName = "Wijaya"
	p1.Age = 28
	p1.isMarried = false
	fmt.Println(p1)
	fmt.Println("Nama Depan:", p1.FirstName)

	p2 := Person{
		FirstName: "Siti",
		LastName:  "Aminah",
		Age:       32,
		isMarried: true,
	}
	fmt.Println(p2)

	p3 := Person{"Rudi", "Hartono", 40, true}
	fmt.Println(p3)

	p4 := Person{FirstName: "Dewi"}
	fmt.Println(p4)

	c1 := Contact{
		Email: "andi.w@example.com",
		Phone: "0812345678",
		HomeAddress: Address{
			Street:  "Jl. Merdeka No. 10",
			City:    "Jakarta",
			ZipCode: "10110",
		},
	}
	fmt.Println(c1)
	fmt.Println("Kota:", c1.HomeAddress.City)

	p5 := &Person{"Bambang", "Pamungkas", 42, true}
	fmt.Println(p5.FirstName)

	point := struct {
		X int
		Y int
	}{
		X: 10,
		Y: 20,
	}
	fmt.Println(point)
}

func ifElseExample() {
	score := 75
	grade := ""

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "E"
	}
	fmt.Println("Nilai:", score, "Grade:", grade)

	if num := 10; num%2 == 0 {
		fmt.Printf("%d adalah genap\n", num)
	} else {
		fmt.Printf("%d adalah ganjil\n", num)
	}

	x := 5
	if x > 0 {
		fmt.Println("x positif")
	}
}

func switchExample() {
	day := "Senin"
	activity := ""

	switch day {
	case "Senin":
		activity = "Meeting awal minggu"
	case "Selasa", "Rabu", "Kamis":
		activity = "Kerja rutin"
	case "Jumat":
		activity = "Review mingguan & persiapan weekend"
	case "Sabtu", "Minggu":
		activity = "Libur!"
	default:
		activity = "Hari tidak valid"
	}
	fmt.Printf("Hari %s: %s\n", day, activity)

	score := 85
	grade := ""
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	default:
		grade = "C atau kurang"
	}
	fmt.Println("Grade:", grade)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	num := 1
	switch num {
	case 1:
		fmt.Println("Satu")
		fallthrough
	case 2:
		fmt.Println("Dua (atau fallthrough dari 1)")
	case 3:
		fmt.Println("Tiga")
	}

	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	default:
		fmt.Printf("Tipe tidak diketahui: %T\n", v)
	}
}

func forLoopExample() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum (C-style):", sum)

	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("n (while-style):", n)

	items := []string{"apel", "pisang", "ceri"}
	fmt.Println("Iterasi Slice:")
	for index, value := range items {
		fmt.Printf("  Index: %d, Value: %s\n", index, value)
	}

	capitals := map[string]string{"Indonesia": "Jakarta", "Jepang": "Tokyo"}
	fmt.Println("\nIterasi Map:")
	for country, capital := range capitals {
		fmt.Printf("  Ibukota %s adalah %s\n", country, capital)
	}

	fmt.Println("\nIterasi String:")
	for i, r := range "Go€" {
		fmt.Printf("  Index byte: %d, Rune: %c\n", i, r)
	}
}

func breakContinueExample() {
	fmt.Println("Contoh break:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("  Berhenti di i=5")
			break
		}
		fmt.Printf("  i = %d\n", i)
	}

	fmt.Println("\nContoh continue (lewati angka genap):")
	for j := 0; j < 6; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Printf("  j = %d (ganjil)\n", j)
	}

	fmt.Println("\nContoh break dengan label:")
OuterLoop:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Printf("  x=%d, y=%d\n", x, y)
			if x == 1 && y == 1 {
				fmt.Println("    Keluar dari OuterLoop")
				break OuterLoop
			}
		}
	}
}

func simpleDeferExample() {
	fmt.Println("--- Contoh Urutan Defer Sederhana ---")
	fmt.Println("Satu")
	defer fmt.Println("Empat (defer pertama, eksekusi terakhir)")
	fmt.Println("Dua")
	defer fmt.Println("Tiga (defer kedua, eksekusi sebelum 'Empat')")
	fmt.Println("Tiga setengah")
}

func exampleDeferArgs() {
	i := 0
	defer fmt.Println("Nilai i saat defer dievaluasi:", i)
	i++
	fmt.Println("Nilai i sebelum return:", i)
}

func sayHello() {
	fmt.Println("Halo!")
}

func greet(name string) string {
	message := "Halo, " + name + "!"
	return message
}

func add(a int, b int) int {
	return a + b
}

func multiply(x, y int, factor float64) float64 {
	return float64(x*y) * factor
}

func divide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("tidak bisa dibagi dengan nol")
	}
	return numerator / denominator, nil
}

func checkLength(s string) (string, bool) {
	if len(s) > 10 {
		return "Terlalu panjang", false
	}
	return "Panjang OK", true
}

func subtractNamedReturn(a, b int) (result int, success bool) {
	if a < b {
		success = false
		result = 0
		return result, success
	}

	result = a - b
	success = true
	return result, success
}

func sumNumbers(label string, numbers ...int) int {
	fmt.Printf("Menerima untuk '%s': %v (tipe: %T)\n", label, numbers, numbers)
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func calculate(x, y int, operation func(int, int) int) int {
	fmt.Printf("Menjalankan operasi pada %d dan %d\n", x, y)
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func functionExamples() {
	sayHello()

	greeting := greet("Pengguna Go")
	fmt.Println(greeting)

	sum := add(5, 3)
	fmt.Println("5 + 3 =", sum)

	resultFloat := multiply(4, 5, 1.5)
	fmt.Println("4 * 5 * 1.5 =", resultFloat)

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("5 / 0 =", result)
	}

	statusMsg, _ := checkLength(" pendek")
	fmt.Println("Status:", statusMsg)

	_, isValid := checkLength("string yang sangat panjang sekali")
	if !isValid {
		fmt.Println("String tidak valid!")
	}

	numStr := "123"
	numInt, convErr := strconv.Atoi(numStr)
	if convErr != nil {
		fmt.Println("Konversi gagal:", convErr)
	} else {
		fmt.Println("Hasil konversi:", numInt)
	}

	numStr = "abc"
	numInt, convErr = strconv.Atoi(numStr)
	if convErr != nil {
		fmt.Println("Konversi gagal:", convErr)
	} else {
		fmt.Println("Hasil konversi:", numInt)
	}

	res, ok := subtractNamedReturn(10, 3)
	fmt.Printf("Hasil: %d, Sukses: %t\n", res, ok)

	res, ok = subtractNamedReturn(5, 8)
	fmt.Printf("Hasil: %d, Sukses: %t\n", res, ok)

	total1 := sumNumbers("Set 1")
	fmt.Println("Total 1:", total1)
	total2 := sumNumbers("Set 2", 1, 2, 3)
	fmt.Println("Total 2:", total2)
	total3 := sumNumbers("Set 3", 10, 20, 30, 40, 50)
	fmt.Println("Total 3:", total3)
	nums := []int{5, 10, 15}
	total4 := sumNumbers("Set 4", nums...)
	fmt.Println("Total 4:", total4)

	var op func(int, int) int
	op = add
	resultOp := op(10, 5)
	fmt.Println("Hasil op (add):", resultOp)

	op = subtract // Dummy subtract needed
	resultOp = op(10, 5)
	fmt.Println("Hasil op (subtract):", resultOp)

	sumResult := calculate(20, 7, add)
	fmt.Println("Hasil calculate (add):", sumResult)
	diffResult := calculate(20, 7, subtract)
	fmt.Println("Hasil calculate (subtract):", diffResult)
	multResult := calculate(5, 6, func(a, b int) int { return a * b })
	fmt.Println("Hasil calculate (multiply anonim):", multResult)

	double := createMultiplier(2)
	triple := createMultiplier(3)
	fmt.Println("Double 5:", double(5))
	fmt.Println("Triple 5:", triple(5))

	func(message string) {
		fmt.Println("Pesan anonim:", message)
	}("Halo langsung!")

	c1 := counter()
	fmt.Println("Counter 1:", c1())
	fmt.Println("Counter 1:", c1())
	fmt.Println("Counter 1:", c1())
	c2 := counter()
	fmt.Println("Counter 2:", c2())
	fmt.Println("Counter 1:", c1())
}

func pointerBasics() {
	x := 100
	fmt.Printf("Nilai x: %d, Alamat x: %p\n", x, &x)

	var p *int
	fmt.Printf("Nilai p (sebelum assignment): %v\n", p)

	p = &x
	fmt.Printf("Nilai p (alamat x): %p\n", p)

	fmt.Printf("Nilai yang ditunjuk p (*p): %d\n", *p)

	*p = 200
	fmt.Printf("Nilai x setelah diubah via p: %d\n", x)

	fmt.Printf("Alamat dari pointer p: %p\n", &p)

	var pp **int
	pp = &p
	fmt.Printf("Nilai pp (alamat p): %p\n", pp)
	fmt.Printf("Nilai yang ditunjuk p (*p) via pp (**pp): %d\n", **pp)

	ptrStr := new(string)
	fmt.Printf("Nilai ptrStr: %p, Nilai *ptrStr: '%s'\n", ptrStr, *ptrStr)
	*ptrStr = "Halo dari new()"
	fmt.Printf("Nilai *ptrStr setelah diubah: '%s'\n", *ptrStr)

	var pNil *int
	if pNil != nil {
		fmt.Println(*pNil)
	} else {
		fmt.Println("pNil adalah nil")
	}
}

func incrementValue(val int) {
	val++
	fmt.Printf("  Nilai di dalam incrementValue: %d\n", val)
}

func incrementPointer(ptr *int) {
	*ptr++
	fmt.Printf("  Nilai di dalam incrementPointer (*ptr): %d\n", *ptr)
}

func pointerArgsExample() {
	num := 10
	fmt.Printf("Nilai num sebelum incrementValue: %d\n", num)
	incrementValue(num)
	fmt.Printf("Nilai num setelah incrementValue: %d\n", num)

	fmt.Printf("\nNilai num sebelum incrementPointer: %d\n", num)
	incrementPointer(&num)
	fmt.Printf("Nilai num setelah incrementPointer: %d\n", num)
}

func pointerForOptionalConfig() {
	defaultTimeout := 30
	cfg1 := Config{Timeout: &defaultTimeout, Retries: 3}
	cfg2 := Config{Retries: 5}

	if cfg1.Timeout != nil {
		fmt.Printf("Cfg1 Timeout: %d\n", *cfg1.Timeout)
	}
	if cfg2.Timeout != nil {
		fmt.Printf("Cfg2 Timeout: %d\n", *cfg2.Timeout)
	} else {
		fmt.Println("Cfg2 Timeout: default")
	}
}

func pointerToStructExample() {
	v := Vertex{1, 2}
	p := &v

	fmt.Println((*p).X)
	fmt.Println(p.X)

	p.X = 100
	fmt.Println(v.X)
	fmt.Println(v)

	p1 := new(Vertex)
	p2 := &Vertex{}
	p3 := &Vertex{X: 1}
	p4 := &Vertex{1, 2}
	fmt.Println(p1, p2, p3, p4)
}

func (p Person) Greet() {
	fmt.Printf("Halo, nama saya %s dan umur saya %d tahun.\n", p.FirstName, p.Age)
}

func (m Manager) DelegateTask() {
	fmt.Printf("%s (%s) mendelegasikan tugas.\n", m.FirstName, m.Department)
}

func (p Point) DistanceFromOrigin() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) Scale(factor float64) {
	p.X = p.X * factor
	p.Y = p.Y * factor
	fmt.Printf("  Di dalam Scale: Point menjadi %v\n", *p)
}

func structMethodEmbeddingExample() {
	m := Manager{
		Person: Person{
			FirstName: "Budi Gunawan",
			Age:  45,
		},
		Contact: Contact{
			Email: "budi.g@company.com",
			Phone: "0811-111-222",
		},
		Department: "Teknologi",
		Level:      5,
	}

	fmt.Println("Nama:", m.FirstName)
	fmt.Println("Email:", m.Email)
	fmt.Println("Departemen:", m.Department)
	fmt.Println("Umur (eksplisit):", m.Person.Age)

	m.Greet()
	m.DelegateTask()

	pt1 := Point{3, 4}
	fmt.Printf("Point pt1: %v\n", pt1)
	dist := pt1.DistanceFromOrigin()
	fmt.Printf("Jarak pt1 dari origin: %.2f\n", dist)
	fmt.Printf("Point pt1 setelah DistanceFromOrigin: %v\n", pt1)

	fmt.Println("---")
	pt1.Scale(2)
	fmt.Printf("Point pt1 setelah Scale(2): %v\n", pt1)

	fmt.Println("---")
	pt2 := &Point{1, 1}
	fmt.Printf("Point pt2 (pointer): %v\n", pt2)
	pt2.Scale(5)
	fmt.Printf("Point pt2 setelah Scale(5): %v\n", pt2)

	dist2 := pt2.DistanceFromOrigin()
	fmt.Printf("Jarak pt2 dari origin: %.2f\n", dist2)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Tipe: %T\n", s)
	fmt.Printf("  Area: %.2f\n", s.Area())
	fmt.Printf("  Perimeter: %.2f\n", s.Perimeter())
}

func interfaceExample() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	fmt.Println("Info Persegi Panjang:")
	PrintShapeInfo(rect)

	fmt.Println("\nInfo Lingkaran:")
	PrintShapeInfo(circ)

	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
		rect,
		circ,
	}

	fmt.Println("\nInfo dari Slice Shapes:")
	totalArea := 0.0
	for _, s := range shapes {
		PrintShapeInfo(s)
		totalArea += s.Area()
	}
	fmt.Printf("\nTotal Area semua bentuk: %.2f\n", totalArea)
}

func describe(i interface{}) {
	fmt.Printf("Nilai: %v, Tipe: %T\n", i, i)
}

func emptyInterfaceExample() {
	var data interface{}

	data = 10
	describe(data)

	data = "Halo Go"
	describe(data)

	data = true
	describe(data)

	data = struct{ Name string }{"Sample"}
	describe(data)

	mixed := []interface{}{1, "dua", false, 3.14, nil}
	for _, v := range mixed {
		describe(v)
	}
}

func process(i interface{}) {
	fmt.Printf("Memproses: %v (%T)\n", i, i)

	strVal, ok := i.(string)
	if ok {
		fmt.Printf("  Ini adalah string! Panjangnya: %d\n", len(strVal))
		return
	}

	intVal, ok := i.(int)
	if ok {
		fmt.Printf("  Ini adalah integer! Nilai kuadrat: %d\n", intVal*intVal)
		return
	}

	fmt.Println("  Tipe tidak dikenali atau tidak ditangani.")
}

func typeAssertionExample() {
	process("Halo Dunia")
	process(42)
	process(true)
}

func describeWithTypeSwitch(i interface{}) {
	fmt.Printf("Mendeskripsikan: %v (%T) -> ", i, i)
	switch v := i.(type) {
	case string:
		fmt.Printf("String dengan panjang %d\n", len(v))
	case int:
		fmt.Printf("Integer, nilainya %d\n", v)
	case bool:
		fmt.Printf("Boolean, nilainya %t\n", v)
	case float64:
		fmt.Printf("Float64, nilainya %f\n", v)
	case nil:
		fmt.Println("Nilai nil")
	default:
		fmt.Printf("Tipe lain: %T\n", v)
	}
}

func typeSwitchExample() {
	describeWithTypeSwitch("Go")
	describeWithTypeSwitch(123)
	describeWithTypeSwitch(false)
	describeWithTypeSwitch(3.14)
	describeWithTypeSwitch(nil)
	describeWithTypeSwitch([]int{1, 2})
}

func errorHandlingConventionExample() {
	filePath := "non_existent_file.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Gagal membuka file '%s': %v\n", filePath, err)
	} else {
		fmt.Printf("Berhasil membuka file: %s\n", file.Name())
		defer file.Close()
	}

	fmt.Println("---")

	numStr := "123a"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Gagal mengkonversi '%s' ke int: %v\n", numStr, err)
	} else {
		fmt.Printf("Hasil konversi: %d\n", num)
	}

	numStr = "456"
	num, err = strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Gagal mengkonversi '%s' ke int: %v\n", numStr, err)
	} else {
		fmt.Printf("Hasil konversi: %d\n", num)
	}
}

func validateInput(input string) error {
	if input == "" {
		return errors.New("input tidak boleh kosong")
	}
	return nil
}

func connectToDB(host string, port int) error {
	if port <= 0 || port > 65535 {
		return fmt.Errorf("port database tidak valid: %d", port)
	}
	fmt.Printf("Mencoba koneksi ke %s:%d...\n", host, port)
	return fmt.Errorf("gagal terkoneksi ke %s:%d (host tidak ditemukan)", host, port)
}

func errorCreationExample() {
	err1 := validateInput("")
	if err1 != nil {
		fmt.Println("Error validasi:", err1)
	}
	err2 := connectToDB("localhost", 80000)
	if err2 != nil {
		fmt.Println("Error koneksi 1:", err2)
	}
	err3 := connectToDB("db.example.com", 5432)
	if err3 != nil {
		fmt.Println("Error koneksi 2:", err3)
	}
}

func (e *ConfigError) Error() string {
	msg := fmt.Sprintf("kesalahan konfigurasi di file '%s'", e.FileName)
	if e.Field != "" {
		msg += fmt.Sprintf(", field '%s'", e.Field)
	}
	return msg
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}

func loadConfigSimplified(filename string, failParsing bool) error {
	if filename == "non_existent_config.yaml" {
		originalErr := os.ErrNotExist
		return fmt.Errorf("gagal memuat konfigurasi: %w", originalErr)
	}

	if failParsing {
		missingField := "database_url"
		parseErr := &ConfigError{FileName: filename, Field: missingField}
		return fmt.Errorf("gagal memuat konfigurasi: %w", parseErr)
	}

	return nil
}

func errorWrappingExample() {
	err := loadConfigSimplified("non_existent_config.yaml", false)
	if err != nil {
		fmt.Println("Error utama:", err)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("  Detail: File konfigurasi tidak ditemukan.")
		}
		var configErr *ConfigError
		if errors.As(err, &configErr) {
			fmt.Printf("  Detail: Kesalahan pada field '%s' di file '%s'\n", configErr.Field, configErr.FileName)
			wrapped := configErr.Unwrap()
			if wrapped != nil {
				fmt.Println("    Error yang dibungkus:", wrapped)
			}
		}
	}

	err = loadConfigSimplified("config.yaml", true)
	if err != nil {
		fmt.Println("\nError utama (parsing):", err)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("  Detail: File konfigurasi tidak ditemukan.")
		}
		var configErr *ConfigError
		if errors.As(err, &configErr) {
			fmt.Printf("  Detail: Kesalahan pada field '%s' di file '%s'\n", configErr.Field, configErr.FileName)
			wrapped := errors.Unwrap(err)
			fmt.Println("    Error yang dibungkus (via errors.Unwrap):", wrapped)
			wrapped = configErr.Unwrap()
			fmt.Println("    Error yang dibungkus (via method Unwrap):", wrapped)
		}
	}
}

func mightPanic(shouldPanic bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC TERDETEKSI (di recover):", r)
		}
	}()

	fmt.Println("Sebelum potensi panic...")
	if shouldPanic {
		panic("Sesuatu yang sangat buruk terjadi!")
	}
	fmt.Println("Setelah potensi panic (tidak akan tercapai jika panic)")
}

func panicRecoverExample() {
	fmt.Println("--- Memanggil mightPanic(false) ---")
	mightPanic(false)
	fmt.Println("mightPanic(false) selesai.")

	fmt.Println("\n--- Memanggil mightPanic(true) ---")
	mightPanic(true)
	fmt.Println("mightPanic(true) selesai (setelah recover).")
}

func say(s string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(10 * time.Millisecond) // Shorter sleep for faster example
		fmt.Printf("Pesan dari '%s': %s - iterasi %d\n", s, s, i)
	}
	fmt.Printf("'%s' selesai.\n", s)
}

func goroutineSimpleExample() {
	fmt.Println("Memulai main goroutine.")
	go say("Halo", 3)
	go say("Dunia", 2)
	fmt.Println("Main goroutine menunggu sejenak...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main goroutine selesai.")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: Memulai\n", id)
	time.Sleep(time.Duration(id) * 20 * time.Millisecond) // Shorter sleep
	fmt.Printf("Worker %d: Selesai\n", id)
}

func waitGroupExample() {
	var wg sync.WaitGroup
	numWorkers := 3
	fmt.Printf("Memulai %d worker...\n", numWorkers)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	fmt.Println("Main: Menunggu semua worker selesai...")
	wg.Wait()
	fmt.Println("Main: Semua worker telah selesai.")
}

func sendMessage(ch chan string, msg string) {
	fmt.Printf("Mengirim: '%s'\n", msg)
	time.Sleep(50 * time.Millisecond) // Shorter sleep
	ch <- msg
	fmt.Printf("Terkirim: '%s'\n", msg)
}

func receiveMessage(ch chan string) {
	fmt.Println("Menunggu pesan...")
	receivedMsg := <-ch
	fmt.Printf("Diterima: '%s'\n", receivedMsg)
}

func unbufferedChannelExample() {
	messageChannel := make(chan string)
	go sendMessage(messageChannel, "Halo Channel!")
	go receiveMessage(messageChannel)
	time.Sleep(200 * time.Millisecond) // Allow goroutines to finish
	fmt.Println("Main selesai.")
}

func bufferedChannelExample() {
	bufferedChan := make(chan int, 2)
	fmt.Println("Mengirim 1 ke buffer...")
	bufferedChan <- 1
	fmt.Println("Mengirim 2 ke buffer...")
	bufferedChan <- 2
	fmt.Println("Menerima dari buffer...")
	val1 := <-bufferedChan
	fmt.Printf("Diterima: %d\n", val1)
	fmt.Println("Menerima dari buffer...")
	val2 := <-bufferedChan
	fmt.Printf("Diterima: %d\n", val2)
}

func produce(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Produsen: Mengirim %d\n", i)
		ch <- i
		time.Sleep(10 * time.Millisecond) // Shorter sleep
	}
	fmt.Println("Produsen: Selesai mengirim, menutup channel.")
	close(ch)
}

func consume(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Konsumen %d: Memulai\n", id)
	for value := range ch {
		fmt.Printf("Konsumen %d: Menerima %d\n", id, value)
		time.Sleep(20 * time.Millisecond) // Shorter sleep
	}
	fmt.Printf("Konsumen %d: Channel ditutup, selesai.\n", id)
}

func rangeCloseChannelExample() {
	dataChan := make(chan int, 3)
	var wg sync.WaitGroup
	go produce(dataChan, 5)
	numConsumers := 2
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consume(i, dataChan, &wg)
	}
	wg.Wait()
	fmt.Println("Main: Semua konsumen selesai.")
}

func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond) // Shorter sleep
		ch1 <- "Pesan dari channel 1"
	}()
	go func() {
		time.Sleep(100 * time.Millisecond) // Shorter sleep
		ch2 <- "Pesan dari channel 2"
	}()

	fmt.Println("Menunggu pesan dari ch1 atau ch2...")
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Diterima:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Diterima:", msg2)
		case <-time.After(300 * time.Millisecond): // Shorter timeout
			fmt.Println("Timeout menunggu pesan!")
			return
		}
	}
	fmt.Println("Selesai menerima dua pesan.")
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func mutexExample() {
	counter := SafeCounter{}
	var wg sync.WaitGroup
	numIncrements := 100

	wg.Add(numIncrements)
	for i := 0; i < numIncrements; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Printf("Nilai counter akhir: %d\n", counter.Value())
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func prepareExpensiveData() string {
    return "some data"
}

func MyFunction(data string) {
    _ = data
}

func SayHello(name string) string {
    return fmt.Sprintf("Hello, %s! Welcome!", name)
}

func SayGoodbye(name string) string {
    return fmt.Sprintf("Goodbye, %s. See you!", name)
}

func exampleTestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

func exampleTestSubtract(t *testing.T) {
	 expected := 1
     actual := Subtract(4,3)
	 if actual != expected {
		t.Fatalf("Subtract(...) failed: expected %v, got %v", expected, actual)
	 }
}

func exampleTestMultiply(t *testing.T) {
	t.Log("Testing multiplication...")
	if false {
		 t.Error("Multiplication failed")
	}
}

func exampleTestAddTableDriven(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positif", 2, 3, 5},
		{"Negatif", -1, -5, -6},
		{"Nol", 0, 10, 10},
		{"Positif Negatif", 5, -3, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(st *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				st.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func exampleBenchmarkMyFunction(b *testing.B) {
	setupData := prepareExpensiveData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MyFunction(setupData)
	}
}

func exampleExampleSayHello() {
	message := SayHello("Gopher")
	fmt.Println(message)
}

func exampleExampleSayGoodbye() {
    fmt.Println(SayGoodbye("Alice"))
    fmt.Println(SayGoodbye("Bob"))
}

func exampleTestMain(m *testing.M) {
	fmt.Println("Melakukan setup global...")
	exitCode := m.Run()
	fmt.Println("Melakukan teardown global...")
	os.Exit(exitCode)
}

func exampleTestSomething(t *testing.T) {
	t.Log("Menjalankan TestSomething...")
}

func exampleTestAnother(t *testing.T) {
	 t.Log("Menjalankan TestAnother...")
}

func subtract(a, b int) int {
    return a-b
}


func main() {
	fmt.Println("--- 2. Hello World ---")
	helloWorldExample()

	fmt.Println("\n--- 3. Dasar: Variabel ---")
	variableExamples()

	fmt.Println("\n--- 3. Dasar: Konstanta ---")
	constantsExample()

	fmt.Println("\n--- 3. Dasar: Tipe Dasar ---")
	basicTypesExample()

	fmt.Println("\n--- 3. Dasar: Array ---")
	arrayExample()

	fmt.Println("\n--- 3. Dasar: Slice ---")
	sliceExample()

	fmt.Println("\n--- 3. Dasar: Map ---")
	mapExample()

	fmt.Println("\n--- 3. Dasar: Struct ---")
	structExample()

	fmt.Println("\n--- 3. Dasar: If/Else ---")
	ifElseExample()

	fmt.Println("\n--- 3. Dasar: Switch ---")
	switchExample()

	fmt.Println("\n--- 3. Dasar: For Loop ---")
	forLoopExample()

	fmt.Println("\n--- 3. Dasar: Break/Continue ---")
	breakContinueExample()

	fmt.Println("\n--- 3. Dasar: Defer (Simple) ---")
	simpleDeferExample()

    fmt.Println("\n--- 3. Dasar: Defer (Args Evaluation) ---")
    exampleDeferArgs()

	fmt.Println("\n--- 4. Fungsi ---")
	functionExamples()

	fmt.Println("\n--- 5. Pointer: Dasar ---")
	pointerBasics()

	fmt.Println("\n--- 5. Pointer: Argumen Fungsi ---")
	pointerArgsExample()

	fmt.Println("\n--- 5. Pointer: Nilai Opsional ---")
	pointerForOptionalConfig()

	fmt.Println("\n--- 5. Pointer: Struct ---")
	pointerToStructExample()

	fmt.Println("\n--- 6. Struct & Method: Embedding & Panggil Method ---")
	structMethodEmbeddingExample()

	fmt.Println("\n--- 7. Interface: Dasar & Polimorfisme ---")
	interfaceExample()

	fmt.Println("\n--- 7. Interface: Kosong ---")
	emptyInterfaceExample()

	fmt.Println("\n--- 7. Interface: Type Assertion ---")
	typeAssertionExample()

	fmt.Println("\n--- 7. Interface: Type Switch ---")
	typeSwitchExample()

	fmt.Println("\n--- 8. Error Handling: Konvensi ---")
	errorHandlingConventionExample()

	fmt.Println("\n--- 8. Error Handling: Pembuatan Error ---")
	errorCreationExample()

	fmt.Println("\n--- 8. Error Handling: Wrapping ---")
	errorWrappingExample()

	fmt.Println("\n--- 8. Error Handling: Panic/Recover ---")
	panicRecoverExample()

	fmt.Println("\n--- 9. Konkurensi: Goroutine Sederhana ---")
	goroutineSimpleExample()

	fmt.Println("\n--- 9. Konkurensi: WaitGroup ---")
	waitGroupExample()

	fmt.Println("\n--- 9. Konkurensi: Channel Tak Terbuffer ---")
	unbufferedChannelExample()

	fmt.Println("\n--- 9. Konkurensi: Channel Terbuffer ---")
	bufferedChannelExample()

	fmt.Println("\n--- 9. Konkurensi: Range/Close Channel ---")
	rangeCloseChannelExample()

	fmt.Println("\n--- 9. Konkurensi: Select ---")
	selectExample()

	fmt.Println("\n--- 9. Konkurensi: Mutex ---")
	mutexExample()

    fmt.Println("\n--- 11. Testing Examples (Simulated in main) ---")
	fmt.Println("Simulating test runs (output not identical to 'go test'):")
    var t testing.T
    exampleTestAdd(&t)
    exampleTestSubtract(&t)
    exampleTestMultiply(&t)
    exampleTestAddTableDriven(&t)
    exampleExampleSayHello()
    exampleExampleSayGoodbye()
    exampleTestSomething(&t)
    exampleTestAnother(&t)
    fmt.Println("Benchmark simulation (no actual benchmark run):")
    var b testing.B
    exampleBenchmarkMyFunction(&b)
    fmt.Println("(Note: TestMain would normally wrap all tests, not shown here)")


	fmt.Println("\n--- Selesai ---")
}
