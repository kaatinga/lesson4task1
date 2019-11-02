package main

import "log"

type CarData struct {
	Model       string
	YearOfIssue int
}

type Cars []CarData

type Info interface {
	List()
}

func (i *CarData) List() {
	log.Println(i.Model)
	log.Println(i.YearOfIssue)
}

func (cars *Cars) List() {
	for _, value := range *cars {
		log.Println(value.Model)
		log.Println(value.YearOfIssue)
	}
}

func (car *CarData) Add(Model string, YearOfIssue int) {
	car.Model = Model
	car.YearOfIssue = YearOfIssue
}

func (cars *Cars) Add(Model string, YearOfIssue int) {
	*cars = append(*cars, CarData{Model, YearOfIssue})
}

func main() {

	var cars Cars
	var something Info

	log.Println("Работаем обычным образом")
	cars = append(cars, CarData{"Ford", 1997})
	cars = append(cars, CarData{"Nissan", 2006})
	for _, value := range cars { // выводим все машины
		log.Println(value.Model)
		log.Println(value.YearOfIssue)
	}
	log.Println(cars[0].Model) // выводим данные одной машины
	log.Println(cars[0].YearOfIssue)
	// типа не очень красиво

	log.Println("Теперь работаем красиво, через метод")
	cars[1].Add("Dodge", 2007) // обновляем данные одной машины зная индекс
	cars.Add("Toyota", 2017)   // или просто добавляем в конец
	cars.List()                // выводим более красиво
	cars[0].List()

	log.Println("Или через интерфейс")

	something = &cars[1]
	something.List()
	something = &cars
	something.List()
}
