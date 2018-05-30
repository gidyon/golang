package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var allCommands = allcommands{
	{name: "/book", format: "admission gender /book", description: "Book a space if availbale for a student"},
	{name: "/unbook", format: "admission /unbook", description: "Unbook a student from the bookings"},
	{name: "/space", format: "/space", description: "Gives availbale spaces if available", do: space},
	{name: "/total", format: "/total", description: "Give total Number of bookings", do: total},
	{name: "/total", format: "/total -m", description: "Give total Number of males bookings", do: total},
	{name: "/total", format: "/total -f", description: "Give total Number of females bookings", do: total},
	{name: "/males", format: "/males", description: "Details of of all males bookings (in json)", do: all},
	{name: "/females", format: "/females", description: "Details of of all females bookings (in json)", do: all},
	{name: "/all", format: "/all", description: "Details of of all (Both males anf female) bookings (in json)", do: all},
}

func space(string) string {
	if len(allBookings) < 25 {
		return fmt.Sprintf("Space Availble is: %d out of %d", (expectedTotal - len(allBookings)), expectedTotal)
	}
	return fmt.Sprintf("Space Filled is: %d/%d", len(allBookings), expectedTotal)
}

func total(gender string) string {
	var counter uint8
	if gender == "total" {
		return fmt.Sprintf("Total Bookings is: %d/%d\n", (expectedTotal - len(allBookings)), expectedTotal)
	}
	var sex string
	switch gender {
	case "total -m":
		sex = "male"
	case "total -f":
		sex = "female"
	}
	for _, v := range allBookings {
		if v.Gender == sex {
			counter++
		}
	}
	return fmt.Sprintf("Total %d bookings is: %d\n", strings.ToUpper(sex), counter)
}

func all(gender string) string {
	var allGenderStudents allbookings
	//var std student
	if gender == "all" {
		allGenderStudents = append(allGenderStudents, allBookings...)
	} else {
		for _, v := range allBookings {
			if v.Gender == gender[:len(gender)-1] {
				allGenderStudents = append(allGenderStudents, v)
			}
		}
	}
	for _, v := range allGenderStudents {
		json.NewEncoder(os.Stdout).Encode(&v)
	}
	return "Done!"
}

type allcommands []commands

type student struct {
	Admission  string    `json:"admission,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	TimeBooked time.Time `json:"timeBooked,omitempty"`
}

const expectedTotal = 24

func (std *student) book() string {
	if len(allBookings) == expectedTotal {
		return fmt.Sprintf("Space filled!: %d/%d", len(allBookings), expectedTotal)
	}
	for _, v := range allBookings {
		if v.Admission == std.Admission {
			return fmt.Sprintf("Admission: %v already booked a ticket @%v", std.Admission, std.TimeBooked)
		}
	}
	allBookings = append(allBookings, *std)
	return fmt.Sprintf("Booking for Admission: %v is successful @%v", std.Admission, std.TimeBooked)
}

type allbookings []student

func unbook(admission string) string {
	for i, v := range allBookings {
		if v.Admission == admission {
			allBookings = append(allBookings[:i], allBookings[i+1:]...)
			return fmt.Sprintf("Admission: %v removed from bookings", admission)
		}
	}
	return fmt.Sprintf("Admission: %v not in bookings", admission)
}

var allBookings allbookings

type commands struct {
	name        string
	format      string
	description string
	do          func(string) string
}

func main() {
	fmt.Println("Welcome. Type /commands to view availbale commands")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		switch scanner.Text() {
		case "/commands":
			for _, v := range allCommands {
				fmt.Println(" - ", v.name)
			}
			fmt.Println("Type [command name] --help to get a description of the command")
			fmt.Println("Example: /book --help")
		case scanner.Text():
			if strings.HasSuffix(scanner.Text(), "/book") {
				std := student{
					Admission:  scanner.Text()[:strings.Index(scanner.Text(), " ")],
					Gender:     scanner.Text()[strings.Index(scanner.Text(), " "):strings.LastIndex(scanner.Text(), " ")],
					TimeBooked: time.Now(),
				}
				fmt.Println(std.book())
				break
			}
			if strings.HasSuffix(scanner.Text(), "/unbook") {
				fmt.Println(unbook(scanner.Text()[:strings.Index(scanner.Text(), " ")]))
				break
			}
			for _, v := range allCommands {
				if v.format == scanner.Text() && v.name != "/book" && v.name != "/unbook" {
					fmt.Print(v.do(scanner.Text()))
					break
				}
				if v.name+" --help" == scanner.Text() {
					fmt.Printf("Command: %v\n", strings.TrimSuffix(scanner.Text(), " --help"))
					fmt.Printf("\tFormat: %v\n", v.format)
					fmt.Printf("\tDescription: %v\n", v.description)
					break
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}
	}
}
