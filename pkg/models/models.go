package models

import "errors"

var (
	ErrNoRecord = errors.New("models: no matching record found")
)

type BathroomStyle struct {
	ID              int
	Title           string
	LongDescription string
}

type RoomType struct {
	ID              int
	ShortCode       string
	LongDescription string
	NumberOfBeds    int
}

type RoomNote struct {
	ID      int
	Content string
}

type Feature struct {
	ID    int
	Title string
}

type Street struct {
	ID    int
	Title string
}

type Room struct {
	RoomNumber          int
	Floor               int
	NumberOfShowerHeads int
	NumberOfSinks       int
	SquareFootage       int
	BathroomStyle       *BathroomStyle
	RoomType            *RoomType
	RoomNotes           []*RoomNote
	Features            []*Feature
	FacingStreets       []*Street
	ConnectingRooms     []int
}
