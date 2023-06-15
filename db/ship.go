package db

type Ship struct {
	SeaVehicle
	DeckVolume int
}

func (ship *Ship) Migrate() error {
	if err := DB.AutoMigrate(ship); err != nil {
		return err
	}

	return nil
}

func (ship *Ship) Get() error {
	if db := DB.First(&ship, "id = ?", ship.Id); db.Error != nil {
		return db.Error
	}
	return nil
}

func (ship *Ship) GetAll() (interface{}, error) {
	var ships []Ship
	if db := DB.Find(&ships); db.Error != nil {
		return nil, db.Error
	}
	return ships, nil
}

func (ship *Ship) Create() error {
	if db := DB.Create(&ship); db.Error != nil {
		return db.Error
	}
	return nil
}

func (ship *Ship) Delete() error {
	if db := DB.Delete(&ship, ship); db.Error != nil {
		return db.Error
	}
	return nil
}

func (ship *Ship) Updates() error {
	if db := DB.Model(&ship).
		Where("id = ?", ship.Id).
		Updates(ship); db.Error != nil {
		return db.Error
	}

	return nil
}

func (ship *Ship) Update(column string, value interface{}) error {
	if db := DB.Model(&ship).
		Where("id = ?", ship.Id).
		Update(column, value); db.Error != nil {
		return db.Error
	}
	return nil
}
