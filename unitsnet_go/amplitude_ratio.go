package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// AmplitudeRatioUnits enumeration
type AmplitudeRatioUnits string

const (
    
        // 
        AmplitudeRatioDecibelVolt AmplitudeRatioUnits = "DecibelVolt"
        // 
        AmplitudeRatioDecibelMicrovolt AmplitudeRatioUnits = "DecibelMicrovolt"
        // 
        AmplitudeRatioDecibelMillivolt AmplitudeRatioUnits = "DecibelMillivolt"
        // 
        AmplitudeRatioDecibelUnloaded AmplitudeRatioUnits = "DecibelUnloaded"
)

// AmplitudeRatioDto represents an AmplitudeRatio
type AmplitudeRatioDto struct {
	Value float64
	Unit  AmplitudeRatioUnits
}

// AmplitudeRatioDtoFactory struct to group related functions
type AmplitudeRatioDtoFactory struct{}

func (udf AmplitudeRatioDtoFactory) FromJSON(data []byte) (*AmplitudeRatioDto, error) {
	a := AmplitudeRatioDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a AmplitudeRatioDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// AmplitudeRatio struct
type AmplitudeRatio struct {
	value       float64
    
    decibel_voltsLazy *float64 
    decibel_microvoltsLazy *float64 
    decibel_millivoltsLazy *float64 
    decibels_unloadedLazy *float64 
}

// AmplitudeRatioFactory struct to group related functions
type AmplitudeRatioFactory struct{}

func (uf AmplitudeRatioFactory) CreateAmplitudeRatio(value float64, unit AmplitudeRatioUnits) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, unit)
}

func (uf AmplitudeRatioFactory) FromDto(dto AmplitudeRatioDto) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(dto.Value, dto.Unit)
}

func (uf AmplitudeRatioFactory) FromDtoJSON(data []byte) (*AmplitudeRatio, error) {
	unitDto, err := AmplitudeRatioDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return AmplitudeRatioFactory{}.FromDto(*unitDto)
}


// FromDecibelVolt creates a new AmplitudeRatio instance from DecibelVolt.
func (uf AmplitudeRatioFactory) FromDecibelVolts(value float64) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, AmplitudeRatioDecibelVolt)
}

// FromDecibelMicrovolt creates a new AmplitudeRatio instance from DecibelMicrovolt.
func (uf AmplitudeRatioFactory) FromDecibelMicrovolts(value float64) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, AmplitudeRatioDecibelMicrovolt)
}

// FromDecibelMillivolt creates a new AmplitudeRatio instance from DecibelMillivolt.
func (uf AmplitudeRatioFactory) FromDecibelMillivolts(value float64) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, AmplitudeRatioDecibelMillivolt)
}

// FromDecibelUnloaded creates a new AmplitudeRatio instance from DecibelUnloaded.
func (uf AmplitudeRatioFactory) FromDecibelsUnloaded(value float64) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, AmplitudeRatioDecibelUnloaded)
}




// newAmplitudeRatio creates a new AmplitudeRatio.
func newAmplitudeRatio(value float64, fromUnit AmplitudeRatioUnits) (*AmplitudeRatio, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &AmplitudeRatio{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AmplitudeRatio in DecibelVolt.
func (a *AmplitudeRatio) BaseValue() float64 {
	return a.value
}


// DecibelVolt returns the value in DecibelVolt.
func (a *AmplitudeRatio) DecibelVolts() float64 {
	if a.decibel_voltsLazy != nil {
		return *a.decibel_voltsLazy
	}
	decibel_volts := a.convertFromBase(AmplitudeRatioDecibelVolt)
	a.decibel_voltsLazy = &decibel_volts
	return decibel_volts
}

// DecibelMicrovolt returns the value in DecibelMicrovolt.
func (a *AmplitudeRatio) DecibelMicrovolts() float64 {
	if a.decibel_microvoltsLazy != nil {
		return *a.decibel_microvoltsLazy
	}
	decibel_microvolts := a.convertFromBase(AmplitudeRatioDecibelMicrovolt)
	a.decibel_microvoltsLazy = &decibel_microvolts
	return decibel_microvolts
}

// DecibelMillivolt returns the value in DecibelMillivolt.
func (a *AmplitudeRatio) DecibelMillivolts() float64 {
	if a.decibel_millivoltsLazy != nil {
		return *a.decibel_millivoltsLazy
	}
	decibel_millivolts := a.convertFromBase(AmplitudeRatioDecibelMillivolt)
	a.decibel_millivoltsLazy = &decibel_millivolts
	return decibel_millivolts
}

// DecibelUnloaded returns the value in DecibelUnloaded.
func (a *AmplitudeRatio) DecibelsUnloaded() float64 {
	if a.decibels_unloadedLazy != nil {
		return *a.decibels_unloadedLazy
	}
	decibels_unloaded := a.convertFromBase(AmplitudeRatioDecibelUnloaded)
	a.decibels_unloadedLazy = &decibels_unloaded
	return decibels_unloaded
}


// ToDto creates an AmplitudeRatioDto representation.
func (a *AmplitudeRatio) ToDto(holdInUnit *AmplitudeRatioUnits) AmplitudeRatioDto {
	if holdInUnit == nil {
		defaultUnit := AmplitudeRatioDecibelVolt // Default value
		holdInUnit = &defaultUnit
	}

	return AmplitudeRatioDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an AmplitudeRatioDto representation.
func (a *AmplitudeRatio) ToDtoJSON(holdInUnit *AmplitudeRatioUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts AmplitudeRatio to a specific unit value.
func (a *AmplitudeRatio) Convert(toUnit AmplitudeRatioUnits) float64 {
	switch toUnit { 
    case AmplitudeRatioDecibelVolt:
		return a.DecibelVolts()
    case AmplitudeRatioDecibelMicrovolt:
		return a.DecibelMicrovolts()
    case AmplitudeRatioDecibelMillivolt:
		return a.DecibelMillivolts()
    case AmplitudeRatioDecibelUnloaded:
		return a.DecibelsUnloaded()
	default:
		return 0
	}
}

func (a *AmplitudeRatio) convertFromBase(toUnit AmplitudeRatioUnits) float64 {
    value := a.value
	switch toUnit { 
	case AmplitudeRatioDecibelVolt:
		return (value) 
	case AmplitudeRatioDecibelMicrovolt:
		return (value + 120) 
	case AmplitudeRatioDecibelMillivolt:
		return (value + 60) 
	case AmplitudeRatioDecibelUnloaded:
		return (value + 2.218487499) 
	default:
		return math.NaN()
	}
}

func (a *AmplitudeRatio) convertToBase(value float64, fromUnit AmplitudeRatioUnits) float64 {
	switch fromUnit { 
	case AmplitudeRatioDecibelVolt:
		return (value) 
	case AmplitudeRatioDecibelMicrovolt:
		return (value - 120) 
	case AmplitudeRatioDecibelMillivolt:
		return (value - 60) 
	case AmplitudeRatioDecibelUnloaded:
		return (value - 2.218487499) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a AmplitudeRatio) String() string {
	return a.ToString(AmplitudeRatioDecibelVolt, 2)
}

// ToString formats the AmplitudeRatio to string.
// fractionalDigits -1 for not mention
func (a *AmplitudeRatio) ToString(unit AmplitudeRatioUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *AmplitudeRatio) getUnitAbbreviation(unit AmplitudeRatioUnits) string {
	switch unit { 
	case AmplitudeRatioDecibelVolt:
		return "dBV" 
	case AmplitudeRatioDecibelMicrovolt:
		return "dBµV" 
	case AmplitudeRatioDecibelMillivolt:
		return "dBmV" 
	case AmplitudeRatioDecibelUnloaded:
		return "dBu" 
	default:
		return ""
	}
}

// Check if the given AmplitudeRatio are equals to the current AmplitudeRatio
func (a *AmplitudeRatio) Equals(other *AmplitudeRatio) bool {
	return a.value == other.BaseValue()
}

// Check if the given AmplitudeRatio are equals to the current AmplitudeRatio
func (a *AmplitudeRatio) CompareTo(other *AmplitudeRatio) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given AmplitudeRatio to the current AmplitudeRatio.
func (a *AmplitudeRatio) Add(other *AmplitudeRatio) *AmplitudeRatio {
	return &AmplitudeRatio{value: a.value + other.BaseValue()}
}

// Subtract the given AmplitudeRatio to the current AmplitudeRatio.
func (a *AmplitudeRatio) Subtract(other *AmplitudeRatio) *AmplitudeRatio {
	return &AmplitudeRatio{value: a.value - other.BaseValue()}
}

// Multiply the given AmplitudeRatio to the current AmplitudeRatio.
func (a *AmplitudeRatio) Multiply(other *AmplitudeRatio) *AmplitudeRatio {
	return &AmplitudeRatio{value: a.value * other.BaseValue()}
}

// Divide the given AmplitudeRatio to the current AmplitudeRatio.
func (a *AmplitudeRatio) Divide(other *AmplitudeRatio) *AmplitudeRatio {
	return &AmplitudeRatio{value: a.value / other.BaseValue()}
}