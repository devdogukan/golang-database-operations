package db

type Truck struct {
	LandVehicle
	CargoAreaVolume int
}

func (truck *Truck) Migrate() error {
	if err := DB.AutoMigrate(truck); err != nil {
		return err
	}

	return nil
}

func (truck *Truck) Get() error {
	if db := DB.First(&truck, "id = ?", truck.Id); db.Error != nil {
		return db.Error
	}

	return nil
}

func (truck *Truck) GetAll() (interface{}, error) {
	var trucks []Truck
	if db := DB.Find(&trucks); db.Error != nil {
		return nil, db.Error
	}
	return trucks, nil
}

func (truck *Truck) Create() error {
	if db := DB.Create(&truck); db.Error != nil {
		return db.Error
	}
	return nil
}

func (truck *Truck) Delete() error {
	if db := DB.Delete(&truck, truck); db.Error != nil {
		return db.Error
	}
	return nil
}

func (truck *Truck) Updates() error {
	if db := DB.Model(&truck).
		Where("id = ?", truck.Id).
		Updates(truck); db.Error != nil {
		return db.Error
	}

	return nil
}

func (truck *Truck) Update(column string, value interface{}) error {
	if db := DB.Model(&truck).
		Where("id = ?", truck.Id).
		Update(column, value); db.Error != nil {
		return db.Error
	}
	return nil
}
