package seed

import "gorm.io/gorm"

// Seed is seed for database
type Seed struct {
	Name string
	Run func(*gorm.DB) error
}

// All is running all seeder
func All() []Seed{
	return []Seed{
		{
			Name: "CreateTodo",
			Run: func(db *gorm.DB) error {
				var err error
				for i := 0; i < 3; i++ {
					err = CreateTodo(db)
					if err != nil {
						break
					}
				}
				return err
			},
		},
	}
}