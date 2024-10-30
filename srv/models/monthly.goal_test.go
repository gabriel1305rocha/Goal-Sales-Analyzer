package models

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
)

// TestMonthlyGoal tests the MonthlyGoal model
func TestMonthlyGoal(t *testing.T) {
	// Set up the in-memory

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Migrate the schema for MonthlyGoal and User models
	err = db.AutoMigrate(&User{}, &MonthlyGoal{})
	assert.NoError(t, err)

	// Create a sample user
	user := User{Name: "John Doe", Email: "johndoe@example.com", Password: "password", Age: 30}
	err = db.Create(&user).Error
	assert.NoError(t, err)

	// Define a MonthlyGoal with CustomMonthYear set to March 2024
	goal := MonthlyGoal{
		MonthYear: CustomMonthYear{Time: time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)},
		GoalOne:   1000.0,
		GoalTwo:   2000.0,
		GoalThree: 3000.0,
		ExtraGoal: 400.0,
		UserID:    user.ID,
	}

	// Insert the MonthlyGoal into the database
	err = db.Create(&goal).Error
	assert.NoError(t, err)

	// Retrieve the goal to verify correct storage
	var retrievedGoal MonthlyGoal
	err = db.Preload("User").First(&retrievedGoal, goal.ID).Error
	assert.NoError(t, err)

	// Validate the fields
	assert.Equal(t, "2024-03", retrievedGoal.MonthYear.Format("2006-01"), "MonthYear should be stored as YYYY-MM")
	assert.Equal(t, goal.GoalOne, retrievedGoal.GoalOne, "GoalOne should match")
	assert.Equal(t, goal.GoalTwo, retrievedGoal.GoalTwo, "GoalTwo should match")
	assert.Equal(t, goal.GoalThree, retrievedGoal.GoalThree, "GoalThree should match")
	assert.Equal(t, goal.ExtraGoal, retrievedGoal.ExtraGoal, "ExtraGoal should match")
	assert.Equal(t, user.ID, retrievedGoal.UserID, "UserID should match")
	assert.Equal(t, user.Name, retrievedGoal.User.Name, "User's Name should match")
}
