package app

import "migtationbot/fsm"

const (
	StateMainMenu             fsm.StateID = "main_menu"
	StateCountryMenu          fsm.StateID = "country_menu"
	StateCountryDetailsMenu   fsm.StateID = "country_details_menu"
	StateTripMenu             fsm.StateID = "trip_menu"
	StateCountryTripMenu      fsm.StateID = "country_trip_menu"
	StateManagerCreateCountry fsm.StateID = "manager_create_country"
	StateCountry              fsm.StateID = "country"
	StateCountryTrip          fsm.StateID = "country_trip"
	StateAccount              fsm.StateID = "account"
	StateFavorite             fsm.StateID = "favorite"
	StateBookmarkDetails      fsm.StateID = "bookmark_details"
)

const (
	CallbackMainMenu           = "main_menu"
	CallbackCountryMenu        = "country_menu"
	CallbackCountryDetailsMenu = "country_details_menu"
	CallbackCountryTripMenu    = "country_trip_menu"
	CallbackCountry            = "country"
	CallbackBack               = "back"
	CallbackAccount            = "account"
	CallbackAddFavorite        = "add_favorite"
	CallbackFavorite           = "favorite"
	CallbackRemoveBookmark     = "remove_bookmark"
	CallbackBookmarkDetails    = "bookmark_details"
)
