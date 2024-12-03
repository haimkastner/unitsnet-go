package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MagneticFieldUnits enumeration
type MagneticFieldUnits string

const (
    
        // 
        MagneticFieldTesla MagneticFieldUnits = "Tesla"
        // 
        MagneticFieldGauss MagneticFieldUnits = "Gauss"
        // 
        MagneticFieldNanotesla MagneticFieldUnits = "Nanotesla"
        // 
        MagneticFieldMicrotesla MagneticFieldUnits = "Microtesla"
        // 
        MagneticFieldMillitesla MagneticFieldUnits = "Millitesla"
        // 
        MagneticFieldMilligauss MagneticFieldUnits = "Milligauss"
)

// MagneticFieldDto represents an MagneticField
type MagneticFieldDto struct {
	Value float64
	Unit  MagneticFieldUnits
}

// MagneticFieldDtoFactory struct to group related functions
type MagneticFieldDtoFactory struct{}

func (udf MagneticFieldDtoFactory) FromJSON(data []byte) (*MagneticFieldDto, error) {
	a := MagneticFieldDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MagneticFieldDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MagneticField struct
type MagneticField struct {
	value       float64
    
    teslasLazy *float64 
    gaussesLazy *float64 
    nanoteslasLazy *float64 
    microteslasLazy *float64 
    milliteslasLazy *float64 
    milligaussesLazy *float64 
}

// MagneticFieldFactory struct to group related functions
type MagneticFieldFactory struct{}

func (uf MagneticFieldFactory) CreateMagneticField(value float64, unit MagneticFieldUnits) (*MagneticField, error) {
	return newMagneticField(value, unit)
}

func (uf MagneticFieldFactory) FromDto(dto MagneticFieldDto) (*MagneticField, error) {
	return newMagneticField(dto.Value, dto.Unit)
}

func (uf MagneticFieldFactory) FromDtoJSON(data []byte) (*MagneticField, error) {
	unitDto, err := MagneticFieldDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MagneticFieldFactory{}.FromDto(*unitDto)
}


// FromTesla creates a new MagneticField instance from Tesla.
func (uf MagneticFieldFactory) FromTeslas(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldTesla)
}

// FromGauss creates a new MagneticField instance from Gauss.
func (uf MagneticFieldFactory) FromGausses(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldGauss)
}

// FromNanotesla creates a new MagneticField instance from Nanotesla.
func (uf MagneticFieldFactory) FromNanoteslas(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldNanotesla)
}

// FromMicrotesla creates a new MagneticField instance from Microtesla.
func (uf MagneticFieldFactory) FromMicroteslas(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldMicrotesla)
}

// FromMillitesla creates a new MagneticField instance from Millitesla.
func (uf MagneticFieldFactory) FromMilliteslas(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldMillitesla)
}

// FromMilligauss creates a new MagneticField instance from Milligauss.
func (uf MagneticFieldFactory) FromMilligausses(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldMilligauss)
}




// newMagneticField creates a new MagneticField.
func newMagneticField(value float64, fromUnit MagneticFieldUnits) (*MagneticField, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MagneticField{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MagneticField in Tesla.
func (a *MagneticField) BaseValue() float64 {
	return a.value
}


// Tesla returns the value in Tesla.
func (a *MagneticField) Teslas() float64 {
	if a.teslasLazy != nil {
		return *a.teslasLazy
	}
	teslas := a.convertFromBase(MagneticFieldTesla)
	a.teslasLazy = &teslas
	return teslas
}

// Gauss returns the value in Gauss.
func (a *MagneticField) Gausses() float64 {
	if a.gaussesLazy != nil {
		return *a.gaussesLazy
	}
	gausses := a.convertFromBase(MagneticFieldGauss)
	a.gaussesLazy = &gausses
	return gausses
}

// Nanotesla returns the value in Nanotesla.
func (a *MagneticField) Nanoteslas() float64 {
	if a.nanoteslasLazy != nil {
		return *a.nanoteslasLazy
	}
	nanoteslas := a.convertFromBase(MagneticFieldNanotesla)
	a.nanoteslasLazy = &nanoteslas
	return nanoteslas
}

// Microtesla returns the value in Microtesla.
func (a *MagneticField) Microteslas() float64 {
	if a.microteslasLazy != nil {
		return *a.microteslasLazy
	}
	microteslas := a.convertFromBase(MagneticFieldMicrotesla)
	a.microteslasLazy = &microteslas
	return microteslas
}

// Millitesla returns the value in Millitesla.
func (a *MagneticField) Milliteslas() float64 {
	if a.milliteslasLazy != nil {
		return *a.milliteslasLazy
	}
	milliteslas := a.convertFromBase(MagneticFieldMillitesla)
	a.milliteslasLazy = &milliteslas
	return milliteslas
}

// Milligauss returns the value in Milligauss.
func (a *MagneticField) Milligausses() float64 {
	if a.milligaussesLazy != nil {
		return *a.milligaussesLazy
	}
	milligausses := a.convertFromBase(MagneticFieldMilligauss)
	a.milligaussesLazy = &milligausses
	return milligausses
}


// ToDto creates an MagneticFieldDto representation.
func (a *MagneticField) ToDto(holdInUnit *MagneticFieldUnits) MagneticFieldDto {
	if holdInUnit == nil {
		defaultUnit := MagneticFieldTesla // Default value
		holdInUnit = &defaultUnit
	}

	return MagneticFieldDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an MagneticFieldDto representation.
func (a *MagneticField) ToDtoJSON(holdInUnit *MagneticFieldUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MagneticField to a specific unit value.
func (a *MagneticField) Convert(toUnit MagneticFieldUnits) float64 {
	switch toUnit { 
    case MagneticFieldTesla:
		return a.Teslas()
    case MagneticFieldGauss:
		return a.Gausses()
    case MagneticFieldNanotesla:
		return a.Nanoteslas()
    case MagneticFieldMicrotesla:
		return a.Microteslas()
    case MagneticFieldMillitesla:
		return a.Milliteslas()
    case MagneticFieldMilligauss:
		return a.Milligausses()
	default:
		return 0
	}
}

func (a *MagneticField) convertFromBase(toUnit MagneticFieldUnits) float64 {
    value := a.value
	switch toUnit { 
	case MagneticFieldTesla:
		return (value) 
	case MagneticFieldGauss:
		return (value * 1e4) 
	case MagneticFieldNanotesla:
		return ((value) / 1e-09) 
	case MagneticFieldMicrotesla:
		return ((value) / 1e-06) 
	case MagneticFieldMillitesla:
		return ((value) / 0.001) 
	case MagneticFieldMilligauss:
		return ((value * 1e4) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *MagneticField) convertToBase(value float64, fromUnit MagneticFieldUnits) float64 {
	switch fromUnit { 
	case MagneticFieldTesla:
		return (value) 
	case MagneticFieldGauss:
		return (value / 1e4) 
	case MagneticFieldNanotesla:
		return ((value) * 1e-09) 
	case MagneticFieldMicrotesla:
		return ((value) * 1e-06) 
	case MagneticFieldMillitesla:
		return ((value) * 0.001) 
	case MagneticFieldMilligauss:
		return ((value / 1e4) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a MagneticField) String() string {
	return a.ToString(MagneticFieldTesla, 2)
}

// ToString formats the MagneticField to string.
// fractionalDigits -1 for not mention
func (a *MagneticField) ToString(unit MagneticFieldUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MagneticField) getUnitAbbreviation(unit MagneticFieldUnits) string {
	switch unit { 
	case MagneticFieldTesla:
		return "T" 
	case MagneticFieldGauss:
		return "G" 
	case MagneticFieldNanotesla:
		return "nT" 
	case MagneticFieldMicrotesla:
		return "μT" 
	case MagneticFieldMillitesla:
		return "mT" 
	case MagneticFieldMilligauss:
		return "mG" 
	default:
		return ""
	}
}

// Check if the given MagneticField are equals to the current MagneticField
func (a *MagneticField) Equals(other *MagneticField) bool {
	return a.value == other.BaseValue()
}

// Check if the given MagneticField are equals to the current MagneticField
func (a *MagneticField) CompareTo(other *MagneticField) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given MagneticField to the current MagneticField.
func (a *MagneticField) Add(other *MagneticField) *MagneticField {
	return &MagneticField{value: a.value + other.BaseValue()}
}

// Subtract the given MagneticField to the current MagneticField.
func (a *MagneticField) Subtract(other *MagneticField) *MagneticField {
	return &MagneticField{value: a.value - other.BaseValue()}
}

// Multiply the given MagneticField to the current MagneticField.
func (a *MagneticField) Multiply(other *MagneticField) *MagneticField {
	return &MagneticField{value: a.value * other.BaseValue()}
}

// Divide the given MagneticField to the current MagneticField.
func (a *MagneticField) Divide(other *MagneticField) *MagneticField {
	return &MagneticField{value: a.value / other.BaseValue()}
}