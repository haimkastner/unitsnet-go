# unitsnet-go

[![unitsnet-go](https://github.com/haimkastner/unitsnet-go/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/haimkastner/unitsnet-go/actions/workflows/build.yml)

 [![Latest Release](https://img.shields.io/github/v/release/haimkastner/unitsnet-go)](https://github.com/haimkastner/unitsnet-go/releases) 


<!-- [![GitHub stars](https://img.shields.io/github/stars/haimkastner/unitsnet-go.svg?style=social&label=Stars)](https://github.com/haimkastner/unitsnet-go/stargazers)  -->
[![License](https://img.shields.io/github/license/haimkastner/unitsnet-go.svg?style=social)](https://github.com/haimkastner/unitsnet-go/blob/master/LICENSE)

The unitsnet-go package provides an efficient way to store unit variables and perform easy conversions to different units when it required. 

It offers support for more than 100 unit types across various unit categories, including pretty-printing, comparison, and arithmetic methods. 

The API is designed to be user-friendly and straightforward to use.

The library is built on top of the [Units.NET](https://github.com/angularsen/UnitsNet) project and leverages their [definitions sources](https://github.com/angularsen/UnitsNet/tree/master/Common/UnitDefinitions) to generate the Golang units.

###### The unitsnet-go package does not require any external dependencies or packages to function.

[Units.NET on other platforms](#unitsnet-on-other-platforms)

## Install package

```bash 
go get github.com/haimkastner/unitsnet-go
```

## Example Usage

```golang
package main

import (
    "log"
    "github.com/haimkastner/unitsnet-go/units"
)

func main() {

    // Create a factory instance
    af := units.AngleFactory{}
    
    angle, _ := af.FromDegrees(180)
    // equals to
    angle, _ := af.CreateAngle(180, units.AngleDegree)

    log.Println(angle.Radians())      // 3.141592653589793
    log.Println(angle.Microradians()) // 3141592.65358979
    log.Println(angle.Gradians())     // 200
    log.Println(angle.Microdegrees()) // 180000000

    // As an alternative, a convert style method are also available
    log.Println(angle.Convert(units.AngleRadian))      // 3.141592653589793
    log.Println(angle.Convert(units.AngleMicroradian)) // 3141592.65358979
    log.Println(angle.Convert(units.AngleGradian))     // 200
    log.Println(angle.Convert(units.AngleMicrodegree)) // 180000000

    // Print the default unit to_string (The default for angle is degrees)
    log.Println(angle)  // 180.00 °

    // Specify unit and fraction digits max length
    log.Println(angle.ToString(units.AngleDegree, 0)) // 180 °
    log.Println(angle.ToString(units.AngleRadian, 3)) // 3.141 rad
}

```

## Additional methods

Check, compare, calculate etc. with unitsnet:

```golang
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
```


## DTO - Data Transfer Object

As UnitsNet provides a convenient way to work within a running service, there are occasions where the data needs to be exposed outside of the service, typically through an API containing the unit value or consumed from an API.

To support this with a clear API schema and make it easy to convert to and from this schema to the specific format, it's recommended to use DTOs and the UnitsNet flavor converters.
```golang
lf := units.LengthFactory{}

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
// Get the json representation of the DTO
lengthJson, _ := lengthDto.ToJSON() // {"value":100.01,"unit":"Meter"}
// Obtain DTO instance from a json representation
lengthDtoFromJson, _ := units.LengthDtoFactory{}.FromJSON(lengthJson)
// Obtain Length unit from a DTO instance
lengthFromJson, _ := lf.FromDto(*lengthDtoFromJson)
log.Println(lengthFromJson) // 100.01 m
```

Check out the OpenAPI [unitsnet-openapi-spec](https://haimkastner.github.io/unitsnet-openapi-spec-example/) example schema.

Also, refer to the detailed discussions on GitHub: [haimkastner/unitsnet-js#31](https://github.com/haimkastner/unitsnet-js/issues/31) & [angularsen/UnitsNet#1378](https://github.com/angularsen/UnitsNet/issues/1378).

## Supported units

[UnitsNet supported Units](Units.md)


## Units.NET on other platforms

Get the same strongly typed units on other platforms, based on the same [unit definitions](/Common/UnitDefinitions).

| Language                   | Name        | Package                                           					 | Repository                                           | Maintainers  |
|----------------------------|-------------|---------------------------------------------------------------------|------------------------------------------------------|--------------|
| C#                         | UnitsNet    | [nuget](https://www.nuget.org/packages/UnitsNet/) 					 | [github](https://github.com/angularsen/UnitsNet)     | @angularsen  |
| JavaScript /<br>TypeScript | unitsnet-js | [npm](https://www.npmjs.com/package/unitsnet-js)  					 | [github](https://github.com/haimkastner/unitsnet-js) | @haimkastner |
| Python                     | unitsnet-py | [pypi](https://pypi.org/project/unitsnet-py)      					 | [github](https://github.com/haimkastner/unitsnet-py) | @haimkastner |
| Golang                     | unitsnet-go | [pkg.go.dev](https://pkg.go.dev/github.com/haimkastner/unitsnet-go) | [github](https://github.com/haimkastner/unitsnet-go) | @haimkastner |

