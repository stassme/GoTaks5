package messagesService

import "gorm.io/gorm"

type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	FindAll() ([]Message, error)
	UpdateMessageByID(id int, message Message) (Message, error)
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	DB *gorm.DB
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.DB.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) FindAll() ([]Message, error) {
	var messages []Message
	err := r.DB.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	return message, r.DB.Model(&Message{}).Where("id = ?", id).Updates(message).Error
}

func (r *messageRepository) DeleteMessageByID(id int) error {
	return r.DB.Delete(&Message{}, id).Error
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{DB: db}
}
