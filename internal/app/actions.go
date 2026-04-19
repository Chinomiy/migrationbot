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
	StateFAQ                fsm.StateID = "faq"
	StateHelp                           = fsm.StateID("help")

	StateNoChange fsm.StateID = "no_change"
)

const (
	// НЕ МЕНЯЮТ СТЕЙТ
	CallbackAddFavorite = "add_favorite"
	CallbackFavorite    = "favorite"
	CallbackBack        = "back"
	// CallbackMainMenu user callback
	CallbackMainMenu           = "main_menu"
	CallbackCountryMenu        = "country_menu"
	CallbackCountryDetailsMenu = "country_details_menu"
	CallbackCountry            = "country"
	CallbackFAQ                = "faq"
	CallbackHelp               = "help"

	CallbackAccount = "account"

	CallbackRemoveBookmark  = "remove_bookmark"
	CallbackBookmarkDetails = "bookmark_details"
)
