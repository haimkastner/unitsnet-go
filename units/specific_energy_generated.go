// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificEnergyUnits defines various units of SpecificEnergy.
type SpecificEnergyUnits string

const (
    
        // 
        SpecificEnergyJoulePerKilogram SpecificEnergyUnits = "JoulePerKilogram"
        // 
        SpecificEnergyMegaJoulePerTonne SpecificEnergyUnits = "MegaJoulePerTonne"
        // 
        SpecificEnergyCaloriePerGram SpecificEnergyUnits = "CaloriePerGram"
        // 
        SpecificEnergyWattHourPerKilogram SpecificEnergyUnits = "WattHourPerKilogram"
        // 
        SpecificEnergyWattDayPerKilogram SpecificEnergyUnits = "WattDayPerKilogram"
        // 
        SpecificEnergyWattDayPerTonne SpecificEnergyUnits = "WattDayPerTonne"
        // 
        SpecificEnergyWattDayPerShortTon SpecificEnergyUnits = "WattDayPerShortTon"
        // 
        SpecificEnergyWattHourPerPound SpecificEnergyUnits = "WattHourPerPound"
        // 
        SpecificEnergyBtuPerPound SpecificEnergyUnits = "BtuPerPound"
        // 
        SpecificEnergyKilojoulePerKilogram SpecificEnergyUnits = "KilojoulePerKilogram"
        // 
        SpecificEnergyMegajoulePerKilogram SpecificEnergyUnits = "MegajoulePerKilogram"
        // 
        SpecificEnergyKilocaloriePerGram SpecificEnergyUnits = "KilocaloriePerGram"
        // 
        SpecificEnergyKilowattHourPerKilogram SpecificEnergyUnits = "KilowattHourPerKilogram"
        // 
        SpecificEnergyMegawattHourPerKilogram SpecificEnergyUnits = "MegawattHourPerKilogram"
        // 
        SpecificEnergyGigawattHourPerKilogram SpecificEnergyUnits = "GigawattHourPerKilogram"
        // 
        SpecificEnergyKilowattDayPerKilogram SpecificEnergyUnits = "KilowattDayPerKilogram"
        // 
        SpecificEnergyMegawattDayPerKilogram SpecificEnergyUnits = "MegawattDayPerKilogram"
        // 
        SpecificEnergyGigawattDayPerKilogram SpecificEnergyUnits = "GigawattDayPerKilogram"
        // 
        SpecificEnergyTerawattDayPerKilogram SpecificEnergyUnits = "TerawattDayPerKilogram"
        // 
        SpecificEnergyKilowattDayPerTonne SpecificEnergyUnits = "KilowattDayPerTonne"
        // 
        SpecificEnergyMegawattDayPerTonne SpecificEnergyUnits = "MegawattDayPerTonne"
        // 
        SpecificEnergyGigawattDayPerTonne SpecificEnergyUnits = "GigawattDayPerTonne"
        // 
        SpecificEnergyTerawattDayPerTonne SpecificEnergyUnits = "TerawattDayPerTonne"
        // 
        SpecificEnergyKilowattDayPerShortTon SpecificEnergyUnits = "KilowattDayPerShortTon"
        // 
        SpecificEnergyMegawattDayPerShortTon SpecificEnergyUnits = "MegawattDayPerShortTon"
        // 
        SpecificEnergyGigawattDayPerShortTon SpecificEnergyUnits = "GigawattDayPerShortTon"
        // 
        SpecificEnergyTerawattDayPerShortTon SpecificEnergyUnits = "TerawattDayPerShortTon"
        // 
        SpecificEnergyKilowattHourPerPound SpecificEnergyUnits = "KilowattHourPerPound"
        // 
        SpecificEnergyMegawattHourPerPound SpecificEnergyUnits = "MegawattHourPerPound"
        // 
        SpecificEnergyGigawattHourPerPound SpecificEnergyUnits = "GigawattHourPerPound"
)

var internalSpecificEnergyUnitsMap = map[SpecificEnergyUnits]bool{
	
	SpecificEnergyJoulePerKilogram: true,
	SpecificEnergyMegaJoulePerTonne: true,
	SpecificEnergyCaloriePerGram: true,
	SpecificEnergyWattHourPerKilogram: true,
	SpecificEnergyWattDayPerKilogram: true,
	SpecificEnergyWattDayPerTonne: true,
	SpecificEnergyWattDayPerShortTon: true,
	SpecificEnergyWattHourPerPound: true,
	SpecificEnergyBtuPerPound: true,
	SpecificEnergyKilojoulePerKilogram: true,
	SpecificEnergyMegajoulePerKilogram: true,
	SpecificEnergyKilocaloriePerGram: true,
	SpecificEnergyKilowattHourPerKilogram: true,
	SpecificEnergyMegawattHourPerKilogram: true,
	SpecificEnergyGigawattHourPerKilogram: true,
	SpecificEnergyKilowattDayPerKilogram: true,
	SpecificEnergyMegawattDayPerKilogram: true,
	SpecificEnergyGigawattDayPerKilogram: true,
	SpecificEnergyTerawattDayPerKilogram: true,
	SpecificEnergyKilowattDayPerTonne: true,
	SpecificEnergyMegawattDayPerTonne: true,
	SpecificEnergyGigawattDayPerTonne: true,
	SpecificEnergyTerawattDayPerTonne: true,
	SpecificEnergyKilowattDayPerShortTon: true,
	SpecificEnergyMegawattDayPerShortTon: true,
	SpecificEnergyGigawattDayPerShortTon: true,
	SpecificEnergyTerawattDayPerShortTon: true,
	SpecificEnergyKilowattHourPerPound: true,
	SpecificEnergyMegawattHourPerPound: true,
	SpecificEnergyGigawattHourPerPound: true,
}

// SpecificEnergyDto represents a SpecificEnergy measurement with a numerical value and its corresponding unit.
type SpecificEnergyDto struct {
    // Value is the numerical representation of the SpecificEnergy.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the SpecificEnergy, as defined in the SpecificEnergyUnits enumeration.
	Unit  SpecificEnergyUnits `json:"unit" validate:"required,oneof=JoulePerKilogram MegaJoulePerTonne CaloriePerGram WattHourPerKilogram WattDayPerKilogram WattDayPerTonne WattDayPerShortTon WattHourPerPound BtuPerPound KilojoulePerKilogram MegajoulePerKilogram KilocaloriePerGram KilowattHourPerKilogram MegawattHourPerKilogram GigawattHourPerKilogram KilowattDayPerKilogram MegawattDayPerKilogram GigawattDayPerKilogram TerawattDayPerKilogram KilowattDayPerTonne MegawattDayPerTonne GigawattDayPerTonne TerawattDayPerTonne KilowattDayPerShortTon MegawattDayPerShortTon GigawattDayPerShortTon TerawattDayPerShortTon KilowattHourPerPound MegawattHourPerPound GigawattHourPerPound"`
}

// SpecificEnergyDtoFactory groups methods for creating and serializing SpecificEnergyDto objects.
type SpecificEnergyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a SpecificEnergyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf SpecificEnergyDtoFactory) FromJSON(data []byte) (*SpecificEnergyDto, error) {
	a := SpecificEnergyDto{}

    // Parse JSON into SpecificEnergyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a SpecificEnergyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a SpecificEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// SpecificEnergy represents a measurement in a of SpecificEnergy.
//
// The SpecificEnergy
type SpecificEnergy struct {
	// value is the base measurement stored internally.
	value       float64
    
    joules_per_kilogramLazy *float64 
    mega_joules_per_tonneLazy *float64 
    calories_per_gramLazy *float64 
    watt_hours_per_kilogramLazy *float64 
    watt_days_per_kilogramLazy *float64 
    watt_days_per_tonneLazy *float64 
    watt_days_per_short_tonLazy *float64 
    watt_hours_per_poundLazy *float64 
    btu_per_poundLazy *float64 
    kilojoules_per_kilogramLazy *float64 
    megajoules_per_kilogramLazy *float64 
    kilocalories_per_gramLazy *float64 
    kilowatt_hours_per_kilogramLazy *float64 
    megawatt_hours_per_kilogramLazy *float64 
    gigawatt_hours_per_kilogramLazy *float64 
    kilowatt_days_per_kilogramLazy *float64 
    megawatt_days_per_kilogramLazy *float64 
    gigawatt_days_per_kilogramLazy *float64 
    terawatt_days_per_kilogramLazy *float64 
    kilowatt_days_per_tonneLazy *float64 
    megawatt_days_per_tonneLazy *float64 
    gigawatt_days_per_tonneLazy *float64 
    terawatt_days_per_tonneLazy *float64 
    kilowatt_days_per_short_tonLazy *float64 
    megawatt_days_per_short_tonLazy *float64 
    gigawatt_days_per_short_tonLazy *float64 
    terawatt_days_per_short_tonLazy *float64 
    kilowatt_hours_per_poundLazy *float64 
    megawatt_hours_per_poundLazy *float64 
    gigawatt_hours_per_poundLazy *float64 
}

// SpecificEnergyFactory groups methods for creating SpecificEnergy instances.
type SpecificEnergyFactory struct{}

// CreateSpecificEnergy creates a new SpecificEnergy instance from the given value and unit.
func (uf SpecificEnergyFactory) CreateSpecificEnergy(value float64, unit SpecificEnergyUnits) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, unit)
}

// FromDto converts a SpecificEnergyDto to a SpecificEnergy instance.
func (uf SpecificEnergyFactory) FromDto(dto SpecificEnergyDto) (*SpecificEnergy, error) {
	return newSpecificEnergy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a SpecificEnergy instance.
func (uf SpecificEnergyFactory) FromDtoJSON(data []byte) (*SpecificEnergy, error) {
	unitDto, err := SpecificEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SpecificEnergyDto from JSON: %w", err)
	}
	return SpecificEnergyFactory{}.FromDto(*unitDto)
}


// FromJoulesPerKilogram creates a new SpecificEnergy instance from a value in JoulesPerKilogram.
func (uf SpecificEnergyFactory) FromJoulesPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyJoulePerKilogram)
}

// FromMegaJoulesPerTonne creates a new SpecificEnergy instance from a value in MegaJoulesPerTonne.
func (uf SpecificEnergyFactory) FromMegaJoulesPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegaJoulePerTonne)
}

// FromCaloriesPerGram creates a new SpecificEnergy instance from a value in CaloriesPerGram.
func (uf SpecificEnergyFactory) FromCaloriesPerGram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyCaloriePerGram)
}

// FromWattHoursPerKilogram creates a new SpecificEnergy instance from a value in WattHoursPerKilogram.
func (uf SpecificEnergyFactory) FromWattHoursPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattHourPerKilogram)
}

// FromWattDaysPerKilogram creates a new SpecificEnergy instance from a value in WattDaysPerKilogram.
func (uf SpecificEnergyFactory) FromWattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattDayPerKilogram)
}

// FromWattDaysPerTonne creates a new SpecificEnergy instance from a value in WattDaysPerTonne.
func (uf SpecificEnergyFactory) FromWattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattDayPerTonne)
}

// FromWattDaysPerShortTon creates a new SpecificEnergy instance from a value in WattDaysPerShortTon.
func (uf SpecificEnergyFactory) FromWattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattDayPerShortTon)
}

// FromWattHoursPerPound creates a new SpecificEnergy instance from a value in WattHoursPerPound.
func (uf SpecificEnergyFactory) FromWattHoursPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattHourPerPound)
}

// FromBtuPerPound creates a new SpecificEnergy instance from a value in BtuPerPound.
func (uf SpecificEnergyFactory) FromBtuPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyBtuPerPound)
}

// FromKilojoulesPerKilogram creates a new SpecificEnergy instance from a value in KilojoulesPerKilogram.
func (uf SpecificEnergyFactory) FromKilojoulesPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilojoulePerKilogram)
}

// FromMegajoulesPerKilogram creates a new SpecificEnergy instance from a value in MegajoulesPerKilogram.
func (uf SpecificEnergyFactory) FromMegajoulesPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegajoulePerKilogram)
}

// FromKilocaloriesPerGram creates a new SpecificEnergy instance from a value in KilocaloriesPerGram.
func (uf SpecificEnergyFactory) FromKilocaloriesPerGram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilocaloriePerGram)
}

// FromKilowattHoursPerKilogram creates a new SpecificEnergy instance from a value in KilowattHoursPerKilogram.
func (uf SpecificEnergyFactory) FromKilowattHoursPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattHourPerKilogram)
}

// FromMegawattHoursPerKilogram creates a new SpecificEnergy instance from a value in MegawattHoursPerKilogram.
func (uf SpecificEnergyFactory) FromMegawattHoursPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattHourPerKilogram)
}

// FromGigawattHoursPerKilogram creates a new SpecificEnergy instance from a value in GigawattHoursPerKilogram.
func (uf SpecificEnergyFactory) FromGigawattHoursPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattHourPerKilogram)
}

// FromKilowattDaysPerKilogram creates a new SpecificEnergy instance from a value in KilowattDaysPerKilogram.
func (uf SpecificEnergyFactory) FromKilowattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattDayPerKilogram)
}

// FromMegawattDaysPerKilogram creates a new SpecificEnergy instance from a value in MegawattDaysPerKilogram.
func (uf SpecificEnergyFactory) FromMegawattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattDayPerKilogram)
}

// FromGigawattDaysPerKilogram creates a new SpecificEnergy instance from a value in GigawattDaysPerKilogram.
func (uf SpecificEnergyFactory) FromGigawattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattDayPerKilogram)
}

// FromTerawattDaysPerKilogram creates a new SpecificEnergy instance from a value in TerawattDaysPerKilogram.
func (uf SpecificEnergyFactory) FromTerawattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyTerawattDayPerKilogram)
}

// FromKilowattDaysPerTonne creates a new SpecificEnergy instance from a value in KilowattDaysPerTonne.
func (uf SpecificEnergyFactory) FromKilowattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattDayPerTonne)
}

// FromMegawattDaysPerTonne creates a new SpecificEnergy instance from a value in MegawattDaysPerTonne.
func (uf SpecificEnergyFactory) FromMegawattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattDayPerTonne)
}

// FromGigawattDaysPerTonne creates a new SpecificEnergy instance from a value in GigawattDaysPerTonne.
func (uf SpecificEnergyFactory) FromGigawattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattDayPerTonne)
}

// FromTerawattDaysPerTonne creates a new SpecificEnergy instance from a value in TerawattDaysPerTonne.
func (uf SpecificEnergyFactory) FromTerawattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyTerawattDayPerTonne)
}

// FromKilowattDaysPerShortTon creates a new SpecificEnergy instance from a value in KilowattDaysPerShortTon.
func (uf SpecificEnergyFactory) FromKilowattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattDayPerShortTon)
}

// FromMegawattDaysPerShortTon creates a new SpecificEnergy instance from a value in MegawattDaysPerShortTon.
func (uf SpecificEnergyFactory) FromMegawattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattDayPerShortTon)
}

// FromGigawattDaysPerShortTon creates a new SpecificEnergy instance from a value in GigawattDaysPerShortTon.
func (uf SpecificEnergyFactory) FromGigawattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattDayPerShortTon)
}

// FromTerawattDaysPerShortTon creates a new SpecificEnergy instance from a value in TerawattDaysPerShortTon.
func (uf SpecificEnergyFactory) FromTerawattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyTerawattDayPerShortTon)
}

// FromKilowattHoursPerPound creates a new SpecificEnergy instance from a value in KilowattHoursPerPound.
func (uf SpecificEnergyFactory) FromKilowattHoursPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattHourPerPound)
}

// FromMegawattHoursPerPound creates a new SpecificEnergy instance from a value in MegawattHoursPerPound.
func (uf SpecificEnergyFactory) FromMegawattHoursPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattHourPerPound)
}

// FromGigawattHoursPerPound creates a new SpecificEnergy instance from a value in GigawattHoursPerPound.
func (uf SpecificEnergyFactory) FromGigawattHoursPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattHourPerPound)
}


// newSpecificEnergy creates a new SpecificEnergy.
func newSpecificEnergy(value float64, fromUnit SpecificEnergyUnits) (*SpecificEnergy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalSpecificEnergyUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in SpecificEnergyUnits", fromUnit)
	}
	a := &SpecificEnergy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificEnergy in JoulePerKilogram unit (the base/default quantity).
func (a *SpecificEnergy) BaseValue() float64 {
	return a.value
}


// JoulesPerKilogram returns the SpecificEnergy value in JoulesPerKilogram.
//
// 
func (a *SpecificEnergy) JoulesPerKilogram() float64 {
	if a.joules_per_kilogramLazy != nil {
		return *a.joules_per_kilogramLazy
	}
	joules_per_kilogram := a.convertFromBase(SpecificEnergyJoulePerKilogram)
	a.joules_per_kilogramLazy = &joules_per_kilogram
	return joules_per_kilogram
}

// MegaJoulesPerTonne returns the SpecificEnergy value in MegaJoulesPerTonne.
//
// 
func (a *SpecificEnergy) MegaJoulesPerTonne() float64 {
	if a.mega_joules_per_tonneLazy != nil {
		return *a.mega_joules_per_tonneLazy
	}
	mega_joules_per_tonne := a.convertFromBase(SpecificEnergyMegaJoulePerTonne)
	a.mega_joules_per_tonneLazy = &mega_joules_per_tonne
	return mega_joules_per_tonne
}

// CaloriesPerGram returns the SpecificEnergy value in CaloriesPerGram.
//
// 
func (a *SpecificEnergy) CaloriesPerGram() float64 {
	if a.calories_per_gramLazy != nil {
		return *a.calories_per_gramLazy
	}
	calories_per_gram := a.convertFromBase(SpecificEnergyCaloriePerGram)
	a.calories_per_gramLazy = &calories_per_gram
	return calories_per_gram
}

// WattHoursPerKilogram returns the SpecificEnergy value in WattHoursPerKilogram.
//
// 
func (a *SpecificEnergy) WattHoursPerKilogram() float64 {
	if a.watt_hours_per_kilogramLazy != nil {
		return *a.watt_hours_per_kilogramLazy
	}
	watt_hours_per_kilogram := a.convertFromBase(SpecificEnergyWattHourPerKilogram)
	a.watt_hours_per_kilogramLazy = &watt_hours_per_kilogram
	return watt_hours_per_kilogram
}

// WattDaysPerKilogram returns the SpecificEnergy value in WattDaysPerKilogram.
//
// 
func (a *SpecificEnergy) WattDaysPerKilogram() float64 {
	if a.watt_days_per_kilogramLazy != nil {
		return *a.watt_days_per_kilogramLazy
	}
	watt_days_per_kilogram := a.convertFromBase(SpecificEnergyWattDayPerKilogram)
	a.watt_days_per_kilogramLazy = &watt_days_per_kilogram
	return watt_days_per_kilogram
}

// WattDaysPerTonne returns the SpecificEnergy value in WattDaysPerTonne.
//
// 
func (a *SpecificEnergy) WattDaysPerTonne() float64 {
	if a.watt_days_per_tonneLazy != nil {
		return *a.watt_days_per_tonneLazy
	}
	watt_days_per_tonne := a.convertFromBase(SpecificEnergyWattDayPerTonne)
	a.watt_days_per_tonneLazy = &watt_days_per_tonne
	return watt_days_per_tonne
}

// WattDaysPerShortTon returns the SpecificEnergy value in WattDaysPerShortTon.
//
// 
func (a *SpecificEnergy) WattDaysPerShortTon() float64 {
	if a.watt_days_per_short_tonLazy != nil {
		return *a.watt_days_per_short_tonLazy
	}
	watt_days_per_short_ton := a.convertFromBase(SpecificEnergyWattDayPerShortTon)
	a.watt_days_per_short_tonLazy = &watt_days_per_short_ton
	return watt_days_per_short_ton
}

// WattHoursPerPound returns the SpecificEnergy value in WattHoursPerPound.
//
// 
func (a *SpecificEnergy) WattHoursPerPound() float64 {
	if a.watt_hours_per_poundLazy != nil {
		return *a.watt_hours_per_poundLazy
	}
	watt_hours_per_pound := a.convertFromBase(SpecificEnergyWattHourPerPound)
	a.watt_hours_per_poundLazy = &watt_hours_per_pound
	return watt_hours_per_pound
}

// BtuPerPound returns the SpecificEnergy value in BtuPerPound.
//
// 
func (a *SpecificEnergy) BtuPerPound() float64 {
	if a.btu_per_poundLazy != nil {
		return *a.btu_per_poundLazy
	}
	btu_per_pound := a.convertFromBase(SpecificEnergyBtuPerPound)
	a.btu_per_poundLazy = &btu_per_pound
	return btu_per_pound
}

// KilojoulesPerKilogram returns the SpecificEnergy value in KilojoulesPerKilogram.
//
// 
func (a *SpecificEnergy) KilojoulesPerKilogram() float64 {
	if a.kilojoules_per_kilogramLazy != nil {
		return *a.kilojoules_per_kilogramLazy
	}
	kilojoules_per_kilogram := a.convertFromBase(SpecificEnergyKilojoulePerKilogram)
	a.kilojoules_per_kilogramLazy = &kilojoules_per_kilogram
	return kilojoules_per_kilogram
}

// MegajoulesPerKilogram returns the SpecificEnergy value in MegajoulesPerKilogram.
//
// 
func (a *SpecificEnergy) MegajoulesPerKilogram() float64 {
	if a.megajoules_per_kilogramLazy != nil {
		return *a.megajoules_per_kilogramLazy
	}
	megajoules_per_kilogram := a.convertFromBase(SpecificEnergyMegajoulePerKilogram)
	a.megajoules_per_kilogramLazy = &megajoules_per_kilogram
	return megajoules_per_kilogram
}

// KilocaloriesPerGram returns the SpecificEnergy value in KilocaloriesPerGram.
//
// 
func (a *SpecificEnergy) KilocaloriesPerGram() float64 {
	if a.kilocalories_per_gramLazy != nil {
		return *a.kilocalories_per_gramLazy
	}
	kilocalories_per_gram := a.convertFromBase(SpecificEnergyKilocaloriePerGram)
	a.kilocalories_per_gramLazy = &kilocalories_per_gram
	return kilocalories_per_gram
}

// KilowattHoursPerKilogram returns the SpecificEnergy value in KilowattHoursPerKilogram.
//
// 
func (a *SpecificEnergy) KilowattHoursPerKilogram() float64 {
	if a.kilowatt_hours_per_kilogramLazy != nil {
		return *a.kilowatt_hours_per_kilogramLazy
	}
	kilowatt_hours_per_kilogram := a.convertFromBase(SpecificEnergyKilowattHourPerKilogram)
	a.kilowatt_hours_per_kilogramLazy = &kilowatt_hours_per_kilogram
	return kilowatt_hours_per_kilogram
}

// MegawattHoursPerKilogram returns the SpecificEnergy value in MegawattHoursPerKilogram.
//
// 
func (a *SpecificEnergy) MegawattHoursPerKilogram() float64 {
	if a.megawatt_hours_per_kilogramLazy != nil {
		return *a.megawatt_hours_per_kilogramLazy
	}
	megawatt_hours_per_kilogram := a.convertFromBase(SpecificEnergyMegawattHourPerKilogram)
	a.megawatt_hours_per_kilogramLazy = &megawatt_hours_per_kilogram
	return megawatt_hours_per_kilogram
}

// GigawattHoursPerKilogram returns the SpecificEnergy value in GigawattHoursPerKilogram.
//
// 
func (a *SpecificEnergy) GigawattHoursPerKilogram() float64 {
	if a.gigawatt_hours_per_kilogramLazy != nil {
		return *a.gigawatt_hours_per_kilogramLazy
	}
	gigawatt_hours_per_kilogram := a.convertFromBase(SpecificEnergyGigawattHourPerKilogram)
	a.gigawatt_hours_per_kilogramLazy = &gigawatt_hours_per_kilogram
	return gigawatt_hours_per_kilogram
}

// KilowattDaysPerKilogram returns the SpecificEnergy value in KilowattDaysPerKilogram.
//
// 
func (a *SpecificEnergy) KilowattDaysPerKilogram() float64 {
	if a.kilowatt_days_per_kilogramLazy != nil {
		return *a.kilowatt_days_per_kilogramLazy
	}
	kilowatt_days_per_kilogram := a.convertFromBase(SpecificEnergyKilowattDayPerKilogram)
	a.kilowatt_days_per_kilogramLazy = &kilowatt_days_per_kilogram
	return kilowatt_days_per_kilogram
}

// MegawattDaysPerKilogram returns the SpecificEnergy value in MegawattDaysPerKilogram.
//
// 
func (a *SpecificEnergy) MegawattDaysPerKilogram() float64 {
	if a.megawatt_days_per_kilogramLazy != nil {
		return *a.megawatt_days_per_kilogramLazy
	}
	megawatt_days_per_kilogram := a.convertFromBase(SpecificEnergyMegawattDayPerKilogram)
	a.megawatt_days_per_kilogramLazy = &megawatt_days_per_kilogram
	return megawatt_days_per_kilogram
}

// GigawattDaysPerKilogram returns the SpecificEnergy value in GigawattDaysPerKilogram.
//
// 
func (a *SpecificEnergy) GigawattDaysPerKilogram() float64 {
	if a.gigawatt_days_per_kilogramLazy != nil {
		return *a.gigawatt_days_per_kilogramLazy
	}
	gigawatt_days_per_kilogram := a.convertFromBase(SpecificEnergyGigawattDayPerKilogram)
	a.gigawatt_days_per_kilogramLazy = &gigawatt_days_per_kilogram
	return gigawatt_days_per_kilogram
}

// TerawattDaysPerKilogram returns the SpecificEnergy value in TerawattDaysPerKilogram.
//
// 
func (a *SpecificEnergy) TerawattDaysPerKilogram() float64 {
	if a.terawatt_days_per_kilogramLazy != nil {
		return *a.terawatt_days_per_kilogramLazy
	}
	terawatt_days_per_kilogram := a.convertFromBase(SpecificEnergyTerawattDayPerKilogram)
	a.terawatt_days_per_kilogramLazy = &terawatt_days_per_kilogram
	return terawatt_days_per_kilogram
}

// KilowattDaysPerTonne returns the SpecificEnergy value in KilowattDaysPerTonne.
//
// 
func (a *SpecificEnergy) KilowattDaysPerTonne() float64 {
	if a.kilowatt_days_per_tonneLazy != nil {
		return *a.kilowatt_days_per_tonneLazy
	}
	kilowatt_days_per_tonne := a.convertFromBase(SpecificEnergyKilowattDayPerTonne)
	a.kilowatt_days_per_tonneLazy = &kilowatt_days_per_tonne
	return kilowatt_days_per_tonne
}

// MegawattDaysPerTonne returns the SpecificEnergy value in MegawattDaysPerTonne.
//
// 
func (a *SpecificEnergy) MegawattDaysPerTonne() float64 {
	if a.megawatt_days_per_tonneLazy != nil {
		return *a.megawatt_days_per_tonneLazy
	}
	megawatt_days_per_tonne := a.convertFromBase(SpecificEnergyMegawattDayPerTonne)
	a.megawatt_days_per_tonneLazy = &megawatt_days_per_tonne
	return megawatt_days_per_tonne
}

// GigawattDaysPerTonne returns the SpecificEnergy value in GigawattDaysPerTonne.
//
// 
func (a *SpecificEnergy) GigawattDaysPerTonne() float64 {
	if a.gigawatt_days_per_tonneLazy != nil {
		return *a.gigawatt_days_per_tonneLazy
	}
	gigawatt_days_per_tonne := a.convertFromBase(SpecificEnergyGigawattDayPerTonne)
	a.gigawatt_days_per_tonneLazy = &gigawatt_days_per_tonne
	return gigawatt_days_per_tonne
}

// TerawattDaysPerTonne returns the SpecificEnergy value in TerawattDaysPerTonne.
//
// 
func (a *SpecificEnergy) TerawattDaysPerTonne() float64 {
	if a.terawatt_days_per_tonneLazy != nil {
		return *a.terawatt_days_per_tonneLazy
	}
	terawatt_days_per_tonne := a.convertFromBase(SpecificEnergyTerawattDayPerTonne)
	a.terawatt_days_per_tonneLazy = &terawatt_days_per_tonne
	return terawatt_days_per_tonne
}

// KilowattDaysPerShortTon returns the SpecificEnergy value in KilowattDaysPerShortTon.
//
// 
func (a *SpecificEnergy) KilowattDaysPerShortTon() float64 {
	if a.kilowatt_days_per_short_tonLazy != nil {
		return *a.kilowatt_days_per_short_tonLazy
	}
	kilowatt_days_per_short_ton := a.convertFromBase(SpecificEnergyKilowattDayPerShortTon)
	a.kilowatt_days_per_short_tonLazy = &kilowatt_days_per_short_ton
	return kilowatt_days_per_short_ton
}

// MegawattDaysPerShortTon returns the SpecificEnergy value in MegawattDaysPerShortTon.
//
// 
func (a *SpecificEnergy) MegawattDaysPerShortTon() float64 {
	if a.megawatt_days_per_short_tonLazy != nil {
		return *a.megawatt_days_per_short_tonLazy
	}
	megawatt_days_per_short_ton := a.convertFromBase(SpecificEnergyMegawattDayPerShortTon)
	a.megawatt_days_per_short_tonLazy = &megawatt_days_per_short_ton
	return megawatt_days_per_short_ton
}

// GigawattDaysPerShortTon returns the SpecificEnergy value in GigawattDaysPerShortTon.
//
// 
func (a *SpecificEnergy) GigawattDaysPerShortTon() float64 {
	if a.gigawatt_days_per_short_tonLazy != nil {
		return *a.gigawatt_days_per_short_tonLazy
	}
	gigawatt_days_per_short_ton := a.convertFromBase(SpecificEnergyGigawattDayPerShortTon)
	a.gigawatt_days_per_short_tonLazy = &gigawatt_days_per_short_ton
	return gigawatt_days_per_short_ton
}

// TerawattDaysPerShortTon returns the SpecificEnergy value in TerawattDaysPerShortTon.
//
// 
func (a *SpecificEnergy) TerawattDaysPerShortTon() float64 {
	if a.terawatt_days_per_short_tonLazy != nil {
		return *a.terawatt_days_per_short_tonLazy
	}
	terawatt_days_per_short_ton := a.convertFromBase(SpecificEnergyTerawattDayPerShortTon)
	a.terawatt_days_per_short_tonLazy = &terawatt_days_per_short_ton
	return terawatt_days_per_short_ton
}

// KilowattHoursPerPound returns the SpecificEnergy value in KilowattHoursPerPound.
//
// 
func (a *SpecificEnergy) KilowattHoursPerPound() float64 {
	if a.kilowatt_hours_per_poundLazy != nil {
		return *a.kilowatt_hours_per_poundLazy
	}
	kilowatt_hours_per_pound := a.convertFromBase(SpecificEnergyKilowattHourPerPound)
	a.kilowatt_hours_per_poundLazy = &kilowatt_hours_per_pound
	return kilowatt_hours_per_pound
}

// MegawattHoursPerPound returns the SpecificEnergy value in MegawattHoursPerPound.
//
// 
func (a *SpecificEnergy) MegawattHoursPerPound() float64 {
	if a.megawatt_hours_per_poundLazy != nil {
		return *a.megawatt_hours_per_poundLazy
	}
	megawatt_hours_per_pound := a.convertFromBase(SpecificEnergyMegawattHourPerPound)
	a.megawatt_hours_per_poundLazy = &megawatt_hours_per_pound
	return megawatt_hours_per_pound
}

// GigawattHoursPerPound returns the SpecificEnergy value in GigawattHoursPerPound.
//
// 
func (a *SpecificEnergy) GigawattHoursPerPound() float64 {
	if a.gigawatt_hours_per_poundLazy != nil {
		return *a.gigawatt_hours_per_poundLazy
	}
	gigawatt_hours_per_pound := a.convertFromBase(SpecificEnergyGigawattHourPerPound)
	a.gigawatt_hours_per_poundLazy = &gigawatt_hours_per_pound
	return gigawatt_hours_per_pound
}


// ToDto creates a SpecificEnergyDto representation from the SpecificEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerKilogram by default.
func (a *SpecificEnergy) ToDto(holdInUnit *SpecificEnergyUnits) SpecificEnergyDto {
	if holdInUnit == nil {
		defaultUnit := SpecificEnergyJoulePerKilogram // Default value
		holdInUnit = &defaultUnit
	}

	return SpecificEnergyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the SpecificEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerKilogram by default.
func (a *SpecificEnergy) ToDtoJSON(holdInUnit *SpecificEnergyUnits) ([]byte, error) {
	// Convert to SpecificEnergyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a SpecificEnergy to a specific unit value.
// The function uses the provided unit type (SpecificEnergyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *SpecificEnergy) Convert(toUnit SpecificEnergyUnits) float64 {
	switch toUnit { 
    case SpecificEnergyJoulePerKilogram:
		return a.JoulesPerKilogram()
    case SpecificEnergyMegaJoulePerTonne:
		return a.MegaJoulesPerTonne()
    case SpecificEnergyCaloriePerGram:
		return a.CaloriesPerGram()
    case SpecificEnergyWattHourPerKilogram:
		return a.WattHoursPerKilogram()
    case SpecificEnergyWattDayPerKilogram:
		return a.WattDaysPerKilogram()
    case SpecificEnergyWattDayPerTonne:
		return a.WattDaysPerTonne()
    case SpecificEnergyWattDayPerShortTon:
		return a.WattDaysPerShortTon()
    case SpecificEnergyWattHourPerPound:
		return a.WattHoursPerPound()
    case SpecificEnergyBtuPerPound:
		return a.BtuPerPound()
    case SpecificEnergyKilojoulePerKilogram:
		return a.KilojoulesPerKilogram()
    case SpecificEnergyMegajoulePerKilogram:
		return a.MegajoulesPerKilogram()
    case SpecificEnergyKilocaloriePerGram:
		return a.KilocaloriesPerGram()
    case SpecificEnergyKilowattHourPerKilogram:
		return a.KilowattHoursPerKilogram()
    case SpecificEnergyMegawattHourPerKilogram:
		return a.MegawattHoursPerKilogram()
    case SpecificEnergyGigawattHourPerKilogram:
		return a.GigawattHoursPerKilogram()
    case SpecificEnergyKilowattDayPerKilogram:
		return a.KilowattDaysPerKilogram()
    case SpecificEnergyMegawattDayPerKilogram:
		return a.MegawattDaysPerKilogram()
    case SpecificEnergyGigawattDayPerKilogram:
		return a.GigawattDaysPerKilogram()
    case SpecificEnergyTerawattDayPerKilogram:
		return a.TerawattDaysPerKilogram()
    case SpecificEnergyKilowattDayPerTonne:
		return a.KilowattDaysPerTonne()
    case SpecificEnergyMegawattDayPerTonne:
		return a.MegawattDaysPerTonne()
    case SpecificEnergyGigawattDayPerTonne:
		return a.GigawattDaysPerTonne()
    case SpecificEnergyTerawattDayPerTonne:
		return a.TerawattDaysPerTonne()
    case SpecificEnergyKilowattDayPerShortTon:
		return a.KilowattDaysPerShortTon()
    case SpecificEnergyMegawattDayPerShortTon:
		return a.MegawattDaysPerShortTon()
    case SpecificEnergyGigawattDayPerShortTon:
		return a.GigawattDaysPerShortTon()
    case SpecificEnergyTerawattDayPerShortTon:
		return a.TerawattDaysPerShortTon()
    case SpecificEnergyKilowattHourPerPound:
		return a.KilowattHoursPerPound()
    case SpecificEnergyMegawattHourPerPound:
		return a.MegawattHoursPerPound()
    case SpecificEnergyGigawattHourPerPound:
		return a.GigawattHoursPerPound()
	default:
		return math.NaN()
	}
}

func (a *SpecificEnergy) convertFromBase(toUnit SpecificEnergyUnits) float64 {
    value := a.value
	switch toUnit { 
	case SpecificEnergyJoulePerKilogram:
		return (value) 
	case SpecificEnergyMegaJoulePerTonne:
		return (value / 1e3) 
	case SpecificEnergyCaloriePerGram:
		return (value / 4.184e3) 
	case SpecificEnergyWattHourPerKilogram:
		return (value / 3.6e3) 
	case SpecificEnergyWattDayPerKilogram:
		return (value / (24 * 3.6e3)) 
	case SpecificEnergyWattDayPerTonne:
		return (value / ((24 * 3.6e3) / 1e3)) 
	case SpecificEnergyWattDayPerShortTon:
		return (value / ((24 * 3.6e3) / 9.0718474e2)) 
	case SpecificEnergyWattHourPerPound:
		return (value / 7.93664e3) 
	case SpecificEnergyBtuPerPound:
		return (value * 0.45359237 / 1055.05585262) 
	case SpecificEnergyKilojoulePerKilogram:
		return ((value) / 1000.0) 
	case SpecificEnergyMegajoulePerKilogram:
		return ((value) / 1000000.0) 
	case SpecificEnergyKilocaloriePerGram:
		return ((value / 4.184e3) / 1000.0) 
	case SpecificEnergyKilowattHourPerKilogram:
		return ((value / 3.6e3) / 1000.0) 
	case SpecificEnergyMegawattHourPerKilogram:
		return ((value / 3.6e3) / 1000000.0) 
	case SpecificEnergyGigawattHourPerKilogram:
		return ((value / 3.6e3) / 1000000000.0) 
	case SpecificEnergyKilowattDayPerKilogram:
		return ((value / (24 * 3.6e3)) / 1000.0) 
	case SpecificEnergyMegawattDayPerKilogram:
		return ((value / (24 * 3.6e3)) / 1000000.0) 
	case SpecificEnergyGigawattDayPerKilogram:
		return ((value / (24 * 3.6e3)) / 1000000000.0) 
	case SpecificEnergyTerawattDayPerKilogram:
		return ((value / (24 * 3.6e3)) / 1000000000000.0) 
	case SpecificEnergyKilowattDayPerTonne:
		return ((value / ((24 * 3.6e3) / 1e3)) / 1000.0) 
	case SpecificEnergyMegawattDayPerTonne:
		return ((value / ((24 * 3.6e3) / 1e3)) / 1000000.0) 
	case SpecificEnergyGigawattDayPerTonne:
		return ((value / ((24 * 3.6e3) / 1e3)) / 1000000000.0) 
	case SpecificEnergyTerawattDayPerTonne:
		return ((value / ((24 * 3.6e3) / 1e3)) / 1000000000000.0) 
	case SpecificEnergyKilowattDayPerShortTon:
		return ((value / ((24 * 3.6e3) / 9.0718474e2)) / 1000.0) 
	case SpecificEnergyMegawattDayPerShortTon:
		return ((value / ((24 * 3.6e3) / 9.0718474e2)) / 1000000.0) 
	case SpecificEnergyGigawattDayPerShortTon:
		return ((value / ((24 * 3.6e3) / 9.0718474e2)) / 1000000000.0) 
	case SpecificEnergyTerawattDayPerShortTon:
		return ((value / ((24 * 3.6e3) / 9.0718474e2)) / 1000000000000.0) 
	case SpecificEnergyKilowattHourPerPound:
		return ((value / 7.93664e3) / 1000.0) 
	case SpecificEnergyMegawattHourPerPound:
		return ((value / 7.93664e3) / 1000000.0) 
	case SpecificEnergyGigawattHourPerPound:
		return ((value / 7.93664e3) / 1000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *SpecificEnergy) convertToBase(value float64, fromUnit SpecificEnergyUnits) float64 {
	switch fromUnit { 
	case SpecificEnergyJoulePerKilogram:
		return (value) 
	case SpecificEnergyMegaJoulePerTonne:
		return (value * 1e3) 
	case SpecificEnergyCaloriePerGram:
		return (value * 4.184e3) 
	case SpecificEnergyWattHourPerKilogram:
		return (value * 3.6e3) 
	case SpecificEnergyWattDayPerKilogram:
		return (value * (24 * 3.6e3)) 
	case SpecificEnergyWattDayPerTonne:
		return (value * ((24 * 3.6e3) / 1e3)) 
	case SpecificEnergyWattDayPerShortTon:
		return (value * ((24 * 3.6e3) / 9.0718474e2)) 
	case SpecificEnergyWattHourPerPound:
		return (value * 7.93664e3) 
	case SpecificEnergyBtuPerPound:
		return (value * 1055.05585262 / 0.45359237) 
	case SpecificEnergyKilojoulePerKilogram:
		return ((value) * 1000.0) 
	case SpecificEnergyMegajoulePerKilogram:
		return ((value) * 1000000.0) 
	case SpecificEnergyKilocaloriePerGram:
		return ((value * 4.184e3) * 1000.0) 
	case SpecificEnergyKilowattHourPerKilogram:
		return ((value * 3.6e3) * 1000.0) 
	case SpecificEnergyMegawattHourPerKilogram:
		return ((value * 3.6e3) * 1000000.0) 
	case SpecificEnergyGigawattHourPerKilogram:
		return ((value * 3.6e3) * 1000000000.0) 
	case SpecificEnergyKilowattDayPerKilogram:
		return ((value * (24 * 3.6e3)) * 1000.0) 
	case SpecificEnergyMegawattDayPerKilogram:
		return ((value * (24 * 3.6e3)) * 1000000.0) 
	case SpecificEnergyGigawattDayPerKilogram:
		return ((value * (24 * 3.6e3)) * 1000000000.0) 
	case SpecificEnergyTerawattDayPerKilogram:
		return ((value * (24 * 3.6e3)) * 1000000000000.0) 
	case SpecificEnergyKilowattDayPerTonne:
		return ((value * ((24 * 3.6e3) / 1e3)) * 1000.0) 
	case SpecificEnergyMegawattDayPerTonne:
		return ((value * ((24 * 3.6e3) / 1e3)) * 1000000.0) 
	case SpecificEnergyGigawattDayPerTonne:
		return ((value * ((24 * 3.6e3) / 1e3)) * 1000000000.0) 
	case SpecificEnergyTerawattDayPerTonne:
		return ((value * ((24 * 3.6e3) / 1e3)) * 1000000000000.0) 
	case SpecificEnergyKilowattDayPerShortTon:
		return ((value * ((24 * 3.6e3) / 9.0718474e2)) * 1000.0) 
	case SpecificEnergyMegawattDayPerShortTon:
		return ((value * ((24 * 3.6e3) / 9.0718474e2)) * 1000000.0) 
	case SpecificEnergyGigawattDayPerShortTon:
		return ((value * ((24 * 3.6e3) / 9.0718474e2)) * 1000000000.0) 
	case SpecificEnergyTerawattDayPerShortTon:
		return ((value * ((24 * 3.6e3) / 9.0718474e2)) * 1000000000000.0) 
	case SpecificEnergyKilowattHourPerPound:
		return ((value * 7.93664e3) * 1000.0) 
	case SpecificEnergyMegawattHourPerPound:
		return ((value * 7.93664e3) * 1000000.0) 
	case SpecificEnergyGigawattHourPerPound:
		return ((value * 7.93664e3) * 1000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the SpecificEnergy in the default unit (JoulePerKilogram),
// formatted to two decimal places.
func (a SpecificEnergy) String() string {
	return a.ToString(SpecificEnergyJoulePerKilogram, 2)
}

// ToString formats the SpecificEnergy value as a string with the specified unit and fractional digits.
// It converts the SpecificEnergy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the SpecificEnergy value will be converted (e.g., JoulePerKilogram))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the SpecificEnergy with the unit abbreviation.
func (a *SpecificEnergy) ToString(unit SpecificEnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetSpecificEnergyAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetSpecificEnergyAbbreviation(unit))
}

// Equals checks if the given SpecificEnergy is equal to the current SpecificEnergy.
//
// Parameters:
//    other: The SpecificEnergy to compare against.
//
// Returns:
//    bool: Returns true if both SpecificEnergy are equal, false otherwise.
func (a *SpecificEnergy) Equals(other *SpecificEnergy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current SpecificEnergy with another SpecificEnergy.
// It returns -1 if the current SpecificEnergy is less than the other SpecificEnergy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The SpecificEnergy to compare against.
//
// Returns:
//    int: -1 if the current SpecificEnergy is less, 1 if greater, and 0 if equal.
func (a *SpecificEnergy) CompareTo(other *SpecificEnergy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given SpecificEnergy to the current SpecificEnergy and returns the result.
// The result is a new SpecificEnergy instance with the sum of the values.
//
// Parameters:
//    other: The SpecificEnergy to add to the current SpecificEnergy.
//
// Returns:
//    *SpecificEnergy: A new SpecificEnergy instance representing the sum of both SpecificEnergy.
func (a *SpecificEnergy) Add(other *SpecificEnergy) *SpecificEnergy {
	return &SpecificEnergy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given SpecificEnergy from the current SpecificEnergy and returns the result.
// The result is a new SpecificEnergy instance with the difference of the values.
//
// Parameters:
//    other: The SpecificEnergy to subtract from the current SpecificEnergy.
//
// Returns:
//    *SpecificEnergy: A new SpecificEnergy instance representing the difference of both SpecificEnergy.
func (a *SpecificEnergy) Subtract(other *SpecificEnergy) *SpecificEnergy {
	return &SpecificEnergy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current SpecificEnergy by the given SpecificEnergy and returns the result.
// The result is a new SpecificEnergy instance with the product of the values.
//
// Parameters:
//    other: The SpecificEnergy to multiply with the current SpecificEnergy.
//
// Returns:
//    *SpecificEnergy: A new SpecificEnergy instance representing the product of both SpecificEnergy.
func (a *SpecificEnergy) Multiply(other *SpecificEnergy) *SpecificEnergy {
	return &SpecificEnergy{value: a.value * other.BaseValue()}
}

// Divide divides the current SpecificEnergy by the given SpecificEnergy and returns the result.
// The result is a new SpecificEnergy instance with the quotient of the values.
//
// Parameters:
//    other: The SpecificEnergy to divide the current SpecificEnergy by.
//
// Returns:
//    *SpecificEnergy: A new SpecificEnergy instance representing the quotient of both SpecificEnergy.
func (a *SpecificEnergy) Divide(other *SpecificEnergy) *SpecificEnergy {
	return &SpecificEnergy{value: a.value / other.BaseValue()}
}

// GetSpecificEnergyAbbreviation gets the unit abbreviation.
func GetSpecificEnergyAbbreviation(unit SpecificEnergyUnits) string {
	switch unit { 
	case SpecificEnergyJoulePerKilogram:
		return "J/kg" 
	case SpecificEnergyMegaJoulePerTonne:
		return "MJ/t" 
	case SpecificEnergyCaloriePerGram:
		return "cal/g" 
	case SpecificEnergyWattHourPerKilogram:
		return "Wh/kg" 
	case SpecificEnergyWattDayPerKilogram:
		return "Wd/kg" 
	case SpecificEnergyWattDayPerTonne:
		return "Wd/t" 
	case SpecificEnergyWattDayPerShortTon:
		return "Wd/ST" 
	case SpecificEnergyWattHourPerPound:
		return "Wh/lbs" 
	case SpecificEnergyBtuPerPound:
		return "btu/lb" 
	case SpecificEnergyKilojoulePerKilogram:
		return "kJ/kg" 
	case SpecificEnergyMegajoulePerKilogram:
		return "MJ/kg" 
	case SpecificEnergyKilocaloriePerGram:
		return "kcal/g" 
	case SpecificEnergyKilowattHourPerKilogram:
		return "kWh/kg" 
	case SpecificEnergyMegawattHourPerKilogram:
		return "MWh/kg" 
	case SpecificEnergyGigawattHourPerKilogram:
		return "GWh/kg" 
	case SpecificEnergyKilowattDayPerKilogram:
		return "kWd/kg" 
	case SpecificEnergyMegawattDayPerKilogram:
		return "MWd/kg" 
	case SpecificEnergyGigawattDayPerKilogram:
		return "GWd/kg" 
	case SpecificEnergyTerawattDayPerKilogram:
		return "TWd/kg" 
	case SpecificEnergyKilowattDayPerTonne:
		return "kWd/t" 
	case SpecificEnergyMegawattDayPerTonne:
		return "MWd/t" 
	case SpecificEnergyGigawattDayPerTonne:
		return "GWd/t" 
	case SpecificEnergyTerawattDayPerTonne:
		return "TWd/t" 
	case SpecificEnergyKilowattDayPerShortTon:
		return "kWd/ST" 
	case SpecificEnergyMegawattDayPerShortTon:
		return "MWd/ST" 
	case SpecificEnergyGigawattDayPerShortTon:
		return "GWd/ST" 
	case SpecificEnergyTerawattDayPerShortTon:
		return "TWd/ST" 
	case SpecificEnergyKilowattHourPerPound:
		return "kWh/lbs" 
	case SpecificEnergyMegawattHourPerPound:
		return "MWh/lbs" 
	case SpecificEnergyGigawattHourPerPound:
		return "GWh/lbs" 
	default:
		return ""
	}
}