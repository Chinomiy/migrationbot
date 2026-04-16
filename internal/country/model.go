package country

type Country struct {
	ID          int
	Code        string
	Name        string
	Description string

	TripTypes TripType
	Content   string
}

type TripType struct {
	Id int
	// [callback]name
	Data map[string]string
}
