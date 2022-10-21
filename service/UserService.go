package service

import (
	"errors"
	"log"
	"mygram/dto"
	"mygram/entity"
	userrepository "mygram/repository/UserRepository"
)

type UserService interface {
	Register(req *dto.UserRegister) (dto.UserResponse, error)
	Login(req *dto.UserLogin) (dto.UserLoginResponse, error)
	Update(userID int64, req dto.UserUpdate) (dto.UserUpdateResponse, error)
	Delete(userID int64) error
}

type userService struct {
	userRepository userrepository.UserRepository
}

func NewUserService(userRepository userrepository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Register(req *dto.UserRegister) (dto.UserResponse, error) {

	var user dto.UserResponse

	// check username
	_, err := s.userRepository.FindByUsername(req.Username)
	if err == nil {

		log.Println("username telah digunakan")
		return user, errors.New("username telah digunakan")

	}

	// check email
	_, err = s.userRepository.FindByEmail(req.Email)

	// hash pasword
	hasPassword := req.Password

	// age minimal 8
	if req.Age < 9 {
		log.Println("age minmal diatas 8")
		return user, errors.New("age minmal diatas 8")
	}

	var entityUser entity.User
	entityUser.Password = hasPassword
	_, lastInserID, err := s.userRepository.Insert(entityUser)
	if err != nil {

		log.Println(err)
		return user, err

	}

	userEntity, err := s.userRepository.FindByID(lastInserID)
	if err != nil {

		log.Println(err)
		return user, err

	}

	user.ID = userEntity.ID
	user.Username = userEntity.Username
	user.Age = userEntity.Age
	user.Email = userEntity.Email

	return user, nil

}
func (s *userService) Login(req *dto.UserLogin) (dto.UserLoginResponse, error) {

	var userLogin dto.UserLoginResponse

	// check email
	userEntity, err := s.userRepository.FindByEmail(req.Email)
	if err != nil {

		log.Println(err)
		return userLogin, errors.New("user not found")

	}
	// comparea password
	userEntity.Password = userEntity.Password

	// generate token
	token := "123213123123213213123123"

	userLogin.Token = token

	return userLogin, nil

}
func (s *userService) Update(userID int64, req dto.UserUpdate) (dto.UserUpdateResponse, error) {

	var user dto.UserUpdateResponse

	// find by id
	userEntity, err := s.userRepository.FindByID(userID)
	if err != nil {

		log.Println(err)
		return user, err

	}

	// check username sudah ada atau belum dengan yg ada di db tapi id tidak sama dengan yg diupdate
	entityUserCheckUsername, _ := s.userRepository.FindByUsername(req.Username)
	if entityUserCheckUsername.ID != userEntity.ID {
		err = errors.New("username telah digunakan ")
		log.Println(err)
		return user, err
	}

	// check email sudah ada atau belum dengan yg ada di db tapi id tidak sama dengan yg diupdate
	entityUserCheckEmail, _ := s.userRepository.FindByEmail(req.Email)
	if entityUserCheckEmail.ID != userEntity.ID {
		err = errors.New("email telah digunakan ")
		log.Println(err)
		return user, err
	}

	var userEntityUpdate entity.User
	userEntityUpdate.Email = req.Email
	userEntityUpdate.Username = req.Username
	// update user
	s.userRepository.Update(userID, userEntityUpdate)

	// find by id
	userEntity, err = s.userRepository.FindByID(userID)
	if err != nil {

		log.Println(err)
		return user, err

	}

	user.ID = userEntity.ID
	user.Username = userEntity.Username
	user.Age = userEntity.Age
	user.Email = userEntity.Email
	user.UpdatedAt = userEntity.UpdatedAt

	return user, nil

}
func (s *userService) Delete(userID int64) error {

	// find by id
	userEntity, err := s.userRepository.FindByID(userID)
	if err != nil {

		log.Println(err)
		return err

	}

	err = s.userRepository.Delete(userEntity.ID)
	if err != nil {

		log.Println(err)
		return err

	}

	return nil

}
