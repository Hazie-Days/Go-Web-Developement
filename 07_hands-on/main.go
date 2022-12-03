package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip string
	Price         float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

type restaurant struct {
	Name,Location,OpenHour1, OpenHour2 string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := restaurants{
		restaurant{
			Name: " Blu Kouzina",
			Location: "Dempsey, East Coast",
			OpenHour1: "11.30am-3.30pm",
			OpenHour2: "5.30-10pm",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:    "Gemista",
							Descrip: "Stuffed vegetables, marinated rice and potatoes",
							Price:   29.80,
						},
						item{
							Name:    "Spanakorizo",
							Descrip: "Spinach cooked with rice, herbs, lemon and EVOO",
							Price:   20.80,
						},
						item{
							Name:    "Watermelon Salad",
							Descrip: "Watermelon, feta and mint",
							Price:   18.80,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:    "Beef Kalamaki",
							Descrip: "Beef skewer, fresh onion, sliced tomato, pita and tzatziki",
							Price:   21.80,
						},
						item{
							Name:    "Meat Platter (Small)",
							Descrip: "2 lamb meatballs, 1 beef kalamaki, 1 organic chicken skewer and 2 lamb chops",
							Price:   58.80,
						},
						item{
							Name:    "French Fries",
							Descrip: "French eat potatoe fingers",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:    "Pasta Bolognese",
							Descrip: "From Italy delicious eating",
							Price:   7.95,
						},
						item{
							Name:    "Steak",
							Descrip: "Dead cow grilled bloody",
							Price:   13.95,
						},
						item{
							Name:    "Bistro Potatoe",
							Descrip: "Bistro bar wood American bacon",
							Price:   6.95,
						},
					},
				},
			},
		},
		restaurant{
			Name: " KOMA Singapore",
			Location: " 2 Bayfront Ave, # B1 - 67, Singapore 018972",
			OpenHour1: "11.30am-3pm",
			OpenHour2: "5pm-12am",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:    "Chawanmushi",
							Descrip: "Egg custard, sea urchin, salmon roe, ankake sauce, yuzu zest",
							Price:   8.00,
						},
						item{
							Name:    "Cheerios",
							Descrip: "American eating food traditional now",
							Price:   3.95,
						},
						item{
							Name:    "Juice Orange",
							Descrip: "Delicious drinking in throat squeezed fresh",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:    "D.i.y. Spicy Tuna",
							Descrip: "Crispy rice, sweet soy, chives",
							Price:   18.00,
						},
						item{
							Name:    "Taro Chip Tuna Tacos",
							Descrip: "Avocado, jalapeño, lemon, cilantro, mixed cabbage salad",
							Price:   21.00,
						},
						item{
							Name:    "French Fries",
							Descrip: "French eat potatoe fingers",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:    "King Crab Dynamite",
							Descrip: "Spicy aioli, tobiko",
							Price:   42.00,
						},
						item{
							Name:    "Steak",
							Descrip: "Dead cow grilled bloody",
							Price:   13.95,
						},
						item{
							Name:    "Bistro Potatoe",
							Descrip: "Bistro bar wood American bacon",
							Price:   6.95,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
