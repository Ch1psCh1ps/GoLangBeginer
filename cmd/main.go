package main

import (
	"encoding/json"
	"log"
)

type Ticket struct {
	Sector      string `json:"sector"`
	Row         int    `json:"row"`
	Seat_number int    `json:"seat_number"`
}

type Sector struct {
	Name    string   `json:"name"`
	Tickets []Ticket `json:"tickets"`
}

func getTickets() []Ticket {
	return []Ticket{
		{
			Sector:      "A",
			Row:         1,
			Seat_number: 1,
		},
		{
			Sector:      "A",
			Row:         1,
			Seat_number: 2,
		},
		{
			Sector:      "A",
			Row:         1,
			Seat_number: 3,
		},
		{
			Sector:      "A",
			Row:         2,
			Seat_number: 1,
		},
		{
			Sector:      "A",
			Row:         2,
			Seat_number: 2,
		},
		{
			Sector:      "A",
			Row:         2,
			Seat_number: 3,
		},
		{
			Sector:      "B",
			Row:         1,
			Seat_number: 1,
		},
		{
			Sector:      "B",
			Row:         1,
			Seat_number: 2,
		},
		{
			Sector:      "B",
			Row:         1,
			Seat_number: 3,
		},
		{
			Sector:      "B",
			Row:         2,
			Seat_number: 1,
		},
		{
			Sector:      "B",
			Row:         2,
			Seat_number: 2,
		},
		{
			Sector:      "B",
			Row:         2,
			Seat_number: 3,
		},
		{
			Sector:      "C",
			Row:         1,
			Seat_number: 1,
		},
		{
			Sector:      "C",
			Row:         1,
			Seat_number: 2,
		},
		{
			Sector:      "C",
			Row:         1,
			Seat_number: 3,
		},
		{
			Sector:      "C",
			Row:         2,
			Seat_number: 1,
		},
		{
			Sector:      "C",
			Row:         2,
			Seat_number: 2,
		},
		{
			Sector:      "C",
			Row:         2,
			Seat_number: 3,
		},
	}
}

func main() {
	tickets := getTickets()

	sectors := []Sector{}

	for _, ticket := range tickets {
		found := false
		for i, sector := range sectors {
			if ticket.Sector == sector.Name {
				found = true
				sectors[i].Tickets = append(sector.Tickets, ticket)
			}
		}

		if !found {
			sectors = append(sectors, Sector{ticket.Sector, []Ticket{ticket}})
		}
	}

	j, _ := json.Marshal(sectors)

	log.Println(string(j))

	/*	if nameSector == "A" {
			groupSector := sectors[0]
			fmt.Println("По запросу //#сектора сектор найдено:", groupSector)
		} else if nameSector == "B" {
			groupSector := sectors[1]
			fmt.Println("По запросу //#сектора сектор найдено:", groupSector)
		} else if nameSector == "C" {
			groupSector := sectors[2]
			fmt.Println("По запросу //#сектора сектор найдено:", groupSector)
		} else {
			fmt.Println("Такого сектора не существует")
		}

		if nameRow == 1 {
			groupRow := rows[0]
			fmt.Println("По запросу //#ряда ряд найдено:", groupRow)
		} else if nameRow == 2 {
			groupRow := rows[1]
			fmt.Println("По запросу//#ряда ряд найдено:", groupRow)
		} else {
			fmt.Println("Такого ряда не существует")
		}*/
}
