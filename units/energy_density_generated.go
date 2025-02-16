// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// EnergyDensityUnits defines various units of EnergyDensity.
type EnergyDensityUnits string

const (
    
        // 
        EnergyDensityJoulePerCubicMeter EnergyDensityUnits = "JoulePerCubicMeter"
        // 
        EnergyDensityWattHourPerCubicMeter EnergyDensityUnits = "WattHourPerCubicMeter"
        // 
        EnergyDensityKilojoulePerCubicMeter EnergyDensityUnits = "KilojoulePerCubicMeter"
        // 
        EnergyDensityMegajoulePerCubicMeter EnergyDensityUnits = "MegajoulePerCubicMeter"
        // 
        EnergyDensityGigajoulePerCubicMeter EnergyDensityUnits = "GigajoulePerCubicMeter"
        // 
        EnergyDensityTerajoulePerCubicMeter EnergyDensityUnits = "TerajoulePerCubicMeter"
        // 
        EnergyDensityPetajoulePerCubicMeter EnergyDensityUnits = "PetajoulePerCubicMeter"
        // 
        EnergyDensityKilowattHourPerCubicMeter EnergyDensityUnits = "KilowattHourPerCubicMeter"
        // 
        EnergyDensityMegawattHourPerCubicMeter EnergyDensityUnits = "MegawattHourPerCubicMeter"
        // 
        EnergyDensityGigawattHourPerCubicMeter EnergyDensityUnits = "GigawattHourPerCubicMeter"
        // 
        EnergyDensityTerawattHourPerCubicMeter EnergyDensityUnits = "TerawattHourPerCubicMeter"
        // 
        EnergyDensityPetawattHourPerCubicMeter EnergyDensityUnits = "PetawattHourPerCubicMeter"
)

// EnergyDensityDto represents a EnergyDensity measurement with a numerical value and its corresponding unit.
type EnergyDensityDto struct {
    // Value is the numerical representation of the EnergyDensity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the EnergyDensity, as defined in the EnergyDensityUnits enumeration.
	Unit  EnergyDensityUnits `json:"unit"`
}

// EnergyDensityDtoFactory groups methods for creating and serializing EnergyDensityDto objects.
type EnergyDensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a EnergyDensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf EnergyDensityDtoFactory) FromJSON(data []byte) (*EnergyDensityDto, error) {
	a := EnergyDensityDto{}

    // Parse JSON into EnergyDensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a EnergyDensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a EnergyDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// EnergyDensity represents a measurement in a of EnergyDensity.
//
// None
type EnergyDensity struct {
	// value is the base measurement stored internally.
	value       float64
    
    joules_per_cubic_meterLazy *float64 
    watt_hours_per_cubic_meterLazy *float64 
    kilojoules_per_cubic_meterLazy *float64 
    megajoules_per_cubic_meterLazy *float64 
    gigajoules_per_cubic_meterLazy *float64 
    terajoules_per_cubic_meterLazy *float64 
    petajoules_per_cubic_meterLazy *float64 
    kilowatt_hours_per_cubic_meterLazy *float64 
    megawatt_hours_per_cubic_meterLazy *float64 
    gigawatt_hours_per_cubic_meterLazy *float64 
    terawatt_hours_per_cubic_meterLazy *float64 
    petawatt_hours_per_cubic_meterLazy *float64 
}

// EnergyDensityFactory groups methods for creating EnergyDensity instances.
type EnergyDensityFactory struct{}

// CreateEnergyDensity creates a new EnergyDensity instance from the given value and unit.
func (uf EnergyDensityFactory) CreateEnergyDensity(value float64, unit EnergyDensityUnits) (*EnergyDensity, error) {
	return newEnergyDensity(value, unit)
}

// FromDto converts a EnergyDensityDto to a EnergyDensity instance.
func (uf EnergyDensityFactory) FromDto(dto EnergyDensityDto) (*EnergyDensity, error) {
	return newEnergyDensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a EnergyDensity instance.
func (uf EnergyDensityFactory) FromDtoJSON(data []byte) (*EnergyDensity, error) {
	unitDto, err := EnergyDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnergyDensityDto from JSON: %w", err)
	}
	return EnergyDensityFactory{}.FromDto(*unitDto)
}


// FromJoulesPerCubicMeter creates a new EnergyDensity instance from a value in JoulesPerCubicMeter.
func (uf EnergyDensityFactory) FromJoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityJoulePerCubicMeter)
}

// FromWattHoursPerCubicMeter creates a new EnergyDensity instance from a value in WattHoursPerCubicMeter.
func (uf EnergyDensityFactory) FromWattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityWattHourPerCubicMeter)
}

// FromKilojoulesPerCubicMeter creates a new EnergyDensity instance from a value in KilojoulesPerCubicMeter.
func (uf EnergyDensityFactory) FromKilojoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityKilojoulePerCubicMeter)
}

// FromMegajoulesPerCubicMeter creates a new EnergyDensity instance from a value in MegajoulesPerCubicMeter.
func (uf EnergyDensityFactory) FromMegajoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityMegajoulePerCubicMeter)
}

// FromGigajoulesPerCubicMeter creates a new EnergyDensity instance from a value in GigajoulesPerCubicMeter.
func (uf EnergyDensityFactory) FromGigajoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityGigajoulePerCubicMeter)
}

// FromTerajoulesPerCubicMeter creates a new EnergyDensity instance from a value in TerajoulesPerCubicMeter.
func (uf EnergyDensityFactory) FromTerajoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityTerajoulePerCubicMeter)
}

// FromPetajoulesPerCubicMeter creates a new EnergyDensity instance from a value in PetajoulesPerCubicMeter.
func (uf EnergyDensityFactory) FromPetajoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityPetajoulePerCubicMeter)
}

// FromKilowattHoursPerCubicMeter creates a new EnergyDensity instance from a value in KilowattHoursPerCubicMeter.
func (uf EnergyDensityFactory) FromKilowattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityKilowattHourPerCubicMeter)
}

// FromMegawattHoursPerCubicMeter creates a new EnergyDensity instance from a value in MegawattHoursPerCubicMeter.
func (uf EnergyDensityFactory) FromMegawattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityMegawattHourPerCubicMeter)
}

// FromGigawattHoursPerCubicMeter creates a new EnergyDensity instance from a value in GigawattHoursPerCubicMeter.
func (uf EnergyDensityFactory) FromGigawattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityGigawattHourPerCubicMeter)
}

// FromTerawattHoursPerCubicMeter creates a new EnergyDensity instance from a value in TerawattHoursPerCubicMeter.
func (uf EnergyDensityFactory) FromTerawattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityTerawattHourPerCubicMeter)
}

// FromPetawattHoursPerCubicMeter creates a new EnergyDensity instance from a value in PetawattHoursPerCubicMeter.
func (uf EnergyDensityFactory) FromPetawattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityPetawattHourPerCubicMeter)
}


// newEnergyDensity creates a new EnergyDensity.
func newEnergyDensity(value float64, fromUnit EnergyDensityUnits) (*EnergyDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &EnergyDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of EnergyDensity in JoulePerCubicMeter unit (the base/default quantity).
func (a *EnergyDensity) BaseValue() float64 {
	return a.value
}


// JoulesPerCubicMeter returns the EnergyDensity value in JoulesPerCubicMeter.
//
// 
func (a *EnergyDensity) JoulesPerCubicMeter() float64 {
	if a.joules_per_cubic_meterLazy != nil {
		return *a.joules_per_cubic_meterLazy
	}
	joules_per_cubic_meter := a.convertFromBase(EnergyDensityJoulePerCubicMeter)
	a.joules_per_cubic_meterLazy = &joules_per_cubic_meter
	return joules_per_cubic_meter
}

// WattHoursPerCubicMeter returns the EnergyDensity value in WattHoursPerCubicMeter.
//
// 
func (a *EnergyDensity) WattHoursPerCubicMeter() float64 {
	if a.watt_hours_per_cubic_meterLazy != nil {
		return *a.watt_hours_per_cubic_meterLazy
	}
	watt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityWattHourPerCubicMeter)
	a.watt_hours_per_cubic_meterLazy = &watt_hours_per_cubic_meter
	return watt_hours_per_cubic_meter
}

// KilojoulesPerCubicMeter returns the EnergyDensity value in KilojoulesPerCubicMeter.
//
// 
func (a *EnergyDensity) KilojoulesPerCubicMeter() float64 {
	if a.kilojoules_per_cubic_meterLazy != nil {
		return *a.kilojoules_per_cubic_meterLazy
	}
	kilojoules_per_cubic_meter := a.convertFromBase(EnergyDensityKilojoulePerCubicMeter)
	a.kilojoules_per_cubic_meterLazy = &kilojoules_per_cubic_meter
	return kilojoules_per_cubic_meter
}

// MegajoulesPerCubicMeter returns the EnergyDensity value in MegajoulesPerCubicMeter.
//
// 
func (a *EnergyDensity) MegajoulesPerCubicMeter() float64 {
	if a.megajoules_per_cubic_meterLazy != nil {
		return *a.megajoules_per_cubic_meterLazy
	}
	megajoules_per_cubic_meter := a.convertFromBase(EnergyDensityMegajoulePerCubicMeter)
	a.megajoules_per_cubic_meterLazy = &megajoules_per_cubic_meter
	return megajoules_per_cubic_meter
}

// GigajoulesPerCubicMeter returns the EnergyDensity value in GigajoulesPerCubicMeter.
//
// 
func (a *EnergyDensity) GigajoulesPerCubicMeter() float64 {
	if a.gigajoules_per_cubic_meterLazy != nil {
		return *a.gigajoules_per_cubic_meterLazy
	}
	gigajoules_per_cubic_meter := a.convertFromBase(EnergyDensityGigajoulePerCubicMeter)
	a.gigajoules_per_cubic_meterLazy = &gigajoules_per_cubic_meter
	return gigajoules_per_cubic_meter
}

// TerajoulesPerCubicMeter returns the EnergyDensity value in TerajoulesPerCubicMeter.
//
// 
func (a *EnergyDensity) TerajoulesPerCubicMeter() float64 {
	if a.terajoules_per_cubic_meterLazy != nil {
		return *a.terajoules_per_cubic_meterLazy
	}
	terajoules_per_cubic_meter := a.convertFromBase(EnergyDensityTerajoulePerCubicMeter)
	a.terajoules_per_cubic_meterLazy = &terajoules_per_cubic_meter
	return terajoules_per_cubic_meter
}

// PetajoulesPerCubicMeter returns the EnergyDensity value in PetajoulesPerCubicMeter.
//
// 
func (a *EnergyDensity) PetajoulesPerCubicMeter() float64 {
	if a.petajoules_per_cubic_meterLazy != nil {
		return *a.petajoules_per_cubic_meterLazy
	}
	petajoules_per_cubic_meter := a.convertFromBase(EnergyDensityPetajoulePerCubicMeter)
	a.petajoules_per_cubic_meterLazy = &petajoules_per_cubic_meter
	return petajoules_per_cubic_meter
}

// KilowattHoursPerCubicMeter returns the EnergyDensity value in KilowattHoursPerCubicMeter.
//
// 
func (a *EnergyDensity) KilowattHoursPerCubicMeter() float64 {
	if a.kilowatt_hours_per_cubic_meterLazy != nil {
		return *a.kilowatt_hours_per_cubic_meterLazy
	}
	kilowatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityKilowattHourPerCubicMeter)
	a.kilowatt_hours_per_cubic_meterLazy = &kilowatt_hours_per_cubic_meter
	return kilowatt_hours_per_cubic_meter
}

// MegawattHoursPerCubicMeter returns the EnergyDensity value in MegawattHoursPerCubicMeter.
//
// 
func (a *EnergyDensity) MegawattHoursPerCubicMeter() float64 {
	if a.megawatt_hours_per_cubic_meterLazy != nil {
		return *a.megawatt_hours_per_cubic_meterLazy
	}
	megawatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityMegawattHourPerCubicMeter)
	a.megawatt_hours_per_cubic_meterLazy = &megawatt_hours_per_cubic_meter
	return megawatt_hours_per_cubic_meter
}

// GigawattHoursPerCubicMeter returns the EnergyDensity value in GigawattHoursPerCubicMeter.
//
// 
func (a *EnergyDensity) GigawattHoursPerCubicMeter() float64 {
	if a.gigawatt_hours_per_cubic_meterLazy != nil {
		return *a.gigawatt_hours_per_cubic_meterLazy
	}
	gigawatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityGigawattHourPerCubicMeter)
	a.gigawatt_hours_per_cubic_meterLazy = &gigawatt_hours_per_cubic_meter
	return gigawatt_hours_per_cubic_meter
}

// TerawattHoursPerCubicMeter returns the EnergyDensity value in TerawattHoursPerCubicMeter.
//
// 
func (a *EnergyDensity) TerawattHoursPerCubicMeter() float64 {
	if a.terawatt_hours_per_cubic_meterLazy != nil {
		return *a.terawatt_hours_per_cubic_meterLazy
	}
	terawatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityTerawattHourPerCubicMeter)
	a.terawatt_hours_per_cubic_meterLazy = &terawatt_hours_per_cubic_meter
	return terawatt_hours_per_cubic_meter
}

// PetawattHoursPerCubicMeter returns the EnergyDensity value in PetawattHoursPerCubicMeter.
//
// 
func (a *EnergyDensity) PetawattHoursPerCubicMeter() float64 {
	if a.petawatt_hours_per_cubic_meterLazy != nil {
		return *a.petawatt_hours_per_cubic_meterLazy
	}
	petawatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityPetawattHourPerCubicMeter)
	a.petawatt_hours_per_cubic_meterLazy = &petawatt_hours_per_cubic_meter
	return petawatt_hours_per_cubic_meter
}


// ToDto creates a EnergyDensityDto representation from the EnergyDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerCubicMeter by default.
func (a *EnergyDensity) ToDto(holdInUnit *EnergyDensityUnits) EnergyDensityDto {
	if holdInUnit == nil {
		defaultUnit := EnergyDensityJoulePerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return EnergyDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the EnergyDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerCubicMeter by default.
func (a *EnergyDensity) ToDtoJSON(holdInUnit *EnergyDensityUnits) ([]byte, error) {
	// Convert to EnergyDensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a EnergyDensity to a specific unit value.
// The function uses the provided unit type (EnergyDensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *EnergyDensity) Convert(toUnit EnergyDensityUnits) float64 {
	switch toUnit { 
    case EnergyDensityJoulePerCubicMeter:
		return a.JoulesPerCubicMeter()
    case EnergyDensityWattHourPerCubicMeter:
		return a.WattHoursPerCubicMeter()
    case EnergyDensityKilojoulePerCubicMeter:
		return a.KilojoulesPerCubicMeter()
    case EnergyDensityMegajoulePerCubicMeter:
		return a.MegajoulesPerCubicMeter()
    case EnergyDensityGigajoulePerCubicMeter:
		return a.GigajoulesPerCubicMeter()
    case EnergyDensityTerajoulePerCubicMeter:
		return a.TerajoulesPerCubicMeter()
    case EnergyDensityPetajoulePerCubicMeter:
		return a.PetajoulesPerCubicMeter()
    case EnergyDensityKilowattHourPerCubicMeter:
		return a.KilowattHoursPerCubicMeter()
    case EnergyDensityMegawattHourPerCubicMeter:
		return a.MegawattHoursPerCubicMeter()
    case EnergyDensityGigawattHourPerCubicMeter:
		return a.GigawattHoursPerCubicMeter()
    case EnergyDensityTerawattHourPerCubicMeter:
		return a.TerawattHoursPerCubicMeter()
    case EnergyDensityPetawattHourPerCubicMeter:
		return a.PetawattHoursPerCubicMeter()
	default:
		return math.NaN()
	}
}

func (a *EnergyDensity) convertFromBase(toUnit EnergyDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case EnergyDensityJoulePerCubicMeter:
		return (value) 
	case EnergyDensityWattHourPerCubicMeter:
		return (value / 3.6e+3) 
	case EnergyDensityKilojoulePerCubicMeter:
		return ((value) / 1000.0) 
	case EnergyDensityMegajoulePerCubicMeter:
		return ((value) / 1000000.0) 
	case EnergyDensityGigajoulePerCubicMeter:
		return ((value) / 1000000000.0) 
	case EnergyDensityTerajoulePerCubicMeter:
		return ((value) / 1000000000000.0) 
	case EnergyDensityPetajoulePerCubicMeter:
		return ((value) / 1000000000000000.0) 
	case EnergyDensityKilowattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000.0) 
	case EnergyDensityMegawattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000000.0) 
	case EnergyDensityGigawattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000000000.0) 
	case EnergyDensityTerawattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000000000000.0) 
	case EnergyDensityPetawattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *EnergyDensity) convertToBase(value float64, fromUnit EnergyDensityUnits) float64 {
	switch fromUnit { 
	case EnergyDensityJoulePerCubicMeter:
		return (value) 
	case EnergyDensityWattHourPerCubicMeter:
		return (value * 3.6e+3) 
	case EnergyDensityKilojoulePerCubicMeter:
		return ((value) * 1000.0) 
	case EnergyDensityMegajoulePerCubicMeter:
		return ((value) * 1000000.0) 
	case EnergyDensityGigajoulePerCubicMeter:
		return ((value) * 1000000000.0) 
	case EnergyDensityTerajoulePerCubicMeter:
		return ((value) * 1000000000000.0) 
	case EnergyDensityPetajoulePerCubicMeter:
		return ((value) * 1000000000000000.0) 
	case EnergyDensityKilowattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000.0) 
	case EnergyDensityMegawattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000000.0) 
	case EnergyDensityGigawattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000000000.0) 
	case EnergyDensityTerawattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000000000000.0) 
	case EnergyDensityPetawattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000000000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the EnergyDensity in the default unit (JoulePerCubicMeter),
// formatted to two decimal places.
func (a EnergyDensity) String() string {
	return a.ToString(EnergyDensityJoulePerCubicMeter, 2)
}

// ToString formats the EnergyDensity value as a string with the specified unit and fractional digits.
// It converts the EnergyDensity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the EnergyDensity value will be converted (e.g., JoulePerCubicMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the EnergyDensity with the unit abbreviation.
func (a *EnergyDensity) ToString(unit EnergyDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetEnergyDensityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetEnergyDensityAbbreviation(unit))
}

// Equals checks if the given EnergyDensity is equal to the current EnergyDensity.
//
// Parameters:
//    other: The EnergyDensity to compare against.
//
// Returns:
//    bool: Returns true if both EnergyDensity are equal, false otherwise.
func (a *EnergyDensity) Equals(other *EnergyDensity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current EnergyDensity with another EnergyDensity.
// It returns -1 if the current EnergyDensity is less than the other EnergyDensity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The EnergyDensity to compare against.
//
// Returns:
//    int: -1 if the current EnergyDensity is less, 1 if greater, and 0 if equal.
func (a *EnergyDensity) CompareTo(other *EnergyDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given EnergyDensity to the current EnergyDensity and returns the result.
// The result is a new EnergyDensity instance with the sum of the values.
//
// Parameters:
//    other: The EnergyDensity to add to the current EnergyDensity.
//
// Returns:
//    *EnergyDensity: A new EnergyDensity instance representing the sum of both EnergyDensity.
func (a *EnergyDensity) Add(other *EnergyDensity) *EnergyDensity {
	return &EnergyDensity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given EnergyDensity from the current EnergyDensity and returns the result.
// The result is a new EnergyDensity instance with the difference of the values.
//
// Parameters:
//    other: The EnergyDensity to subtract from the current EnergyDensity.
//
// Returns:
//    *EnergyDensity: A new EnergyDensity instance representing the difference of both EnergyDensity.
func (a *EnergyDensity) Subtract(other *EnergyDensity) *EnergyDensity {
	return &EnergyDensity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current EnergyDensity by the given EnergyDensity and returns the result.
// The result is a new EnergyDensity instance with the product of the values.
//
// Parameters:
//    other: The EnergyDensity to multiply with the current EnergyDensity.
//
// Returns:
//    *EnergyDensity: A new EnergyDensity instance representing the product of both EnergyDensity.
func (a *EnergyDensity) Multiply(other *EnergyDensity) *EnergyDensity {
	return &EnergyDensity{value: a.value * other.BaseValue()}
}

// Divide divides the current EnergyDensity by the given EnergyDensity and returns the result.
// The result is a new EnergyDensity instance with the quotient of the values.
//
// Parameters:
//    other: The EnergyDensity to divide the current EnergyDensity by.
//
// Returns:
//    *EnergyDensity: A new EnergyDensity instance representing the quotient of both EnergyDensity.
func (a *EnergyDensity) Divide(other *EnergyDensity) *EnergyDensity {
	return &EnergyDensity{value: a.value / other.BaseValue()}
}

// GetEnergyDensityAbbreviation gets the unit abbreviation.
func GetEnergyDensityAbbreviation(unit EnergyDensityUnits) string {
	switch unit { 
	case EnergyDensityJoulePerCubicMeter:
		return "J/m³" 
	case EnergyDensityWattHourPerCubicMeter:
		return "Wh/m³" 
	case EnergyDensityKilojoulePerCubicMeter:
		return "kJ/m³" 
	case EnergyDensityMegajoulePerCubicMeter:
		return "MJ/m³" 
	case EnergyDensityGigajoulePerCubicMeter:
		return "GJ/m³" 
	case EnergyDensityTerajoulePerCubicMeter:
		return "TJ/m³" 
	case EnergyDensityPetajoulePerCubicMeter:
		return "PJ/m³" 
	case EnergyDensityKilowattHourPerCubicMeter:
		return "kWh/m³" 
	case EnergyDensityMegawattHourPerCubicMeter:
		return "MWh/m³" 
	case EnergyDensityGigawattHourPerCubicMeter:
		return "GWh/m³" 
	case EnergyDensityTerawattHourPerCubicMeter:
		return "TWh/m³" 
	case EnergyDensityPetawattHourPerCubicMeter:
		return "PWh/m³" 
	default:
		return ""
	}
}