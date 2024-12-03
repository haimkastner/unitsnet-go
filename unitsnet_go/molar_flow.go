package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MolarFlowUnits enumeration
type MolarFlowUnits string

const (
    
        // 
        MolarFlowMolePerSecond MolarFlowUnits = "MolePerSecond"
        // 
        MolarFlowMolePerMinute MolarFlowUnits = "MolePerMinute"
        // 
        MolarFlowMolePerHour MolarFlowUnits = "MolePerHour"
        // 
        MolarFlowPoundMolePerSecond MolarFlowUnits = "PoundMolePerSecond"
        // 
        MolarFlowPoundMolePerMinute MolarFlowUnits = "PoundMolePerMinute"
        // 
        MolarFlowPoundMolePerHour MolarFlowUnits = "PoundMolePerHour"
        // 
        MolarFlowKilomolePerSecond MolarFlowUnits = "KilomolePerSecond"
        // 
        MolarFlowKilomolePerMinute MolarFlowUnits = "KilomolePerMinute"
        // 
        MolarFlowKilomolePerHour MolarFlowUnits = "KilomolePerHour"
)

// MolarFlowDto represents an MolarFlow
type MolarFlowDto struct {
	Value float64
	Unit  MolarFlowUnits
}

// MolarFlowDtoFactory struct to group related functions
type MolarFlowDtoFactory struct{}

func (udf MolarFlowDtoFactory) FromJSON(data []byte) (*MolarFlowDto, error) {
	a := MolarFlowDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MolarFlowDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MolarFlow struct
type MolarFlow struct {
	value       float64
    
    moles_per_secondLazy *float64 
    moles_per_minuteLazy *float64 
    moles_per_hourLazy *float64 
    pound_moles_per_secondLazy *float64 
    pound_moles_per_minuteLazy *float64 
    pound_moles_per_hourLazy *float64 
    kilomoles_per_secondLazy *float64 
    kilomoles_per_minuteLazy *float64 
    kilomoles_per_hourLazy *float64 
}

// MolarFlowFactory struct to group related functions
type MolarFlowFactory struct{}

func (uf MolarFlowFactory) CreateMolarFlow(value float64, unit MolarFlowUnits) (*MolarFlow, error) {
	return newMolarFlow(value, unit)
}

func (uf MolarFlowFactory) FromDto(dto MolarFlowDto) (*MolarFlow, error) {
	return newMolarFlow(dto.Value, dto.Unit)
}

func (uf MolarFlowFactory) FromDtoJSON(data []byte) (*MolarFlow, error) {
	unitDto, err := MolarFlowDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MolarFlowFactory{}.FromDto(*unitDto)
}


// FromMolePerSecond creates a new MolarFlow instance from MolePerSecond.
func (uf MolarFlowFactory) FromMolesPerSecond(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowMolePerSecond)
}

// FromMolePerMinute creates a new MolarFlow instance from MolePerMinute.
func (uf MolarFlowFactory) FromMolesPerMinute(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowMolePerMinute)
}

// FromMolePerHour creates a new MolarFlow instance from MolePerHour.
func (uf MolarFlowFactory) FromMolesPerHour(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowMolePerHour)
}

// FromPoundMolePerSecond creates a new MolarFlow instance from PoundMolePerSecond.
func (uf MolarFlowFactory) FromPoundMolesPerSecond(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowPoundMolePerSecond)
}

// FromPoundMolePerMinute creates a new MolarFlow instance from PoundMolePerMinute.
func (uf MolarFlowFactory) FromPoundMolesPerMinute(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowPoundMolePerMinute)
}

// FromPoundMolePerHour creates a new MolarFlow instance from PoundMolePerHour.
func (uf MolarFlowFactory) FromPoundMolesPerHour(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowPoundMolePerHour)
}

// FromKilomolePerSecond creates a new MolarFlow instance from KilomolePerSecond.
func (uf MolarFlowFactory) FromKilomolesPerSecond(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowKilomolePerSecond)
}

// FromKilomolePerMinute creates a new MolarFlow instance from KilomolePerMinute.
func (uf MolarFlowFactory) FromKilomolesPerMinute(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowKilomolePerMinute)
}

// FromKilomolePerHour creates a new MolarFlow instance from KilomolePerHour.
func (uf MolarFlowFactory) FromKilomolesPerHour(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowKilomolePerHour)
}




// newMolarFlow creates a new MolarFlow.
func newMolarFlow(value float64, fromUnit MolarFlowUnits) (*MolarFlow, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MolarFlow{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MolarFlow in MolePerSecond.
func (a *MolarFlow) BaseValue() float64 {
	return a.value
}


// MolePerSecond returns the value in MolePerSecond.
func (a *MolarFlow) MolesPerSecond() float64 {
	if a.moles_per_secondLazy != nil {
		return *a.moles_per_secondLazy
	}
	moles_per_second := a.convertFromBase(MolarFlowMolePerSecond)
	a.moles_per_secondLazy = &moles_per_second
	return moles_per_second
}

// MolePerMinute returns the value in MolePerMinute.
func (a *MolarFlow) MolesPerMinute() float64 {
	if a.moles_per_minuteLazy != nil {
		return *a.moles_per_minuteLazy
	}
	moles_per_minute := a.convertFromBase(MolarFlowMolePerMinute)
	a.moles_per_minuteLazy = &moles_per_minute
	return moles_per_minute
}

// MolePerHour returns the value in MolePerHour.
func (a *MolarFlow) MolesPerHour() float64 {
	if a.moles_per_hourLazy != nil {
		return *a.moles_per_hourLazy
	}
	moles_per_hour := a.convertFromBase(MolarFlowMolePerHour)
	a.moles_per_hourLazy = &moles_per_hour
	return moles_per_hour
}

// PoundMolePerSecond returns the value in PoundMolePerSecond.
func (a *MolarFlow) PoundMolesPerSecond() float64 {
	if a.pound_moles_per_secondLazy != nil {
		return *a.pound_moles_per_secondLazy
	}
	pound_moles_per_second := a.convertFromBase(MolarFlowPoundMolePerSecond)
	a.pound_moles_per_secondLazy = &pound_moles_per_second
	return pound_moles_per_second
}

// PoundMolePerMinute returns the value in PoundMolePerMinute.
func (a *MolarFlow) PoundMolesPerMinute() float64 {
	if a.pound_moles_per_minuteLazy != nil {
		return *a.pound_moles_per_minuteLazy
	}
	pound_moles_per_minute := a.convertFromBase(MolarFlowPoundMolePerMinute)
	a.pound_moles_per_minuteLazy = &pound_moles_per_minute
	return pound_moles_per_minute
}

// PoundMolePerHour returns the value in PoundMolePerHour.
func (a *MolarFlow) PoundMolesPerHour() float64 {
	if a.pound_moles_per_hourLazy != nil {
		return *a.pound_moles_per_hourLazy
	}
	pound_moles_per_hour := a.convertFromBase(MolarFlowPoundMolePerHour)
	a.pound_moles_per_hourLazy = &pound_moles_per_hour
	return pound_moles_per_hour
}

// KilomolePerSecond returns the value in KilomolePerSecond.
func (a *MolarFlow) KilomolesPerSecond() float64 {
	if a.kilomoles_per_secondLazy != nil {
		return *a.kilomoles_per_secondLazy
	}
	kilomoles_per_second := a.convertFromBase(MolarFlowKilomolePerSecond)
	a.kilomoles_per_secondLazy = &kilomoles_per_second
	return kilomoles_per_second
}

// KilomolePerMinute returns the value in KilomolePerMinute.
func (a *MolarFlow) KilomolesPerMinute() float64 {
	if a.kilomoles_per_minuteLazy != nil {
		return *a.kilomoles_per_minuteLazy
	}
	kilomoles_per_minute := a.convertFromBase(MolarFlowKilomolePerMinute)
	a.kilomoles_per_minuteLazy = &kilomoles_per_minute
	return kilomoles_per_minute
}

// KilomolePerHour returns the value in KilomolePerHour.
func (a *MolarFlow) KilomolesPerHour() float64 {
	if a.kilomoles_per_hourLazy != nil {
		return *a.kilomoles_per_hourLazy
	}
	kilomoles_per_hour := a.convertFromBase(MolarFlowKilomolePerHour)
	a.kilomoles_per_hourLazy = &kilomoles_per_hour
	return kilomoles_per_hour
}


// ToDto creates an MolarFlowDto representation.
func (a *MolarFlow) ToDto(holdInUnit *MolarFlowUnits) MolarFlowDto {
	if holdInUnit == nil {
		defaultUnit := MolarFlowMolePerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return MolarFlowDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an MolarFlowDto representation.
func (a *MolarFlow) ToDtoJSON(holdInUnit *MolarFlowUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MolarFlow to a specific unit value.
func (a *MolarFlow) Convert(toUnit MolarFlowUnits) float64 {
	switch toUnit { 
    case MolarFlowMolePerSecond:
		return a.MolesPerSecond()
    case MolarFlowMolePerMinute:
		return a.MolesPerMinute()
    case MolarFlowMolePerHour:
		return a.MolesPerHour()
    case MolarFlowPoundMolePerSecond:
		return a.PoundMolesPerSecond()
    case MolarFlowPoundMolePerMinute:
		return a.PoundMolesPerMinute()
    case MolarFlowPoundMolePerHour:
		return a.PoundMolesPerHour()
    case MolarFlowKilomolePerSecond:
		return a.KilomolesPerSecond()
    case MolarFlowKilomolePerMinute:
		return a.KilomolesPerMinute()
    case MolarFlowKilomolePerHour:
		return a.KilomolesPerHour()
	default:
		return 0
	}
}

func (a *MolarFlow) convertFromBase(toUnit MolarFlowUnits) float64 {
    value := a.value
	switch toUnit { 
	case MolarFlowMolePerSecond:
		return (value) 
	case MolarFlowMolePerMinute:
		return (value * 60) 
	case MolarFlowMolePerHour:
		return (value * 3600) 
	case MolarFlowPoundMolePerSecond:
		return (value / 453.59237) 
	case MolarFlowPoundMolePerMinute:
		return ((value / 453.59237) * 60) 
	case MolarFlowPoundMolePerHour:
		return ((value / 453.59237) * 3600) 
	case MolarFlowKilomolePerSecond:
		return ((value) / 1000.0) 
	case MolarFlowKilomolePerMinute:
		return ((value * 60) / 1000.0) 
	case MolarFlowKilomolePerHour:
		return ((value * 3600) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *MolarFlow) convertToBase(value float64, fromUnit MolarFlowUnits) float64 {
	switch fromUnit { 
	case MolarFlowMolePerSecond:
		return (value) 
	case MolarFlowMolePerMinute:
		return (value / 60) 
	case MolarFlowMolePerHour:
		return (value / 3600) 
	case MolarFlowPoundMolePerSecond:
		return (value * 453.59237) 
	case MolarFlowPoundMolePerMinute:
		return ((value * 453.59237) / 60) 
	case MolarFlowPoundMolePerHour:
		return ((value * 453.59237) / 3600) 
	case MolarFlowKilomolePerSecond:
		return ((value) * 1000.0) 
	case MolarFlowKilomolePerMinute:
		return ((value / 60) * 1000.0) 
	case MolarFlowKilomolePerHour:
		return ((value / 3600) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a MolarFlow) String() string {
	return a.ToString(MolarFlowMolePerSecond, 2)
}

// ToString formats the MolarFlow to string.
// fractionalDigits -1 for not mention
func (a *MolarFlow) ToString(unit MolarFlowUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MolarFlow) getUnitAbbreviation(unit MolarFlowUnits) string {
	switch unit { 
	case MolarFlowMolePerSecond:
		return "mol/s" 
	case MolarFlowMolePerMinute:
		return "mol/min" 
	case MolarFlowMolePerHour:
		return "kmol/h" 
	case MolarFlowPoundMolePerSecond:
		return "lbmol/s" 
	case MolarFlowPoundMolePerMinute:
		return "lbmol/min" 
	case MolarFlowPoundMolePerHour:
		return "lbmol/h" 
	case MolarFlowKilomolePerSecond:
		return "kmol/s" 
	case MolarFlowKilomolePerMinute:
		return "kmol/min" 
	case MolarFlowKilomolePerHour:
		return "kkmol/h" 
	default:
		return ""
	}
}

// Check if the given MolarFlow are equals to the current MolarFlow
func (a *MolarFlow) Equals(other *MolarFlow) bool {
	return a.value == other.BaseValue()
}

// Check if the given MolarFlow are equals to the current MolarFlow
func (a *MolarFlow) CompareTo(other *MolarFlow) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given MolarFlow to the current MolarFlow.
func (a *MolarFlow) Add(other *MolarFlow) *MolarFlow {
	return &MolarFlow{value: a.value + other.BaseValue()}
}

// Subtract the given MolarFlow to the current MolarFlow.
func (a *MolarFlow) Subtract(other *MolarFlow) *MolarFlow {
	return &MolarFlow{value: a.value - other.BaseValue()}
}

// Multiply the given MolarFlow to the current MolarFlow.
func (a *MolarFlow) Multiply(other *MolarFlow) *MolarFlow {
	return &MolarFlow{value: a.value * other.BaseValue()}
}

// Divide the given MolarFlow to the current MolarFlow.
func (a *MolarFlow) Divide(other *MolarFlow) *MolarFlow {
	return &MolarFlow{value: a.value / other.BaseValue()}
}