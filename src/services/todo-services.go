package services

import "sinaugo/src/repository"

// TodoService -> the propose of user repository is handling query for user entity
type TodoService struct {
	TodoRepository repository.TodoRepository
}

// TService -> todo services instance
func TService() TodoService {
	return TodoService{
		TodoRepository: repository.URepository(),
	}
}

// GetTodos -> get todos service logic
func (s *TodoService) GetTodos() []repository.Todos {
	todos := s.TodoRepository.GetUsers()
	return todos
}