package models

import "github.com/Evelon44/MyTests/GormAPI/storage"

func AddNewToDo(t *ToDo) error {
	return storage.DB.Create(t).Error
}

func GetAllToDo(t *[]ToDo) error {
	return storage.DB.Find(t).Error
}

func GetToDoById(t *ToDo, id int) error {
	return storage.DB.Where("id = ?", id).First(t).Error
}

func DeleteToDoById(t *ToDo, id int) error {
	return storage.DB.Where("id = ?", id).Delete(t).Error
}

func UpdateToDoById(t *ToDo, id int) error {
	return storage.DB.Where("id = ?", id).Updates(t).Error
}
