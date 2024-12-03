// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RadioactivityUnits enumeration
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

// RadioactivityDto represents an Radioactivity
type RadioactivityDto struct {
	Value float64
	Unit  RadioactivityUnits
}

// RadioactivityDtoFactory struct to group related functions
type RadioactivityDtoFactory struct{}

func (udf RadioactivityDtoFactory) FromJSON(data []byte) (*RadioactivityDto, error) {
	a := RadioactivityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RadioactivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Radioactivity struct
type Radioactivity struct {
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

// RadioactivityFactory struct to group related functions
type RadioactivityFactory struct{}

func (uf RadioactivityFactory) CreateRadioactivity(value float64, unit RadioactivityUnits) (*Radioactivity, error) {
	return newRadioactivity(value, unit)
}

func (uf RadioactivityFactory) FromDto(dto RadioactivityDto) (*Radioactivity, error) {
	return newRadioactivity(dto.Value, dto.Unit)
}

func (uf RadioactivityFactory) FromDtoJSON(data []byte) (*Radioactivity, error) {
	unitDto, err := RadioactivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RadioactivityFactory{}.FromDto(*unitDto)
}


// FromBecquerel creates a new Radioactivity instance from Becquerel.
func (uf RadioactivityFactory) FromBecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityBecquerel)
}

// FromCurie creates a new Radioactivity instance from Curie.
func (uf RadioactivityFactory) FromCuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityCurie)
}

// FromRutherford creates a new Radioactivity instance from Rutherford.
func (uf RadioactivityFactory) FromRutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityRutherford)
}

// FromPicobecquerel creates a new Radioactivity instance from Picobecquerel.
func (uf RadioactivityFactory) FromPicobecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityPicobecquerel)
}

// FromNanobecquerel creates a new Radioactivity instance from Nanobecquerel.
func (uf RadioactivityFactory) FromNanobecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityNanobecquerel)
}

// FromMicrobecquerel creates a new Radioactivity instance from Microbecquerel.
func (uf RadioactivityFactory) FromMicrobecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMicrobecquerel)
}

// FromMillibecquerel creates a new Radioactivity instance from Millibecquerel.
func (uf RadioactivityFactory) FromMillibecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMillibecquerel)
}

// FromKilobecquerel creates a new Radioactivity instance from Kilobecquerel.
func (uf RadioactivityFactory) FromKilobecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityKilobecquerel)
}

// FromMegabecquerel creates a new Radioactivity instance from Megabecquerel.
func (uf RadioactivityFactory) FromMegabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMegabecquerel)
}

// FromGigabecquerel creates a new Radioactivity instance from Gigabecquerel.
func (uf RadioactivityFactory) FromGigabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityGigabecquerel)
}

// FromTerabecquerel creates a new Radioactivity instance from Terabecquerel.
func (uf RadioactivityFactory) FromTerabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityTerabecquerel)
}

// FromPetabecquerel creates a new Radioactivity instance from Petabecquerel.
func (uf RadioactivityFactory) FromPetabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityPetabecquerel)
}

// FromExabecquerel creates a new Radioactivity instance from Exabecquerel.
func (uf RadioactivityFactory) FromExabecquerels(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityExabecquerel)
}

// FromPicocurie creates a new Radioactivity instance from Picocurie.
func (uf RadioactivityFactory) FromPicocuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityPicocurie)
}

// FromNanocurie creates a new Radioactivity instance from Nanocurie.
func (uf RadioactivityFactory) FromNanocuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityNanocurie)
}

// FromMicrocurie creates a new Radioactivity instance from Microcurie.
func (uf RadioactivityFactory) FromMicrocuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMicrocurie)
}

// FromMillicurie creates a new Radioactivity instance from Millicurie.
func (uf RadioactivityFactory) FromMillicuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMillicurie)
}

// FromKilocurie creates a new Radioactivity instance from Kilocurie.
func (uf RadioactivityFactory) FromKilocuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityKilocurie)
}

// FromMegacurie creates a new Radioactivity instance from Megacurie.
func (uf RadioactivityFactory) FromMegacuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMegacurie)
}

// FromGigacurie creates a new Radioactivity instance from Gigacurie.
func (uf RadioactivityFactory) FromGigacuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityGigacurie)
}

// FromTeracurie creates a new Radioactivity instance from Teracurie.
func (uf RadioactivityFactory) FromTeracuries(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityTeracurie)
}

// FromPicorutherford creates a new Radioactivity instance from Picorutherford.
func (uf RadioactivityFactory) FromPicorutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityPicorutherford)
}

// FromNanorutherford creates a new Radioactivity instance from Nanorutherford.
func (uf RadioactivityFactory) FromNanorutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityNanorutherford)
}

// FromMicrorutherford creates a new Radioactivity instance from Microrutherford.
func (uf RadioactivityFactory) FromMicrorutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMicrorutherford)
}

// FromMillirutherford creates a new Radioactivity instance from Millirutherford.
func (uf RadioactivityFactory) FromMillirutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMillirutherford)
}

// FromKilorutherford creates a new Radioactivity instance from Kilorutherford.
func (uf RadioactivityFactory) FromKilorutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityKilorutherford)
}

// FromMegarutherford creates a new Radioactivity instance from Megarutherford.
func (uf RadioactivityFactory) FromMegarutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityMegarutherford)
}

// FromGigarutherford creates a new Radioactivity instance from Gigarutherford.
func (uf RadioactivityFactory) FromGigarutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityGigarutherford)
}

// FromTerarutherford creates a new Radioactivity instance from Terarutherford.
func (uf RadioactivityFactory) FromTerarutherfords(value float64) (*Radioactivity, error) {
	return newRadioactivity(value, RadioactivityTerarutherford)
}




// newRadioactivity creates a new Radioactivity.
func newRadioactivity(value float64, fromUnit RadioactivityUnits) (*Radioactivity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Radioactivity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Radioactivity in Becquerel.
func (a *Radioactivity) BaseValue() float64 {
	return a.value
}


// Becquerel returns the value in Becquerel.
func (a *Radioactivity) Becquerels() float64 {
	if a.becquerelsLazy != nil {
		return *a.becquerelsLazy
	}
	becquerels := a.convertFromBase(RadioactivityBecquerel)
	a.becquerelsLazy = &becquerels
	return becquerels
}

// Curie returns the value in Curie.
func (a *Radioactivity) Curies() float64 {
	if a.curiesLazy != nil {
		return *a.curiesLazy
	}
	curies := a.convertFromBase(RadioactivityCurie)
	a.curiesLazy = &curies
	return curies
}

// Rutherford returns the value in Rutherford.
func (a *Radioactivity) Rutherfords() float64 {
	if a.rutherfordsLazy != nil {
		return *a.rutherfordsLazy
	}
	rutherfords := a.convertFromBase(RadioactivityRutherford)
	a.rutherfordsLazy = &rutherfords
	return rutherfords
}

// Picobecquerel returns the value in Picobecquerel.
func (a *Radioactivity) Picobecquerels() float64 {
	if a.picobecquerelsLazy != nil {
		return *a.picobecquerelsLazy
	}
	picobecquerels := a.convertFromBase(RadioactivityPicobecquerel)
	a.picobecquerelsLazy = &picobecquerels
	return picobecquerels
}

// Nanobecquerel returns the value in Nanobecquerel.
func (a *Radioactivity) Nanobecquerels() float64 {
	if a.nanobecquerelsLazy != nil {
		return *a.nanobecquerelsLazy
	}
	nanobecquerels := a.convertFromBase(RadioactivityNanobecquerel)
	a.nanobecquerelsLazy = &nanobecquerels
	return nanobecquerels
}

// Microbecquerel returns the value in Microbecquerel.
func (a *Radioactivity) Microbecquerels() float64 {
	if a.microbecquerelsLazy != nil {
		return *a.microbecquerelsLazy
	}
	microbecquerels := a.convertFromBase(RadioactivityMicrobecquerel)
	a.microbecquerelsLazy = &microbecquerels
	return microbecquerels
}

// Millibecquerel returns the value in Millibecquerel.
func (a *Radioactivity) Millibecquerels() float64 {
	if a.millibecquerelsLazy != nil {
		return *a.millibecquerelsLazy
	}
	millibecquerels := a.convertFromBase(RadioactivityMillibecquerel)
	a.millibecquerelsLazy = &millibecquerels
	return millibecquerels
}

// Kilobecquerel returns the value in Kilobecquerel.
func (a *Radioactivity) Kilobecquerels() float64 {
	if a.kilobecquerelsLazy != nil {
		return *a.kilobecquerelsLazy
	}
	kilobecquerels := a.convertFromBase(RadioactivityKilobecquerel)
	a.kilobecquerelsLazy = &kilobecquerels
	return kilobecquerels
}

// Megabecquerel returns the value in Megabecquerel.
func (a *Radioactivity) Megabecquerels() float64 {
	if a.megabecquerelsLazy != nil {
		return *a.megabecquerelsLazy
	}
	megabecquerels := a.convertFromBase(RadioactivityMegabecquerel)
	a.megabecquerelsLazy = &megabecquerels
	return megabecquerels
}

// Gigabecquerel returns the value in Gigabecquerel.
func (a *Radioactivity) Gigabecquerels() float64 {
	if a.gigabecquerelsLazy != nil {
		return *a.gigabecquerelsLazy
	}
	gigabecquerels := a.convertFromBase(RadioactivityGigabecquerel)
	a.gigabecquerelsLazy = &gigabecquerels
	return gigabecquerels
}

// Terabecquerel returns the value in Terabecquerel.
func (a *Radioactivity) Terabecquerels() float64 {
	if a.terabecquerelsLazy != nil {
		return *a.terabecquerelsLazy
	}
	terabecquerels := a.convertFromBase(RadioactivityTerabecquerel)
	a.terabecquerelsLazy = &terabecquerels
	return terabecquerels
}

// Petabecquerel returns the value in Petabecquerel.
func (a *Radioactivity) Petabecquerels() float64 {
	if a.petabecquerelsLazy != nil {
		return *a.petabecquerelsLazy
	}
	petabecquerels := a.convertFromBase(RadioactivityPetabecquerel)
	a.petabecquerelsLazy = &petabecquerels
	return petabecquerels
}

// Exabecquerel returns the value in Exabecquerel.
func (a *Radioactivity) Exabecquerels() float64 {
	if a.exabecquerelsLazy != nil {
		return *a.exabecquerelsLazy
	}
	exabecquerels := a.convertFromBase(RadioactivityExabecquerel)
	a.exabecquerelsLazy = &exabecquerels
	return exabecquerels
}

// Picocurie returns the value in Picocurie.
func (a *Radioactivity) Picocuries() float64 {
	if a.picocuriesLazy != nil {
		return *a.picocuriesLazy
	}
	picocuries := a.convertFromBase(RadioactivityPicocurie)
	a.picocuriesLazy = &picocuries
	return picocuries
}

// Nanocurie returns the value in Nanocurie.
func (a *Radioactivity) Nanocuries() float64 {
	if a.nanocuriesLazy != nil {
		return *a.nanocuriesLazy
	}
	nanocuries := a.convertFromBase(RadioactivityNanocurie)
	a.nanocuriesLazy = &nanocuries
	return nanocuries
}

// Microcurie returns the value in Microcurie.
func (a *Radioactivity) Microcuries() float64 {
	if a.microcuriesLazy != nil {
		return *a.microcuriesLazy
	}
	microcuries := a.convertFromBase(RadioactivityMicrocurie)
	a.microcuriesLazy = &microcuries
	return microcuries
}

// Millicurie returns the value in Millicurie.
func (a *Radioactivity) Millicuries() float64 {
	if a.millicuriesLazy != nil {
		return *a.millicuriesLazy
	}
	millicuries := a.convertFromBase(RadioactivityMillicurie)
	a.millicuriesLazy = &millicuries
	return millicuries
}

// Kilocurie returns the value in Kilocurie.
func (a *Radioactivity) Kilocuries() float64 {
	if a.kilocuriesLazy != nil {
		return *a.kilocuriesLazy
	}
	kilocuries := a.convertFromBase(RadioactivityKilocurie)
	a.kilocuriesLazy = &kilocuries
	return kilocuries
}

// Megacurie returns the value in Megacurie.
func (a *Radioactivity) Megacuries() float64 {
	if a.megacuriesLazy != nil {
		return *a.megacuriesLazy
	}
	megacuries := a.convertFromBase(RadioactivityMegacurie)
	a.megacuriesLazy = &megacuries
	return megacuries
}

// Gigacurie returns the value in Gigacurie.
func (a *Radioactivity) Gigacuries() float64 {
	if a.gigacuriesLazy != nil {
		return *a.gigacuriesLazy
	}
	gigacuries := a.convertFromBase(RadioactivityGigacurie)
	a.gigacuriesLazy = &gigacuries
	return gigacuries
}

// Teracurie returns the value in Teracurie.
func (a *Radioactivity) Teracuries() float64 {
	if a.teracuriesLazy != nil {
		return *a.teracuriesLazy
	}
	teracuries := a.convertFromBase(RadioactivityTeracurie)
	a.teracuriesLazy = &teracuries
	return teracuries
}

// Picorutherford returns the value in Picorutherford.
func (a *Radioactivity) Picorutherfords() float64 {
	if a.picorutherfordsLazy != nil {
		return *a.picorutherfordsLazy
	}
	picorutherfords := a.convertFromBase(RadioactivityPicorutherford)
	a.picorutherfordsLazy = &picorutherfords
	return picorutherfords
}

// Nanorutherford returns the value in Nanorutherford.
func (a *Radioactivity) Nanorutherfords() float64 {
	if a.nanorutherfordsLazy != nil {
		return *a.nanorutherfordsLazy
	}
	nanorutherfords := a.convertFromBase(RadioactivityNanorutherford)
	a.nanorutherfordsLazy = &nanorutherfords
	return nanorutherfords
}

// Microrutherford returns the value in Microrutherford.
func (a *Radioactivity) Microrutherfords() float64 {
	if a.microrutherfordsLazy != nil {
		return *a.microrutherfordsLazy
	}
	microrutherfords := a.convertFromBase(RadioactivityMicrorutherford)
	a.microrutherfordsLazy = &microrutherfords
	return microrutherfords
}

// Millirutherford returns the value in Millirutherford.
func (a *Radioactivity) Millirutherfords() float64 {
	if a.millirutherfordsLazy != nil {
		return *a.millirutherfordsLazy
	}
	millirutherfords := a.convertFromBase(RadioactivityMillirutherford)
	a.millirutherfordsLazy = &millirutherfords
	return millirutherfords
}

// Kilorutherford returns the value in Kilorutherford.
func (a *Radioactivity) Kilorutherfords() float64 {
	if a.kilorutherfordsLazy != nil {
		return *a.kilorutherfordsLazy
	}
	kilorutherfords := a.convertFromBase(RadioactivityKilorutherford)
	a.kilorutherfordsLazy = &kilorutherfords
	return kilorutherfords
}

// Megarutherford returns the value in Megarutherford.
func (a *Radioactivity) Megarutherfords() float64 {
	if a.megarutherfordsLazy != nil {
		return *a.megarutherfordsLazy
	}
	megarutherfords := a.convertFromBase(RadioactivityMegarutherford)
	a.megarutherfordsLazy = &megarutherfords
	return megarutherfords
}

// Gigarutherford returns the value in Gigarutherford.
func (a *Radioactivity) Gigarutherfords() float64 {
	if a.gigarutherfordsLazy != nil {
		return *a.gigarutherfordsLazy
	}
	gigarutherfords := a.convertFromBase(RadioactivityGigarutherford)
	a.gigarutherfordsLazy = &gigarutherfords
	return gigarutherfords
}

// Terarutherford returns the value in Terarutherford.
func (a *Radioactivity) Terarutherfords() float64 {
	if a.terarutherfordsLazy != nil {
		return *a.terarutherfordsLazy
	}
	terarutherfords := a.convertFromBase(RadioactivityTerarutherford)
	a.terarutherfordsLazy = &terarutherfords
	return terarutherfords
}


// ToDto creates an RadioactivityDto representation.
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

// ToDtoJSON creates an RadioactivityDto representation.
func (a *Radioactivity) ToDtoJSON(holdInUnit *RadioactivityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Radioactivity to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Radioactivity) String() string {
	return a.ToString(RadioactivityBecquerel, 2)
}

// ToString formats the Radioactivity to string.
// fractionalDigits -1 for not mention
func (a *Radioactivity) ToString(unit RadioactivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Radioactivity) getUnitAbbreviation(unit RadioactivityUnits) string {
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

// Check if the given Radioactivity are equals to the current Radioactivity
func (a *Radioactivity) Equals(other *Radioactivity) bool {
	return a.value == other.BaseValue()
}

// Check if the given Radioactivity are equals to the current Radioactivity
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

// Add the given Radioactivity to the current Radioactivity.
func (a *Radioactivity) Add(other *Radioactivity) *Radioactivity {
	return &Radioactivity{value: a.value + other.BaseValue()}
}

// Subtract the given Radioactivity to the current Radioactivity.
func (a *Radioactivity) Subtract(other *Radioactivity) *Radioactivity {
	return &Radioactivity{value: a.value - other.BaseValue()}
}

// Multiply the given Radioactivity to the current Radioactivity.
func (a *Radioactivity) Multiply(other *Radioactivity) *Radioactivity {
	return &Radioactivity{value: a.value * other.BaseValue()}
}

// Divide the given Radioactivity to the current Radioactivity.
func (a *Radioactivity) Divide(other *Radioactivity) *Radioactivity {
	return &Radioactivity{value: a.value / other.BaseValue()}
}