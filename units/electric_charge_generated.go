// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricChargeUnits enumeration
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

// ElectricChargeDto represents an ElectricCharge
type ElectricChargeDto struct {
	Value float64
	Unit  ElectricChargeUnits
}

// ElectricChargeDtoFactory struct to group related functions
type ElectricChargeDtoFactory struct{}

func (udf ElectricChargeDtoFactory) FromJSON(data []byte) (*ElectricChargeDto, error) {
	a := ElectricChargeDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricChargeDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricCharge struct
type ElectricCharge struct {
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

// ElectricChargeFactory struct to group related functions
type ElectricChargeFactory struct{}

func (uf ElectricChargeFactory) CreateElectricCharge(value float64, unit ElectricChargeUnits) (*ElectricCharge, error) {
	return newElectricCharge(value, unit)
}

func (uf ElectricChargeFactory) FromDto(dto ElectricChargeDto) (*ElectricCharge, error) {
	return newElectricCharge(dto.Value, dto.Unit)
}

func (uf ElectricChargeFactory) FromDtoJSON(data []byte) (*ElectricCharge, error) {
	unitDto, err := ElectricChargeDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricChargeFactory{}.FromDto(*unitDto)
}


// FromCoulomb creates a new ElectricCharge instance from Coulomb.
func (uf ElectricChargeFactory) FromCoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeCoulomb)
}

// FromAmpereHour creates a new ElectricCharge instance from AmpereHour.
func (uf ElectricChargeFactory) FromAmpereHours(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeAmpereHour)
}

// FromPicocoulomb creates a new ElectricCharge instance from Picocoulomb.
func (uf ElectricChargeFactory) FromPicocoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargePicocoulomb)
}

// FromNanocoulomb creates a new ElectricCharge instance from Nanocoulomb.
func (uf ElectricChargeFactory) FromNanocoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeNanocoulomb)
}

// FromMicrocoulomb creates a new ElectricCharge instance from Microcoulomb.
func (uf ElectricChargeFactory) FromMicrocoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeMicrocoulomb)
}

// FromMillicoulomb creates a new ElectricCharge instance from Millicoulomb.
func (uf ElectricChargeFactory) FromMillicoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeMillicoulomb)
}

// FromKilocoulomb creates a new ElectricCharge instance from Kilocoulomb.
func (uf ElectricChargeFactory) FromKilocoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeKilocoulomb)
}

// FromMegacoulomb creates a new ElectricCharge instance from Megacoulomb.
func (uf ElectricChargeFactory) FromMegacoulombs(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeMegacoulomb)
}

// FromMilliampereHour creates a new ElectricCharge instance from MilliampereHour.
func (uf ElectricChargeFactory) FromMilliampereHours(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeMilliampereHour)
}

// FromKiloampereHour creates a new ElectricCharge instance from KiloampereHour.
func (uf ElectricChargeFactory) FromKiloampereHours(value float64) (*ElectricCharge, error) {
	return newElectricCharge(value, ElectricChargeKiloampereHour)
}

// FromMegaampereHour creates a new ElectricCharge instance from MegaampereHour.
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

// BaseValue returns the base value of ElectricCharge in Coulomb.
func (a *ElectricCharge) BaseValue() float64 {
	return a.value
}


// Coulomb returns the value in Coulomb.
func (a *ElectricCharge) Coulombs() float64 {
	if a.coulombsLazy != nil {
		return *a.coulombsLazy
	}
	coulombs := a.convertFromBase(ElectricChargeCoulomb)
	a.coulombsLazy = &coulombs
	return coulombs
}

// AmpereHour returns the value in AmpereHour.
func (a *ElectricCharge) AmpereHours() float64 {
	if a.ampere_hoursLazy != nil {
		return *a.ampere_hoursLazy
	}
	ampere_hours := a.convertFromBase(ElectricChargeAmpereHour)
	a.ampere_hoursLazy = &ampere_hours
	return ampere_hours
}

// Picocoulomb returns the value in Picocoulomb.
func (a *ElectricCharge) Picocoulombs() float64 {
	if a.picocoulombsLazy != nil {
		return *a.picocoulombsLazy
	}
	picocoulombs := a.convertFromBase(ElectricChargePicocoulomb)
	a.picocoulombsLazy = &picocoulombs
	return picocoulombs
}

// Nanocoulomb returns the value in Nanocoulomb.
func (a *ElectricCharge) Nanocoulombs() float64 {
	if a.nanocoulombsLazy != nil {
		return *a.nanocoulombsLazy
	}
	nanocoulombs := a.convertFromBase(ElectricChargeNanocoulomb)
	a.nanocoulombsLazy = &nanocoulombs
	return nanocoulombs
}

// Microcoulomb returns the value in Microcoulomb.
func (a *ElectricCharge) Microcoulombs() float64 {
	if a.microcoulombsLazy != nil {
		return *a.microcoulombsLazy
	}
	microcoulombs := a.convertFromBase(ElectricChargeMicrocoulomb)
	a.microcoulombsLazy = &microcoulombs
	return microcoulombs
}

// Millicoulomb returns the value in Millicoulomb.
func (a *ElectricCharge) Millicoulombs() float64 {
	if a.millicoulombsLazy != nil {
		return *a.millicoulombsLazy
	}
	millicoulombs := a.convertFromBase(ElectricChargeMillicoulomb)
	a.millicoulombsLazy = &millicoulombs
	return millicoulombs
}

// Kilocoulomb returns the value in Kilocoulomb.
func (a *ElectricCharge) Kilocoulombs() float64 {
	if a.kilocoulombsLazy != nil {
		return *a.kilocoulombsLazy
	}
	kilocoulombs := a.convertFromBase(ElectricChargeKilocoulomb)
	a.kilocoulombsLazy = &kilocoulombs
	return kilocoulombs
}

// Megacoulomb returns the value in Megacoulomb.
func (a *ElectricCharge) Megacoulombs() float64 {
	if a.megacoulombsLazy != nil {
		return *a.megacoulombsLazy
	}
	megacoulombs := a.convertFromBase(ElectricChargeMegacoulomb)
	a.megacoulombsLazy = &megacoulombs
	return megacoulombs
}

// MilliampereHour returns the value in MilliampereHour.
func (a *ElectricCharge) MilliampereHours() float64 {
	if a.milliampere_hoursLazy != nil {
		return *a.milliampere_hoursLazy
	}
	milliampere_hours := a.convertFromBase(ElectricChargeMilliampereHour)
	a.milliampere_hoursLazy = &milliampere_hours
	return milliampere_hours
}

// KiloampereHour returns the value in KiloampereHour.
func (a *ElectricCharge) KiloampereHours() float64 {
	if a.kiloampere_hoursLazy != nil {
		return *a.kiloampere_hoursLazy
	}
	kiloampere_hours := a.convertFromBase(ElectricChargeKiloampereHour)
	a.kiloampere_hoursLazy = &kiloampere_hours
	return kiloampere_hours
}

// MegaampereHour returns the value in MegaampereHour.
func (a *ElectricCharge) MegaampereHours() float64 {
	if a.megaampere_hoursLazy != nil {
		return *a.megaampere_hoursLazy
	}
	megaampere_hours := a.convertFromBase(ElectricChargeMegaampereHour)
	a.megaampere_hoursLazy = &megaampere_hours
	return megaampere_hours
}


// ToDto creates an ElectricChargeDto representation.
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

// ToDtoJSON creates an ElectricChargeDto representation.
func (a *ElectricCharge) ToDtoJSON(holdInUnit *ElectricChargeUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricCharge to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a ElectricCharge) String() string {
	return a.ToString(ElectricChargeCoulomb, 2)
}

// ToString formats the ElectricCharge to string.
// fractionalDigits -1 for not mention
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

// Check if the given ElectricCharge are equals to the current ElectricCharge
func (a *ElectricCharge) Equals(other *ElectricCharge) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricCharge are equals to the current ElectricCharge
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

// Add the given ElectricCharge to the current ElectricCharge.
func (a *ElectricCharge) Add(other *ElectricCharge) *ElectricCharge {
	return &ElectricCharge{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricCharge to the current ElectricCharge.
func (a *ElectricCharge) Subtract(other *ElectricCharge) *ElectricCharge {
	return &ElectricCharge{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricCharge to the current ElectricCharge.
func (a *ElectricCharge) Multiply(other *ElectricCharge) *ElectricCharge {
	return &ElectricCharge{value: a.value * other.BaseValue()}
}

// Divide the given ElectricCharge to the current ElectricCharge.
func (a *ElectricCharge) Divide(other *ElectricCharge) *ElectricCharge {
	return &ElectricCharge{value: a.value / other.BaseValue()}
}