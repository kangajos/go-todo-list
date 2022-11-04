package services

import (
	"errors"
	"time"

	"github.com/kangajos/go-todo-list.git/models"
	"github.com/kangajos/go-todo-list.git/request"
)

type ActivityServiceImpl interface {
	FindAll() ([]models.Activity, error)
	FindByID(ID int) (*models.Activity, error)
	Create(request.ActivityRequest) (*models.Activity, error)
	Update(request.ActivityRequest) (*models.Activity, error)
	Delete(ID int) error
}

type ActivityService struct {
	activityRepository *models.ActivityRepository
}

// init
func (as *ActivityService) New(tr *models.ActivityRepository) *ActivityService {
	as.activityRepository = tr
	return as
}

// get all activity
func (as ActivityService) FindAll() ([]models.Activity, error) {
	return as.activityRepository.FindAll()
}

// get all activity
func (as ActivityService) FindByID(ID int) (*models.Activity, error) {
	data, _ := as.activityRepository.FindByID(ID)
	if data.ID == 0 {
		return nil, errors.New("Activitiy ID #{ID} Not Found")
	}
	return data, nil
}

// create activity
func (as ActivityService) Create(activityRequest request.ActivityRequest) (*models.Activity, error) {
	activity := models.Activity{}
	activity.Title = activityRequest.Title
	activity.Email = activityRequest.Email
	return as.activityRepository.Create(activity)
}

// update activity
func (as ActivityService) Update(activityRequest request.ActivityRequest) (*models.Activity, error) {
	activity, err := as.activityRepository.FindByID(activityRequest.ID)
	if err != nil {
		return nil, err
	}
	activity.Title = activityRequest.Title
	activity.Email = activityRequest.Email
	activity.UpdatedAt = time.Now()
	return as.activityRepository.Update(activity)
}

// delete activity
func (as ActivityService) Delete(ID int) error {
	data, _ := as.activityRepository.FindByID(ID)
	if data.ID == 0 {
		return errors.New("Activitiy ID #{ID} Not Found")
	}

	as.activityRepository.Delete(ID)
	return nil
}
