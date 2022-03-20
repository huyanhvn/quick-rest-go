package model

// Item is a single item in the database.
// Its primary key is auto-populated by gORM.
type Item struct {
	ID    int     `json:"id" gorm:"primary_key"`
	Attr1 string  `json:"attr1"`
	Attr2 float64 `json:"attr2"`
}

// CreateItem is request body for POST /items.
type CreateItem struct {
	Attr1 string  `json:"attr1" binding:"required"`
	Attr2 float64 `json:"attr2" binding:"required"`
}

// UpdateItem is request body for PATCH /items/:id.
type UpdateItem struct {
	Attr1 string  `json:"attr1"`
	Attr2 float64 `json:"attr2"`
}

// Update function example.
func Update(id int, i UpdateItem) error {
	var item Item
	if result := DB.First(&item, id); result.Error != nil {
		return result.Error
	}
	if i.Attr1 != "" {
		item.Attr1 = i.Attr1
	}
	if i.Attr2 != 0 {
		item.Attr2 = i.Attr2
	}
	if result := DB.Save(&item); result.Error != nil {
		return result.Error
	}
	return nil
}

// Create function example.
func Create(i CreateItem) error {
	if result := DB.Create(&Item{Attr1: i.Attr1, Attr2: i.Attr2}); result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete function example.
func Delete(id int) error {
	var item Item
	if result := DB.First(&item, id); result.Error != nil {
		return result.Error
	}
	if result := DB.Delete(&item); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllItems function example.
func GetAllItems() ([]Item, error) {
	var items []Item
	if result := DB.Find(&items); result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// GetItemByID function example.
func GetItemByID(id int) (Item, error) {
	var item Item
	if result := DB.First(&item, id); result.Error != nil {
		return Item{}, result.Error
	}
	return item, nil
}
