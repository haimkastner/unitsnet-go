package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricSurfaceChargeDensityUnits enumeration
type ElectricSurfaceChargeDensityUnits string

const (
    
        // 
        ElectricSurfaceChargeDensityCoulombPerSquareMeter ElectricSurfaceChargeDensityUnits = "CoulombPerSquareMeter"
        // 
        ElectricSurfaceChargeDensityCoulombPerSquareCentimeter ElectricSurfaceChargeDensityUnits = "CoulombPerSquareCentimeter"
        // 
        ElectricSurfaceChargeDensityCoulombPerSquareInch ElectricSurfaceChargeDensityUnits = "CoulombPerSquareInch"
)

// ElectricSurfaceChargeDensityDto represents an ElectricSurfaceChargeDensity
type ElectricSurfaceChargeDensityDto struct {
	Value float64
	Unit  ElectricSurfaceChargeDensityUnits
}

// ElectricSurfaceChargeDensityDtoFactory struct to group related functions
type ElectricSurfaceChargeDensityDtoFactory struct{}

func (udf ElectricSurfaceChargeDensityDtoFactory) FromJSON(data []byte) (*ElectricSurfaceChargeDensityDto, error) {
	a := ElectricSurfaceChargeDensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricSurfaceChargeDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricSurfaceChargeDensity struct
type ElectricSurfaceChargeDensity struct {
	value       float64
    
    coulombs_per_square_meterLazy *float64 
    coulombs_per_square_centimeterLazy *float64 
    coulombs_per_square_inchLazy *float64 
}

// ElectricSurfaceChargeDensityFactory struct to group related functions
type ElectricSurfaceChargeDensityFactory struct{}

func (uf ElectricSurfaceChargeDensityFactory) CreateElectricSurfaceChargeDensity(value float64, unit ElectricSurfaceChargeDensityUnits) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(value, unit)
}

func (uf ElectricSurfaceChargeDensityFactory) FromDto(dto ElectricSurfaceChargeDensityDto) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(dto.Value, dto.Unit)
}

func (uf ElectricSurfaceChargeDensityFactory) FromDtoJSON(data []byte) (*ElectricSurfaceChargeDensity, error) {
	unitDto, err := ElectricSurfaceChargeDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricSurfaceChargeDensityFactory{}.FromDto(*unitDto)
}


// FromCoulombPerSquareMeter creates a new ElectricSurfaceChargeDensity instance from CoulombPerSquareMeter.
func (uf ElectricSurfaceChargeDensityFactory) FromCoulombsPerSquareMeter(value float64) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(value, ElectricSurfaceChargeDensityCoulombPerSquareMeter)
}

// FromCoulombPerSquareCentimeter creates a new ElectricSurfaceChargeDensity instance from CoulombPerSquareCentimeter.
func (uf ElectricSurfaceChargeDensityFactory) FromCoulombsPerSquareCentimeter(value float64) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(value, ElectricSurfaceChargeDensityCoulombPerSquareCentimeter)
}

// FromCoulombPerSquareInch creates a new ElectricSurfaceChargeDensity instance from CoulombPerSquareInch.
func (uf ElectricSurfaceChargeDensityFactory) FromCoulombsPerSquareInch(value float64) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(value, ElectricSurfaceChargeDensityCoulombPerSquareInch)
}




// newElectricSurfaceChargeDensity creates a new ElectricSurfaceChargeDensity.
func newElectricSurfaceChargeDensity(value float64, fromUnit ElectricSurfaceChargeDensityUnits) (*ElectricSurfaceChargeDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricSurfaceChargeDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricSurfaceChargeDensity in CoulombPerSquareMeter.
func (a *ElectricSurfaceChargeDensity) BaseValue() float64 {
	return a.value
}


// CoulombPerSquareMeter returns the value in CoulombPerSquareMeter.
func (a *ElectricSurfaceChargeDensity) CoulombsPerSquareMeter() float64 {
	if a.coulombs_per_square_meterLazy != nil {
		return *a.coulombs_per_square_meterLazy
	}
	coulombs_per_square_meter := a.convertFromBase(ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	a.coulombs_per_square_meterLazy = &coulombs_per_square_meter
	return coulombs_per_square_meter
}

// CoulombPerSquareCentimeter returns the value in CoulombPerSquareCentimeter.
func (a *ElectricSurfaceChargeDensity) CoulombsPerSquareCentimeter() float64 {
	if a.coulombs_per_square_centimeterLazy != nil {
		return *a.coulombs_per_square_centimeterLazy
	}
	coulombs_per_square_centimeter := a.convertFromBase(ElectricSurfaceChargeDensityCoulombPerSquareCentimeter)
	a.coulombs_per_square_centimeterLazy = &coulombs_per_square_centimeter
	return coulombs_per_square_centimeter
}

// CoulombPerSquareInch returns the value in CoulombPerSquareInch.
func (a *ElectricSurfaceChargeDensity) CoulombsPerSquareInch() float64 {
	if a.coulombs_per_square_inchLazy != nil {
		return *a.coulombs_per_square_inchLazy
	}
	coulombs_per_square_inch := a.convertFromBase(ElectricSurfaceChargeDensityCoulombPerSquareInch)
	a.coulombs_per_square_inchLazy = &coulombs_per_square_inch
	return coulombs_per_square_inch
}


// ToDto creates an ElectricSurfaceChargeDensityDto representation.
func (a *ElectricSurfaceChargeDensity) ToDto(holdInUnit *ElectricSurfaceChargeDensityUnits) ElectricSurfaceChargeDensityDto {
	if holdInUnit == nil {
		defaultUnit := ElectricSurfaceChargeDensityCoulombPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricSurfaceChargeDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricSurfaceChargeDensityDto representation.
func (a *ElectricSurfaceChargeDensity) ToDtoJSON(holdInUnit *ElectricSurfaceChargeDensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricSurfaceChargeDensity to a specific unit value.
func (a *ElectricSurfaceChargeDensity) Convert(toUnit ElectricSurfaceChargeDensityUnits) float64 {
	switch toUnit { 
    case ElectricSurfaceChargeDensityCoulombPerSquareMeter:
		return a.CoulombsPerSquareMeter()
    case ElectricSurfaceChargeDensityCoulombPerSquareCentimeter:
		return a.CoulombsPerSquareCentimeter()
    case ElectricSurfaceChargeDensityCoulombPerSquareInch:
		return a.CoulombsPerSquareInch()
	default:
		return 0
	}
}

func (a *ElectricSurfaceChargeDensity) convertFromBase(toUnit ElectricSurfaceChargeDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricSurfaceChargeDensityCoulombPerSquareMeter:
		return (value) 
	case ElectricSurfaceChargeDensityCoulombPerSquareCentimeter:
		return (value / 1.0e4) 
	case ElectricSurfaceChargeDensityCoulombPerSquareInch:
		return (value / 1.5500031000062000e3) 
	default:
		return math.NaN()
	}
}

func (a *ElectricSurfaceChargeDensity) convertToBase(value float64, fromUnit ElectricSurfaceChargeDensityUnits) float64 {
	switch fromUnit { 
	case ElectricSurfaceChargeDensityCoulombPerSquareMeter:
		return (value) 
	case ElectricSurfaceChargeDensityCoulombPerSquareCentimeter:
		return (value * 1.0e4) 
	case ElectricSurfaceChargeDensityCoulombPerSquareInch:
		return (value * 1.5500031000062000e3) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricSurfaceChargeDensity) String() string {
	return a.ToString(ElectricSurfaceChargeDensityCoulombPerSquareMeter, 2)
}

// ToString formats the ElectricSurfaceChargeDensity to string.
// fractionalDigits -1 for not mention
func (a *ElectricSurfaceChargeDensity) ToString(unit ElectricSurfaceChargeDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricSurfaceChargeDensity) getUnitAbbreviation(unit ElectricSurfaceChargeDensityUnits) string {
	switch unit { 
	case ElectricSurfaceChargeDensityCoulombPerSquareMeter:
		return "C/m²" 
	case ElectricSurfaceChargeDensityCoulombPerSquareCentimeter:
		return "C/cm²" 
	case ElectricSurfaceChargeDensityCoulombPerSquareInch:
		return "C/in²" 
	default:
		return ""
	}
}

// Check if the given ElectricSurfaceChargeDensity are equals to the current ElectricSurfaceChargeDensity
func (a *ElectricSurfaceChargeDensity) Equals(other *ElectricSurfaceChargeDensity) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricSurfaceChargeDensity are equals to the current ElectricSurfaceChargeDensity
func (a *ElectricSurfaceChargeDensity) CompareTo(other *ElectricSurfaceChargeDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricSurfaceChargeDensity to the current ElectricSurfaceChargeDensity.
func (a *ElectricSurfaceChargeDensity) Add(other *ElectricSurfaceChargeDensity) *ElectricSurfaceChargeDensity {
	return &ElectricSurfaceChargeDensity{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricSurfaceChargeDensity to the current ElectricSurfaceChargeDensity.
func (a *ElectricSurfaceChargeDensity) Subtract(other *ElectricSurfaceChargeDensity) *ElectricSurfaceChargeDensity {
	return &ElectricSurfaceChargeDensity{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricSurfaceChargeDensity to the current ElectricSurfaceChargeDensity.
func (a *ElectricSurfaceChargeDensity) Multiply(other *ElectricSurfaceChargeDensity) *ElectricSurfaceChargeDensity {
	return &ElectricSurfaceChargeDensity{value: a.value * other.BaseValue()}
}

// Divide the given ElectricSurfaceChargeDensity to the current ElectricSurfaceChargeDensity.
func (a *ElectricSurfaceChargeDensity) Divide(other *ElectricSurfaceChargeDensity) *ElectricSurfaceChargeDensity {
	return &ElectricSurfaceChargeDensity{value: a.value / other.BaseValue()}
}