package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LinearDensityUnits enumeration
type LinearDensityUnits string

const (
    
        // 
        LinearDensityGramPerMillimeter LinearDensityUnits = "GramPerMillimeter"
        // 
        LinearDensityGramPerCentimeter LinearDensityUnits = "GramPerCentimeter"
        // 
        LinearDensityGramPerMeter LinearDensityUnits = "GramPerMeter"
        // 
        LinearDensityPoundPerInch LinearDensityUnits = "PoundPerInch"
        // 
        LinearDensityPoundPerFoot LinearDensityUnits = "PoundPerFoot"
        // 
        LinearDensityGramPerFoot LinearDensityUnits = "GramPerFoot"
        // 
        LinearDensityMicrogramPerMillimeter LinearDensityUnits = "MicrogramPerMillimeter"
        // 
        LinearDensityMilligramPerMillimeter LinearDensityUnits = "MilligramPerMillimeter"
        // 
        LinearDensityKilogramPerMillimeter LinearDensityUnits = "KilogramPerMillimeter"
        // 
        LinearDensityMicrogramPerCentimeter LinearDensityUnits = "MicrogramPerCentimeter"
        // 
        LinearDensityMilligramPerCentimeter LinearDensityUnits = "MilligramPerCentimeter"
        // 
        LinearDensityKilogramPerCentimeter LinearDensityUnits = "KilogramPerCentimeter"
        // 
        LinearDensityMicrogramPerMeter LinearDensityUnits = "MicrogramPerMeter"
        // 
        LinearDensityMilligramPerMeter LinearDensityUnits = "MilligramPerMeter"
        // 
        LinearDensityKilogramPerMeter LinearDensityUnits = "KilogramPerMeter"
        // 
        LinearDensityMicrogramPerFoot LinearDensityUnits = "MicrogramPerFoot"
        // 
        LinearDensityMilligramPerFoot LinearDensityUnits = "MilligramPerFoot"
        // 
        LinearDensityKilogramPerFoot LinearDensityUnits = "KilogramPerFoot"
)

// LinearDensityDto represents an LinearDensity
type LinearDensityDto struct {
	Value float64
	Unit  LinearDensityUnits
}

// LinearDensityDtoFactory struct to group related functions
type LinearDensityDtoFactory struct{}

func (udf LinearDensityDtoFactory) FromJSON(data []byte) (*LinearDensityDto, error) {
	a := LinearDensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a LinearDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// LinearDensity struct
type LinearDensity struct {
	value       float64
    
    grams_per_millimeterLazy *float64 
    grams_per_centimeterLazy *float64 
    grams_per_meterLazy *float64 
    pounds_per_inchLazy *float64 
    pounds_per_footLazy *float64 
    grams_per_footLazy *float64 
    micrograms_per_millimeterLazy *float64 
    milligrams_per_millimeterLazy *float64 
    kilograms_per_millimeterLazy *float64 
    micrograms_per_centimeterLazy *float64 
    milligrams_per_centimeterLazy *float64 
    kilograms_per_centimeterLazy *float64 
    micrograms_per_meterLazy *float64 
    milligrams_per_meterLazy *float64 
    kilograms_per_meterLazy *float64 
    micrograms_per_footLazy *float64 
    milligrams_per_footLazy *float64 
    kilograms_per_footLazy *float64 
}

// LinearDensityFactory struct to group related functions
type LinearDensityFactory struct{}

func (uf LinearDensityFactory) CreateLinearDensity(value float64, unit LinearDensityUnits) (*LinearDensity, error) {
	return newLinearDensity(value, unit)
}

func (uf LinearDensityFactory) FromDto(dto LinearDensityDto) (*LinearDensity, error) {
	return newLinearDensity(dto.Value, dto.Unit)
}

func (uf LinearDensityFactory) FromDtoJSON(data []byte) (*LinearDensity, error) {
	unitDto, err := LinearDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return LinearDensityFactory{}.FromDto(*unitDto)
}


// FromGramPerMillimeter creates a new LinearDensity instance from GramPerMillimeter.
func (uf LinearDensityFactory) FromGramsPerMillimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityGramPerMillimeter)
}

// FromGramPerCentimeter creates a new LinearDensity instance from GramPerCentimeter.
func (uf LinearDensityFactory) FromGramsPerCentimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityGramPerCentimeter)
}

// FromGramPerMeter creates a new LinearDensity instance from GramPerMeter.
func (uf LinearDensityFactory) FromGramsPerMeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityGramPerMeter)
}

// FromPoundPerInch creates a new LinearDensity instance from PoundPerInch.
func (uf LinearDensityFactory) FromPoundsPerInch(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityPoundPerInch)
}

// FromPoundPerFoot creates a new LinearDensity instance from PoundPerFoot.
func (uf LinearDensityFactory) FromPoundsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityPoundPerFoot)
}

// FromGramPerFoot creates a new LinearDensity instance from GramPerFoot.
func (uf LinearDensityFactory) FromGramsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityGramPerFoot)
}

// FromMicrogramPerMillimeter creates a new LinearDensity instance from MicrogramPerMillimeter.
func (uf LinearDensityFactory) FromMicrogramsPerMillimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMicrogramPerMillimeter)
}

// FromMilligramPerMillimeter creates a new LinearDensity instance from MilligramPerMillimeter.
func (uf LinearDensityFactory) FromMilligramsPerMillimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMilligramPerMillimeter)
}

// FromKilogramPerMillimeter creates a new LinearDensity instance from KilogramPerMillimeter.
func (uf LinearDensityFactory) FromKilogramsPerMillimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityKilogramPerMillimeter)
}

// FromMicrogramPerCentimeter creates a new LinearDensity instance from MicrogramPerCentimeter.
func (uf LinearDensityFactory) FromMicrogramsPerCentimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMicrogramPerCentimeter)
}

// FromMilligramPerCentimeter creates a new LinearDensity instance from MilligramPerCentimeter.
func (uf LinearDensityFactory) FromMilligramsPerCentimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMilligramPerCentimeter)
}

// FromKilogramPerCentimeter creates a new LinearDensity instance from KilogramPerCentimeter.
func (uf LinearDensityFactory) FromKilogramsPerCentimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityKilogramPerCentimeter)
}

// FromMicrogramPerMeter creates a new LinearDensity instance from MicrogramPerMeter.
func (uf LinearDensityFactory) FromMicrogramsPerMeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMicrogramPerMeter)
}

// FromMilligramPerMeter creates a new LinearDensity instance from MilligramPerMeter.
func (uf LinearDensityFactory) FromMilligramsPerMeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMilligramPerMeter)
}

// FromKilogramPerMeter creates a new LinearDensity instance from KilogramPerMeter.
func (uf LinearDensityFactory) FromKilogramsPerMeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityKilogramPerMeter)
}

// FromMicrogramPerFoot creates a new LinearDensity instance from MicrogramPerFoot.
func (uf LinearDensityFactory) FromMicrogramsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMicrogramPerFoot)
}

// FromMilligramPerFoot creates a new LinearDensity instance from MilligramPerFoot.
func (uf LinearDensityFactory) FromMilligramsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMilligramPerFoot)
}

// FromKilogramPerFoot creates a new LinearDensity instance from KilogramPerFoot.
func (uf LinearDensityFactory) FromKilogramsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityKilogramPerFoot)
}




// newLinearDensity creates a new LinearDensity.
func newLinearDensity(value float64, fromUnit LinearDensityUnits) (*LinearDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &LinearDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of LinearDensity in KilogramPerMeter.
func (a *LinearDensity) BaseValue() float64 {
	return a.value
}


// GramPerMillimeter returns the value in GramPerMillimeter.
func (a *LinearDensity) GramsPerMillimeter() float64 {
	if a.grams_per_millimeterLazy != nil {
		return *a.grams_per_millimeterLazy
	}
	grams_per_millimeter := a.convertFromBase(LinearDensityGramPerMillimeter)
	a.grams_per_millimeterLazy = &grams_per_millimeter
	return grams_per_millimeter
}

// GramPerCentimeter returns the value in GramPerCentimeter.
func (a *LinearDensity) GramsPerCentimeter() float64 {
	if a.grams_per_centimeterLazy != nil {
		return *a.grams_per_centimeterLazy
	}
	grams_per_centimeter := a.convertFromBase(LinearDensityGramPerCentimeter)
	a.grams_per_centimeterLazy = &grams_per_centimeter
	return grams_per_centimeter
}

// GramPerMeter returns the value in GramPerMeter.
func (a *LinearDensity) GramsPerMeter() float64 {
	if a.grams_per_meterLazy != nil {
		return *a.grams_per_meterLazy
	}
	grams_per_meter := a.convertFromBase(LinearDensityGramPerMeter)
	a.grams_per_meterLazy = &grams_per_meter
	return grams_per_meter
}

// PoundPerInch returns the value in PoundPerInch.
func (a *LinearDensity) PoundsPerInch() float64 {
	if a.pounds_per_inchLazy != nil {
		return *a.pounds_per_inchLazy
	}
	pounds_per_inch := a.convertFromBase(LinearDensityPoundPerInch)
	a.pounds_per_inchLazy = &pounds_per_inch
	return pounds_per_inch
}

// PoundPerFoot returns the value in PoundPerFoot.
func (a *LinearDensity) PoundsPerFoot() float64 {
	if a.pounds_per_footLazy != nil {
		return *a.pounds_per_footLazy
	}
	pounds_per_foot := a.convertFromBase(LinearDensityPoundPerFoot)
	a.pounds_per_footLazy = &pounds_per_foot
	return pounds_per_foot
}

// GramPerFoot returns the value in GramPerFoot.
func (a *LinearDensity) GramsPerFoot() float64 {
	if a.grams_per_footLazy != nil {
		return *a.grams_per_footLazy
	}
	grams_per_foot := a.convertFromBase(LinearDensityGramPerFoot)
	a.grams_per_footLazy = &grams_per_foot
	return grams_per_foot
}

// MicrogramPerMillimeter returns the value in MicrogramPerMillimeter.
func (a *LinearDensity) MicrogramsPerMillimeter() float64 {
	if a.micrograms_per_millimeterLazy != nil {
		return *a.micrograms_per_millimeterLazy
	}
	micrograms_per_millimeter := a.convertFromBase(LinearDensityMicrogramPerMillimeter)
	a.micrograms_per_millimeterLazy = &micrograms_per_millimeter
	return micrograms_per_millimeter
}

// MilligramPerMillimeter returns the value in MilligramPerMillimeter.
func (a *LinearDensity) MilligramsPerMillimeter() float64 {
	if a.milligrams_per_millimeterLazy != nil {
		return *a.milligrams_per_millimeterLazy
	}
	milligrams_per_millimeter := a.convertFromBase(LinearDensityMilligramPerMillimeter)
	a.milligrams_per_millimeterLazy = &milligrams_per_millimeter
	return milligrams_per_millimeter
}

// KilogramPerMillimeter returns the value in KilogramPerMillimeter.
func (a *LinearDensity) KilogramsPerMillimeter() float64 {
	if a.kilograms_per_millimeterLazy != nil {
		return *a.kilograms_per_millimeterLazy
	}
	kilograms_per_millimeter := a.convertFromBase(LinearDensityKilogramPerMillimeter)
	a.kilograms_per_millimeterLazy = &kilograms_per_millimeter
	return kilograms_per_millimeter
}

// MicrogramPerCentimeter returns the value in MicrogramPerCentimeter.
func (a *LinearDensity) MicrogramsPerCentimeter() float64 {
	if a.micrograms_per_centimeterLazy != nil {
		return *a.micrograms_per_centimeterLazy
	}
	micrograms_per_centimeter := a.convertFromBase(LinearDensityMicrogramPerCentimeter)
	a.micrograms_per_centimeterLazy = &micrograms_per_centimeter
	return micrograms_per_centimeter
}

// MilligramPerCentimeter returns the value in MilligramPerCentimeter.
func (a *LinearDensity) MilligramsPerCentimeter() float64 {
	if a.milligrams_per_centimeterLazy != nil {
		return *a.milligrams_per_centimeterLazy
	}
	milligrams_per_centimeter := a.convertFromBase(LinearDensityMilligramPerCentimeter)
	a.milligrams_per_centimeterLazy = &milligrams_per_centimeter
	return milligrams_per_centimeter
}

// KilogramPerCentimeter returns the value in KilogramPerCentimeter.
func (a *LinearDensity) KilogramsPerCentimeter() float64 {
	if a.kilograms_per_centimeterLazy != nil {
		return *a.kilograms_per_centimeterLazy
	}
	kilograms_per_centimeter := a.convertFromBase(LinearDensityKilogramPerCentimeter)
	a.kilograms_per_centimeterLazy = &kilograms_per_centimeter
	return kilograms_per_centimeter
}

// MicrogramPerMeter returns the value in MicrogramPerMeter.
func (a *LinearDensity) MicrogramsPerMeter() float64 {
	if a.micrograms_per_meterLazy != nil {
		return *a.micrograms_per_meterLazy
	}
	micrograms_per_meter := a.convertFromBase(LinearDensityMicrogramPerMeter)
	a.micrograms_per_meterLazy = &micrograms_per_meter
	return micrograms_per_meter
}

// MilligramPerMeter returns the value in MilligramPerMeter.
func (a *LinearDensity) MilligramsPerMeter() float64 {
	if a.milligrams_per_meterLazy != nil {
		return *a.milligrams_per_meterLazy
	}
	milligrams_per_meter := a.convertFromBase(LinearDensityMilligramPerMeter)
	a.milligrams_per_meterLazy = &milligrams_per_meter
	return milligrams_per_meter
}

// KilogramPerMeter returns the value in KilogramPerMeter.
func (a *LinearDensity) KilogramsPerMeter() float64 {
	if a.kilograms_per_meterLazy != nil {
		return *a.kilograms_per_meterLazy
	}
	kilograms_per_meter := a.convertFromBase(LinearDensityKilogramPerMeter)
	a.kilograms_per_meterLazy = &kilograms_per_meter
	return kilograms_per_meter
}

// MicrogramPerFoot returns the value in MicrogramPerFoot.
func (a *LinearDensity) MicrogramsPerFoot() float64 {
	if a.micrograms_per_footLazy != nil {
		return *a.micrograms_per_footLazy
	}
	micrograms_per_foot := a.convertFromBase(LinearDensityMicrogramPerFoot)
	a.micrograms_per_footLazy = &micrograms_per_foot
	return micrograms_per_foot
}

// MilligramPerFoot returns the value in MilligramPerFoot.
func (a *LinearDensity) MilligramsPerFoot() float64 {
	if a.milligrams_per_footLazy != nil {
		return *a.milligrams_per_footLazy
	}
	milligrams_per_foot := a.convertFromBase(LinearDensityMilligramPerFoot)
	a.milligrams_per_footLazy = &milligrams_per_foot
	return milligrams_per_foot
}

// KilogramPerFoot returns the value in KilogramPerFoot.
func (a *LinearDensity) KilogramsPerFoot() float64 {
	if a.kilograms_per_footLazy != nil {
		return *a.kilograms_per_footLazy
	}
	kilograms_per_foot := a.convertFromBase(LinearDensityKilogramPerFoot)
	a.kilograms_per_footLazy = &kilograms_per_foot
	return kilograms_per_foot
}


// ToDto creates an LinearDensityDto representation.
func (a *LinearDensity) ToDto(holdInUnit *LinearDensityUnits) LinearDensityDto {
	if holdInUnit == nil {
		defaultUnit := LinearDensityKilogramPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return LinearDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an LinearDensityDto representation.
func (a *LinearDensity) ToDtoJSON(holdInUnit *LinearDensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts LinearDensity to a specific unit value.
func (a *LinearDensity) Convert(toUnit LinearDensityUnits) float64 {
	switch toUnit { 
    case LinearDensityGramPerMillimeter:
		return a.GramsPerMillimeter()
    case LinearDensityGramPerCentimeter:
		return a.GramsPerCentimeter()
    case LinearDensityGramPerMeter:
		return a.GramsPerMeter()
    case LinearDensityPoundPerInch:
		return a.PoundsPerInch()
    case LinearDensityPoundPerFoot:
		return a.PoundsPerFoot()
    case LinearDensityGramPerFoot:
		return a.GramsPerFoot()
    case LinearDensityMicrogramPerMillimeter:
		return a.MicrogramsPerMillimeter()
    case LinearDensityMilligramPerMillimeter:
		return a.MilligramsPerMillimeter()
    case LinearDensityKilogramPerMillimeter:
		return a.KilogramsPerMillimeter()
    case LinearDensityMicrogramPerCentimeter:
		return a.MicrogramsPerCentimeter()
    case LinearDensityMilligramPerCentimeter:
		return a.MilligramsPerCentimeter()
    case LinearDensityKilogramPerCentimeter:
		return a.KilogramsPerCentimeter()
    case LinearDensityMicrogramPerMeter:
		return a.MicrogramsPerMeter()
    case LinearDensityMilligramPerMeter:
		return a.MilligramsPerMeter()
    case LinearDensityKilogramPerMeter:
		return a.KilogramsPerMeter()
    case LinearDensityMicrogramPerFoot:
		return a.MicrogramsPerFoot()
    case LinearDensityMilligramPerFoot:
		return a.MilligramsPerFoot()
    case LinearDensityKilogramPerFoot:
		return a.KilogramsPerFoot()
	default:
		return 0
	}
}

func (a *LinearDensity) convertFromBase(toUnit LinearDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case LinearDensityGramPerMillimeter:
		return (value) 
	case LinearDensityGramPerCentimeter:
		return (value / 1e-1) 
	case LinearDensityGramPerMeter:
		return (value / 1e-3) 
	case LinearDensityPoundPerInch:
		return (value * 5.5997415e-2) 
	case LinearDensityPoundPerFoot:
		return (value / 1.48816394) 
	case LinearDensityGramPerFoot:
		return (value / ( 1e-3 / 0.3048 )) 
	case LinearDensityMicrogramPerMillimeter:
		return ((value) / 1e-06) 
	case LinearDensityMilligramPerMillimeter:
		return ((value) / 0.001) 
	case LinearDensityKilogramPerMillimeter:
		return ((value) / 1000.0) 
	case LinearDensityMicrogramPerCentimeter:
		return ((value / 1e-1) / 1e-06) 
	case LinearDensityMilligramPerCentimeter:
		return ((value / 1e-1) / 0.001) 
	case LinearDensityKilogramPerCentimeter:
		return ((value / 1e-1) / 1000.0) 
	case LinearDensityMicrogramPerMeter:
		return ((value / 1e-3) / 1e-06) 
	case LinearDensityMilligramPerMeter:
		return ((value / 1e-3) / 0.001) 
	case LinearDensityKilogramPerMeter:
		return ((value / 1e-3) / 1000.0) 
	case LinearDensityMicrogramPerFoot:
		return ((value / ( 1e-3 / 0.3048 )) / 1e-06) 
	case LinearDensityMilligramPerFoot:
		return ((value / ( 1e-3 / 0.3048 )) / 0.001) 
	case LinearDensityKilogramPerFoot:
		return ((value / ( 1e-3 / 0.3048 )) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *LinearDensity) convertToBase(value float64, fromUnit LinearDensityUnits) float64 {
	switch fromUnit { 
	case LinearDensityGramPerMillimeter:
		return (value) 
	case LinearDensityGramPerCentimeter:
		return (value * 1e-1) 
	case LinearDensityGramPerMeter:
		return (value * 1e-3) 
	case LinearDensityPoundPerInch:
		return (value / 5.5997415e-2) 
	case LinearDensityPoundPerFoot:
		return (value * 1.48816394) 
	case LinearDensityGramPerFoot:
		return (value * ( 1e-3 / 0.3048 )) 
	case LinearDensityMicrogramPerMillimeter:
		return ((value) * 1e-06) 
	case LinearDensityMilligramPerMillimeter:
		return ((value) * 0.001) 
	case LinearDensityKilogramPerMillimeter:
		return ((value) * 1000.0) 
	case LinearDensityMicrogramPerCentimeter:
		return ((value * 1e-1) * 1e-06) 
	case LinearDensityMilligramPerCentimeter:
		return ((value * 1e-1) * 0.001) 
	case LinearDensityKilogramPerCentimeter:
		return ((value * 1e-1) * 1000.0) 
	case LinearDensityMicrogramPerMeter:
		return ((value * 1e-3) * 1e-06) 
	case LinearDensityMilligramPerMeter:
		return ((value * 1e-3) * 0.001) 
	case LinearDensityKilogramPerMeter:
		return ((value * 1e-3) * 1000.0) 
	case LinearDensityMicrogramPerFoot:
		return ((value * ( 1e-3 / 0.3048 )) * 1e-06) 
	case LinearDensityMilligramPerFoot:
		return ((value * ( 1e-3 / 0.3048 )) * 0.001) 
	case LinearDensityKilogramPerFoot:
		return ((value * ( 1e-3 / 0.3048 )) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a LinearDensity) String() string {
	return a.ToString(LinearDensityKilogramPerMeter, 2)
}

// ToString formats the LinearDensity to string.
// fractionalDigits -1 for not mention
func (a *LinearDensity) ToString(unit LinearDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *LinearDensity) getUnitAbbreviation(unit LinearDensityUnits) string {
	switch unit { 
	case LinearDensityGramPerMillimeter:
		return "g/mm" 
	case LinearDensityGramPerCentimeter:
		return "g/cm" 
	case LinearDensityGramPerMeter:
		return "g/m" 
	case LinearDensityPoundPerInch:
		return "lb/in" 
	case LinearDensityPoundPerFoot:
		return "lb/ft" 
	case LinearDensityGramPerFoot:
		return "g/ft" 
	case LinearDensityMicrogramPerMillimeter:
		return "μg/mm" 
	case LinearDensityMilligramPerMillimeter:
		return "mg/mm" 
	case LinearDensityKilogramPerMillimeter:
		return "kg/mm" 
	case LinearDensityMicrogramPerCentimeter:
		return "μg/cm" 
	case LinearDensityMilligramPerCentimeter:
		return "mg/cm" 
	case LinearDensityKilogramPerCentimeter:
		return "kg/cm" 
	case LinearDensityMicrogramPerMeter:
		return "μg/m" 
	case LinearDensityMilligramPerMeter:
		return "mg/m" 
	case LinearDensityKilogramPerMeter:
		return "kg/m" 
	case LinearDensityMicrogramPerFoot:
		return "μg/ft" 
	case LinearDensityMilligramPerFoot:
		return "mg/ft" 
	case LinearDensityKilogramPerFoot:
		return "kg/ft" 
	default:
		return ""
	}
}

// Check if the given LinearDensity are equals to the current LinearDensity
func (a *LinearDensity) Equals(other *LinearDensity) bool {
	return a.value == other.BaseValue()
}

// Check if the given LinearDensity are equals to the current LinearDensity
func (a *LinearDensity) CompareTo(other *LinearDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given LinearDensity to the current LinearDensity.
func (a *LinearDensity) Add(other *LinearDensity) *LinearDensity {
	return &LinearDensity{value: a.value + other.BaseValue()}
}

// Subtract the given LinearDensity to the current LinearDensity.
func (a *LinearDensity) Subtract(other *LinearDensity) *LinearDensity {
	return &LinearDensity{value: a.value - other.BaseValue()}
}

// Multiply the given LinearDensity to the current LinearDensity.
func (a *LinearDensity) Multiply(other *LinearDensity) *LinearDensity {
	return &LinearDensity{value: a.value * other.BaseValue()}
}

// Divide the given LinearDensity to the current LinearDensity.
func (a *LinearDensity) Divide(other *LinearDensity) *LinearDensity {
	return &LinearDensity{value: a.value / other.BaseValue()}
}