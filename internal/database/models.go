package database

import "database/sql"

type Models struct {
	Users UserModel
	Events EventModel
	Attendees AttendeeModel
}


func NewModels(db *sql.DB) Models {
    return Models{
		Users: UserModel{DB: db},
		Events: EventModel{DB: db},
		Attendees: AttendeeModel{DB: db},
	}  
}



//basically we are creating a struct called Models that will hold all of our database models. Each model will have a reference to the database connection, which we can use to perform queries and operations on the database. The NewModels function is a constructor that takes a database connection as an argument and returns an instance of the Models struct with all the models initialized.
