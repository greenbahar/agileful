package domain

// Repository represent database operations on User information
type UserInfoInterface interface {
	Close()
	FindByID(id string) (*UserInfoModel, error)
	Find() ([]*UserInfoModel, error)
	Create(user *UserInfoModel) error
	Update(user *UserInfoModel) error
	Delete(id string) error
}

// UserInfoModel represents the user information model
type UserInfoModel struct {
	ID       int		`db:"id"`
	Name     string		`db:"namee"`
	Email    string		`db:"email"`
	PassWord string		`db:"password"`
}

// UserInfoModel represents the user information model
type ExecutionTimeModel struct {
	ID      	 int		`db:"id"`
	Query   	 string		`db:"query"`
	TimeSpent    string		`db:"time_spent"`
}