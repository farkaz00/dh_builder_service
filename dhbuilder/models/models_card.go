package models

type CardRealm string

const (
	Realmless        CardRealm = "REALMLESS"
	DistortedShadows CardRealm = "DISTORTED_SHADOWS"
	PiercingWinds    CardRealm = "PIERCING_WINDS"
	HostilePlains    CardRealm = "HOSTILE_PLAINS"
	GloomyWaters     CardRealm = "GLOOMY_WATERS"
)

// Card defines the data used to represent DH cards
type Card struct {
	ID           string    `json:"id"`
	ManaCost     int       `json:"mana_cost"`
	Name         string    `json:"name"`
	Effect       string    `json:"description"`
	Image        string    `json:"image"`
	Realm        CardRealm `json:"card_realm"`
	LimitPerDeck int       `json:"limit_per_deck"`
}
