/* Create a method that calls another method */
package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	BirthDate MyBirthDate
}

type MyBirthDate string

func (p Person) GenerateInfo() {
	age := p.BirthDate.GetAge()
	fmt.Println("Person Name:" + p.Name)
	fmt.Printf("Person Age: %v", age)
}

func (s MyBirthDate) GetAge() int {
	layout := "2006-01-02T15:04:05.000Z"
	birthTime, err := time.Parse(layout, string(s))
	nowTime := time.Now()
	if err != nil {
		return -1

	}
	duration := nowTime.Sub(birthTime)
	return int(duration.Hours() / 24 / 365)
}

func main() {
	p := Person{
		Name:      "Nguyen Van Son",
		BirthDate: "1998-05-05T16:01:05.123Z",
	}
	p.GenerateInfo()
}

// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strconv"
// )

// type Car struct {
//     width  float64
//     height float64
//     length float64
// }

// func (c *Car) GetRealVolume() float64 {
//     return c.width * c.height * c.length
// }

// func main() {
//     car := &Car{}
//     reader := bufio.NewReader(os.Stdin)
//     fmt.Println("Enter car's width:")
//     widthString, _ := reader.ReadString('\n')
//     fmt.Println("Enter car's height:")
//     heightString, _ := reader.ReadString('\n')
//     fmt.Println("Enter car's length:")
//     lengthString, _ := reader.ReadString('\n')
//     if v, err := strconv.ParseFloat(widthString, 64); err == nil {
//         car.width = v
//     }
//     if v, err := strconv.ParseFloat(heightString, 64); err == nil {
//         car.height = v
//     }
//     if v, err := strconv.ParseFloat(lengthString, 64); err == nil {
//         car.length = v
//     }
//     fmt.Println("Car Volumn:")
//     fmt.Println(car.GetRealVolume())

// }
