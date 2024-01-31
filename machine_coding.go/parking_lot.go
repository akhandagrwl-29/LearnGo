package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ParkingLot struct {
	parkingLotID  string
	numFloors     int
	slotsPerFloor int
	slots         map[string]map[int][]int
	ticketCounter int
}

func NewParkingLot(parkingLotID string, numFloors, slotsPerFloor int) *ParkingLot {
	return &ParkingLot{
		parkingLotID:  parkingLotID,
		numFloors:     numFloors,
		slotsPerFloor: slotsPerFloor,
		slots:         make(map[string]map[int][]int),
		ticketCounter: 1,
	}
}

func (pl *ParkingLot) getVehicleType(floor, slot int) string {
	if slot == 1 {
		return "TRUCK"
	} else if slot >= 2 && slot <= 3 {
		return "BIKE"
	}
	return "CAR"
}

func (pl *ParkingLot) createParkingLot() {
	fmt.Printf("Created parking lot with %d floors and %d slots per floor\n", pl.numFloors, pl.slotsPerFloor)
}

func (pl *ParkingLot) parkVehicle(vehicleType, regNo, color string) {
	if _, exists := pl.slots[vehicleType]; !exists {
		fmt.Println("Invalid Vehicle Type")
		return
	}

	for floor := 1; floor <= pl.numFloors; floor++ {
		if slots, exists := pl.slots[vehicleType][floor]; exists && len(slots) > 0 {
			slot := slots[0]
			pl.slots[vehicleType][floor] = slots[1:]
			ticketID := fmt.Sprintf("%s_%d_%d", pl.parkingLotID, floor, slot)
			fmt.Printf("Parked vehicle. Ticket ID: %s\n", ticketID)
			return
		}
	}

	fmt.Println("Parking Lot Full")
}

func (pl *ParkingLot) unparkVehicle(ticketID string) {
	parts := strings.Split(ticketID, "_")
	if len(parts) != 3 || parts[0] != pl.parkingLotID {
		fmt.Println("Invalid Ticket")
		return
	}

	floor, _ := strconv.Atoi(parts[1])
	slot, _ := strconv.Atoi(parts[2])
	vehicleType := pl.getVehicleType(floor, slot)

	if slots, exists := pl.slots[vehicleType][floor]; exists {
		pl.slots[vehicleType][floor] = append(slots, slot)
		fmt.Printf("Unparked vehicle with Ticket ID: %s\n", ticketID)
	} else {
		fmt.Println("Invalid Ticket")
	}
}

func (pl *ParkingLot) displaySlots(displayType, vehicleType string) {
	if _, exists := pl.slots[vehicleType]; !exists {
		fmt.Println("Invalid Vehicle Type")
		return
	}

	for floor := 1; floor <= pl.numFloors; floor++ {
		if slots, exists := pl.slots[vehicleType][floor]; exists {
			switch displayType {
			case "free_count":
				fmt.Printf("No. of free slots for %s on Floor %d: %d\n", vehicleType, floor, len(slots))
			case "free_slots":
				fmt.Printf("Free slots for %s on Floor %d: %s\n", vehicleType, floor, strings.Join(strings.Fields(fmt.Sprint(slots)), ","))
			case "occupied_slots":
				occupiedSlots := make([]int, 0, pl.slotsPerFloor-len(slots))
				for i := 1; i <= pl.slotsPerFloor; i++ {
					if !contains(slots, i) {
						occupiedSlots = append(occupiedSlots, i)
					}
				}
				fmt.Printf("Occupied slots for %s on Floor %d: %s\n", vehicleType, floor, strings.Join(strings.Fields(fmt.Sprint(occupiedSlots)), ","))
			}
		}
	}
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	var parkingLot *ParkingLot

	for {
		var command string
		fmt.Scan(&command)

		switch command {
		case "create_parking_lot":
			var parkingLotID string
			var numFloors, slotsPerFloor int
			fmt.Scan(&parkingLotID, &numFloors, &slotsPerFloor)
			parkingLot = NewParkingLot(parkingLotID, numFloors, slotsPerFloor)
			parkingLot.createParkingLot()

		case "park_vehicle":
			var vehicleType, regNo, color string
			fmt.Scan(&vehicleType, &regNo, &color)
			parkingLot.parkVehicle(vehicleType, regNo, color)

		case "unpark_vehicle":
			var ticketID string
			fmt.Scan(&ticketID)
			parkingLot.unparkVehicle(ticketID)

		case "display":
			var displayType, vehicleType string
			fmt.Scan(&displayType, &vehicleType)
			parkingLot.displaySlots(displayType, vehicleType)

		case "exit":
			return
		}
	}
}
