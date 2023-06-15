package main

import (
	model "example/db"
	"fmt"
)

func main() {
	model.Initialize()

	// ----------------------------- Migrate ---------------------------------
	if err := model.Migrate(&model.Truck{}); err != nil {
		fmt.Println(err)
	}

	if err := model.Migrate(&model.Ship{}); err != nil {
		fmt.Println(err)
	}

	// ------------------------------ Create --------------------------------
	truck1 := model.Truck{
		LandVehicle: model.LandVehicle{
			Vehicle: model.Vehicle{
				Id:       "31-49",
				Brand:    "Meta",
				Engine:   890,
				Capacity: 4,
			},
			Whell: 5,
		},
		CargoAreaVolume: 2000,
	}

	if err := model.Create(&truck1); err != nil {
		fmt.Println(err.Error())
	}

	// ----------------------------- Get ---------------------------------
	err := model.Get(&truck1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(truck1)

	// ----------------------------- GetAll ---------------------------------
	var truck2 model.Truck

	trucks, err := model.GetAll(&truck2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(trucks)

	// ----------------------------- Updates ---------------------------------
	truck3 := model.Truck{
		LandVehicle: model.LandVehicle{
			Vehicle: model.Vehicle{
				Id:       truck1.Id,
				Brand:    "Meta",
				Engine:   890,
				Capacity: 4,
			},
			Whell: 5,
		},
		CargoAreaVolume: 2000,
	}

	if err := model.Updates(&truck3); err != nil {
		fmt.Println(err)
	}

	fmt.Println(truck3)

	// ----------------------------- Update ---------------------------------
	var truck5 model.Truck
	truck5.Id = "68-69"

	if err := model.Update(&truck5, "brand", "Mercedes"); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(truck5)

	// ----------------------------- Delete ---------------------------------
	var truck4 model.Truck
	truck4.Id = "31-49"

	if err := model.Delete(&truck4); err != nil {
		fmt.Println(err)
	}

	fmt.Println(truck4)

}
