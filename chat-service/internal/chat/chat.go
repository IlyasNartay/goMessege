package chat

import (
	"log"
)

// Message represents a chat message
type Message struct {
	ID        string
	Sender    string
	Recipient string
	Content   string
}

// ChatService provides chat functionality
type ChatService interface {
	SendMessage(message *Message) error
	GetMessagesByRecipient(recipient string) ([]*Message, error)
}

// ChatRepository provides methods to interact with the database for chat messages
type ChatRepository interface {
	SaveMessage(message *Message) error
	GetMessagesByRecipient(recipient string) ([]*Message, error)
}

// ChatServiceImpl implements ChatService interface
type ChatServiceImpl struct {
	repo   ChatRepository
	logger *log.Logger
}

// NewChatService creates a new instance of ChatServiceImpl
func NewChatService(repo ChatRepository, logger *log.Logger) *ChatServiceImpl {
	return &ChatServiceImpl{repo: repo, logger: logger}
}

// SendMessage sends a message
func (s *ChatServiceImpl) SendMessage(message *Message) error {
	// Save message to the repository
	err := s.repo.SaveMessage(message)
	if err != nil {
		s.logger.Printf("Failed to save message: %v", err)
		return err
	}
	return nil
}

// GetMessagesByRecipient retrieves messages for a recipient
func (s *ChatServiceImpl) GetMessagesByRecipient(recipient string) ([]*Message, error) {
	// Retrieve messages from the repository
	messages, err := s.repo.GetMessagesByRecipient(recipient)
	if err != nil {
		s.logger.Printf("Failed to get messages for recipient %s: %v", recipient, err)
		return nil, err
	}
	return messages, nil
}
