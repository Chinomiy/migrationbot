package app

import "migtationbot/fsm"

const (
	// StateMainMenu стейты юзера главная менюшка закладки
	StateMainMenu           fsm.StateID = "main_menu"
	StateCountryMenu        fsm.StateID = "country_menu"         // Меню со списком СТРАН
	StateCountryDetailsMenu fsm.StateID = "country_details_menu" // Меню страны + выбран тип поездки
	StateCountry            fsm.StateID = "country"              // меню страны с выбором типа поездки
	StateAccount            fsm.StateID = "account"
	StateFavorite           fsm.StateID = "favorite"
	StateBookmarkDetails    fsm.StateID = "bookmark_details"

	StateNoChange fsm.StateID = "no_change"

	// StateManagerMenu менеджерские стетйы
	StateManagerMenu          fsm.StateID = "manager_menu"
	StateManagerCreateCountry fsm.StateID = "manager_create_country"
)

const (
	// CallbackMainMenu user callback
	CallbackMainMenu           = "main_menu"
	CallbackCountryMenu        = "country_menu"
	CallbackCountryDetailsMenu = "country_details_menu"
	CallbackCountry            = "country"
	CallbackBack               = "back"
	CallbackAccount            = "account"
	CallbackAddFavorite        = "add_favorite"
	CallbackFavorite           = "favorite"
	CallbackRemoveBookmark     = "remove_bookmark"
	CallbackBookmarkDetails    = "bookmark_details"

	// CallbackManagerMenu manager callback
	CallbackManagerMenu          = "manager_menu"
	CallbackManagerCreateCountry = "create_country"
)
