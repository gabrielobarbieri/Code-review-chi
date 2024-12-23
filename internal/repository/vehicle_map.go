package repository

import (
	"app/internal"
	"errors"
)

var (
	ErrVehiclesNotFound = errors.New("no vehicles found")
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) FindByBrandAndYearInterval(brand string, startYear, endYear int) (v []internal.Vehicle, err error) {
	var vehicles []internal.Vehicle

	for _, v := range r.db {
		if v.Brand == brand {
			if v.FabricationYear > startYear && v.FabricationYear < endYear {
				vehicles = append(vehicles, v)
			}
		}
	}

	if len(vehicles) == 0 {
		return nil, ErrVehiclesNotFound
	}

	return vehicles, nil
}
