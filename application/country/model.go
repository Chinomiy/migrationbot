package country

type Country struct {
	ID          string
	Code        string
	Name        string
	Description string
	// callback = name
	TripTypes TripType
	Content   string
}

type TripType struct {
	Data map[string]string
}
