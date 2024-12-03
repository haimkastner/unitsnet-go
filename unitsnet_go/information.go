package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// InformationUnits enumeration
type InformationUnits string

const (
    
        // 
        InformationByte InformationUnits = "Byte"
        // 
        InformationBit InformationUnits = "Bit"
        // 
        InformationKilobyte InformationUnits = "Kilobyte"
        // 
        InformationMegabyte InformationUnits = "Megabyte"
        // 
        InformationGigabyte InformationUnits = "Gigabyte"
        // 
        InformationTerabyte InformationUnits = "Terabyte"
        // 
        InformationPetabyte InformationUnits = "Petabyte"
        // 
        InformationExabyte InformationUnits = "Exabyte"
        // 
        InformationKibibyte InformationUnits = "Kibibyte"
        // 
        InformationMebibyte InformationUnits = "Mebibyte"
        // 
        InformationGibibyte InformationUnits = "Gibibyte"
        // 
        InformationTebibyte InformationUnits = "Tebibyte"
        // 
        InformationPebibyte InformationUnits = "Pebibyte"
        // 
        InformationExbibyte InformationUnits = "Exbibyte"
        // 
        InformationKilobit InformationUnits = "Kilobit"
        // 
        InformationMegabit InformationUnits = "Megabit"
        // 
        InformationGigabit InformationUnits = "Gigabit"
        // 
        InformationTerabit InformationUnits = "Terabit"
        // 
        InformationPetabit InformationUnits = "Petabit"
        // 
        InformationExabit InformationUnits = "Exabit"
        // 
        InformationKibibit InformationUnits = "Kibibit"
        // 
        InformationMebibit InformationUnits = "Mebibit"
        // 
        InformationGibibit InformationUnits = "Gibibit"
        // 
        InformationTebibit InformationUnits = "Tebibit"
        // 
        InformationPebibit InformationUnits = "Pebibit"
        // 
        InformationExbibit InformationUnits = "Exbibit"
)

// InformationDto represents an Information
type InformationDto struct {
	Value float64
	Unit  InformationUnits
}

// InformationDtoFactory struct to group related functions
type InformationDtoFactory struct{}

func (udf InformationDtoFactory) FromJSON(data []byte) (*InformationDto, error) {
	a := InformationDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a InformationDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Information struct
type Information struct {
	value       float64
    
    bytesLazy *float64 
    bitsLazy *float64 
    kilobytesLazy *float64 
    megabytesLazy *float64 
    gigabytesLazy *float64 
    terabytesLazy *float64 
    petabytesLazy *float64 
    exabytesLazy *float64 
    kibibytesLazy *float64 
    mebibytesLazy *float64 
    gibibytesLazy *float64 
    tebibytesLazy *float64 
    pebibytesLazy *float64 
    exbibytesLazy *float64 
    kilobitsLazy *float64 
    megabitsLazy *float64 
    gigabitsLazy *float64 
    terabitsLazy *float64 
    petabitsLazy *float64 
    exabitsLazy *float64 
    kibibitsLazy *float64 
    mebibitsLazy *float64 
    gibibitsLazy *float64 
    tebibitsLazy *float64 
    pebibitsLazy *float64 
    exbibitsLazy *float64 
}

// InformationFactory struct to group related functions
type InformationFactory struct{}

func (uf InformationFactory) CreateInformation(value float64, unit InformationUnits) (*Information, error) {
	return newInformation(value, unit)
}

func (uf InformationFactory) FromDto(dto InformationDto) (*Information, error) {
	return newInformation(dto.Value, dto.Unit)
}

func (uf InformationFactory) FromDtoJSON(data []byte) (*Information, error) {
	unitDto, err := InformationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return InformationFactory{}.FromDto(*unitDto)
}


// FromByte creates a new Information instance from Byte.
func (uf InformationFactory) FromBytes(value float64) (*Information, error) {
	return newInformation(value, InformationByte)
}

// FromBit creates a new Information instance from Bit.
func (uf InformationFactory) FromBits(value float64) (*Information, error) {
	return newInformation(value, InformationBit)
}

// FromKilobyte creates a new Information instance from Kilobyte.
func (uf InformationFactory) FromKilobytes(value float64) (*Information, error) {
	return newInformation(value, InformationKilobyte)
}

// FromMegabyte creates a new Information instance from Megabyte.
func (uf InformationFactory) FromMegabytes(value float64) (*Information, error) {
	return newInformation(value, InformationMegabyte)
}

// FromGigabyte creates a new Information instance from Gigabyte.
func (uf InformationFactory) FromGigabytes(value float64) (*Information, error) {
	return newInformation(value, InformationGigabyte)
}

// FromTerabyte creates a new Information instance from Terabyte.
func (uf InformationFactory) FromTerabytes(value float64) (*Information, error) {
	return newInformation(value, InformationTerabyte)
}

// FromPetabyte creates a new Information instance from Petabyte.
func (uf InformationFactory) FromPetabytes(value float64) (*Information, error) {
	return newInformation(value, InformationPetabyte)
}

// FromExabyte creates a new Information instance from Exabyte.
func (uf InformationFactory) FromExabytes(value float64) (*Information, error) {
	return newInformation(value, InformationExabyte)
}

// FromKibibyte creates a new Information instance from Kibibyte.
func (uf InformationFactory) FromKibibytes(value float64) (*Information, error) {
	return newInformation(value, InformationKibibyte)
}

// FromMebibyte creates a new Information instance from Mebibyte.
func (uf InformationFactory) FromMebibytes(value float64) (*Information, error) {
	return newInformation(value, InformationMebibyte)
}

// FromGibibyte creates a new Information instance from Gibibyte.
func (uf InformationFactory) FromGibibytes(value float64) (*Information, error) {
	return newInformation(value, InformationGibibyte)
}

// FromTebibyte creates a new Information instance from Tebibyte.
func (uf InformationFactory) FromTebibytes(value float64) (*Information, error) {
	return newInformation(value, InformationTebibyte)
}

// FromPebibyte creates a new Information instance from Pebibyte.
func (uf InformationFactory) FromPebibytes(value float64) (*Information, error) {
	return newInformation(value, InformationPebibyte)
}

// FromExbibyte creates a new Information instance from Exbibyte.
func (uf InformationFactory) FromExbibytes(value float64) (*Information, error) {
	return newInformation(value, InformationExbibyte)
}

// FromKilobit creates a new Information instance from Kilobit.
func (uf InformationFactory) FromKilobits(value float64) (*Information, error) {
	return newInformation(value, InformationKilobit)
}

// FromMegabit creates a new Information instance from Megabit.
func (uf InformationFactory) FromMegabits(value float64) (*Information, error) {
	return newInformation(value, InformationMegabit)
}

// FromGigabit creates a new Information instance from Gigabit.
func (uf InformationFactory) FromGigabits(value float64) (*Information, error) {
	return newInformation(value, InformationGigabit)
}

// FromTerabit creates a new Information instance from Terabit.
func (uf InformationFactory) FromTerabits(value float64) (*Information, error) {
	return newInformation(value, InformationTerabit)
}

// FromPetabit creates a new Information instance from Petabit.
func (uf InformationFactory) FromPetabits(value float64) (*Information, error) {
	return newInformation(value, InformationPetabit)
}

// FromExabit creates a new Information instance from Exabit.
func (uf InformationFactory) FromExabits(value float64) (*Information, error) {
	return newInformation(value, InformationExabit)
}

// FromKibibit creates a new Information instance from Kibibit.
func (uf InformationFactory) FromKibibits(value float64) (*Information, error) {
	return newInformation(value, InformationKibibit)
}

// FromMebibit creates a new Information instance from Mebibit.
func (uf InformationFactory) FromMebibits(value float64) (*Information, error) {
	return newInformation(value, InformationMebibit)
}

// FromGibibit creates a new Information instance from Gibibit.
func (uf InformationFactory) FromGibibits(value float64) (*Information, error) {
	return newInformation(value, InformationGibibit)
}

// FromTebibit creates a new Information instance from Tebibit.
func (uf InformationFactory) FromTebibits(value float64) (*Information, error) {
	return newInformation(value, InformationTebibit)
}

// FromPebibit creates a new Information instance from Pebibit.
func (uf InformationFactory) FromPebibits(value float64) (*Information, error) {
	return newInformation(value, InformationPebibit)
}

// FromExbibit creates a new Information instance from Exbibit.
func (uf InformationFactory) FromExbibits(value float64) (*Information, error) {
	return newInformation(value, InformationExbibit)
}




// newInformation creates a new Information.
func newInformation(value float64, fromUnit InformationUnits) (*Information, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Information{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Information in Bit.
func (a *Information) BaseValue() float64 {
	return a.value
}


// Byte returns the value in Byte.
func (a *Information) Bytes() float64 {
	if a.bytesLazy != nil {
		return *a.bytesLazy
	}
	bytes := a.convertFromBase(InformationByte)
	a.bytesLazy = &bytes
	return bytes
}

// Bit returns the value in Bit.
func (a *Information) Bits() float64 {
	if a.bitsLazy != nil {
		return *a.bitsLazy
	}
	bits := a.convertFromBase(InformationBit)
	a.bitsLazy = &bits
	return bits
}

// Kilobyte returns the value in Kilobyte.
func (a *Information) Kilobytes() float64 {
	if a.kilobytesLazy != nil {
		return *a.kilobytesLazy
	}
	kilobytes := a.convertFromBase(InformationKilobyte)
	a.kilobytesLazy = &kilobytes
	return kilobytes
}

// Megabyte returns the value in Megabyte.
func (a *Information) Megabytes() float64 {
	if a.megabytesLazy != nil {
		return *a.megabytesLazy
	}
	megabytes := a.convertFromBase(InformationMegabyte)
	a.megabytesLazy = &megabytes
	return megabytes
}

// Gigabyte returns the value in Gigabyte.
func (a *Information) Gigabytes() float64 {
	if a.gigabytesLazy != nil {
		return *a.gigabytesLazy
	}
	gigabytes := a.convertFromBase(InformationGigabyte)
	a.gigabytesLazy = &gigabytes
	return gigabytes
}

// Terabyte returns the value in Terabyte.
func (a *Information) Terabytes() float64 {
	if a.terabytesLazy != nil {
		return *a.terabytesLazy
	}
	terabytes := a.convertFromBase(InformationTerabyte)
	a.terabytesLazy = &terabytes
	return terabytes
}

// Petabyte returns the value in Petabyte.
func (a *Information) Petabytes() float64 {
	if a.petabytesLazy != nil {
		return *a.petabytesLazy
	}
	petabytes := a.convertFromBase(InformationPetabyte)
	a.petabytesLazy = &petabytes
	return petabytes
}

// Exabyte returns the value in Exabyte.
func (a *Information) Exabytes() float64 {
	if a.exabytesLazy != nil {
		return *a.exabytesLazy
	}
	exabytes := a.convertFromBase(InformationExabyte)
	a.exabytesLazy = &exabytes
	return exabytes
}

// Kibibyte returns the value in Kibibyte.
func (a *Information) Kibibytes() float64 {
	if a.kibibytesLazy != nil {
		return *a.kibibytesLazy
	}
	kibibytes := a.convertFromBase(InformationKibibyte)
	a.kibibytesLazy = &kibibytes
	return kibibytes
}

// Mebibyte returns the value in Mebibyte.
func (a *Information) Mebibytes() float64 {
	if a.mebibytesLazy != nil {
		return *a.mebibytesLazy
	}
	mebibytes := a.convertFromBase(InformationMebibyte)
	a.mebibytesLazy = &mebibytes
	return mebibytes
}

// Gibibyte returns the value in Gibibyte.
func (a *Information) Gibibytes() float64 {
	if a.gibibytesLazy != nil {
		return *a.gibibytesLazy
	}
	gibibytes := a.convertFromBase(InformationGibibyte)
	a.gibibytesLazy = &gibibytes
	return gibibytes
}

// Tebibyte returns the value in Tebibyte.
func (a *Information) Tebibytes() float64 {
	if a.tebibytesLazy != nil {
		return *a.tebibytesLazy
	}
	tebibytes := a.convertFromBase(InformationTebibyte)
	a.tebibytesLazy = &tebibytes
	return tebibytes
}

// Pebibyte returns the value in Pebibyte.
func (a *Information) Pebibytes() float64 {
	if a.pebibytesLazy != nil {
		return *a.pebibytesLazy
	}
	pebibytes := a.convertFromBase(InformationPebibyte)
	a.pebibytesLazy = &pebibytes
	return pebibytes
}

// Exbibyte returns the value in Exbibyte.
func (a *Information) Exbibytes() float64 {
	if a.exbibytesLazy != nil {
		return *a.exbibytesLazy
	}
	exbibytes := a.convertFromBase(InformationExbibyte)
	a.exbibytesLazy = &exbibytes
	return exbibytes
}

// Kilobit returns the value in Kilobit.
func (a *Information) Kilobits() float64 {
	if a.kilobitsLazy != nil {
		return *a.kilobitsLazy
	}
	kilobits := a.convertFromBase(InformationKilobit)
	a.kilobitsLazy = &kilobits
	return kilobits
}

// Megabit returns the value in Megabit.
func (a *Information) Megabits() float64 {
	if a.megabitsLazy != nil {
		return *a.megabitsLazy
	}
	megabits := a.convertFromBase(InformationMegabit)
	a.megabitsLazy = &megabits
	return megabits
}

// Gigabit returns the value in Gigabit.
func (a *Information) Gigabits() float64 {
	if a.gigabitsLazy != nil {
		return *a.gigabitsLazy
	}
	gigabits := a.convertFromBase(InformationGigabit)
	a.gigabitsLazy = &gigabits
	return gigabits
}

// Terabit returns the value in Terabit.
func (a *Information) Terabits() float64 {
	if a.terabitsLazy != nil {
		return *a.terabitsLazy
	}
	terabits := a.convertFromBase(InformationTerabit)
	a.terabitsLazy = &terabits
	return terabits
}

// Petabit returns the value in Petabit.
func (a *Information) Petabits() float64 {
	if a.petabitsLazy != nil {
		return *a.petabitsLazy
	}
	petabits := a.convertFromBase(InformationPetabit)
	a.petabitsLazy = &petabits
	return petabits
}

// Exabit returns the value in Exabit.
func (a *Information) Exabits() float64 {
	if a.exabitsLazy != nil {
		return *a.exabitsLazy
	}
	exabits := a.convertFromBase(InformationExabit)
	a.exabitsLazy = &exabits
	return exabits
}

// Kibibit returns the value in Kibibit.
func (a *Information) Kibibits() float64 {
	if a.kibibitsLazy != nil {
		return *a.kibibitsLazy
	}
	kibibits := a.convertFromBase(InformationKibibit)
	a.kibibitsLazy = &kibibits
	return kibibits
}

// Mebibit returns the value in Mebibit.
func (a *Information) Mebibits() float64 {
	if a.mebibitsLazy != nil {
		return *a.mebibitsLazy
	}
	mebibits := a.convertFromBase(InformationMebibit)
	a.mebibitsLazy = &mebibits
	return mebibits
}

// Gibibit returns the value in Gibibit.
func (a *Information) Gibibits() float64 {
	if a.gibibitsLazy != nil {
		return *a.gibibitsLazy
	}
	gibibits := a.convertFromBase(InformationGibibit)
	a.gibibitsLazy = &gibibits
	return gibibits
}

// Tebibit returns the value in Tebibit.
func (a *Information) Tebibits() float64 {
	if a.tebibitsLazy != nil {
		return *a.tebibitsLazy
	}
	tebibits := a.convertFromBase(InformationTebibit)
	a.tebibitsLazy = &tebibits
	return tebibits
}

// Pebibit returns the value in Pebibit.
func (a *Information) Pebibits() float64 {
	if a.pebibitsLazy != nil {
		return *a.pebibitsLazy
	}
	pebibits := a.convertFromBase(InformationPebibit)
	a.pebibitsLazy = &pebibits
	return pebibits
}

// Exbibit returns the value in Exbibit.
func (a *Information) Exbibits() float64 {
	if a.exbibitsLazy != nil {
		return *a.exbibitsLazy
	}
	exbibits := a.convertFromBase(InformationExbibit)
	a.exbibitsLazy = &exbibits
	return exbibits
}


// ToDto creates an InformationDto representation.
func (a *Information) ToDto(holdInUnit *InformationUnits) InformationDto {
	if holdInUnit == nil {
		defaultUnit := InformationBit // Default value
		holdInUnit = &defaultUnit
	}

	return InformationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an InformationDto representation.
func (a *Information) ToDtoJSON(holdInUnit *InformationUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Information to a specific unit value.
func (a *Information) Convert(toUnit InformationUnits) float64 {
	switch toUnit { 
    case InformationByte:
		return a.Bytes()
    case InformationBit:
		return a.Bits()
    case InformationKilobyte:
		return a.Kilobytes()
    case InformationMegabyte:
		return a.Megabytes()
    case InformationGigabyte:
		return a.Gigabytes()
    case InformationTerabyte:
		return a.Terabytes()
    case InformationPetabyte:
		return a.Petabytes()
    case InformationExabyte:
		return a.Exabytes()
    case InformationKibibyte:
		return a.Kibibytes()
    case InformationMebibyte:
		return a.Mebibytes()
    case InformationGibibyte:
		return a.Gibibytes()
    case InformationTebibyte:
		return a.Tebibytes()
    case InformationPebibyte:
		return a.Pebibytes()
    case InformationExbibyte:
		return a.Exbibytes()
    case InformationKilobit:
		return a.Kilobits()
    case InformationMegabit:
		return a.Megabits()
    case InformationGigabit:
		return a.Gigabits()
    case InformationTerabit:
		return a.Terabits()
    case InformationPetabit:
		return a.Petabits()
    case InformationExabit:
		return a.Exabits()
    case InformationKibibit:
		return a.Kibibits()
    case InformationMebibit:
		return a.Mebibits()
    case InformationGibibit:
		return a.Gibibits()
    case InformationTebibit:
		return a.Tebibits()
    case InformationPebibit:
		return a.Pebibits()
    case InformationExbibit:
		return a.Exbibits()
	default:
		return 0
	}
}

func (a *Information) convertFromBase(toUnit InformationUnits) float64 {
    value := a.value
	switch toUnit { 
	case InformationByte:
		return (value / 8) 
	case InformationBit:
		return (value) 
	case InformationKilobyte:
		return ((value / 8) / 1000.0) 
	case InformationMegabyte:
		return ((value / 8) / 1000000.0) 
	case InformationGigabyte:
		return ((value / 8) / 1000000000.0) 
	case InformationTerabyte:
		return ((value / 8) / 1000000000000.0) 
	case InformationPetabyte:
		return ((value / 8) / 1000000000000000.0) 
	case InformationExabyte:
		return ((value / 8) / 1e+18) 
	case InformationKibibyte:
		return ((value / 8) / 1024) 
	case InformationMebibyte:
		return ((value / 8) / 1048576) 
	case InformationGibibyte:
		return ((value / 8) / 1073741824) 
	case InformationTebibyte:
		return ((value / 8) / 1099511627776) 
	case InformationPebibyte:
		return ((value / 8) / 1125899906842624) 
	case InformationExbibyte:
		return ((value / 8) / 1152921504606846976) 
	case InformationKilobit:
		return ((value) / 1000.0) 
	case InformationMegabit:
		return ((value) / 1000000.0) 
	case InformationGigabit:
		return ((value) / 1000000000.0) 
	case InformationTerabit:
		return ((value) / 1000000000000.0) 
	case InformationPetabit:
		return ((value) / 1000000000000000.0) 
	case InformationExabit:
		return ((value) / 1e+18) 
	case InformationKibibit:
		return ((value) / 1024) 
	case InformationMebibit:
		return ((value) / 1048576) 
	case InformationGibibit:
		return ((value) / 1073741824) 
	case InformationTebibit:
		return ((value) / 1099511627776) 
	case InformationPebibit:
		return ((value) / 1125899906842624) 
	case InformationExbibit:
		return ((value) / 1152921504606846976) 
	default:
		return math.NaN()
	}
}

func (a *Information) convertToBase(value float64, fromUnit InformationUnits) float64 {
	switch fromUnit { 
	case InformationByte:
		return (value * 8) 
	case InformationBit:
		return (value) 
	case InformationKilobyte:
		return ((value * 8) * 1000.0) 
	case InformationMegabyte:
		return ((value * 8) * 1000000.0) 
	case InformationGigabyte:
		return ((value * 8) * 1000000000.0) 
	case InformationTerabyte:
		return ((value * 8) * 1000000000000.0) 
	case InformationPetabyte:
		return ((value * 8) * 1000000000000000.0) 
	case InformationExabyte:
		return ((value * 8) * 1e+18) 
	case InformationKibibyte:
		return ((value * 8) * 1024) 
	case InformationMebibyte:
		return ((value * 8) * 1048576) 
	case InformationGibibyte:
		return ((value * 8) * 1073741824) 
	case InformationTebibyte:
		return ((value * 8) * 1099511627776) 
	case InformationPebibyte:
		return ((value * 8) * 1125899906842624) 
	case InformationExbibyte:
		return ((value * 8) * 1152921504606846976) 
	case InformationKilobit:
		return ((value) * 1000.0) 
	case InformationMegabit:
		return ((value) * 1000000.0) 
	case InformationGigabit:
		return ((value) * 1000000000.0) 
	case InformationTerabit:
		return ((value) * 1000000000000.0) 
	case InformationPetabit:
		return ((value) * 1000000000000000.0) 
	case InformationExabit:
		return ((value) * 1e+18) 
	case InformationKibibit:
		return ((value) * 1024) 
	case InformationMebibit:
		return ((value) * 1048576) 
	case InformationGibibit:
		return ((value) * 1073741824) 
	case InformationTebibit:
		return ((value) * 1099511627776) 
	case InformationPebibit:
		return ((value) * 1125899906842624) 
	case InformationExbibit:
		return ((value) * 1152921504606846976) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Information) String() string {
	return a.ToString(InformationBit, 2)
}

// ToString formats the Information to string.
// fractionalDigits -1 for not mention
func (a *Information) ToString(unit InformationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Information) getUnitAbbreviation(unit InformationUnits) string {
	switch unit { 
	case InformationByte:
		return "B" 
	case InformationBit:
		return "b" 
	case InformationKilobyte:
		return "kB" 
	case InformationMegabyte:
		return "MB" 
	case InformationGigabyte:
		return "GB" 
	case InformationTerabyte:
		return "TB" 
	case InformationPetabyte:
		return "PB" 
	case InformationExabyte:
		return "EB" 
	case InformationKibibyte:
		return "KiBB" 
	case InformationMebibyte:
		return "MiBB" 
	case InformationGibibyte:
		return "GiBB" 
	case InformationTebibyte:
		return "TiBB" 
	case InformationPebibyte:
		return "PiBB" 
	case InformationExbibyte:
		return "EiBB" 
	case InformationKilobit:
		return "kb" 
	case InformationMegabit:
		return "Mb" 
	case InformationGigabit:
		return "Gb" 
	case InformationTerabit:
		return "Tb" 
	case InformationPetabit:
		return "Pb" 
	case InformationExabit:
		return "Eb" 
	case InformationKibibit:
		return "KiBb" 
	case InformationMebibit:
		return "MiBb" 
	case InformationGibibit:
		return "GiBb" 
	case InformationTebibit:
		return "TiBb" 
	case InformationPebibit:
		return "PiBb" 
	case InformationExbibit:
		return "EiBb" 
	default:
		return ""
	}
}

// Check if the given Information are equals to the current Information
func (a *Information) Equals(other *Information) bool {
	return a.value == other.BaseValue()
}

// Check if the given Information are equals to the current Information
func (a *Information) CompareTo(other *Information) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Information to the current Information.
func (a *Information) Add(other *Information) *Information {
	return &Information{value: a.value + other.BaseValue()}
}

// Subtract the given Information to the current Information.
func (a *Information) Subtract(other *Information) *Information {
	return &Information{value: a.value - other.BaseValue()}
}

// Multiply the given Information to the current Information.
func (a *Information) Multiply(other *Information) *Information {
	return &Information{value: a.value * other.BaseValue()}
}

// Divide the given Information to the current Information.
func (a *Information) Divide(other *Information) *Information {
	return &Information{value: a.value / other.BaseValue()}
}