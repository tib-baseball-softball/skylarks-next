package dp

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
)

const (
	LocationCollection = "locations"
)

var _ core.RecordProxy = (*Location)(nil)

// Location RecordProxy for collection `locations`.
type Location struct {
	core.BaseRecordProxy
}

func (l *Location) CollectionName() string {
	return LocationCollection
}

func (l *Location) ID() string {
	return l.Id
}

func (l *Location) BSMID() int {
	return l.GetInt("bsm_id")
}

func (l *Location) SetBSMID(bsmID int) {
	l.Set("bsm_id", bsmID)
}

// Name contains mostly useless data like "Baseball"
// there is also `internal_name` in the record model to have a field that automatic imports do not overwrite
func (l *Location) Name() string {
	return l.GetString("name")
}

func (l *Location) SetName(name string) {
	l.Set("name", name)
}

func (l *Location) InternalName() string {
	return l.GetString("internal_name")
}

func (l *Location) SetInternalName(internalName string) {
	l.Set("internal_name", internalName)
}

func (l *Location) Description() string {
	return l.GetString("description")
}

func (l *Location) SetDescription(description string) {
	l.Set("description", description)
}

// AddressAddon for many BSM datasets contains the actual name of the field
func (l *Location) AddressAddon() string {
	return l.GetString("address_addon")
}

func (l *Location) SetAddressAddon(addressAddon string) {
	l.Set("address_addon", addressAddon)
}

func (l *Location) Street() string {
	return l.GetString("street")
}

func (l *Location) SetStreet(street string) {
	l.Set("street", street)
}

func (l *Location) PostalCode() string {
	return l.GetString("postal_code")
}

func (l *Location) SetPostalCode(postalCode string) {
	l.Set("postal_code", postalCode)
}

func (l *Location) Country() string {
	return l.GetString("country")
}

func (l *Location) SetCountry(country string) {
	l.Set("country", country)
}

func (l *Location) City() string {
	return l.GetString("city")
}

func (l *Location) SetCity(city string) {
	l.Set("city", city)
}

func (l *Location) Longitude() float64 {
	return l.GetFloat("longitude")
}

func (l *Location) SetLongitude(longitude float64) {
	l.Set("longitude", longitude)
}

func (l *Location) Latitude() float64 {
	return l.GetFloat("latitude")
}

func (l *Location) SetLatitude(latitude float64) {
	l.Set("latitude", latitude)
}

func (l *Location) SpectatorTotal() int {
	return l.GetInt("spectator_total")
}

func (l *Location) SetSpectatorTotal(spectatorTotal int) {
	l.Set("spectator_total", spectatorTotal)
}

func (l *Location) SpectatorSeats() int {
	return l.GetInt("spectator_seats")
}

func (l *Location) SetSpectatorSeats(spectatorSeats int) {
	l.Set("spectator_seats", spectatorSeats)
}

func (l *Location) OtherInformation() string {
	return l.GetString("other_information")
}

func (l *Location) SetOtherInformation(otherInformation string) {
	l.Set("other_information", otherInformation)
}

func (l *Location) GroundRules() string {
	return l.GetString("groundrules")
}

func (l *Location) SetGroundRules(groundRules string) {
	l.Set("groundrules", groundRules)
}

func (l *Location) HumanCountry() string {
	return l.GetString("human_country")
}

func (l *Location) SetHumanCountry(humanCountry string) {
	l.Set("human_country", humanCountry)
}

func (l *Location) PhotoURL() string {
	return l.GetString("photo_url")
}

func (l *Location) SetPhotoURL(photoURL string) {
	l.Set("photo_url", photoURL)
}

func (l *Location) Club() string {
	return l.GetString("club")
}

func (l *Location) SetClub(club string) {
	l.Set("club", club)
}

func (l *Location) String() string {
	return l.GetCalendarFormatted()
}

func (l *Location) GetCalendarFormatted() string {
	return fmt.Sprintf("%s (%s), %s %s, %s, %s, Lat: %f, Lng: %f", l.Name(), l.AddressAddon(), l.Street(), l.PostalCode(), l.City(), l.HumanCountry(), l.Latitude(), l.Longitude())
}

type LocationDTO struct {
	Name         string  `json:"name"`
	AddressAddon string  `json:"address_addon"`
	InternalName string  `json:"internal_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Street       string  `json:"street"`
	PostalCode   string  `json:"postal_code"`
	City         string  `json:"city"`
	Country      string  `json:"country"`
	BSMID        int     `json:"bsm_id"`
}

func (l *Location) ToDTO() *LocationDTO {
	return &LocationDTO{
		Name:         l.Name(),
		AddressAddon: l.AddressAddon(),
		InternalName: l.InternalName(),
		Latitude:     l.Latitude(),
		Longitude:    l.Longitude(),
		Street:       l.Street(),
		PostalCode:   l.PostalCode(),
		City:         l.City(),
		Country:      l.Country(),
		BSMID:        l.BSMID(),
	}
}
