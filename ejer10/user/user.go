package user

import "time"

type User struct {
	Id     int
	Name   string
	Date   time.Time
	Status bool
}

func (this *User) AltaUser(id int, name string, date time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.Date = date
	this.Status = status
}
