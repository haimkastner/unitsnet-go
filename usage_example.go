package main

import (
	"log"

	"github.com/haimkastner/unitsnet-go/units"
)

func main() {

	// Create a factory instance
	af := units.AngleFactory{}

	angle, _ := af.FromDegrees(180)
	angle2, _ := af.CreateAngle(180, units.AngleDegree)

	log.Println(angle.Radians())      // 3.141592653589793
	log.Println(angle.Microradians()) // 3141592.65358979
	log.Println(angle.Gradians())     // 200
	log.Println(angle.Microdegrees()) // 180000000

	log.Println(angle.Convert(units.AngleRadian))      // 3.141592653589793
	log.Println(angle.Convert(units.AngleMicroradian)) // 3141592.65358979
	log.Println(angle.Convert(units.AngleGradian))     // 200
	log.Println(angle.Convert(units.AngleMicrodegree)) // 180000000

	log.Println(angle)                                // 180.00 °
	log.Println(angle.ToString(units.AngleDegree, 0)) // 180 °
	log.Println(angle.ToString(units.AngleRadian, 3)) // 3.141 rad

	lf := units.LengthFactory{}

	length1, _ := lf.FromMeters(10)
	length2, _ := lf.FromDecimeters(100)
	length3, _ := lf.FromMeters(3)

	// 'Equals' method
	log.Println(length1.Equals(length2)) // true
	log.Println(length1.Equals(length3)) // false

	// 'CompareTo' method
	log.Println(length1.CompareTo(length3)) // 1;
	log.Println(length3.CompareTo(length1)) // -1;
	log.Println(length2.CompareTo(length1)) // 0;

	// Arithmetics methods
	results1 := length1.Add(length3)
	results2 := length1.Subtract(length3)
	results3 := length1.Multiply(length3)
	results4 := length1.Divide(length3)
	log.Println(results1) // 13.00 m
	log.Println(results2) // 7.00 m
	log.Println(results3) // 30.00 m
	log.Println(results4) // 3.33 m

	log.Println(angle2) // 180.00 °

	length, _ := lf.FromMeters(100.01)
	// Obtain the DTO object as json, represented by the default unit - meter
	lengthDtoJson, _ := length.ToDtoJSON(nil)
	log.Println(string(lengthDtoJson)) // {"value":100.01,"unit":"Meter"}

	// Obtain the same value but represent DTO in KM
	lengthKilometer := units.LengthKilometer // Default value
	lengthDtoRepresentsInKmJson, _ := length.ToDtoJSON(&lengthKilometer)
	log.Println(string(lengthDtoRepresentsInKmJson)) // {"value":0.10001,"unit":"Kilometer"}

	// Load JSON to DTO, and load
	lengthFromMetersDto, _ := lf.FromDtoJSON(lengthDtoJson)
	log.Println(lengthFromMetersDto) // 100.01 m
	// The exact same value as
	lengthFromKmDto, _ := lf.FromDtoJSON(lengthDtoRepresentsInKmJson)
	log.Println(lengthFromKmDto) // 100.01 m

	// Additionally, it supports creating and handling as a DTO instance, as well as creating and converting to/from JSON.

	// Get a DTO instance from a Length instance
	lengthDto := length.ToDto(nil)
	// # Get the json representation of the DTO
	lengthJson, _ := lengthDto.ToJSON() // {"value":100.01,"unit":"Meter"}
	// # Obtain DTO instance from a json representation
	lengthDtoFromJson, _ := units.LengthDtoFactory{}.FromJSON(lengthJson)
	// # Obtain Length unit from a DTO instance
	lengthFromJson, _ := lf.FromDto(*lengthDtoFromJson)
	log.Println(lengthFromJson) // 100.01 m
}
