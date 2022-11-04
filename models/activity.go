package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        int             `json:"id"`
	Title     string          `json:"title"`
	Email     string          `json:"email"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

type ActivityImpl interface {
	FindAll() ([]Activity, error)
	FindByID(ID int) (*Activity, error)
	Create(Activity) (*Activity, error)
	Update(*Activity) (*Activity, error)
	Delete(ID int) error
}

type ActivityRepository struct {
	db gorm.DB
}

func (tr *ActivityRepository) New(db gorm.DB) *ActivityRepository {
	tr.db = db
	return tr
}

// get all Activitys
func (tr ActivityRepository) FindAll() ([]Activity, error) {

	activities := []Activity{}
	if err := tr.db.Model(&Activity{}).Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

// get Activity by ID
func (tr ActivityRepository) FindByID(ID int) (*Activity, error) {

	activity := Activity{}
	if err := tr.db.Model(&Activity{}).Where("id = ?", ID).Find(&activity).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

// create Activity
func (tr ActivityRepository) Create(activity Activity) (*Activity, error) {

	if err := tr.db.Create(&activity).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

// update Activity by ID
func (tr ActivityRepository) Update(activity *Activity) (*Activity, error) {

	if err := tr.db.Save(&activity).Error; err != nil {
		return nil, err
	}
	return activity, nil
}

// delete Activity by ID
func (tr ActivityRepository) Delete(ID int) error {

	if err := tr.db.Where("id = ?", ID).Delete(&Activity{}).Error; err != nil {
		return err
	}
	return nil
}
