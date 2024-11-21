package main

import (
	"fmt"
	"sync"
)

/* Task: Simulating a Ticket Booking System

You're building a ticket booking system for a movie theater. The system should:
    Handle multiple users trying to book tickets at the same time.
    Each user should either successfully book a ticket or be told that the tickets are sold out.

Requirements:
    The theater has a limited number of tickets (e.g., 5 tickets).
    Multiple users can try to book tickets concurrently.
    If a ticket is successfully booked, it should reduce the number of available tickets.
    If there are no tickets left, the user should receive a "Sold Out" message.
*/

var ticketsAvailable = 5
var mu sync.Mutex

func bookTicket(user string, wg *sync.WaitGroup) {
	defer wg.Done()

	if ticketsAvailable <= 0 {
		fmt.Printf("%s: Sold Out!\n", user)
		return
	}

	mu.Lock()

	if ticketsAvailable > 0 {
		ticketsAvailable--
		fmt.Printf("%s booked a ticket! Tickets remaining: %d\n", user, ticketsAvailable)
	} else {
		fmt.Printf("Sold Out")
	}

	mu.Unlock()

}

func main() {

	var wg sync.WaitGroup

	users := []string{"User1", "User2", "User3", "User4", "User5", "User6", "User7"}

	for _, user := range users {
		wg.Add(1)
		go bookTicket(user, &wg)
	}

	wg.Wait()
}
