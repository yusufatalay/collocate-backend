package models

import (
	"time"

	firestore "cloud.google.com/go/firestore/apiv1beta1"
)

type ClientModel struct {
	CL *firestore.Client
}

// Models is the wrapper for the firestore client
type Models struct {
	CL ClientModel
}

// NewModels returns models with client pool
func NewModels(cl *firestore.Client) Models {
	return Models{
		CL: ClientModel{
			CL: cl,
		},
	}
}

type BoardItem struct {
	Status int8   `json:"Status"`
	Text   string `json:"Text"`
}

type BoardSection struct {
	SectionItems []BoardItem `json:"SectionItems"`
	SectionType  int         `json:"SectionType"`
}

type Day struct {
	Date       time.Time `json:"Date"`
	TaskDone   []Task    `json:"TaskDone"`
	WorkAmount int       `json:"WorkAmount"`
}

type Task struct {
	IsDone   bool   `json:"IsDone"`
	TaskText string `json:"TaskText"`
}

type TaskCard struct {
	Header     string `json:"Header"`
	IsComplete bool   `json:"IsComplete"`
	Tasks      []Task `json:"Tasks"`
}

type Preferences struct {
	Language         int8   `json:"Language"`
	PomodoroRestTime uint16 `json:"PomodoroRestTime"`
	PomodoroWorkTime uint16 `json:"PomodoroWorkTime"`
	Theme            uint8  `json:"Theme"`
}

type User struct {
	Email           string `json:"E-mail"`
	Name            string `json:"Name"`
	ProfilePhotoUrl string `json:"ProfilePhotoUrl"`
	UserPreference  string `json:"UserPreference"` // reference to Preferences document like /Preferences/<preference_document_id>
}

func getArrayFromInterface(i interface{}) []Task {
	var result []Task
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		result = append(result, v.(Task))
	}
	return result
}
