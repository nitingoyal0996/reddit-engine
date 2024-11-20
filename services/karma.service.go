package services

import "github.com/nitingoyal0996/reddit-clone/repositories"

type KarmaService struct {
	userRepo 	repositories.UserRepository
}

func NewKarmaService(userRepo repositories.UserRepository) *KarmaService {
	return &KarmaService{
		userRepo: userRepo,
	}
}

func (s *KarmaService) UpdateKarma(userId uint, amount int) error {
	return s.userRepo.UpdateKarma(userId, amount)
}