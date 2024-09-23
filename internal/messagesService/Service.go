package messagesService

type MessageService struct {
	repo MessageRepository
}

func NewMessageService(service MessageRepository) *MessageService {
	return &MessageService{repo: service}
}

func (r *MessageService) CreateMessage(message Message) (Message, error) {
	return r.repo.CreateMessage(message)
}

func (r *MessageService) FindAll() ([]Message, error) {
	return r.repo.FindAll()
}

func (r *MessageService) UpdateMessageByID(id int, message Message) (Message, error) {
	return r.repo.UpdateMessageByID(id, message)
}

func (r *MessageService) DeleteMessageByID(id int) error {
	return r.repo.DeleteMessageByID(id)
}
