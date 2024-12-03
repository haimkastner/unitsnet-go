// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RotationalStiffnessUnits enumeration
type RotationalStiffnessUnits string

const (
    
        // 
        RotationalStiffnessNewtonMeterPerRadian RotationalStiffnessUnits = "NewtonMeterPerRadian"
        // 
        RotationalStiffnessPoundForceFootPerDegrees RotationalStiffnessUnits = "PoundForceFootPerDegrees"
        // 
        RotationalStiffnessKilopoundForceFootPerDegrees RotationalStiffnessUnits = "KilopoundForceFootPerDegrees"
        // 
        RotationalStiffnessNewtonMillimeterPerDegree RotationalStiffnessUnits = "NewtonMillimeterPerDegree"
        // 
        RotationalStiffnessNewtonMeterPerDegree RotationalStiffnessUnits = "NewtonMeterPerDegree"
        // 
        RotationalStiffnessNewtonMillimeterPerRadian RotationalStiffnessUnits = "NewtonMillimeterPerRadian"
        // 
        RotationalStiffnessPoundForceFeetPerRadian RotationalStiffnessUnits = "PoundForceFeetPerRadian"
        // 
        RotationalStiffnessKilonewtonMeterPerRadian RotationalStiffnessUnits = "KilonewtonMeterPerRadian"
        // 
        RotationalStiffnessMeganewtonMeterPerRadian RotationalStiffnessUnits = "MeganewtonMeterPerRadian"
        // 
        RotationalStiffnessNanonewtonMillimeterPerDegree RotationalStiffnessUnits = "NanonewtonMillimeterPerDegree"
        // 
        RotationalStiffnessMicronewtonMillimeterPerDegree RotationalStiffnessUnits = "MicronewtonMillimeterPerDegree"
        // 
        RotationalStiffnessMillinewtonMillimeterPerDegree RotationalStiffnessUnits = "MillinewtonMillimeterPerDegree"
        // 
        RotationalStiffnessCentinewtonMillimeterPerDegree RotationalStiffnessUnits = "CentinewtonMillimeterPerDegree"
        // 
        RotationalStiffnessDecinewtonMillimeterPerDegree RotationalStiffnessUnits = "DecinewtonMillimeterPerDegree"
        // 
        RotationalStiffnessDecanewtonMillimeterPerDegree RotationalStiffnessUnits = "DecanewtonMillimeterPerDegree"
        // 
        RotationalStiffnessKilonewtonMillimeterPerDegree RotationalStiffnessUnits = "KilonewtonMillimeterPerDegree"
        // 
        RotationalStiffnessMeganewtonMillimeterPerDegree RotationalStiffnessUnits = "MeganewtonMillimeterPerDegree"
        // 
        RotationalStiffnessNanonewtonMeterPerDegree RotationalStiffnessUnits = "NanonewtonMeterPerDegree"
        // 
        RotationalStiffnessMicronewtonMeterPerDegree RotationalStiffnessUnits = "MicronewtonMeterPerDegree"
        // 
        RotationalStiffnessMillinewtonMeterPerDegree RotationalStiffnessUnits = "MillinewtonMeterPerDegree"
        // 
        RotationalStiffnessCentinewtonMeterPerDegree RotationalStiffnessUnits = "CentinewtonMeterPerDegree"
        // 
        RotationalStiffnessDecinewtonMeterPerDegree RotationalStiffnessUnits = "DecinewtonMeterPerDegree"
        // 
        RotationalStiffnessDecanewtonMeterPerDegree RotationalStiffnessUnits = "DecanewtonMeterPerDegree"
        // 
        RotationalStiffnessKilonewtonMeterPerDegree RotationalStiffnessUnits = "KilonewtonMeterPerDegree"
        // 
        RotationalStiffnessMeganewtonMeterPerDegree RotationalStiffnessUnits = "MeganewtonMeterPerDegree"
        // 
        RotationalStiffnessNanonewtonMillimeterPerRadian RotationalStiffnessUnits = "NanonewtonMillimeterPerRadian"
        // 
        RotationalStiffnessMicronewtonMillimeterPerRadian RotationalStiffnessUnits = "MicronewtonMillimeterPerRadian"
        // 
        RotationalStiffnessMillinewtonMillimeterPerRadian RotationalStiffnessUnits = "MillinewtonMillimeterPerRadian"
        // 
        RotationalStiffnessCentinewtonMillimeterPerRadian RotationalStiffnessUnits = "CentinewtonMillimeterPerRadian"
        // 
        RotationalStiffnessDecinewtonMillimeterPerRadian RotationalStiffnessUnits = "DecinewtonMillimeterPerRadian"
        // 
        RotationalStiffnessDecanewtonMillimeterPerRadian RotationalStiffnessUnits = "DecanewtonMillimeterPerRadian"
        // 
        RotationalStiffnessKilonewtonMillimeterPerRadian RotationalStiffnessUnits = "KilonewtonMillimeterPerRadian"
        // 
        RotationalStiffnessMeganewtonMillimeterPerRadian RotationalStiffnessUnits = "MeganewtonMillimeterPerRadian"
)

// RotationalStiffnessDto represents an RotationalStiffness
type RotationalStiffnessDto struct {
	Value float64
	Unit  RotationalStiffnessUnits
}

// RotationalStiffnessDtoFactory struct to group related functions
type RotationalStiffnessDtoFactory struct{}

func (udf RotationalStiffnessDtoFactory) FromJSON(data []byte) (*RotationalStiffnessDto, error) {
	a := RotationalStiffnessDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RotationalStiffnessDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// RotationalStiffness struct
type RotationalStiffness struct {
	value       float64
    
    newton_meters_per_radianLazy *float64 
    pound_force_feet_per_degreesLazy *float64 
    kilopound_force_feet_per_degreesLazy *float64 
    newton_millimeters_per_degreeLazy *float64 
    newton_meters_per_degreeLazy *float64 
    newton_millimeters_per_radianLazy *float64 
    pound_force_feet_per_radianLazy *float64 
    kilonewton_meters_per_radianLazy *float64 
    meganewton_meters_per_radianLazy *float64 
    nanonewton_millimeters_per_degreeLazy *float64 
    micronewton_millimeters_per_degreeLazy *float64 
    millinewton_millimeters_per_degreeLazy *float64 
    centinewton_millimeters_per_degreeLazy *float64 
    decinewton_millimeters_per_degreeLazy *float64 
    decanewton_millimeters_per_degreeLazy *float64 
    kilonewton_millimeters_per_degreeLazy *float64 
    meganewton_millimeters_per_degreeLazy *float64 
    nanonewton_meters_per_degreeLazy *float64 
    micronewton_meters_per_degreeLazy *float64 
    millinewton_meters_per_degreeLazy *float64 
    centinewton_meters_per_degreeLazy *float64 
    decinewton_meters_per_degreeLazy *float64 
    decanewton_meters_per_degreeLazy *float64 
    kilonewton_meters_per_degreeLazy *float64 
    meganewton_meters_per_degreeLazy *float64 
    nanonewton_millimeters_per_radianLazy *float64 
    micronewton_millimeters_per_radianLazy *float64 
    millinewton_millimeters_per_radianLazy *float64 
    centinewton_millimeters_per_radianLazy *float64 
    decinewton_millimeters_per_radianLazy *float64 
    decanewton_millimeters_per_radianLazy *float64 
    kilonewton_millimeters_per_radianLazy *float64 
    meganewton_millimeters_per_radianLazy *float64 
}

// RotationalStiffnessFactory struct to group related functions
type RotationalStiffnessFactory struct{}

func (uf RotationalStiffnessFactory) CreateRotationalStiffness(value float64, unit RotationalStiffnessUnits) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, unit)
}

func (uf RotationalStiffnessFactory) FromDto(dto RotationalStiffnessDto) (*RotationalStiffness, error) {
	return newRotationalStiffness(dto.Value, dto.Unit)
}

func (uf RotationalStiffnessFactory) FromDtoJSON(data []byte) (*RotationalStiffness, error) {
	unitDto, err := RotationalStiffnessDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RotationalStiffnessFactory{}.FromDto(*unitDto)
}


// FromNewtonMeterPerRadian creates a new RotationalStiffness instance from NewtonMeterPerRadian.
func (uf RotationalStiffnessFactory) FromNewtonMetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNewtonMeterPerRadian)
}

// FromPoundForceFootPerDegrees creates a new RotationalStiffness instance from PoundForceFootPerDegrees.
func (uf RotationalStiffnessFactory) FromPoundForceFeetPerDegrees(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessPoundForceFootPerDegrees)
}

// FromKilopoundForceFootPerDegrees creates a new RotationalStiffness instance from KilopoundForceFootPerDegrees.
func (uf RotationalStiffnessFactory) FromKilopoundForceFeetPerDegrees(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilopoundForceFootPerDegrees)
}

// FromNewtonMillimeterPerDegree creates a new RotationalStiffness instance from NewtonMillimeterPerDegree.
func (uf RotationalStiffnessFactory) FromNewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNewtonMillimeterPerDegree)
}

// FromNewtonMeterPerDegree creates a new RotationalStiffness instance from NewtonMeterPerDegree.
func (uf RotationalStiffnessFactory) FromNewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNewtonMeterPerDegree)
}

// FromNewtonMillimeterPerRadian creates a new RotationalStiffness instance from NewtonMillimeterPerRadian.
func (uf RotationalStiffnessFactory) FromNewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNewtonMillimeterPerRadian)
}

// FromPoundForceFeetPerRadian creates a new RotationalStiffness instance from PoundForceFeetPerRadian.
func (uf RotationalStiffnessFactory) FromPoundForceFeetPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessPoundForceFeetPerRadian)
}

// FromKilonewtonMeterPerRadian creates a new RotationalStiffness instance from KilonewtonMeterPerRadian.
func (uf RotationalStiffnessFactory) FromKilonewtonMetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilonewtonMeterPerRadian)
}

// FromMeganewtonMeterPerRadian creates a new RotationalStiffness instance from MeganewtonMeterPerRadian.
func (uf RotationalStiffnessFactory) FromMeganewtonMetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMeganewtonMeterPerRadian)
}

// FromNanonewtonMillimeterPerDegree creates a new RotationalStiffness instance from NanonewtonMillimeterPerDegree.
func (uf RotationalStiffnessFactory) FromNanonewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNanonewtonMillimeterPerDegree)
}

// FromMicronewtonMillimeterPerDegree creates a new RotationalStiffness instance from MicronewtonMillimeterPerDegree.
func (uf RotationalStiffnessFactory) FromMicronewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMicronewtonMillimeterPerDegree)
}

// FromMillinewtonMillimeterPerDegree creates a new RotationalStiffness instance from MillinewtonMillimeterPerDegree.
func (uf RotationalStiffnessFactory) FromMillinewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMillinewtonMillimeterPerDegree)
}

// FromCentinewtonMillimeterPerDegree creates a new RotationalStiffness instance from CentinewtonMillimeterPerDegree.
func (uf RotationalStiffnessFactory) FromCentinewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessCentinewtonMillimeterPerDegree)
}

// FromDecinewtonMillimeterPerDegree creates a new RotationalStiffness instance from DecinewtonMillimeterPerDegree.
func (uf RotationalStiffnessFactory) FromDecinewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecinewtonMillimeterPerDegree)
}

// FromDecanewtonMillimeterPerDegree creates a new RotationalStiffness instance from DecanewtonMillimeterPerDegree.
func (uf RotationalStiffnessFactory) FromDecanewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecanewtonMillimeterPerDegree)
}

// FromKilonewtonMillimeterPerDegree creates a new RotationalStiffness instance from KilonewtonMillimeterPerDegree.
func (uf RotationalStiffnessFactory) FromKilonewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilonewtonMillimeterPerDegree)
}

// FromMeganewtonMillimeterPerDegree creates a new RotationalStiffness instance from MeganewtonMillimeterPerDegree.
func (uf RotationalStiffnessFactory) FromMeganewtonMillimetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMeganewtonMillimeterPerDegree)
}

// FromNanonewtonMeterPerDegree creates a new RotationalStiffness instance from NanonewtonMeterPerDegree.
func (uf RotationalStiffnessFactory) FromNanonewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNanonewtonMeterPerDegree)
}

// FromMicronewtonMeterPerDegree creates a new RotationalStiffness instance from MicronewtonMeterPerDegree.
func (uf RotationalStiffnessFactory) FromMicronewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMicronewtonMeterPerDegree)
}

// FromMillinewtonMeterPerDegree creates a new RotationalStiffness instance from MillinewtonMeterPerDegree.
func (uf RotationalStiffnessFactory) FromMillinewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMillinewtonMeterPerDegree)
}

// FromCentinewtonMeterPerDegree creates a new RotationalStiffness instance from CentinewtonMeterPerDegree.
func (uf RotationalStiffnessFactory) FromCentinewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessCentinewtonMeterPerDegree)
}

// FromDecinewtonMeterPerDegree creates a new RotationalStiffness instance from DecinewtonMeterPerDegree.
func (uf RotationalStiffnessFactory) FromDecinewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecinewtonMeterPerDegree)
}

// FromDecanewtonMeterPerDegree creates a new RotationalStiffness instance from DecanewtonMeterPerDegree.
func (uf RotationalStiffnessFactory) FromDecanewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecanewtonMeterPerDegree)
}

// FromKilonewtonMeterPerDegree creates a new RotationalStiffness instance from KilonewtonMeterPerDegree.
func (uf RotationalStiffnessFactory) FromKilonewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilonewtonMeterPerDegree)
}

// FromMeganewtonMeterPerDegree creates a new RotationalStiffness instance from MeganewtonMeterPerDegree.
func (uf RotationalStiffnessFactory) FromMeganewtonMetersPerDegree(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMeganewtonMeterPerDegree)
}

// FromNanonewtonMillimeterPerRadian creates a new RotationalStiffness instance from NanonewtonMillimeterPerRadian.
func (uf RotationalStiffnessFactory) FromNanonewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessNanonewtonMillimeterPerRadian)
}

// FromMicronewtonMillimeterPerRadian creates a new RotationalStiffness instance from MicronewtonMillimeterPerRadian.
func (uf RotationalStiffnessFactory) FromMicronewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMicronewtonMillimeterPerRadian)
}

// FromMillinewtonMillimeterPerRadian creates a new RotationalStiffness instance from MillinewtonMillimeterPerRadian.
func (uf RotationalStiffnessFactory) FromMillinewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMillinewtonMillimeterPerRadian)
}

// FromCentinewtonMillimeterPerRadian creates a new RotationalStiffness instance from CentinewtonMillimeterPerRadian.
func (uf RotationalStiffnessFactory) FromCentinewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessCentinewtonMillimeterPerRadian)
}

// FromDecinewtonMillimeterPerRadian creates a new RotationalStiffness instance from DecinewtonMillimeterPerRadian.
func (uf RotationalStiffnessFactory) FromDecinewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecinewtonMillimeterPerRadian)
}

// FromDecanewtonMillimeterPerRadian creates a new RotationalStiffness instance from DecanewtonMillimeterPerRadian.
func (uf RotationalStiffnessFactory) FromDecanewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessDecanewtonMillimeterPerRadian)
}

// FromKilonewtonMillimeterPerRadian creates a new RotationalStiffness instance from KilonewtonMillimeterPerRadian.
func (uf RotationalStiffnessFactory) FromKilonewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessKilonewtonMillimeterPerRadian)
}

// FromMeganewtonMillimeterPerRadian creates a new RotationalStiffness instance from MeganewtonMillimeterPerRadian.
func (uf RotationalStiffnessFactory) FromMeganewtonMillimetersPerRadian(value float64) (*RotationalStiffness, error) {
	return newRotationalStiffness(value, RotationalStiffnessMeganewtonMillimeterPerRadian)
}




// newRotationalStiffness creates a new RotationalStiffness.
func newRotationalStiffness(value float64, fromUnit RotationalStiffnessUnits) (*RotationalStiffness, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RotationalStiffness{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RotationalStiffness in NewtonMeterPerRadian.
func (a *RotationalStiffness) BaseValue() float64 {
	return a.value
}


// NewtonMeterPerRadian returns the value in NewtonMeterPerRadian.
func (a *RotationalStiffness) NewtonMetersPerRadian() float64 {
	if a.newton_meters_per_radianLazy != nil {
		return *a.newton_meters_per_radianLazy
	}
	newton_meters_per_radian := a.convertFromBase(RotationalStiffnessNewtonMeterPerRadian)
	a.newton_meters_per_radianLazy = &newton_meters_per_radian
	return newton_meters_per_radian
}

// PoundForceFootPerDegrees returns the value in PoundForceFootPerDegrees.
func (a *RotationalStiffness) PoundForceFeetPerDegrees() float64 {
	if a.pound_force_feet_per_degreesLazy != nil {
		return *a.pound_force_feet_per_degreesLazy
	}
	pound_force_feet_per_degrees := a.convertFromBase(RotationalStiffnessPoundForceFootPerDegrees)
	a.pound_force_feet_per_degreesLazy = &pound_force_feet_per_degrees
	return pound_force_feet_per_degrees
}

// KilopoundForceFootPerDegrees returns the value in KilopoundForceFootPerDegrees.
func (a *RotationalStiffness) KilopoundForceFeetPerDegrees() float64 {
	if a.kilopound_force_feet_per_degreesLazy != nil {
		return *a.kilopound_force_feet_per_degreesLazy
	}
	kilopound_force_feet_per_degrees := a.convertFromBase(RotationalStiffnessKilopoundForceFootPerDegrees)
	a.kilopound_force_feet_per_degreesLazy = &kilopound_force_feet_per_degrees
	return kilopound_force_feet_per_degrees
}

// NewtonMillimeterPerDegree returns the value in NewtonMillimeterPerDegree.
func (a *RotationalStiffness) NewtonMillimetersPerDegree() float64 {
	if a.newton_millimeters_per_degreeLazy != nil {
		return *a.newton_millimeters_per_degreeLazy
	}
	newton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessNewtonMillimeterPerDegree)
	a.newton_millimeters_per_degreeLazy = &newton_millimeters_per_degree
	return newton_millimeters_per_degree
}

// NewtonMeterPerDegree returns the value in NewtonMeterPerDegree.
func (a *RotationalStiffness) NewtonMetersPerDegree() float64 {
	if a.newton_meters_per_degreeLazy != nil {
		return *a.newton_meters_per_degreeLazy
	}
	newton_meters_per_degree := a.convertFromBase(RotationalStiffnessNewtonMeterPerDegree)
	a.newton_meters_per_degreeLazy = &newton_meters_per_degree
	return newton_meters_per_degree
}

// NewtonMillimeterPerRadian returns the value in NewtonMillimeterPerRadian.
func (a *RotationalStiffness) NewtonMillimetersPerRadian() float64 {
	if a.newton_millimeters_per_radianLazy != nil {
		return *a.newton_millimeters_per_radianLazy
	}
	newton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessNewtonMillimeterPerRadian)
	a.newton_millimeters_per_radianLazy = &newton_millimeters_per_radian
	return newton_millimeters_per_radian
}

// PoundForceFeetPerRadian returns the value in PoundForceFeetPerRadian.
func (a *RotationalStiffness) PoundForceFeetPerRadian() float64 {
	if a.pound_force_feet_per_radianLazy != nil {
		return *a.pound_force_feet_per_radianLazy
	}
	pound_force_feet_per_radian := a.convertFromBase(RotationalStiffnessPoundForceFeetPerRadian)
	a.pound_force_feet_per_radianLazy = &pound_force_feet_per_radian
	return pound_force_feet_per_radian
}

// KilonewtonMeterPerRadian returns the value in KilonewtonMeterPerRadian.
func (a *RotationalStiffness) KilonewtonMetersPerRadian() float64 {
	if a.kilonewton_meters_per_radianLazy != nil {
		return *a.kilonewton_meters_per_radianLazy
	}
	kilonewton_meters_per_radian := a.convertFromBase(RotationalStiffnessKilonewtonMeterPerRadian)
	a.kilonewton_meters_per_radianLazy = &kilonewton_meters_per_radian
	return kilonewton_meters_per_radian
}

// MeganewtonMeterPerRadian returns the value in MeganewtonMeterPerRadian.
func (a *RotationalStiffness) MeganewtonMetersPerRadian() float64 {
	if a.meganewton_meters_per_radianLazy != nil {
		return *a.meganewton_meters_per_radianLazy
	}
	meganewton_meters_per_radian := a.convertFromBase(RotationalStiffnessMeganewtonMeterPerRadian)
	a.meganewton_meters_per_radianLazy = &meganewton_meters_per_radian
	return meganewton_meters_per_radian
}

// NanonewtonMillimeterPerDegree returns the value in NanonewtonMillimeterPerDegree.
func (a *RotationalStiffness) NanonewtonMillimetersPerDegree() float64 {
	if a.nanonewton_millimeters_per_degreeLazy != nil {
		return *a.nanonewton_millimeters_per_degreeLazy
	}
	nanonewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessNanonewtonMillimeterPerDegree)
	a.nanonewton_millimeters_per_degreeLazy = &nanonewton_millimeters_per_degree
	return nanonewton_millimeters_per_degree
}

// MicronewtonMillimeterPerDegree returns the value in MicronewtonMillimeterPerDegree.
func (a *RotationalStiffness) MicronewtonMillimetersPerDegree() float64 {
	if a.micronewton_millimeters_per_degreeLazy != nil {
		return *a.micronewton_millimeters_per_degreeLazy
	}
	micronewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessMicronewtonMillimeterPerDegree)
	a.micronewton_millimeters_per_degreeLazy = &micronewton_millimeters_per_degree
	return micronewton_millimeters_per_degree
}

// MillinewtonMillimeterPerDegree returns the value in MillinewtonMillimeterPerDegree.
func (a *RotationalStiffness) MillinewtonMillimetersPerDegree() float64 {
	if a.millinewton_millimeters_per_degreeLazy != nil {
		return *a.millinewton_millimeters_per_degreeLazy
	}
	millinewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessMillinewtonMillimeterPerDegree)
	a.millinewton_millimeters_per_degreeLazy = &millinewton_millimeters_per_degree
	return millinewton_millimeters_per_degree
}

// CentinewtonMillimeterPerDegree returns the value in CentinewtonMillimeterPerDegree.
func (a *RotationalStiffness) CentinewtonMillimetersPerDegree() float64 {
	if a.centinewton_millimeters_per_degreeLazy != nil {
		return *a.centinewton_millimeters_per_degreeLazy
	}
	centinewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessCentinewtonMillimeterPerDegree)
	a.centinewton_millimeters_per_degreeLazy = &centinewton_millimeters_per_degree
	return centinewton_millimeters_per_degree
}

// DecinewtonMillimeterPerDegree returns the value in DecinewtonMillimeterPerDegree.
func (a *RotationalStiffness) DecinewtonMillimetersPerDegree() float64 {
	if a.decinewton_millimeters_per_degreeLazy != nil {
		return *a.decinewton_millimeters_per_degreeLazy
	}
	decinewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessDecinewtonMillimeterPerDegree)
	a.decinewton_millimeters_per_degreeLazy = &decinewton_millimeters_per_degree
	return decinewton_millimeters_per_degree
}

// DecanewtonMillimeterPerDegree returns the value in DecanewtonMillimeterPerDegree.
func (a *RotationalStiffness) DecanewtonMillimetersPerDegree() float64 {
	if a.decanewton_millimeters_per_degreeLazy != nil {
		return *a.decanewton_millimeters_per_degreeLazy
	}
	decanewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessDecanewtonMillimeterPerDegree)
	a.decanewton_millimeters_per_degreeLazy = &decanewton_millimeters_per_degree
	return decanewton_millimeters_per_degree
}

// KilonewtonMillimeterPerDegree returns the value in KilonewtonMillimeterPerDegree.
func (a *RotationalStiffness) KilonewtonMillimetersPerDegree() float64 {
	if a.kilonewton_millimeters_per_degreeLazy != nil {
		return *a.kilonewton_millimeters_per_degreeLazy
	}
	kilonewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessKilonewtonMillimeterPerDegree)
	a.kilonewton_millimeters_per_degreeLazy = &kilonewton_millimeters_per_degree
	return kilonewton_millimeters_per_degree
}

// MeganewtonMillimeterPerDegree returns the value in MeganewtonMillimeterPerDegree.
func (a *RotationalStiffness) MeganewtonMillimetersPerDegree() float64 {
	if a.meganewton_millimeters_per_degreeLazy != nil {
		return *a.meganewton_millimeters_per_degreeLazy
	}
	meganewton_millimeters_per_degree := a.convertFromBase(RotationalStiffnessMeganewtonMillimeterPerDegree)
	a.meganewton_millimeters_per_degreeLazy = &meganewton_millimeters_per_degree
	return meganewton_millimeters_per_degree
}

// NanonewtonMeterPerDegree returns the value in NanonewtonMeterPerDegree.
func (a *RotationalStiffness) NanonewtonMetersPerDegree() float64 {
	if a.nanonewton_meters_per_degreeLazy != nil {
		return *a.nanonewton_meters_per_degreeLazy
	}
	nanonewton_meters_per_degree := a.convertFromBase(RotationalStiffnessNanonewtonMeterPerDegree)
	a.nanonewton_meters_per_degreeLazy = &nanonewton_meters_per_degree
	return nanonewton_meters_per_degree
}

// MicronewtonMeterPerDegree returns the value in MicronewtonMeterPerDegree.
func (a *RotationalStiffness) MicronewtonMetersPerDegree() float64 {
	if a.micronewton_meters_per_degreeLazy != nil {
		return *a.micronewton_meters_per_degreeLazy
	}
	micronewton_meters_per_degree := a.convertFromBase(RotationalStiffnessMicronewtonMeterPerDegree)
	a.micronewton_meters_per_degreeLazy = &micronewton_meters_per_degree
	return micronewton_meters_per_degree
}

// MillinewtonMeterPerDegree returns the value in MillinewtonMeterPerDegree.
func (a *RotationalStiffness) MillinewtonMetersPerDegree() float64 {
	if a.millinewton_meters_per_degreeLazy != nil {
		return *a.millinewton_meters_per_degreeLazy
	}
	millinewton_meters_per_degree := a.convertFromBase(RotationalStiffnessMillinewtonMeterPerDegree)
	a.millinewton_meters_per_degreeLazy = &millinewton_meters_per_degree
	return millinewton_meters_per_degree
}

// CentinewtonMeterPerDegree returns the value in CentinewtonMeterPerDegree.
func (a *RotationalStiffness) CentinewtonMetersPerDegree() float64 {
	if a.centinewton_meters_per_degreeLazy != nil {
		return *a.centinewton_meters_per_degreeLazy
	}
	centinewton_meters_per_degree := a.convertFromBase(RotationalStiffnessCentinewtonMeterPerDegree)
	a.centinewton_meters_per_degreeLazy = &centinewton_meters_per_degree
	return centinewton_meters_per_degree
}

// DecinewtonMeterPerDegree returns the value in DecinewtonMeterPerDegree.
func (a *RotationalStiffness) DecinewtonMetersPerDegree() float64 {
	if a.decinewton_meters_per_degreeLazy != nil {
		return *a.decinewton_meters_per_degreeLazy
	}
	decinewton_meters_per_degree := a.convertFromBase(RotationalStiffnessDecinewtonMeterPerDegree)
	a.decinewton_meters_per_degreeLazy = &decinewton_meters_per_degree
	return decinewton_meters_per_degree
}

// DecanewtonMeterPerDegree returns the value in DecanewtonMeterPerDegree.
func (a *RotationalStiffness) DecanewtonMetersPerDegree() float64 {
	if a.decanewton_meters_per_degreeLazy != nil {
		return *a.decanewton_meters_per_degreeLazy
	}
	decanewton_meters_per_degree := a.convertFromBase(RotationalStiffnessDecanewtonMeterPerDegree)
	a.decanewton_meters_per_degreeLazy = &decanewton_meters_per_degree
	return decanewton_meters_per_degree
}

// KilonewtonMeterPerDegree returns the value in KilonewtonMeterPerDegree.
func (a *RotationalStiffness) KilonewtonMetersPerDegree() float64 {
	if a.kilonewton_meters_per_degreeLazy != nil {
		return *a.kilonewton_meters_per_degreeLazy
	}
	kilonewton_meters_per_degree := a.convertFromBase(RotationalStiffnessKilonewtonMeterPerDegree)
	a.kilonewton_meters_per_degreeLazy = &kilonewton_meters_per_degree
	return kilonewton_meters_per_degree
}

// MeganewtonMeterPerDegree returns the value in MeganewtonMeterPerDegree.
func (a *RotationalStiffness) MeganewtonMetersPerDegree() float64 {
	if a.meganewton_meters_per_degreeLazy != nil {
		return *a.meganewton_meters_per_degreeLazy
	}
	meganewton_meters_per_degree := a.convertFromBase(RotationalStiffnessMeganewtonMeterPerDegree)
	a.meganewton_meters_per_degreeLazy = &meganewton_meters_per_degree
	return meganewton_meters_per_degree
}

// NanonewtonMillimeterPerRadian returns the value in NanonewtonMillimeterPerRadian.
func (a *RotationalStiffness) NanonewtonMillimetersPerRadian() float64 {
	if a.nanonewton_millimeters_per_radianLazy != nil {
		return *a.nanonewton_millimeters_per_radianLazy
	}
	nanonewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessNanonewtonMillimeterPerRadian)
	a.nanonewton_millimeters_per_radianLazy = &nanonewton_millimeters_per_radian
	return nanonewton_millimeters_per_radian
}

// MicronewtonMillimeterPerRadian returns the value in MicronewtonMillimeterPerRadian.
func (a *RotationalStiffness) MicronewtonMillimetersPerRadian() float64 {
	if a.micronewton_millimeters_per_radianLazy != nil {
		return *a.micronewton_millimeters_per_radianLazy
	}
	micronewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessMicronewtonMillimeterPerRadian)
	a.micronewton_millimeters_per_radianLazy = &micronewton_millimeters_per_radian
	return micronewton_millimeters_per_radian
}

// MillinewtonMillimeterPerRadian returns the value in MillinewtonMillimeterPerRadian.
func (a *RotationalStiffness) MillinewtonMillimetersPerRadian() float64 {
	if a.millinewton_millimeters_per_radianLazy != nil {
		return *a.millinewton_millimeters_per_radianLazy
	}
	millinewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessMillinewtonMillimeterPerRadian)
	a.millinewton_millimeters_per_radianLazy = &millinewton_millimeters_per_radian
	return millinewton_millimeters_per_radian
}

// CentinewtonMillimeterPerRadian returns the value in CentinewtonMillimeterPerRadian.
func (a *RotationalStiffness) CentinewtonMillimetersPerRadian() float64 {
	if a.centinewton_millimeters_per_radianLazy != nil {
		return *a.centinewton_millimeters_per_radianLazy
	}
	centinewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessCentinewtonMillimeterPerRadian)
	a.centinewton_millimeters_per_radianLazy = &centinewton_millimeters_per_radian
	return centinewton_millimeters_per_radian
}

// DecinewtonMillimeterPerRadian returns the value in DecinewtonMillimeterPerRadian.
func (a *RotationalStiffness) DecinewtonMillimetersPerRadian() float64 {
	if a.decinewton_millimeters_per_radianLazy != nil {
		return *a.decinewton_millimeters_per_radianLazy
	}
	decinewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessDecinewtonMillimeterPerRadian)
	a.decinewton_millimeters_per_radianLazy = &decinewton_millimeters_per_radian
	return decinewton_millimeters_per_radian
}

// DecanewtonMillimeterPerRadian returns the value in DecanewtonMillimeterPerRadian.
func (a *RotationalStiffness) DecanewtonMillimetersPerRadian() float64 {
	if a.decanewton_millimeters_per_radianLazy != nil {
		return *a.decanewton_millimeters_per_radianLazy
	}
	decanewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessDecanewtonMillimeterPerRadian)
	a.decanewton_millimeters_per_radianLazy = &decanewton_millimeters_per_radian
	return decanewton_millimeters_per_radian
}

// KilonewtonMillimeterPerRadian returns the value in KilonewtonMillimeterPerRadian.
func (a *RotationalStiffness) KilonewtonMillimetersPerRadian() float64 {
	if a.kilonewton_millimeters_per_radianLazy != nil {
		return *a.kilonewton_millimeters_per_radianLazy
	}
	kilonewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessKilonewtonMillimeterPerRadian)
	a.kilonewton_millimeters_per_radianLazy = &kilonewton_millimeters_per_radian
	return kilonewton_millimeters_per_radian
}

// MeganewtonMillimeterPerRadian returns the value in MeganewtonMillimeterPerRadian.
func (a *RotationalStiffness) MeganewtonMillimetersPerRadian() float64 {
	if a.meganewton_millimeters_per_radianLazy != nil {
		return *a.meganewton_millimeters_per_radianLazy
	}
	meganewton_millimeters_per_radian := a.convertFromBase(RotationalStiffnessMeganewtonMillimeterPerRadian)
	a.meganewton_millimeters_per_radianLazy = &meganewton_millimeters_per_radian
	return meganewton_millimeters_per_radian
}


// ToDto creates an RotationalStiffnessDto representation.
func (a *RotationalStiffness) ToDto(holdInUnit *RotationalStiffnessUnits) RotationalStiffnessDto {
	if holdInUnit == nil {
		defaultUnit := RotationalStiffnessNewtonMeterPerRadian // Default value
		holdInUnit = &defaultUnit
	}

	return RotationalStiffnessDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an RotationalStiffnessDto representation.
func (a *RotationalStiffness) ToDtoJSON(holdInUnit *RotationalStiffnessUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts RotationalStiffness to a specific unit value.
func (a *RotationalStiffness) Convert(toUnit RotationalStiffnessUnits) float64 {
	switch toUnit { 
    case RotationalStiffnessNewtonMeterPerRadian:
		return a.NewtonMetersPerRadian()
    case RotationalStiffnessPoundForceFootPerDegrees:
		return a.PoundForceFeetPerDegrees()
    case RotationalStiffnessKilopoundForceFootPerDegrees:
		return a.KilopoundForceFeetPerDegrees()
    case RotationalStiffnessNewtonMillimeterPerDegree:
		return a.NewtonMillimetersPerDegree()
    case RotationalStiffnessNewtonMeterPerDegree:
		return a.NewtonMetersPerDegree()
    case RotationalStiffnessNewtonMillimeterPerRadian:
		return a.NewtonMillimetersPerRadian()
    case RotationalStiffnessPoundForceFeetPerRadian:
		return a.PoundForceFeetPerRadian()
    case RotationalStiffnessKilonewtonMeterPerRadian:
		return a.KilonewtonMetersPerRadian()
    case RotationalStiffnessMeganewtonMeterPerRadian:
		return a.MeganewtonMetersPerRadian()
    case RotationalStiffnessNanonewtonMillimeterPerDegree:
		return a.NanonewtonMillimetersPerDegree()
    case RotationalStiffnessMicronewtonMillimeterPerDegree:
		return a.MicronewtonMillimetersPerDegree()
    case RotationalStiffnessMillinewtonMillimeterPerDegree:
		return a.MillinewtonMillimetersPerDegree()
    case RotationalStiffnessCentinewtonMillimeterPerDegree:
		return a.CentinewtonMillimetersPerDegree()
    case RotationalStiffnessDecinewtonMillimeterPerDegree:
		return a.DecinewtonMillimetersPerDegree()
    case RotationalStiffnessDecanewtonMillimeterPerDegree:
		return a.DecanewtonMillimetersPerDegree()
    case RotationalStiffnessKilonewtonMillimeterPerDegree:
		return a.KilonewtonMillimetersPerDegree()
    case RotationalStiffnessMeganewtonMillimeterPerDegree:
		return a.MeganewtonMillimetersPerDegree()
    case RotationalStiffnessNanonewtonMeterPerDegree:
		return a.NanonewtonMetersPerDegree()
    case RotationalStiffnessMicronewtonMeterPerDegree:
		return a.MicronewtonMetersPerDegree()
    case RotationalStiffnessMillinewtonMeterPerDegree:
		return a.MillinewtonMetersPerDegree()
    case RotationalStiffnessCentinewtonMeterPerDegree:
		return a.CentinewtonMetersPerDegree()
    case RotationalStiffnessDecinewtonMeterPerDegree:
		return a.DecinewtonMetersPerDegree()
    case RotationalStiffnessDecanewtonMeterPerDegree:
		return a.DecanewtonMetersPerDegree()
    case RotationalStiffnessKilonewtonMeterPerDegree:
		return a.KilonewtonMetersPerDegree()
    case RotationalStiffnessMeganewtonMeterPerDegree:
		return a.MeganewtonMetersPerDegree()
    case RotationalStiffnessNanonewtonMillimeterPerRadian:
		return a.NanonewtonMillimetersPerRadian()
    case RotationalStiffnessMicronewtonMillimeterPerRadian:
		return a.MicronewtonMillimetersPerRadian()
    case RotationalStiffnessMillinewtonMillimeterPerRadian:
		return a.MillinewtonMillimetersPerRadian()
    case RotationalStiffnessCentinewtonMillimeterPerRadian:
		return a.CentinewtonMillimetersPerRadian()
    case RotationalStiffnessDecinewtonMillimeterPerRadian:
		return a.DecinewtonMillimetersPerRadian()
    case RotationalStiffnessDecanewtonMillimeterPerRadian:
		return a.DecanewtonMillimetersPerRadian()
    case RotationalStiffnessKilonewtonMillimeterPerRadian:
		return a.KilonewtonMillimetersPerRadian()
    case RotationalStiffnessMeganewtonMillimeterPerRadian:
		return a.MeganewtonMillimetersPerRadian()
	default:
		return 0
	}
}

func (a *RotationalStiffness) convertFromBase(toUnit RotationalStiffnessUnits) float64 {
    value := a.value
	switch toUnit { 
	case RotationalStiffnessNewtonMeterPerRadian:
		return (value) 
	case RotationalStiffnessPoundForceFootPerDegrees:
		return (value / 77.6826) 
	case RotationalStiffnessKilopoundForceFootPerDegrees:
		return (value / 77682.6) 
	case RotationalStiffnessNewtonMillimeterPerDegree:
		return (value / 180 * math.Pi * 1000) 
	case RotationalStiffnessNewtonMeterPerDegree:
		return (value / (180 / math.Pi)) 
	case RotationalStiffnessNewtonMillimeterPerRadian:
		return (value * 1000) 
	case RotationalStiffnessPoundForceFeetPerRadian:
		return (value / 1.3558179483314) 
	case RotationalStiffnessKilonewtonMeterPerRadian:
		return ((value) / 1000.0) 
	case RotationalStiffnessMeganewtonMeterPerRadian:
		return ((value) / 1000000.0) 
	case RotationalStiffnessNanonewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 1e-09) 
	case RotationalStiffnessMicronewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 1e-06) 
	case RotationalStiffnessMillinewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 0.001) 
	case RotationalStiffnessCentinewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 0.01) 
	case RotationalStiffnessDecinewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 0.1) 
	case RotationalStiffnessDecanewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 10.0) 
	case RotationalStiffnessKilonewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 1000.0) 
	case RotationalStiffnessMeganewtonMillimeterPerDegree:
		return ((value / 180 * math.Pi * 1000) / 1000000.0) 
	case RotationalStiffnessNanonewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 1e-09) 
	case RotationalStiffnessMicronewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 1e-06) 
	case RotationalStiffnessMillinewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 0.001) 
	case RotationalStiffnessCentinewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 0.01) 
	case RotationalStiffnessDecinewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 0.1) 
	case RotationalStiffnessDecanewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 10.0) 
	case RotationalStiffnessKilonewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 1000.0) 
	case RotationalStiffnessMeganewtonMeterPerDegree:
		return ((value / (180 / math.Pi)) / 1000000.0) 
	case RotationalStiffnessNanonewtonMillimeterPerRadian:
		return ((value * 1000) / 1e-09) 
	case RotationalStiffnessMicronewtonMillimeterPerRadian:
		return ((value * 1000) / 1e-06) 
	case RotationalStiffnessMillinewtonMillimeterPerRadian:
		return ((value * 1000) / 0.001) 
	case RotationalStiffnessCentinewtonMillimeterPerRadian:
		return ((value * 1000) / 0.01) 
	case RotationalStiffnessDecinewtonMillimeterPerRadian:
		return ((value * 1000) / 0.1) 
	case RotationalStiffnessDecanewtonMillimeterPerRadian:
		return ((value * 1000) / 10.0) 
	case RotationalStiffnessKilonewtonMillimeterPerRadian:
		return ((value * 1000) / 1000.0) 
	case RotationalStiffnessMeganewtonMillimeterPerRadian:
		return ((value * 1000) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *RotationalStiffness) convertToBase(value float64, fromUnit RotationalStiffnessUnits) float64 {
	switch fromUnit { 
	case RotationalStiffnessNewtonMeterPerRadian:
		return (value) 
	case RotationalStiffnessPoundForceFootPerDegrees:
		return (value * 77.6826) 
	case RotationalStiffnessKilopoundForceFootPerDegrees:
		return (value * 77682.6) 
	case RotationalStiffnessNewtonMillimeterPerDegree:
		return (value * 180 / math.Pi * 0.001) 
	case RotationalStiffnessNewtonMeterPerDegree:
		return (value * (180 / math.Pi)) 
	case RotationalStiffnessNewtonMillimeterPerRadian:
		return (value * 0.001) 
	case RotationalStiffnessPoundForceFeetPerRadian:
		return (value * 1.3558179483314) 
	case RotationalStiffnessKilonewtonMeterPerRadian:
		return ((value) * 1000.0) 
	case RotationalStiffnessMeganewtonMeterPerRadian:
		return ((value) * 1000000.0) 
	case RotationalStiffnessNanonewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 1e-09) 
	case RotationalStiffnessMicronewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 1e-06) 
	case RotationalStiffnessMillinewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 0.001) 
	case RotationalStiffnessCentinewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 0.01) 
	case RotationalStiffnessDecinewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 0.1) 
	case RotationalStiffnessDecanewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 10.0) 
	case RotationalStiffnessKilonewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 1000.0) 
	case RotationalStiffnessMeganewtonMillimeterPerDegree:
		return ((value * 180 / math.Pi * 0.001) * 1000000.0) 
	case RotationalStiffnessNanonewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 1e-09) 
	case RotationalStiffnessMicronewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 1e-06) 
	case RotationalStiffnessMillinewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 0.001) 
	case RotationalStiffnessCentinewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 0.01) 
	case RotationalStiffnessDecinewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 0.1) 
	case RotationalStiffnessDecanewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 10.0) 
	case RotationalStiffnessKilonewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 1000.0) 
	case RotationalStiffnessMeganewtonMeterPerDegree:
		return ((value * (180 / math.Pi)) * 1000000.0) 
	case RotationalStiffnessNanonewtonMillimeterPerRadian:
		return ((value * 0.001) * 1e-09) 
	case RotationalStiffnessMicronewtonMillimeterPerRadian:
		return ((value * 0.001) * 1e-06) 
	case RotationalStiffnessMillinewtonMillimeterPerRadian:
		return ((value * 0.001) * 0.001) 
	case RotationalStiffnessCentinewtonMillimeterPerRadian:
		return ((value * 0.001) * 0.01) 
	case RotationalStiffnessDecinewtonMillimeterPerRadian:
		return ((value * 0.001) * 0.1) 
	case RotationalStiffnessDecanewtonMillimeterPerRadian:
		return ((value * 0.001) * 10.0) 
	case RotationalStiffnessKilonewtonMillimeterPerRadian:
		return ((value * 0.001) * 1000.0) 
	case RotationalStiffnessMeganewtonMillimeterPerRadian:
		return ((value * 0.001) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a RotationalStiffness) String() string {
	return a.ToString(RotationalStiffnessNewtonMeterPerRadian, 2)
}

// ToString formats the RotationalStiffness to string.
// fractionalDigits -1 for not mention
func (a *RotationalStiffness) ToString(unit RotationalStiffnessUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RotationalStiffness) getUnitAbbreviation(unit RotationalStiffnessUnits) string {
	switch unit { 
	case RotationalStiffnessNewtonMeterPerRadian:
		return "N·m/rad" 
	case RotationalStiffnessPoundForceFootPerDegrees:
		return "lbf·ft/deg" 
	case RotationalStiffnessKilopoundForceFootPerDegrees:
		return "kipf·ft/°" 
	case RotationalStiffnessNewtonMillimeterPerDegree:
		return "N·mm/deg" 
	case RotationalStiffnessNewtonMeterPerDegree:
		return "N·m/deg" 
	case RotationalStiffnessNewtonMillimeterPerRadian:
		return "N·mm/rad" 
	case RotationalStiffnessPoundForceFeetPerRadian:
		return "lbf·ft/rad" 
	case RotationalStiffnessKilonewtonMeterPerRadian:
		return "kN·m/rad" 
	case RotationalStiffnessMeganewtonMeterPerRadian:
		return "MN·m/rad" 
	case RotationalStiffnessNanonewtonMillimeterPerDegree:
		return "nN·mm/deg" 
	case RotationalStiffnessMicronewtonMillimeterPerDegree:
		return "μN·mm/deg" 
	case RotationalStiffnessMillinewtonMillimeterPerDegree:
		return "mN·mm/deg" 
	case RotationalStiffnessCentinewtonMillimeterPerDegree:
		return "cN·mm/deg" 
	case RotationalStiffnessDecinewtonMillimeterPerDegree:
		return "dN·mm/deg" 
	case RotationalStiffnessDecanewtonMillimeterPerDegree:
		return "daN·mm/deg" 
	case RotationalStiffnessKilonewtonMillimeterPerDegree:
		return "kN·mm/deg" 
	case RotationalStiffnessMeganewtonMillimeterPerDegree:
		return "MN·mm/deg" 
	case RotationalStiffnessNanonewtonMeterPerDegree:
		return "nN·m/deg" 
	case RotationalStiffnessMicronewtonMeterPerDegree:
		return "μN·m/deg" 
	case RotationalStiffnessMillinewtonMeterPerDegree:
		return "mN·m/deg" 
	case RotationalStiffnessCentinewtonMeterPerDegree:
		return "cN·m/deg" 
	case RotationalStiffnessDecinewtonMeterPerDegree:
		return "dN·m/deg" 
	case RotationalStiffnessDecanewtonMeterPerDegree:
		return "daN·m/deg" 
	case RotationalStiffnessKilonewtonMeterPerDegree:
		return "kN·m/deg" 
	case RotationalStiffnessMeganewtonMeterPerDegree:
		return "MN·m/deg" 
	case RotationalStiffnessNanonewtonMillimeterPerRadian:
		return "nN·mm/rad" 
	case RotationalStiffnessMicronewtonMillimeterPerRadian:
		return "μN·mm/rad" 
	case RotationalStiffnessMillinewtonMillimeterPerRadian:
		return "mN·mm/rad" 
	case RotationalStiffnessCentinewtonMillimeterPerRadian:
		return "cN·mm/rad" 
	case RotationalStiffnessDecinewtonMillimeterPerRadian:
		return "dN·mm/rad" 
	case RotationalStiffnessDecanewtonMillimeterPerRadian:
		return "daN·mm/rad" 
	case RotationalStiffnessKilonewtonMillimeterPerRadian:
		return "kN·mm/rad" 
	case RotationalStiffnessMeganewtonMillimeterPerRadian:
		return "MN·mm/rad" 
	default:
		return ""
	}
}

// Check if the given RotationalStiffness are equals to the current RotationalStiffness
func (a *RotationalStiffness) Equals(other *RotationalStiffness) bool {
	return a.value == other.BaseValue()
}

// Check if the given RotationalStiffness are equals to the current RotationalStiffness
func (a *RotationalStiffness) CompareTo(other *RotationalStiffness) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given RotationalStiffness to the current RotationalStiffness.
func (a *RotationalStiffness) Add(other *RotationalStiffness) *RotationalStiffness {
	return &RotationalStiffness{value: a.value + other.BaseValue()}
}

// Subtract the given RotationalStiffness to the current RotationalStiffness.
func (a *RotationalStiffness) Subtract(other *RotationalStiffness) *RotationalStiffness {
	return &RotationalStiffness{value: a.value - other.BaseValue()}
}

// Multiply the given RotationalStiffness to the current RotationalStiffness.
func (a *RotationalStiffness) Multiply(other *RotationalStiffness) *RotationalStiffness {
	return &RotationalStiffness{value: a.value * other.BaseValue()}
}

// Divide the given RotationalStiffness to the current RotationalStiffness.
func (a *RotationalStiffness) Divide(other *RotationalStiffness) *RotationalStiffness {
	return &RotationalStiffness{value: a.value / other.BaseValue()}
}