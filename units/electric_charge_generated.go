// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricChargeUnits defines various units of ElectricCharge.
type ElectricChargeUnits string

const (
    
        // 
        ElectricChargeCoulomb ElectricChargeUnits = "Coulomb"
        // 
        ElectricChargeAmpereHour ElectricChargeUnits = "AmpereHour"
        // 
        ElectricChargePicocoulomb ElectricChargeUnits = "Picocoulomb"
        // 
        ElectricChargeNanocoulomb ElectricChargeUnits = "Nanocoulomb"
        // 
        ElectricChargeMicrocoulomb ElectricChargeUnits = "Microcoulomb"
        // 
        ElectricChargeMillicoulomb ElectricChargeUnits = "Millicoulomb"
        // 
        ElectricChargeKilocoulomb ElectricChargeUnits = "Kilocoulomb"
        // 
        ElectricChargeMegacoulomb ElectricChargeUnits = "Megacoulomb"
        // 
        ElectricChargeMilliampereHour ElectricChargeUnits = "MilliampereHour"
        // 
        ElectricChargeKiloampereHour ElectricChargeUnits = "KiloampereHour"
        // 
        ElectricChargeMegaampereHour ElectricChargeUnits = "MegaampereHour"
)

// ElectricChargeDto represents a ElectricCharge measurement with a numerical value and its corresponding unit.
type ElectricChargeDto struct {
    // Value is the numerical representation of the ElectricCharge.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricCharge, as defined in the ElectricChargeUnits enumeration.
	Unit  ElectricChargeUnits `json:"unit"`
}

// ElectricChargeDtoFactory groups methods for creating and serializing ElectricChargeDto objects.
type ElectricChargeDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricChargeDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricChargeDtoFactory) FromJSON(data []byte) (*ElectricChargeDto, error) {
	a := ElectricChargeDto{}

    // Parse JSON into ElectricChargeDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricChargeDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricChargeDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricCharge represents a measurement in a of ElectricCharge.
//
// Electric charge is the physical property of matter that causes it to experience a force when placed in an electromagnetic field.
type ElectricCharge struct {
	// value is the base measurement stored internally.
	value       float64
    
    coulombsLazy *float64 
    ampere_hoursLazy *float64 
    picocoulombsLazy *float64 
    nanocoulombsLazy *float64 
    microcoulombsLazy *float64 
    millicoulombsLazy *float64 
    kilocoulombsLazy *float64 
    megacoulombsLazy *float64 
    milliampere_hoursLazy *float64 
    kiloampere_hoursLazy *float64 
    megaampere_hoursLazy *float64 
}

// ElectricChargeFactory groups methods for creating ElectricCharge instances.
type ElectricChargeFactory struct{}

// CreateElectricCharge creates a new ElectricCharge instance from the given value and unit.
func (uf ElectricChargeFactory) CreateElectricCharge(value float64, unit ElectricChargeUnits) (*ElectricCharge, error) {
	return newElectricCharge(value, unit)
}

// FromDto converts a ElectricChargeDto to a ElectricCharge instance.
func (uf ElectricChargeFactory) FromDto(dto ElectricChargeDto) (*ElectricCharge, error) {
	return newElectricCharge(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricCharge instance.
func (uf ElectricChargeFactory) FromDtoJSON(data []byte) (*ElectricCharge, error) {
	unitDto, err := ElectricChargeDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricChargeDto from JSON: %w", err)
	}
	return ElectricChargeFactory{}.FromDto(*unitDto)
}


// FromCoulombs creates a new ElectricCharge instance from a value in Coulombs.
func (uf ElectricChargeFactory) FromCoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeCoulomb)
}

// FromAmpereHours creates a new ElectricCharge instance from a value in AmpereHours.
func (uf ElectricChargeFactory) FromAmpereHours(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeAmpereHour)
}

// FromPicocoulombs creates a new ElectricCharge instance from a value in Picocoulombs.
func (uf ElectricChargeFactory) FromPicocoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargePicocoulomb)
}

// FromNanocoulombs creates a new ElectricCharge instance from a value in Nanocoulombs.
func (uf ElectricChargeFactory) FromNanocoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeNanocoulomb)
}

// FromMicrocoulombs creates a new ElectricCharge instance from a value in Microcoulombs.
func (uf ElectricChargeFactory) FromMicrocoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeMicrocoulomb)
}

// FromMillicoulombs creates a new ElectricCharge instance from a value in Millicoulombs.
func (uf ElectricChargeFactory) FromMillicoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeMillicoulomb)
}

// FromKilocoulombs creates a new ElectricCharge instance from a value in Kilocoulombs.
func (uf ElectricChargeFactory) FromKilocoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeKilocoulomb)
}

// FromMegacoulombs creates a new ElectricCharge instance from a value in Megacoulombs.
func (uf ElectricChargeFactory) FromMegacoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeMegacoulomb)
}

// FromMilliampereHours creates a new ElectricCharge instance from a value in MilliampereHours.
func (uf ElectricChargeFactory) FromMilliampereHours(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeMilliampereHour)
}

// FromKiloampereHours creates a new ElectricCharge instance from a value in KiloampereHours.
func (uf ElectricChargeFactory) FromKiloampereHours(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeKiloampereHour)
}

// FromMegaampereHours creates a new ElectricCharge instance from a value in MegaampereHours.
func (uf ElectricChargeFactory) FromMegaampereHours(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeMegaampereHour)
}


// newElectricCharge creates a new ElectricCharge.
func newElectricCharge(value float64, fromUnit ElectricChargeUnits) (*ElectricCharge, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricCharge{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricCharge in Coulomb unit (the base/default quantity).
func (a *ElectricCharge) BaseValue() float64 {
	return a.value
}


// Coulombs returns the ElectricCharge value in Coulombs.
//
// 
func (a *ElectricCharge) Coulombs() float64 {
	if a.coulombsLazy != nil {
		return *a.coulombsLazy
	}
	coulombs := a.convertFromBase(ElectricChargeCoulomb)
	a.coulombsLazy = &coulombs
	return coulombs
}

// AmpereHours returns the ElectricCharge value in AmpereHours.
//
// 
func (a *ElectricCharge) AmpereHours() float64 {
	if a.ampere_hoursLazy != nil {
		return *a.ampere_hoursLazy
	}
	ampere_hours := a.convertFromBase(ElectricChargeAmpereHour)
	a.ampere_hoursLazy = &ampere_hours
	return ampere_hours
}

// Picocoulombs returns the ElectricCharge value in Picocoulombs.
//
// 
func (a *ElectricCharge) Picocoulombs() float64 {
	if a.picocoulombsLazy != nil {
		return *a.picocoulombsLazy
	}
	picocoulombs := a.convertFromBase(ElectricChargePicocoulomb)
	a.picocoulombsLazy = &picocoulombs
	return picocoulombs
}

// Nanocoulombs returns the ElectricCharge value in Nanocoulombs.
//
// 
func (a *ElectricCharge) Nanocoulombs() float64 {
	if a.nanocoulombsLazy != nil {
		return *a.nanocoulombsLazy
	}
	nanocoulombs := a.convertFromBase(ElectricChargeNanocoulomb)
	a.nanocoulombsLazy = &nanocoulombs
	return nanocoulombs
}

// Microcoulombs returns the ElectricCharge value in Microcoulombs.
//
// 
func (a *ElectricCharge) Microcoulombs() float64 {
	if a.microcoulombsLazy != nil {
		return *a.microcoulombsLazy
	}
	microcoulombs := a.convertFromBase(ElectricChargeMicrocoulomb)
	a.microcoulombsLazy = &microcoulombs
	return microcoulombs
}

// Millicoulombs returns the ElectricCharge value in Millicoulombs.
//
// 
func (a *ElectricCharge) Millicoulombs() float64 {
	if a.millicoulombsLazy != nil {
		return *a.millicoulombsLazy
	}
	millicoulombs := a.convertFromBase(ElectricChargeMillicoulomb)
	a.millicoulombsLazy = &millicoulombs
	return millicoulombs
}

// Kilocoulombs returns the ElectricCharge value in Kilocoulombs.
//
// 
func (a *ElectricCharge) Kilocoulombs() float64 {
	if a.kilocoulombsLazy != nil {
		return *a.kilocoulombsLazy
	}
	kilocoulombs := a.convertFromBase(ElectricChargeKilocoulomb)
	a.kilocoulombsLazy = &kilocoulombs
	return kilocoulombs
}

// Megacoulombs returns the ElectricCharge value in Megacoulombs.
//
// 
func (a *ElectricCharge) Megacoulombs() float64 {
	if a.megacoulombsLazy != nil {
		return *a.megacoulombsLazy
	}
	megacoulombs := a.convertFromBase(ElectricChargeMegacoulomb)
	a.megacoulombsLazy = &megacoulombs
	return megacoulombs
}

// MilliampereHours returns the ElectricCharge value in MilliampereHours.
//
// 
func (a *ElectricCharge) MilliampereHours() float64 {
	if a.milliampere_hoursLazy != nil {
		return *a.milliampere_hoursLazy
	}
	milliampere_hours := a.convertFromBase(ElectricChargeMilliampereHour)
	a.milliampere_hoursLazy = &milliampere_hours
	return milliampere_hours
}

// KiloampereHours returns the ElectricCharge value in KiloampereHours.
//
// 
func (a *ElectricCharge) KiloampereHours() float64 {
	if a.kiloampere_hoursLazy != nil {
		return *a.kiloampere_hoursLazy
	}
	kiloampere_hours := a.convertFromBase(ElectricChargeKiloampereHour)
	a.kiloampere_hoursLazy = &kiloampere_hours
	return kiloampere_hours
}

// MegaampereHours returns the ElectricCharge value in MegaampereHours.
//
// 
func (a *ElectricCharge) MegaampereHours() float64 {
	if a.megaampere_hoursLazy != nil {
		return *a.megaampere_hoursLazy
	}
	megaampere_hours := a.convertFromBase(ElectricChargeMegaampereHour)
	a.megaampere_hoursLazy = &megaampere_hours
	return megaampere_hours
}


// ToDto creates a ElectricChargeDto representation from the ElectricCharge instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Coulomb by default.
func (a *ElectricCharge) ToDto(holdInUnit *ElectricChargeUnits) ElectricChargeDto {
	if holdInUnit == nil {
		defaultUnit := ElectricChargeCoulomb // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricChargeDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricCharge instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Coulomb by default.
func (a *ElectricCharge) ToDtoJSON(holdInUnit *ElectricChargeUnits) ([]byte, error) {
	// Convert to ElectricChargeDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricCharge to a specific unit value.
// The function uses the provided unit type (ElectricChargeUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricCharge) Convert(toUnit ElectricChargeUnits) float64 {
	switch toUnit { 
    case ElectricChargeCoulomb:
		return a.Coulombs()
    case ElectricChargeAmpereHour:
		return a.AmpereHours()
    case ElectricChargePicocoulomb:
		return a.Picocoulombs()
    case ElectricChargeNanocoulomb:
		return a.Nanocoulombs()
    case ElectricChargeMicrocoulomb:
		return a.Microcoulombs()
    case ElectricChargeMillicoulomb:
		return a.Millicoulombs()
    case ElectricChargeKilocoulomb:
		return a.Kilocoulombs()
    case ElectricChargeMegacoulomb:
		return a.Megacoulombs()
    case ElectricChargeMilliampereHour:
		return a.MilliampereHours()
    case ElectricChargeKiloampereHour:
		return a.KiloampereHours()
    case ElectricChargeMegaampereHour:
		return a.MegaampereHours()
	default:
		return math.NaN()
	}
}

func (a *ElectricCharge) convertFromBase(toUnit ElectricChargeUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricChargeCoulomb:
		return (value) 
	case ElectricChargeAmpereHour:
		return (value * 2.77777777777e-4) 
	case ElectricChargePicocoulomb:
		return ((value) / 1e-12) 
	case ElectricChargeNanocoulomb:
		return ((value) / 1e-09) 
	case ElectricChargeMicrocoulomb:
		return ((value) / 1e-06) 
	case ElectricChargeMillicoulomb:
		return ((value) / 0.001) 
	case ElectricChargeKilocoulomb:
		return ((value) / 1000.0) 
	case ElectricChargeMegacoulomb:
		return ((value) / 1000000.0) 
	case ElectricChargeMilliampereHour:
		return ((value * 2.77777777777e-4) / 0.001) 
	case ElectricChargeKiloampereHour:
		return ((value * 2.77777777777e-4) / 1000.0) 
	case ElectricChargeMegaampereHour:
		return ((value * 2.77777777777e-4) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricCharge) convertToBase(value float64, fromUnit ElectricChargeUnits) float64 {
	switch fromUnit { 
	case ElectricChargeCoulomb:
		return (value) 
	case ElectricChargeAmpereHour:
		return (value / 2.77777777777e-4) 
	case ElectricChargePicocoulomb:
		return ((value) * 1e-12) 
	case ElectricChargeNanocoulomb:
		return ((value) * 1e-09) 
	case ElectricChargeMicrocoulomb:
		return ((value) * 1e-06) 
	case ElectricChargeMillicoulomb:
		return ((value) * 0.001) 
	case ElectricChargeKilocoulomb:
		return ((value) * 1000.0) 
	case ElectricChargeMegacoulomb:
		return ((value) * 1000000.0) 
	case ElectricChargeMilliampereHour:
		return ((value / 2.77777777777e-4) * 0.001) 
	case ElectricChargeKiloampereHour:
		return ((value / 2.77777777777e-4) * 1000.0) 
	case ElectricChargeMegaampereHour:
		return ((value / 2.77777777777e-4) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricCharge in the default unit (Coulomb),
// formatted to two decimal places.
func (a ElectricCharge) String() string {
	return a.ToString(ElectricChargeCoulomb, 2)
}

// ToString formats the ElectricCharge value as a string with the specified unit and fractional digits.
// It converts the ElectricCharge to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricCharge value will be converted (e.g., Coulomb))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricCharge with the unit abbreviation.
func (a *ElectricCharge) ToString(unit ElectricChargeUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricCharge) getUnitAbbreviation(unit ElectricChargeUnits) string {
	switch unit { 
	case ElectricChargeCoulomb:
		return "C" 
	case ElectricChargeAmpereHour:
		return "A-h" 
	case ElectricChargePicocoulomb:
		return "pC" 
	case ElectricChargeNanocoulomb:
		return "nC" 
	case ElectricChargeMicrocoulomb:
		return "μC" 
	case ElectricChargeMillicoulomb:
		return "mC" 
	case ElectricChargeKilocoulomb:
		return "kC" 
	case ElectricChargeMegacoulomb:
		return "MC" 
	case ElectricChargeMilliampereHour:
		return "mA-h" 
	case ElectricChargeKiloampereHour:
		return "kA-h" 
	case ElectricChargeMegaampereHour:
		return "MA-h" 
	default:
		return ""
	}
}

// Equals checks if the given ElectricCharge is equal to the current ElectricCharge.
//
// Parameters:
//    other: The ElectricCharge to compare against.
//
// Returns:
//    bool: Returns true if both ElectricCharge are equal, false otherwise.
func (a *ElectricCharge) Equals(other *ElectricCharge) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricCharge with another ElectricCharge.
// It returns -1 if the current ElectricCharge is less than the other ElectricCharge, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricCharge to compare against.
//
// Returns:
//    int: -1 if the current ElectricCharge is less, 1 if greater, and 0 if equal.
func (a *ElectricCharge) CompareTo(other *ElectricCharge) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricCharge to the current ElectricCharge and returns the result.
// The result is a new ElectricCharge instance with the sum of the values.
//
// Parameters:
//    other: The ElectricCharge to add to the current ElectricCharge.
//
// Returns:
//    *ElectricCharge: A new ElectricCharge instance representing the sum of both ElectricCharge.
func (a *ElectricCharge) Add(other *ElectricCharge) *ElectricCharge {
	return &ElectricCharge{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricCharge from the current ElectricCharge and returns the result.
// The result is a new ElectricCharge instance with the difference of the values.
//
// Parameters:
//    other: The ElectricCharge to subtract from the current ElectricCharge.
//
// Returns:
//    *ElectricCharge: A new ElectricCharge instance representing the difference of both ElectricCharge.
func (a *ElectricCharge) Subtract(other *ElectricCharge) *ElectricCharge {
	return &ElectricCharge{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricCharge by the given ElectricCharge and returns the result.
// The result is a new ElectricCharge instance with the product of the values.
//
// Parameters:
//    other: The ElectricCharge to multiply with the current ElectricCharge.
//
// Returns:
//    *ElectricCharge: A new ElectricCharge instance representing the product of both ElectricCharge.
func (a *ElectricCharge) Multiply(other *ElectricCharge) *ElectricCharge {
	return &ElectricCharge{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricCharge by the given ElectricCharge and returns the result.
// The result is a new ElectricCharge instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricCharge to divide the current ElectricCharge by.
//
// Returns:
//    *ElectricCharge: A new ElectricCharge instance representing the quotient of both ElectricCharge.
func (a *ElectricCharge) Divide(other *ElectricCharge) *ElectricCharge {
	return &ElectricCharge{value: a.value / other.BaseValue()}
}