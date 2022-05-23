package service

import (
	"log"

	"github.com/Omar-Temirgali/final-exam-go-lang/internal/domain/dto"
	"github.com/Omar-Temirgali/final-exam-go-lang/internal/domain/models"
	"github.com/Omar-Temirgali/final-exam-go-lang/internal/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) models.User
	Profile(userID string) models.User
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}
