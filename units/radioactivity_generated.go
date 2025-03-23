// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RadioactivityUnits defines various units of Radioactivity.
type RadioactivityUnits string

const (
    
        // Activity of a quantity of radioactive material in which one nucleus decays per second.
        RadioactivityBecquerel RadioactivityUnits = "Becquerel"
        // 
        RadioactivityCurie RadioactivityUnits = "Curie"
        // Activity of a quantity of radioactive material in which one million nuclei decay per second.
        RadioactivityRutherford RadioactivityUnits = "Rutherford"
        // 
        RadioactivityPicobecquerel RadioactivityUnits = "Picobecquerel"
        // 
        RadioactivityNanobecquerel RadioactivityUnits = "Nanobecquerel"
        // 
        RadioactivityMicrobecquerel RadioactivityUnits = "Microbecquerel"
        // 
        RadioactivityMillibecquerel RadioactivityUnits = "Millibecquerel"
        // 
        RadioactivityKilobecquerel RadioactivityUnits = "Kilobecquerel"
        // 
        RadioactivityMegabecquerel RadioactivityUnits = "Megabecquerel"
        // 
        RadioactivityGigabecquerel RadioactivityUnits = "Gigabecquerel"
        // 
        RadioactivityTerabecquerel RadioactivityUnits = "Terabecquerel"
        // 
        RadioactivityPetabecquerel RadioactivityUnits = "Petabecquerel"
        // 
        RadioactivityExabecquerel RadioactivityUnits = "Exabecquerel"
        // 
        RadioactivityPicocurie RadioactivityUnits = "Picocurie"
        // 
        RadioactivityNanocurie RadioactivityUnits = "Nanocurie"
        // 
        RadioactivityMicrocurie RadioactivityUnits = "Microcurie"
        // 
        RadioactivityMillicurie RadioactivityUnits = "Millicurie"
        // 
        RadioactivityKilocurie RadioactivityUnits = "Kilocurie"
        // 
        RadioactivityMegacurie RadioactivityUnits = "Megacurie"
        // 
        RadioactivityGigacurie RadioactivityUnits = "Gigacurie"
        // 
        RadioactivityTeracurie RadioactivityUnits = "Teracurie"
        // 
        RadioactivityPicorutherford RadioactivityUnits = "Picorutherford"
        // 
        RadioactivityNanorutherford RadioactivityUnits = "Nanorutherford"
        // 
        RadioactivityMicrorutherford RadioactivityUnits = "Microrutherford"
        // 
        RadioactivityMillirutherford RadioactivityUnits = "Millirutherford"
        // 
        RadioactivityKilorutherford RadioactivityUnits = "Kilorutherford"
        // 
        RadioactivityMegarutherford RadioactivityUnits = "Megarutherford"
        // 
        RadioactivityGigarutherford RadioactivityUnits = "Gigarutherford"
        // 
        RadioactivityTerarutherford RadioactivityUnits = "Terarutherford"
)

var internalRadioactivityUnitsMap = map[RadioactivityUnits]bool{
	
	RadioactivityBecquerel: true,
	RadioactivityCurie: true,
	RadioactivityRutherford: true,
	RadioactivityPicobecquerel: true,
	RadioactivityNanobecquerel: true,
	RadioactivityMicrobecquerel: true,
	RadioactivityMillibecquerel: true,
	RadioactivityKilobecquerel: true,
	RadioactivityMegabecquerel: true,
	RadioactivityGigabecquerel: true,
	RadioactivityTerabecquerel: true,
	RadioactivityPetabecquerel: true,
	RadioactivityExabecquerel: true,
	RadioactivityPicocurie: true,
	RadioactivityNanocurie: true,
	RadioactivityMicrocurie: true,
	RadioactivityMillicurie: true,
	RadioactivityKilocurie: true,
	RadioactivityMegacurie: true,
	RadioactivityGigacurie: true,
	RadioactivityTeracurie: true,
	RadioactivityPicorutherford: true,
	RadioactivityNanorutherford: true,
	RadioactivityMicrorutherford: true,
	RadioactivityMillirutherford: true,
	RadioactivityKilorutherford: true,
	RadioactivityMegarutherford: true,
	RadioactivityGigarutherford: true,
	RadioactivityTerarutherford: true,
}

// RadioactivityDto represents a Radioactivity measurement with a numerical value and its corresponding unit.
type RadioactivityDto struct {
    // Value is the numerical representation of the Radioactivity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Radioactivity, as defined in the RadioactivityUnits enumeration.
	Unit  RadioactivityUnits `json:"unit" validate:"required,oneof=Becquerel Curie Rutherford Picobecquerel Nanobecquerel Microbecquerel Millibecquerel Kilobecquerel Megabecquerel Gigabecquerel Terabecquerel Petabecquerel Exabecquerel Picocurie Nanocurie Microcurie Millicurie Kilocurie Megacurie Gigacurie Teracurie Picorutherford Nanorutherford Microrutherford Millirutherford Kilorutherford Megarutherford Gigarutherford Terarutherford"`
}

// RadioactivityDtoFactory groups methods for creating and serializing RadioactivityDto objects.
type RadioactivityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RadioactivityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RadioactivityDtoFactory) FromJSON(data []byte) (*RadioactivityDto, error) {
	a := RadioactivityDto{}

    // Parse JSON into RadioactivityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a RadioactivityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RadioactivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Radioactivity represents a measurement in a of Radioactivity.
//
// Amount of ionizing radiation released when an element spontaneously emits energy as a result of the radioactive decay of an unstable atom per unit time.
type Radioactivity struct {
	// value is the base measurement stored internally.
	value       float64
    
    becquerelsLazy *float64 
    curiesLazy *float64 
    rutherfordsLazy *float64 
    picobecquerelsLazy *float64 
    nanobecquerelsLazy *float64 
    microbecquerelsLazy *float64 
    millibecquerelsLazy *float64 
    kilobecquerelsLazy *float64 
    megabecquerelsLazy *float64 
    gigabecquerelsLazy *float64 
    terabecquerelsLazy *float64 
    petabecquerelsLazy *float64 
    exabecquerelsLazy *float64 
    picocuriesLazy *float64 
    nanocuriesLazy *float64 
    microcuriesLazy *float64 
    millicuriesLazy *float64 
    kilocuriesLazy *float64 
    megacuriesLazy *float64 
    gigacuriesLazy *float64 
    teracuriesLazy *float64 
    picorutherfordsLazy *float64 
    nanorutherfordsLazy *float64 
    microrutherfordsLazy *float64 
    millirutherfordsLazy *float64 
    kilorutherfordsLazy *float64 
    megarutherfordsLazy *float64 
    gigarutherfordsLazy *float64 
    terarutherfordsLazy *float64 
}

// RadioactivityFactory groups methods for creating Radioactivity instances.
type RadioactivityFactory struct{}

// CreateRadioactivity creates a new Radioactivity instance from the given value and unit.
func (uf RadioactivityFactory) CreateRadioactivity(value float64, unit RadioactivityUnits) (*Radioactivity, error) {
	return newRadioactivity(value, unit)
}

// FromDto converts a RadioactivityDto to a Radioactivity instance.
func (uf RadioactivityFactory) FromDto(dto RadioactivityDto) (*Radioactivity, error) {
	return newRadioactivity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Radioactivity instance.
func (uf RadioactivityFactory) FromDtoJSON(data []byte) (*Radioactivity, error) {
	unitDto, err := RadioactivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RadioactivityDto from JSON: %w", err)
	}
	return RadioactivityFactory{}.FromDto(*unitDto)
}


// FromBecquerels creates a new Radioactivity instance from a value in Becquerels.
func (uf RadioactivityFactory) FromBecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityBecquerel)
}

// FromCuries creates a new Radioactivity instance from a value in Curies.
func (uf RadioactivityFactory) FromCuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityCurie)
}

// FromRutherfords creates a new Radioactivity instance from a value in Rutherfords.
func (uf RadioactivityFactory) FromRutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityRutherford)
}

// FromPicobecquerels creates a new Radioactivity instance from a value in Picobecquerels.
func (uf RadioactivityFactory) FromPicobecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityPicobecquerel)
}

// FromNanobecquerels creates a new Radioactivity instance from a value in Nanobecquerels.
func (uf RadioactivityFactory) FromNanobecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityNanobecquerel)
}

// FromMicrobecquerels creates a new Radioactivity instance from a value in Microbecquerels.
func (uf RadioactivityFactory) FromMicrobecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMicrobecquerel)
}

// FromMillibecquerels creates a new Radioactivity instance from a value in Millibecquerels.
func (uf RadioactivityFactory) FromMillibecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMillibecquerel)
}

// FromKilobecquerels creates a new Radioactivity instance from a value in Kilobecquerels.
func (uf RadioactivityFactory) FromKilobecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityKilobecquerel)
}

// FromMegabecquerels creates a new Radioactivity instance from a value in Megabecquerels.
func (uf RadioactivityFactory) FromMegabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMegabecquerel)
}

// FromGigabecquerels creates a new Radioactivity instance from a value in Gigabecquerels.
func (uf RadioactivityFactory) FromGigabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityGigabecquerel)
}

// FromTerabecquerels creates a new Radioactivity instance from a value in Terabecquerels.
func (uf RadioactivityFactory) FromTerabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityTerabecquerel)
}

// FromPetabecquerels creates a new Radioactivity instance from a value in Petabecquerels.
func (uf RadioactivityFactory) FromPetabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityPetabecquerel)
}

// FromExabecquerels creates a new Radioactivity instance from a value in Exabecquerels.
func (uf RadioactivityFactory) FromExabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityExabecquerel)
}

// FromPicocuries creates a new Radioactivity instance from a value in Picocuries.
func (uf RadioactivityFactory) FromPicocuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityPicocurie)
}

// FromNanocuries creates a new Radioactivity instance from a value in Nanocuries.
func (uf RadioactivityFactory) FromNanocuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityNanocurie)
}

// FromMicrocuries creates a new Radioactivity instance from a value in Microcuries.
func (uf RadioactivityFactory) FromMicrocuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMicrocurie)
}

// FromMillicuries creates a new Radioactivity instance from a value in Millicuries.
func (uf RadioactivityFactory) FromMillicuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMillicurie)
}

// FromKilocuries creates a new Radioactivity instance from a value in Kilocuries.
func (uf RadioactivityFactory) FromKilocuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityKilocurie)
}

// FromMegacuries creates a new Radioactivity instance from a value in Megacuries.
func (uf RadioactivityFactory) FromMegacuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMegacurie)
}

// FromGigacuries creates a new Radioactivity instance from a value in Gigacuries.
func (uf RadioactivityFactory) FromGigacuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityGigacurie)
}

// FromTeracuries creates a new Radioactivity instance from a value in Teracuries.
func (uf RadioactivityFactory) FromTeracuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityTeracurie)
}

// FromPicorutherfords creates a new Radioactivity instance from a value in Picorutherfords.
func (uf RadioactivityFactory) FromPicorutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityPicorutherford)
}

// FromNanorutherfords creates a new Radioactivity instance from a value in Nanorutherfords.
func (uf RadioactivityFactory) FromNanorutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityNanorutherford)
}

// FromMicrorutherfords creates a new Radioactivity instance from a value in Microrutherfords.
func (uf RadioactivityFactory) FromMicrorutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMicrorutherford)
}

// FromMillirutherfords creates a new Radioactivity instance from a value in Millirutherfords.
func (uf RadioactivityFactory) FromMillirutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMillirutherford)
}

// FromKilorutherfords creates a new Radioactivity instance from a value in Kilorutherfords.
func (uf RadioactivityFactory) FromKilorutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityKilorutherford)
}

// FromMegarutherfords creates a new Radioactivity instance from a value in Megarutherfords.
func (uf RadioactivityFactory) FromMegarutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMegarutherford)
}

// FromGigarutherfords creates a new Radioactivity instance from a value in Gigarutherfords.
func (uf RadioactivityFactory) FromGigarutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityGigarutherford)
}

// FromTerarutherfords creates a new Radioactivity instance from a value in Terarutherfords.
func (uf RadioactivityFactory) FromTerarutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityTerarutherford)
}


// newRadioactivity creates a new Radioactivity.
func newRadioactivity(value float64, fromUnit RadioactivityUnits) (*Radioactivity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalRadioactivityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in RadioactivityUnits", fromUnit)
	}
	a := &Radioactivity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Radioactivity in Becquerel unit (the base/default quantity).
func (a *Radioactivity) BaseValue() float64 {
	return a.value
}


// Becquerels returns the Radioactivity value in Becquerels.
//
// Activity of a quantity of radioactive material in which one nucleus decays per second.
func (a *Radioactivity) Becquerels() float64 {
	if a.becquerelsLazy != nil {
		return *a.becquerelsLazy
	}
	becquerels := a.convertFromBase(RadioactivityBecquerel)
	a.becquerelsLazy = &becquerels
	return becquerels
}

// Curies returns the Radioactivity value in Curies.
//
// 
func (a *Radioactivity) Curies() float64 {
	if a.curiesLazy != nil {
		return *a.curiesLazy
	}
	curies := a.convertFromBase(RadioactivityCurie)
	a.curiesLazy = &curies
	return curies
}

// Rutherfords returns the Radioactivity value in Rutherfords.
//
// Activity of a quantity of radioactive material in which one million nuclei decay per second.
func (a *Radioactivity) Rutherfords() float64 {
	if a.rutherfordsLazy != nil {
		return *a.rutherfordsLazy
	}
	rutherfords := a.convertFromBase(RadioactivityRutherford)
	a.rutherfordsLazy = &rutherfords
	return rutherfords
}

// Picobecquerels returns the Radioactivity value in Picobecquerels.
//
// 
func (a *Radioactivity) Picobecquerels() float64 {
	if a.picobecquerelsLazy != nil {
		return *a.picobecquerelsLazy
	}
	picobecquerels := a.convertFromBase(RadioactivityPicobecquerel)
	a.picobecquerelsLazy = &picobecquerels
	return picobecquerels
}

// Nanobecquerels returns the Radioactivity value in Nanobecquerels.
//
// 
func (a *Radioactivity) Nanobecquerels() float64 {
	if a.nanobecquerelsLazy != nil {
		return *a.nanobecquerelsLazy
	}
	nanobecquerels := a.convertFromBase(RadioactivityNanobecquerel)
	a.nanobecquerelsLazy = &nanobecquerels
	return nanobecquerels
}

// Microbecquerels returns the Radioactivity value in Microbecquerels.
//
// 
func (a *Radioactivity) Microbecquerels() float64 {
	if a.microbecquerelsLazy != nil {
		return *a.microbecquerelsLazy
	}
	microbecquerels := a.convertFromBase(RadioactivityMicrobecquerel)
	a.microbecquerelsLazy = &microbecquerels
	return microbecquerels
}

// Millibecquerels returns the Radioactivity value in Millibecquerels.
//
// 
func (a *Radioactivity) Millibecquerels() float64 {
	if a.millibecquerelsLazy != nil {
		return *a.millibecquerelsLazy
	}
	millibecquerels := a.convertFromBase(RadioactivityMillibecquerel)
	a.millibecquerelsLazy = &millibecquerels
	return millibecquerels
}

// Kilobecquerels returns the Radioactivity value in Kilobecquerels.
//
// 
func (a *Radioactivity) Kilobecquerels() float64 {
	if a.kilobecquerelsLazy != nil {
		return *a.kilobecquerelsLazy
	}
	kilobecquerels := a.convertFromBase(RadioactivityKilobecquerel)
	a.kilobecquerelsLazy = &kilobecquerels
	return kilobecquerels
}

// Megabecquerels returns the Radioactivity value in Megabecquerels.
//
// 
func (a *Radioactivity) Megabecquerels() float64 {
	if a.megabecquerelsLazy != nil {
		return *a.megabecquerelsLazy
	}
	megabecquerels := a.convertFromBase(RadioactivityMegabecquerel)
	a.megabecquerelsLazy = &megabecquerels
	return megabecquerels
}

// Gigabecquerels returns the Radioactivity value in Gigabecquerels.
//
// 
func (a *Radioactivity) Gigabecquerels() float64 {
	if a.gigabecquerelsLazy != nil {
		return *a.gigabecquerelsLazy
	}
	gigabecquerels := a.convertFromBase(RadioactivityGigabecquerel)
	a.gigabecquerelsLazy = &gigabecquerels
	return gigabecquerels
}

// Terabecquerels returns the Radioactivity value in Terabecquerels.
//
// 
func (a *Radioactivity) Terabecquerels() float64 {
	if a.terabecquerelsLazy != nil {
		return *a.terabecquerelsLazy
	}
	terabecquerels := a.convertFromBase(RadioactivityTerabecquerel)
	a.terabecquerelsLazy = &terabecquerels
	return terabecquerels
}

// Petabecquerels returns the Radioactivity value in Petabecquerels.
//
// 
func (a *Radioactivity) Petabecquerels() float64 {
	if a.petabecquerelsLazy != nil {
		return *a.petabecquerelsLazy
	}
	petabecquerels := a.convertFromBase(RadioactivityPetabecquerel)
	a.petabecquerelsLazy = &petabecquerels
	return petabecquerels
}

// Exabecquerels returns the Radioactivity value in Exabecquerels.
//
// 
func (a *Radioactivity) Exabecquerels() float64 {
	if a.exabecquerelsLazy != nil {
		return *a.exabecquerelsLazy
	}
	exabecquerels := a.convertFromBase(RadioactivityExabecquerel)
	a.exabecquerelsLazy = &exabecquerels
	return exabecquerels
}

// Picocuries returns the Radioactivity value in Picocuries.
//
// 
func (a *Radioactivity) Picocuries() float64 {
	if a.picocuriesLazy != nil {
		return *a.picocuriesLazy
	}
	picocuries := a.convertFromBase(RadioactivityPicocurie)
	a.picocuriesLazy = &picocuries
	return picocuries
}

// Nanocuries returns the Radioactivity value in Nanocuries.
//
// 
func (a *Radioactivity) Nanocuries() float64 {
	if a.nanocuriesLazy != nil {
		return *a.nanocuriesLazy
	}
	nanocuries := a.convertFromBase(RadioactivityNanocurie)
	a.nanocuriesLazy = &nanocuries
	return nanocuries
}

// Microcuries returns the Radioactivity value in Microcuries.
//
// 
func (a *Radioactivity) Microcuries() float64 {
	if a.microcuriesLazy != nil {
		return *a.microcuriesLazy
	}
	microcuries := a.convertFromBase(RadioactivityMicrocurie)
	a.microcuriesLazy = &microcuries
	return microcuries
}

// Millicuries returns the Radioactivity value in Millicuries.
//
// 
func (a *Radioactivity) Millicuries() float64 {
	if a.millicuriesLazy != nil {
		return *a.millicuriesLazy
	}
	millicuries := a.convertFromBase(RadioactivityMillicurie)
	a.millicuriesLazy = &millicuries
	return millicuries
}

// Kilocuries returns the Radioactivity value in Kilocuries.
//
// 
func (a *Radioactivity) Kilocuries() float64 {
	if a.kilocuriesLazy != nil {
		return *a.kilocuriesLazy
	}
	kilocuries := a.convertFromBase(RadioactivityKilocurie)
	a.kilocuriesLazy = &kilocuries
	return kilocuries
}

// Megacuries returns the Radioactivity value in Megacuries.
//
// 
func (a *Radioactivity) Megacuries() float64 {
	if a.megacuriesLazy != nil {
		return *a.megacuriesLazy
	}
	megacuries := a.convertFromBase(RadioactivityMegacurie)
	a.megacuriesLazy = &megacuries
	return megacuries
}

// Gigacuries returns the Radioactivity value in Gigacuries.
//
// 
func (a *Radioactivity) Gigacuries() float64 {
	if a.gigacuriesLazy != nil {
		return *a.gigacuriesLazy
	}
	gigacuries := a.convertFromBase(RadioactivityGigacurie)
	a.gigacuriesLazy = &gigacuries
	return gigacuries
}

// Teracuries returns the Radioactivity value in Teracuries.
//
// 
func (a *Radioactivity) Teracuries() float64 {
	if a.teracuriesLazy != nil {
		return *a.teracuriesLazy
	}
	teracuries := a.convertFromBase(RadioactivityTeracurie)
	a.teracuriesLazy = &teracuries
	return teracuries
}

// Picorutherfords returns the Radioactivity value in Picorutherfords.
//
// 
func (a *Radioactivity) Picorutherfords() float64 {
	if a.picorutherfordsLazy != nil {
		return *a.picorutherfordsLazy
	}
	picorutherfords := a.convertFromBase(RadioactivityPicorutherford)
	a.picorutherfordsLazy = &picorutherfords
	return picorutherfords
}

// Nanorutherfords returns the Radioactivity value in Nanorutherfords.
//
// 
func (a *Radioactivity) Nanorutherfords() float64 {
	if a.nanorutherfordsLazy != nil {
		return *a.nanorutherfordsLazy
	}
	nanorutherfords := a.convertFromBase(RadioactivityNanorutherford)
	a.nanorutherfordsLazy = &nanorutherfords
	return nanorutherfords
}

// Microrutherfords returns the Radioactivity value in Microrutherfords.
//
// 
func (a *Radioactivity) Microrutherfords() float64 {
	if a.microrutherfordsLazy != nil {
		return *a.microrutherfordsLazy
	}
	microrutherfords := a.convertFromBase(RadioactivityMicrorutherford)
	a.microrutherfordsLazy = &microrutherfords
	return microrutherfords
}

// Millirutherfords returns the Radioactivity value in Millirutherfords.
//
// 
func (a *Radioactivity) Millirutherfords() float64 {
	if a.millirutherfordsLazy != nil {
		return *a.millirutherfordsLazy
	}
	millirutherfords := a.convertFromBase(RadioactivityMillirutherford)
	a.millirutherfordsLazy = &millirutherfords
	return millirutherfords
}

// Kilorutherfords returns the Radioactivity value in Kilorutherfords.
//
// 
func (a *Radioactivity) Kilorutherfords() float64 {
	if a.kilorutherfordsLazy != nil {
		return *a.kilorutherfordsLazy
	}
	kilorutherfords := a.convertFromBase(RadioactivityKilorutherford)
	a.kilorutherfordsLazy = &kilorutherfords
	return kilorutherfords
}

// Megarutherfords returns the Radioactivity value in Megarutherfords.
//
// 
func (a *Radioactivity) Megarutherfords() float64 {
	if a.megarutherfordsLazy != nil {
		return *a.megarutherfordsLazy
	}
	megarutherfords := a.convertFromBase(RadioactivityMegarutherford)
	a.megarutherfordsLazy = &megarutherfords
	return megarutherfords
}

// Gigarutherfords returns the Radioactivity value in Gigarutherfords.
//
// 
func (a *Radioactivity) Gigarutherfords() float64 {
	if a.gigarutherfordsLazy != nil {
		return *a.gigarutherfordsLazy
	}
	gigarutherfords := a.convertFromBase(RadioactivityGigarutherford)
	a.gigarutherfordsLazy = &gigarutherfords
	return gigarutherfords
}

// Terarutherfords returns the Radioactivity value in Terarutherfords.
//
// 
func (a *Radioactivity) Terarutherfords() float64 {
	if a.terarutherfordsLazy != nil {
		return *a.terarutherfordsLazy
	}
	terarutherfords := a.convertFromBase(RadioactivityTerarutherford)
	a.terarutherfordsLazy = &terarutherfords
	return terarutherfords
}


// ToDto creates a RadioactivityDto representation from the Radioactivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Becquerel by default.
func (a *Radioactivity) ToDto(holdInUnit *RadioactivityUnits) RadioactivityDto {
	if holdInUnit == nil {
		defaultUnit := RadioactivityBecquerel // Default value
		holdInUnit = &defaultUnit
	}

	return RadioactivityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Radioactivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Becquerel by default.
func (a *Radioactivity) ToDtoJSON(holdInUnit *RadioactivityUnits) ([]byte, error) {
	// Convert to RadioactivityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Radioactivity to a specific unit value.
// The function uses the provided unit type (RadioactivityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Radioactivity) Convert(toUnit RadioactivityUnits) float64 {
	switch toUnit { 
    case RadioactivityBecquerel:
		return a.Becquerels()
    case RadioactivityCurie:
		return a.Curies()
    case RadioactivityRutherford:
		return a.Rutherfords()
    case RadioactivityPicobecquerel:
		return a.Picobecquerels()
    case RadioactivityNanobecquerel:
		return a.Nanobecquerels()
    case RadioactivityMicrobecquerel:
		return a.Microbecquerels()
    case RadioactivityMillibecquerel:
		return a.Millibecquerels()
    case RadioactivityKilobecquerel:
		return a.Kilobecquerels()
    case RadioactivityMegabecquerel:
		return a.Megabecquerels()
    case RadioactivityGigabecquerel:
		return a.Gigabecquerels()
    case RadioactivityTerabecquerel:
		return a.Terabecquerels()
    case RadioactivityPetabecquerel:
		return a.Petabecquerels()
    case RadioactivityExabecquerel:
		return a.Exabecquerels()
    case RadioactivityPicocurie:
		return a.Picocuries()
    case RadioactivityNanocurie:
		return a.Nanocuries()
    case RadioactivityMicrocurie:
		return a.Microcuries()
    case RadioactivityMillicurie:
		return a.Millicuries()
    case RadioactivityKilocurie:
		return a.Kilocuries()
    case RadioactivityMegacurie:
		return a.Megacuries()
    case RadioactivityGigacurie:
		return a.Gigacuries()
    case RadioactivityTeracurie:
		return a.Teracuries()
    case RadioactivityPicorutherford:
		return a.Picorutherfords()
    case RadioactivityNanorutherford:
		return a.Nanorutherfords()
    case RadioactivityMicrorutherford:
		return a.Microrutherfords()
    case RadioactivityMillirutherford:
		return a.Millirutherfords()
    case RadioactivityKilorutherford:
		return a.Kilorutherfords()
    case RadioactivityMegarutherford:
		return a.Megarutherfords()
    case RadioactivityGigarutherford:
		return a.Gigarutherfords()
    case RadioactivityTerarutherford:
		return a.Terarutherfords()
	default:
		return math.NaN()
	}
}

func (a *Radioactivity) convertFromBase(toUnit RadioactivityUnits) float64 {
    value := a.value
	switch toUnit { 
	case RadioactivityBecquerel:
		return (value) 
	case RadioactivityCurie:
		return (value / 3.7e10) 
	case RadioactivityRutherford:
		return (value / 1e6) 
	case RadioactivityPicobecquerel:
		return ((value) / 1e-12) 
	case RadioactivityNanobecquerel:
		return ((value) / 1e-09) 
	case RadioactivityMicrobecquerel:
		return ((value) / 1e-06) 
	case RadioactivityMillibecquerel:
		return ((value) / 0.001) 
	case RadioactivityKilobecquerel:
		return ((value) / 1000.0) 
	case RadioactivityMegabecquerel:
		return ((value) / 1000000.0) 
	case RadioactivityGigabecquerel:
		return ((value) / 1000000000.0) 
	case RadioactivityTerabecquerel:
		return ((value) / 1000000000000.0) 
	case RadioactivityPetabecquerel:
		return ((value) / 1000000000000000.0) 
	case RadioactivityExabecquerel:
		return ((value) / 1e+18) 
	case RadioactivityPicocurie:
		return ((value / 3.7e10) / 1e-12) 
	case RadioactivityNanocurie:
		return ((value / 3.7e10) / 1e-09) 
	case RadioactivityMicrocurie:
		return ((value / 3.7e10) / 1e-06) 
	case RadioactivityMillicurie:
		return ((value / 3.7e10) / 0.001) 
	case RadioactivityKilocurie:
		return ((value / 3.7e10) / 1000.0) 
	case RadioactivityMegacurie:
		return ((value / 3.7e10) / 1000000.0) 
	case RadioactivityGigacurie:
		return ((value / 3.7e10) / 1000000000.0) 
	case RadioactivityTeracurie:
		return ((value / 3.7e10) / 1000000000000.0) 
	case RadioactivityPicorutherford:
		return ((value / 1e6) / 1e-12) 
	case RadioactivityNanorutherford:
		return ((value / 1e6) / 1e-09) 
	case RadioactivityMicrorutherford:
		return ((value / 1e6) / 1e-06) 
	case RadioactivityMillirutherford:
		return ((value / 1e6) / 0.001) 
	case RadioactivityKilorutherford:
		return ((value / 1e6) / 1000.0) 
	case RadioactivityMegarutherford:
		return ((value / 1e6) / 1000000.0) 
	case RadioactivityGigarutherford:
		return ((value / 1e6) / 1000000000.0) 
	case RadioactivityTerarutherford:
		return ((value / 1e6) / 1000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Radioactivity) convertToBase(value float64, fromUnit RadioactivityUnits) float64 {
	switch fromUnit { 
	case RadioactivityBecquerel:
		return (value) 
	case RadioactivityCurie:
		return (value * 3.7e10) 
	case RadioactivityRutherford:
		return (value * 1e6) 
	case RadioactivityPicobecquerel:
		return ((value) * 1e-12) 
	case RadioactivityNanobecquerel:
		return ((value) * 1e-09) 
	case RadioactivityMicrobecquerel:
		return ((value) * 1e-06) 
	case RadioactivityMillibecquerel:
		return ((value) * 0.001) 
	case RadioactivityKilobecquerel:
		return ((value) * 1000.0) 
	case RadioactivityMegabecquerel:
		return ((value) * 1000000.0) 
	case RadioactivityGigabecquerel:
		return ((value) * 1000000000.0) 
	case RadioactivityTerabecquerel:
		return ((value) * 1000000000000.0) 
	case RadioactivityPetabecquerel:
		return ((value) * 1000000000000000.0) 
	case RadioactivityExabecquerel:
		return ((value) * 1e+18) 
	case RadioactivityPicocurie:
		return ((value * 3.7e10) * 1e-12) 
	case RadioactivityNanocurie:
		return ((value * 3.7e10) * 1e-09) 
	case RadioactivityMicrocurie:
		return ((value * 3.7e10) * 1e-06) 
	case RadioactivityMillicurie:
		return ((value * 3.7e10) * 0.001) 
	case RadioactivityKilocurie:
		return ((value * 3.7e10) * 1000.0) 
	case RadioactivityMegacurie:
		return ((value * 3.7e10) * 1000000.0) 
	case RadioactivityGigacurie:
		return ((value * 3.7e10) * 1000000000.0) 
	case RadioactivityTeracurie:
		return ((value * 3.7e10) * 1000000000000.0) 
	case RadioactivityPicorutherford:
		return ((value * 1e6) * 1e-12) 
	case RadioactivityNanorutherford:
		return ((value * 1e6) * 1e-09) 
	case RadioactivityMicrorutherford:
		return ((value * 1e6) * 1e-06) 
	case RadioactivityMillirutherford:
		return ((value * 1e6) * 0.001) 
	case RadioactivityKilorutherford:
		return ((value * 1e6) * 1000.0) 
	case RadioactivityMegarutherford:
		return ((value * 1e6) * 1000000.0) 
	case RadioactivityGigarutherford:
		return ((value * 1e6) * 1000000000.0) 
	case RadioactivityTerarutherford:
		return ((value * 1e6) * 1000000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Radioactivity in the default unit (Becquerel),
// formatted to two decimal places.
func (a Radioactivity) String() string {
	return a.ToString(RadioactivityBecquerel, 2)
}

// ToString formats the Radioactivity value as a string with the specified unit and fractional digits.
// It converts the Radioactivity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Radioactivity value will be converted (e.g., Becquerel))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Radioactivity with the unit abbreviation.
func (a *Radioactivity) ToString(unit RadioactivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetRadioactivityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetRadioactivityAbbreviation(unit))
}

// Equals checks if the given Radioactivity is equal to the current Radioactivity.
//
// Parameters:
//    other: The Radioactivity to compare against.
//
// Returns:
//    bool: Returns true if both Radioactivity are equal, false otherwise.
func (a *Radioactivity) Equals(other *Radioactivity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Radioactivity with another Radioactivity.
// It returns -1 if the current Radioactivity is less than the other Radioactivity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Radioactivity to compare against.
//
// Returns:
//    int: -1 if the current Radioactivity is less, 1 if greater, and 0 if equal.
func (a *Radioactivity) CompareTo(other *Radioactivity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Radioactivity to the current Radioactivity and returns the result.
// The result is a new Radioactivity instance with the sum of the values.
//
// Parameters:
//    other: The Radioactivity to add to the current Radioactivity.
//
// Returns:
//    *Radioactivity: A new Radioactivity instance representing the sum of both Radioactivity.
func (a *Radioactivity) Add(other *Radioactivity) *Radioactivity {
	return &Radioactivity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Radioactivity from the current Radioactivity and returns the result.
// The result is a new Radioactivity instance with the difference of the values.
//
// Parameters:
//    other: The Radioactivity to subtract from the current Radioactivity.
//
// Returns:
//    *Radioactivity: A new Radioactivity instance representing the difference of both Radioactivity.
func (a *Radioactivity) Subtract(other *Radioactivity) *Radioactivity {
	return &Radioactivity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Radioactivity by the given Radioactivity and returns the result.
// The result is a new Radioactivity instance with the product of the values.
//
// Parameters:
//    other: The Radioactivity to multiply with the current Radioactivity.
//
// Returns:
//    *Radioactivity: A new Radioactivity instance representing the product of both Radioactivity.
func (a *Radioactivity) Multiply(other *Radioactivity) *Radioactivity {
	return &Radioactivity{value: a.value * other.BaseValue()}
}

// Divide divides the current Radioactivity by the given Radioactivity and returns the result.
// The result is a new Radioactivity instance with the quotient of the values.
//
// Parameters:
//    other: The Radioactivity to divide the current Radioactivity by.
//
// Returns:
//    *Radioactivity: A new Radioactivity instance representing the quotient of both Radioactivity.
func (a *Radioactivity) Divide(other *Radioactivity) *Radioactivity {
	return &Radioactivity{value: a.value / other.BaseValue()}
}

// GetRadioactivityAbbreviation gets the unit abbreviation.
func GetRadioactivityAbbreviation(unit RadioactivityUnits) string {
	switch unit { 
	case RadioactivityBecquerel:
		return "Bq" 
	case RadioactivityCurie:
		return "Ci" 
	case RadioactivityRutherford:
		return "Rd" 
	case RadioactivityPicobecquerel:
		return "pBq" 
	case RadioactivityNanobecquerel:
		return "nBq" 
	case RadioactivityMicrobecquerel:
		return "μBq" 
	case RadioactivityMillibecquerel:
		return "mBq" 
	case RadioactivityKilobecquerel:
		return "kBq" 
	case RadioactivityMegabecquerel:
		return "MBq" 
	case RadioactivityGigabecquerel:
		return "GBq" 
	case RadioactivityTerabecquerel:
		return "TBq" 
	case RadioactivityPetabecquerel:
		return "PBq" 
	case RadioactivityExabecquerel:
		return "EBq" 
	case RadioactivityPicocurie:
		return "pCi" 
	case RadioactivityNanocurie:
		return "nCi" 
	case RadioactivityMicrocurie:
		return "μCi" 
	case RadioactivityMillicurie:
		return "mCi" 
	case RadioactivityKilocurie:
		return "kCi" 
	case RadioactivityMegacurie:
		return "MCi" 
	case RadioactivityGigacurie:
		return "GCi" 
	case RadioactivityTeracurie:
		return "TCi" 
	case RadioactivityPicorutherford:
		return "pRd" 
	case RadioactivityNanorutherford:
		return "nRd" 
	case RadioactivityMicrorutherford:
		return "μRd" 
	case RadioactivityMillirutherford:
		return "mRd" 
	case RadioactivityKilorutherford:
		return "kRd" 
	case RadioactivityMegarutherford:
		return "MRd" 
	case RadioactivityGigarutherford:
		return "GRd" 
	case RadioactivityTerarutherford:
		return "TRd" 
	default:
		return ""
	}
}