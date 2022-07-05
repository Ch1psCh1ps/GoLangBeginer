package main

import (
	"fmt"
)

type Ticket struct {
	sector      string
	row         int
	seat_number int
}

type Sector struct {
	name    string
	tickets []Ticket
}

func getTickets() []Ticket {
	return []Ticket{
		{
			sector:      "A",
			row:         1,
			seat_number: 1,
		},
		{
			sector:      "A",
			row:         1,
			seat_number: 2,
		},
		{
			sector:      "A",
			row:         1,
			seat_number: 3,
		},
		{
			sector:      "A",
			row:         2,
			seat_number: 1,
		},
		{
			sector:      "A",
			row:         2,
			seat_number: 2,
		},
		{
			sector:      "A",
			row:         2,
			seat_number: 3,
		},
		{
			sector:      "B",
			row:         1,
			seat_number: 1,
		},
		{
			sector:      "B",
			row:         1,
			seat_number: 2,
		},
		{
			sector:      "B",
			row:         1,
			seat_number: 3,
		},
		{
			sector:      "B",
			row:         2,
			seat_number: 1,
		},
		{
			sector:      "B",
			row:         2,
			seat_number: 2,
		},
		{
			sector:      "B",
			row:         2,
			seat_number: 3,
		},
		{
			sector:      "C",
			row:         1,
			seat_number: 1,
		},
		{
			sector:      "C",
			row:         1,
			seat_number: 2,
		},
		{
			sector:      "C",
			row:         1,
			seat_number: 3,
		},
		{
			sector:      "C",
			row:         2,
			seat_number: 1,
		},
		{
			sector:      "C",
			row:         2,
			seat_number: 2,
		},
		{
			sector:      "C",
			row:         2,
			seat_number: 3,
		},
	}
}

func main() {
	tickets := getTickets()

	sectors := []Sector{}

	for _, ticket := range tickets {
		found := false
		for i, sector := range sectors {
			if ticket.sector == sector.name {
				found = true
				sectors[i].tickets = append(sector.tickets, ticket)
			}
		}

		if !found {
			sectors = append(sectors, Sector{ticket.sector, []Ticket{ticket}})
		}
	}

	fmt.Println(sectors)

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
