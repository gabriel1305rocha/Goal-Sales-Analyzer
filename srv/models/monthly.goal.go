package models

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/* 

How to use
goal := MonthlyGoal{
	MonthYear: CustomMonthYear{Time: time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)},
}
db.Create(&goal) 

*/

// MonthlyGoal defines the monthly goal with year and month
type MonthlyGoal struct {
	gorm.Model
	MonthYear CustomMonthYear `gorm:"type:date"` // Sets the data type as 'date' in PostgreSQL
	GoalOne   float64
	GoalTwo   float64
	GoalThree float64
	ExtraGoal float64
	UserID    uint // Fk to User
	User      User `gorm:"foreignKey:UserID"` // Relation with User
}

// CustomMonthYear implements a custom type for month/year
type CustomMonthYear struct {
	time.Time
}

// MarshalJSON formats the output as "YYYY-MM"
func (c CustomMonthYear) MarshalJSON() ([]byte, error) {
	formatted := c.Format("\"2006-01\"")
	return []byte(formatted), nil
}

// UnmarshalJSON allows input in the "YYYY-MM" format
func (c *CustomMonthYear) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse("\"2006-01\"", string(data))
	if err != nil {
		return err
	}
	// Sets to the first day of the month
	c.Time = time.Date(parsedTime.Year(), parsedTime.Month(), 1, 0, 0, 0, 0, time.UTC)
	return nil
}

// Scan reads the value from the database
func (c *CustomMonthYear) Scan(value interface{}) error {
	if value == nil {
		*c = CustomMonthYear{Time: time.Time{}}
		return nil
	}
	switch t := value.(type) {
	case time.Time:
		*c = CustomMonthYear{Time: time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)}
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into CustomMonthYear", value)
	}
}

// Value writes the value to the database
func (c CustomMonthYear) Value() (driver.Value, error) {
	return c.Time, nil
}
