package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificEnergyUnits enumeration
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

// SpecificEnergyDto represents an SpecificEnergy
type SpecificEnergyDto struct {
	Value float64
	Unit  SpecificEnergyUnits
}

// SpecificEnergyDtoFactory struct to group related functions
type SpecificEnergyDtoFactory struct{}

func (udf SpecificEnergyDtoFactory) FromJSON(data []byte) (*SpecificEnergyDto, error) {
	a := SpecificEnergyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a SpecificEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// SpecificEnergy struct
type SpecificEnergy struct {
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

// SpecificEnergyFactory struct to group related functions
type SpecificEnergyFactory struct{}

func (uf SpecificEnergyFactory) CreateSpecificEnergy(value float64, unit SpecificEnergyUnits) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, unit)
}

func (uf SpecificEnergyFactory) FromDto(dto SpecificEnergyDto) (*SpecificEnergy, error) {
	return newSpecificEnergy(dto.Value, dto.Unit)
}

func (uf SpecificEnergyFactory) FromDtoJSON(data []byte) (*SpecificEnergy, error) {
	unitDto, err := SpecificEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return SpecificEnergyFactory{}.FromDto(*unitDto)
}


// FromJoulePerKilogram creates a new SpecificEnergy instance from JoulePerKilogram.
func (uf SpecificEnergyFactory) FromJoulesPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyJoulePerKilogram)
}

// FromMegaJoulePerTonne creates a new SpecificEnergy instance from MegaJoulePerTonne.
func (uf SpecificEnergyFactory) FromMegaJoulesPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegaJoulePerTonne)
}

// FromCaloriePerGram creates a new SpecificEnergy instance from CaloriePerGram.
func (uf SpecificEnergyFactory) FromCaloriesPerGram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyCaloriePerGram)
}

// FromWattHourPerKilogram creates a new SpecificEnergy instance from WattHourPerKilogram.
func (uf SpecificEnergyFactory) FromWattHoursPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattHourPerKilogram)
}

// FromWattDayPerKilogram creates a new SpecificEnergy instance from WattDayPerKilogram.
func (uf SpecificEnergyFactory) FromWattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattDayPerKilogram)
}

// FromWattDayPerTonne creates a new SpecificEnergy instance from WattDayPerTonne.
func (uf SpecificEnergyFactory) FromWattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattDayPerTonne)
}

// FromWattDayPerShortTon creates a new SpecificEnergy instance from WattDayPerShortTon.
func (uf SpecificEnergyFactory) FromWattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattDayPerShortTon)
}

// FromWattHourPerPound creates a new SpecificEnergy instance from WattHourPerPound.
func (uf SpecificEnergyFactory) FromWattHoursPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyWattHourPerPound)
}

// FromBtuPerPound creates a new SpecificEnergy instance from BtuPerPound.
func (uf SpecificEnergyFactory) FromBtuPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyBtuPerPound)
}

// FromKilojoulePerKilogram creates a new SpecificEnergy instance from KilojoulePerKilogram.
func (uf SpecificEnergyFactory) FromKilojoulesPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilojoulePerKilogram)
}

// FromMegajoulePerKilogram creates a new SpecificEnergy instance from MegajoulePerKilogram.
func (uf SpecificEnergyFactory) FromMegajoulesPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegajoulePerKilogram)
}

// FromKilocaloriePerGram creates a new SpecificEnergy instance from KilocaloriePerGram.
func (uf SpecificEnergyFactory) FromKilocaloriesPerGram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilocaloriePerGram)
}

// FromKilowattHourPerKilogram creates a new SpecificEnergy instance from KilowattHourPerKilogram.
func (uf SpecificEnergyFactory) FromKilowattHoursPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattHourPerKilogram)
}

// FromMegawattHourPerKilogram creates a new SpecificEnergy instance from MegawattHourPerKilogram.
func (uf SpecificEnergyFactory) FromMegawattHoursPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattHourPerKilogram)
}

// FromGigawattHourPerKilogram creates a new SpecificEnergy instance from GigawattHourPerKilogram.
func (uf SpecificEnergyFactory) FromGigawattHoursPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattHourPerKilogram)
}

// FromKilowattDayPerKilogram creates a new SpecificEnergy instance from KilowattDayPerKilogram.
func (uf SpecificEnergyFactory) FromKilowattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattDayPerKilogram)
}

// FromMegawattDayPerKilogram creates a new SpecificEnergy instance from MegawattDayPerKilogram.
func (uf SpecificEnergyFactory) FromMegawattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattDayPerKilogram)
}

// FromGigawattDayPerKilogram creates a new SpecificEnergy instance from GigawattDayPerKilogram.
func (uf SpecificEnergyFactory) FromGigawattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattDayPerKilogram)
}

// FromTerawattDayPerKilogram creates a new SpecificEnergy instance from TerawattDayPerKilogram.
func (uf SpecificEnergyFactory) FromTerawattDaysPerKilogram(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyTerawattDayPerKilogram)
}

// FromKilowattDayPerTonne creates a new SpecificEnergy instance from KilowattDayPerTonne.
func (uf SpecificEnergyFactory) FromKilowattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattDayPerTonne)
}

// FromMegawattDayPerTonne creates a new SpecificEnergy instance from MegawattDayPerTonne.
func (uf SpecificEnergyFactory) FromMegawattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattDayPerTonne)
}

// FromGigawattDayPerTonne creates a new SpecificEnergy instance from GigawattDayPerTonne.
func (uf SpecificEnergyFactory) FromGigawattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattDayPerTonne)
}

// FromTerawattDayPerTonne creates a new SpecificEnergy instance from TerawattDayPerTonne.
func (uf SpecificEnergyFactory) FromTerawattDaysPerTonne(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyTerawattDayPerTonne)
}

// FromKilowattDayPerShortTon creates a new SpecificEnergy instance from KilowattDayPerShortTon.
func (uf SpecificEnergyFactory) FromKilowattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattDayPerShortTon)
}

// FromMegawattDayPerShortTon creates a new SpecificEnergy instance from MegawattDayPerShortTon.
func (uf SpecificEnergyFactory) FromMegawattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattDayPerShortTon)
}

// FromGigawattDayPerShortTon creates a new SpecificEnergy instance from GigawattDayPerShortTon.
func (uf SpecificEnergyFactory) FromGigawattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattDayPerShortTon)
}

// FromTerawattDayPerShortTon creates a new SpecificEnergy instance from TerawattDayPerShortTon.
func (uf SpecificEnergyFactory) FromTerawattDaysPerShortTon(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyTerawattDayPerShortTon)
}

// FromKilowattHourPerPound creates a new SpecificEnergy instance from KilowattHourPerPound.
func (uf SpecificEnergyFactory) FromKilowattHoursPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyKilowattHourPerPound)
}

// FromMegawattHourPerPound creates a new SpecificEnergy instance from MegawattHourPerPound.
func (uf SpecificEnergyFactory) FromMegawattHoursPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyMegawattHourPerPound)
}

// FromGigawattHourPerPound creates a new SpecificEnergy instance from GigawattHourPerPound.
func (uf SpecificEnergyFactory) FromGigawattHoursPerPound(value float64) (*SpecificEnergy, error) {
	return newSpecificEnergy(value, SpecificEnergyGigawattHourPerPound)
}




// newSpecificEnergy creates a new SpecificEnergy.
func newSpecificEnergy(value float64, fromUnit SpecificEnergyUnits) (*SpecificEnergy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &SpecificEnergy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificEnergy in JoulePerKilogram.
func (a *SpecificEnergy) BaseValue() float64 {
	return a.value
}


// JoulePerKilogram returns the value in JoulePerKilogram.
func (a *SpecificEnergy) JoulesPerKilogram() float64 {
	if a.joules_per_kilogramLazy != nil {
		return *a.joules_per_kilogramLazy
	}
	joules_per_kilogram := a.convertFromBase(SpecificEnergyJoulePerKilogram)
	a.joules_per_kilogramLazy = &joules_per_kilogram
	return joules_per_kilogram
}

// MegaJoulePerTonne returns the value in MegaJoulePerTonne.
func (a *SpecificEnergy) MegaJoulesPerTonne() float64 {
	if a.mega_joules_per_tonneLazy != nil {
		return *a.mega_joules_per_tonneLazy
	}
	mega_joules_per_tonne := a.convertFromBase(SpecificEnergyMegaJoulePerTonne)
	a.mega_joules_per_tonneLazy = &mega_joules_per_tonne
	return mega_joules_per_tonne
}

// CaloriePerGram returns the value in CaloriePerGram.
func (a *SpecificEnergy) CaloriesPerGram() float64 {
	if a.calories_per_gramLazy != nil {
		return *a.calories_per_gramLazy
	}
	calories_per_gram := a.convertFromBase(SpecificEnergyCaloriePerGram)
	a.calories_per_gramLazy = &calories_per_gram
	return calories_per_gram
}

// WattHourPerKilogram returns the value in WattHourPerKilogram.
func (a *SpecificEnergy) WattHoursPerKilogram() float64 {
	if a.watt_hours_per_kilogramLazy != nil {
		return *a.watt_hours_per_kilogramLazy
	}
	watt_hours_per_kilogram := a.convertFromBase(SpecificEnergyWattHourPerKilogram)
	a.watt_hours_per_kilogramLazy = &watt_hours_per_kilogram
	return watt_hours_per_kilogram
}

// WattDayPerKilogram returns the value in WattDayPerKilogram.
func (a *SpecificEnergy) WattDaysPerKilogram() float64 {
	if a.watt_days_per_kilogramLazy != nil {
		return *a.watt_days_per_kilogramLazy
	}
	watt_days_per_kilogram := a.convertFromBase(SpecificEnergyWattDayPerKilogram)
	a.watt_days_per_kilogramLazy = &watt_days_per_kilogram
	return watt_days_per_kilogram
}

// WattDayPerTonne returns the value in WattDayPerTonne.
func (a *SpecificEnergy) WattDaysPerTonne() float64 {
	if a.watt_days_per_tonneLazy != nil {
		return *a.watt_days_per_tonneLazy
	}
	watt_days_per_tonne := a.convertFromBase(SpecificEnergyWattDayPerTonne)
	a.watt_days_per_tonneLazy = &watt_days_per_tonne
	return watt_days_per_tonne
}

// WattDayPerShortTon returns the value in WattDayPerShortTon.
func (a *SpecificEnergy) WattDaysPerShortTon() float64 {
	if a.watt_days_per_short_tonLazy != nil {
		return *a.watt_days_per_short_tonLazy
	}
	watt_days_per_short_ton := a.convertFromBase(SpecificEnergyWattDayPerShortTon)
	a.watt_days_per_short_tonLazy = &watt_days_per_short_ton
	return watt_days_per_short_ton
}

// WattHourPerPound returns the value in WattHourPerPound.
func (a *SpecificEnergy) WattHoursPerPound() float64 {
	if a.watt_hours_per_poundLazy != nil {
		return *a.watt_hours_per_poundLazy
	}
	watt_hours_per_pound := a.convertFromBase(SpecificEnergyWattHourPerPound)
	a.watt_hours_per_poundLazy = &watt_hours_per_pound
	return watt_hours_per_pound
}

// BtuPerPound returns the value in BtuPerPound.
func (a *SpecificEnergy) BtuPerPound() float64 {
	if a.btu_per_poundLazy != nil {
		return *a.btu_per_poundLazy
	}
	btu_per_pound := a.convertFromBase(SpecificEnergyBtuPerPound)
	a.btu_per_poundLazy = &btu_per_pound
	return btu_per_pound
}

// KilojoulePerKilogram returns the value in KilojoulePerKilogram.
func (a *SpecificEnergy) KilojoulesPerKilogram() float64 {
	if a.kilojoules_per_kilogramLazy != nil {
		return *a.kilojoules_per_kilogramLazy
	}
	kilojoules_per_kilogram := a.convertFromBase(SpecificEnergyKilojoulePerKilogram)
	a.kilojoules_per_kilogramLazy = &kilojoules_per_kilogram
	return kilojoules_per_kilogram
}

// MegajoulePerKilogram returns the value in MegajoulePerKilogram.
func (a *SpecificEnergy) MegajoulesPerKilogram() float64 {
	if a.megajoules_per_kilogramLazy != nil {
		return *a.megajoules_per_kilogramLazy
	}
	megajoules_per_kilogram := a.convertFromBase(SpecificEnergyMegajoulePerKilogram)
	a.megajoules_per_kilogramLazy = &megajoules_per_kilogram
	return megajoules_per_kilogram
}

// KilocaloriePerGram returns the value in KilocaloriePerGram.
func (a *SpecificEnergy) KilocaloriesPerGram() float64 {
	if a.kilocalories_per_gramLazy != nil {
		return *a.kilocalories_per_gramLazy
	}
	kilocalories_per_gram := a.convertFromBase(SpecificEnergyKilocaloriePerGram)
	a.kilocalories_per_gramLazy = &kilocalories_per_gram
	return kilocalories_per_gram
}

// KilowattHourPerKilogram returns the value in KilowattHourPerKilogram.
func (a *SpecificEnergy) KilowattHoursPerKilogram() float64 {
	if a.kilowatt_hours_per_kilogramLazy != nil {
		return *a.kilowatt_hours_per_kilogramLazy
	}
	kilowatt_hours_per_kilogram := a.convertFromBase(SpecificEnergyKilowattHourPerKilogram)
	a.kilowatt_hours_per_kilogramLazy = &kilowatt_hours_per_kilogram
	return kilowatt_hours_per_kilogram
}

// MegawattHourPerKilogram returns the value in MegawattHourPerKilogram.
func (a *SpecificEnergy) MegawattHoursPerKilogram() float64 {
	if a.megawatt_hours_per_kilogramLazy != nil {
		return *a.megawatt_hours_per_kilogramLazy
	}
	megawatt_hours_per_kilogram := a.convertFromBase(SpecificEnergyMegawattHourPerKilogram)
	a.megawatt_hours_per_kilogramLazy = &megawatt_hours_per_kilogram
	return megawatt_hours_per_kilogram
}

// GigawattHourPerKilogram returns the value in GigawattHourPerKilogram.
func (a *SpecificEnergy) GigawattHoursPerKilogram() float64 {
	if a.gigawatt_hours_per_kilogramLazy != nil {
		return *a.gigawatt_hours_per_kilogramLazy
	}
	gigawatt_hours_per_kilogram := a.convertFromBase(SpecificEnergyGigawattHourPerKilogram)
	a.gigawatt_hours_per_kilogramLazy = &gigawatt_hours_per_kilogram
	return gigawatt_hours_per_kilogram
}

// KilowattDayPerKilogram returns the value in KilowattDayPerKilogram.
func (a *SpecificEnergy) KilowattDaysPerKilogram() float64 {
	if a.kilowatt_days_per_kilogramLazy != nil {
		return *a.kilowatt_days_per_kilogramLazy
	}
	kilowatt_days_per_kilogram := a.convertFromBase(SpecificEnergyKilowattDayPerKilogram)
	a.kilowatt_days_per_kilogramLazy = &kilowatt_days_per_kilogram
	return kilowatt_days_per_kilogram
}

// MegawattDayPerKilogram returns the value in MegawattDayPerKilogram.
func (a *SpecificEnergy) MegawattDaysPerKilogram() float64 {
	if a.megawatt_days_per_kilogramLazy != nil {
		return *a.megawatt_days_per_kilogramLazy
	}
	megawatt_days_per_kilogram := a.convertFromBase(SpecificEnergyMegawattDayPerKilogram)
	a.megawatt_days_per_kilogramLazy = &megawatt_days_per_kilogram
	return megawatt_days_per_kilogram
}

// GigawattDayPerKilogram returns the value in GigawattDayPerKilogram.
func (a *SpecificEnergy) GigawattDaysPerKilogram() float64 {
	if a.gigawatt_days_per_kilogramLazy != nil {
		return *a.gigawatt_days_per_kilogramLazy
	}
	gigawatt_days_per_kilogram := a.convertFromBase(SpecificEnergyGigawattDayPerKilogram)
	a.gigawatt_days_per_kilogramLazy = &gigawatt_days_per_kilogram
	return gigawatt_days_per_kilogram
}

// TerawattDayPerKilogram returns the value in TerawattDayPerKilogram.
func (a *SpecificEnergy) TerawattDaysPerKilogram() float64 {
	if a.terawatt_days_per_kilogramLazy != nil {
		return *a.terawatt_days_per_kilogramLazy
	}
	terawatt_days_per_kilogram := a.convertFromBase(SpecificEnergyTerawattDayPerKilogram)
	a.terawatt_days_per_kilogramLazy = &terawatt_days_per_kilogram
	return terawatt_days_per_kilogram
}

// KilowattDayPerTonne returns the value in KilowattDayPerTonne.
func (a *SpecificEnergy) KilowattDaysPerTonne() float64 {
	if a.kilowatt_days_per_tonneLazy != nil {
		return *a.kilowatt_days_per_tonneLazy
	}
	kilowatt_days_per_tonne := a.convertFromBase(SpecificEnergyKilowattDayPerTonne)
	a.kilowatt_days_per_tonneLazy = &kilowatt_days_per_tonne
	return kilowatt_days_per_tonne
}

// MegawattDayPerTonne returns the value in MegawattDayPerTonne.
func (a *SpecificEnergy) MegawattDaysPerTonne() float64 {
	if a.megawatt_days_per_tonneLazy != nil {
		return *a.megawatt_days_per_tonneLazy
	}
	megawatt_days_per_tonne := a.convertFromBase(SpecificEnergyMegawattDayPerTonne)
	a.megawatt_days_per_tonneLazy = &megawatt_days_per_tonne
	return megawatt_days_per_tonne
}

// GigawattDayPerTonne returns the value in GigawattDayPerTonne.
func (a *SpecificEnergy) GigawattDaysPerTonne() float64 {
	if a.gigawatt_days_per_tonneLazy != nil {
		return *a.gigawatt_days_per_tonneLazy
	}
	gigawatt_days_per_tonne := a.convertFromBase(SpecificEnergyGigawattDayPerTonne)
	a.gigawatt_days_per_tonneLazy = &gigawatt_days_per_tonne
	return gigawatt_days_per_tonne
}

// TerawattDayPerTonne returns the value in TerawattDayPerTonne.
func (a *SpecificEnergy) TerawattDaysPerTonne() float64 {
	if a.terawatt_days_per_tonneLazy != nil {
		return *a.terawatt_days_per_tonneLazy
	}
	terawatt_days_per_tonne := a.convertFromBase(SpecificEnergyTerawattDayPerTonne)
	a.terawatt_days_per_tonneLazy = &terawatt_days_per_tonne
	return terawatt_days_per_tonne
}

// KilowattDayPerShortTon returns the value in KilowattDayPerShortTon.
func (a *SpecificEnergy) KilowattDaysPerShortTon() float64 {
	if a.kilowatt_days_per_short_tonLazy != nil {
		return *a.kilowatt_days_per_short_tonLazy
	}
	kilowatt_days_per_short_ton := a.convertFromBase(SpecificEnergyKilowattDayPerShortTon)
	a.kilowatt_days_per_short_tonLazy = &kilowatt_days_per_short_ton
	return kilowatt_days_per_short_ton
}

// MegawattDayPerShortTon returns the value in MegawattDayPerShortTon.
func (a *SpecificEnergy) MegawattDaysPerShortTon() float64 {
	if a.megawatt_days_per_short_tonLazy != nil {
		return *a.megawatt_days_per_short_tonLazy
	}
	megawatt_days_per_short_ton := a.convertFromBase(SpecificEnergyMegawattDayPerShortTon)
	a.megawatt_days_per_short_tonLazy = &megawatt_days_per_short_ton
	return megawatt_days_per_short_ton
}

// GigawattDayPerShortTon returns the value in GigawattDayPerShortTon.
func (a *SpecificEnergy) GigawattDaysPerShortTon() float64 {
	if a.gigawatt_days_per_short_tonLazy != nil {
		return *a.gigawatt_days_per_short_tonLazy
	}
	gigawatt_days_per_short_ton := a.convertFromBase(SpecificEnergyGigawattDayPerShortTon)
	a.gigawatt_days_per_short_tonLazy = &gigawatt_days_per_short_ton
	return gigawatt_days_per_short_ton
}

// TerawattDayPerShortTon returns the value in TerawattDayPerShortTon.
func (a *SpecificEnergy) TerawattDaysPerShortTon() float64 {
	if a.terawatt_days_per_short_tonLazy != nil {
		return *a.terawatt_days_per_short_tonLazy
	}
	terawatt_days_per_short_ton := a.convertFromBase(SpecificEnergyTerawattDayPerShortTon)
	a.terawatt_days_per_short_tonLazy = &terawatt_days_per_short_ton
	return terawatt_days_per_short_ton
}

// KilowattHourPerPound returns the value in KilowattHourPerPound.
func (a *SpecificEnergy) KilowattHoursPerPound() float64 {
	if a.kilowatt_hours_per_poundLazy != nil {
		return *a.kilowatt_hours_per_poundLazy
	}
	kilowatt_hours_per_pound := a.convertFromBase(SpecificEnergyKilowattHourPerPound)
	a.kilowatt_hours_per_poundLazy = &kilowatt_hours_per_pound
	return kilowatt_hours_per_pound
}

// MegawattHourPerPound returns the value in MegawattHourPerPound.
func (a *SpecificEnergy) MegawattHoursPerPound() float64 {
	if a.megawatt_hours_per_poundLazy != nil {
		return *a.megawatt_hours_per_poundLazy
	}
	megawatt_hours_per_pound := a.convertFromBase(SpecificEnergyMegawattHourPerPound)
	a.megawatt_hours_per_poundLazy = &megawatt_hours_per_pound
	return megawatt_hours_per_pound
}

// GigawattHourPerPound returns the value in GigawattHourPerPound.
func (a *SpecificEnergy) GigawattHoursPerPound() float64 {
	if a.gigawatt_hours_per_poundLazy != nil {
		return *a.gigawatt_hours_per_poundLazy
	}
	gigawatt_hours_per_pound := a.convertFromBase(SpecificEnergyGigawattHourPerPound)
	a.gigawatt_hours_per_poundLazy = &gigawatt_hours_per_pound
	return gigawatt_hours_per_pound
}


// ToDto creates an SpecificEnergyDto representation.
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

// ToDtoJSON creates an SpecificEnergyDto representation.
func (a *SpecificEnergy) ToDtoJSON(holdInUnit *SpecificEnergyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts SpecificEnergy to a specific unit value.
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
		return 0
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
		return (value / 2326.000075362) 
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
		return (value * 2326.000075362) 
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

// Implement the String() method for AngleDto
func (a SpecificEnergy) String() string {
	return a.ToString(SpecificEnergyJoulePerKilogram, 2)
}

// ToString formats the SpecificEnergy to string.
// fractionalDigits -1 for not mention
func (a *SpecificEnergy) ToString(unit SpecificEnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *SpecificEnergy) getUnitAbbreviation(unit SpecificEnergyUnits) string {
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

// Check if the given SpecificEnergy are equals to the current SpecificEnergy
func (a *SpecificEnergy) Equals(other *SpecificEnergy) bool {
	return a.value == other.BaseValue()
}

// Check if the given SpecificEnergy are equals to the current SpecificEnergy
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

// Add the given SpecificEnergy to the current SpecificEnergy.
func (a *SpecificEnergy) Add(other *SpecificEnergy) *SpecificEnergy {
	return &SpecificEnergy{value: a.value + other.BaseValue()}
}

// Subtract the given SpecificEnergy to the current SpecificEnergy.
func (a *SpecificEnergy) Subtract(other *SpecificEnergy) *SpecificEnergy {
	return &SpecificEnergy{value: a.value - other.BaseValue()}
}

// Multiply the given SpecificEnergy to the current SpecificEnergy.
func (a *SpecificEnergy) Multiply(other *SpecificEnergy) *SpecificEnergy {
	return &SpecificEnergy{value: a.value * other.BaseValue()}
}

// Divide the given SpecificEnergy to the current SpecificEnergy.
func (a *SpecificEnergy) Divide(other *SpecificEnergy) *SpecificEnergy {
	return &SpecificEnergy{value: a.value / other.BaseValue()}
}