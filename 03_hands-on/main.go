// 1. Create a data structure to pass to a template which
// * contains information about California hotels including Name, Address, City, Zip, Region
// * region can be: Southern, Central, Northern
// * can hold an unlimited number of hotels

package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region  string
}
type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
				hotel{
						Name: "The Mission Inn Hotel",
						Address: "3649 Mission Inn Avenue",
						City: "RiverSide",
						Zip: "92501",
						Region: "Eastern",
					},
					hotel{
						Name: "Casa Del Mar Inn",
						Address: "18 Bath Street",
						City: "Santa Barbara",
						Zip:"93101",
						Region: "Western",
						},
						hotel{
							Name: "Hotel Yosemite View Lodge",
							Address: "11136 Highway 140",
							City: "El Portal",
							Zip:"95318",
							Region: "northern",
						},
					}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}