package vehicles

import "fmt"

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicleType   = "CAR"
	TruckVehicleType = "TRUCK"
	MotorVehicleType = "MOTORCYCLE"
	BikeVehicleType  = "BIKE"
)

type Car struct {
	Time int
}

func (c *Car) Distance() float64 {
	return (float64(c.Time) / 60) * 130
}

func (c *Bike) Distance() float64 {
	return (float64(c.Time) / 60) * 40
}

type MotorCicle struct {
	Time int
}

func (c *MotorCicle) Distance() float64 {
	return (float64(c.Time) / 60) * 100
}

type Truck struct {
	Time int
}

func (c *Truck) Distance() float64 {
	return float64(c.Time) * 60
}

type Bike struct {
	Time int
}

func New(vehicle string, time int) (Vehicle, error) {
	switch vehicle {
	case CarVehicleType:
		return &Car{Time: time}, nil
	case TruckVehicleType:
		return &Truck{Time: time}, nil
	case MotorVehicleType:
		return &MotorCicle{Time: time}, nil
	case BikeVehicleType:
		return &Bike{Time: time}, nil
	default:
		return nil, fmt.Errorf("invalid vehicle type")
	}
}
