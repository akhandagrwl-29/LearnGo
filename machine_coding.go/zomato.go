package main

// // Restaurant struct represents a restaurant
// type Restaurant struct {
// 	ID         string
// 	Name       string
// 	City       string
// 	Area       string
// 	Cuisine    string
// 	CostForTwo int
// 	Slots      map[time.Time]bool
// }

// // User struct represents a user
// type User struct {
// 	ID   string
// 	Name string
// 	// Add more user details if needed
// }

// // Booking struct represents a booking
// type Booking struct {
// 	ID           string
// 	UserID       string
// 	RestaurantID string
// 	BookingTime  time.Time
// 	NoOfPeople   int
// }

// var restaurantsDB = make(map[string]Restaurant)
// var usersDB = make(map[string]User)
// var bookingsDB = make(map[string]Booking)

// func registerRestaurant(name, city, area, cuisine string, costForTwo int) string {
// 	restaurantID := generateID()
// 	restaurant := Restaurant{
// 		ID:         restaurantID,
// 		Name:       name,
// 		City:       city,
// 		Area:       area,
// 		Cuisine:    cuisine,
// 		CostForTwo: costForTwo,
// 		Slots:      make(map[time.Time]bool),
// 	}
// 	restaurantsDB[restaurantID] = restaurant
// 	return restaurantID
// }

// func updateTimeSlots(restaurantID string, slots []time.Time) {
// 	restaurant, exists := restaurantsDB[restaurantID]
// 	if !exists {
// 		fmt.Println("Restaurant not found")
// 		return
// 	}

// 	for _, slot := range slots {
// 		restaurant.Slots[slot] = true
// 	}

// 	restaurantsDB[restaurantID] = restaurant
// }

// func searchRestaurant(loggedInUserID, city, area, cuisine string, costForTwo int) []Restaurant {
// 	var result []Restaurant
// 	for _, restaurant := range restaurantsDB {
// 		// Add search criteria based on the parameters
// 		if city != "" && restaurant.City != city {
// 			continue
// 		}
// 		if area != "" && restaurant.Area != area {
// 			continue
// 		}
// 		if cuisine != "" && restaurant.Cuisine != cuisine {
// 			continue
// 		}
// 		if costForTwo > 0 && restaurant.CostForTwo > costForTwo {
// 			continue
// 		}
// 		// Add more search criteria if needed

// 		result = append(result, restaurant)
// 	}
// 	return result
// }

// func bookTable(restaurantID, userID string, bookingTime time.Time, noOfPeople int) (string, error) {
// 	restaurant, exists := restaurantsDB[restaurantID]
// 	if !exists {
// 		return "", fmt.Errorf("Restaurant not found")
// 	}

// 	if !restaurant.Slots[bookingTime] {
// 		return "", fmt.Errorf("Slot not available")
// 	}

// 	bookingID := generateID()
// 	booking := Booking{
// 		ID:           bookingID,
// 		UserID:       userID,
// 		RestaurantID: restaurantID,
// 		BookingTime:  bookingTime,
// 		NoOfPeople:   noOfPeople,
// 	}

// 	bookingsDB[bookingID] = booking
// 	delete(restaurant.Slots, bookingTime)
// 	restaurantsDB[restaurantID] = restaurant

// 	return bookingID, nil
// }

// func generateID() string {
// 	return fmt.Sprintf("%d", time.Now().UnixNano())
// }

// func main() {
// 	// Example Usage
// 	restaurantID := registerRestaurant("Restaurant A", "City A", "Area A", "Cuisine A", 50)
// 	updateTimeSlots(restaurantID, []time.Time{
// 		time.Now().Add(time.Hour * 24),
// 		time.Now().Add(time.Hour * 48),
// 	})
// 	t1 := time.Now()
// 	t2 := time.Now()

// 	if t1 == t2 {
// 		fmt.Println("T1 and T2 are equal", t1, t2)
// 	}

// 	userID := generateID()
// 	bookingID, err := bookTable(restaurantID, userID, time.Now().Add(time.Hour*24), 2)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Booking successful. Booking ID:", bookingID)
// 	}
// }
