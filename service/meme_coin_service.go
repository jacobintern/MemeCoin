package service

import (
	"github.com/jacobintern/meme_coin/controller/command"
	"github.com/jacobintern/meme_coin/controller/response"
	"github.com/jacobintern/meme_coin/pkg/errors"
	"github.com/jacobintern/meme_coin/pkg/util"
	"github.com/jacobintern/meme_coin/repository"
	"github.com/jacobintern/meme_coin/service/adapter_model.go"
)

type IMemeCoinService interface {
	Create(command command.CreateMemeCoinCommand) error
	Get(id string) (response.MemeCoinInfoResp, error)
	Patch(command command.PatchMemeCoinCommand) error
	Delete(id string) error
	Poke(id string) error
}

type memeCoinService struct {
	memeCoinRepo repository.IMemeCoinRepository
}

func NewMemeCoinService(memeCoinRepo repository.IMemeCoinRepository) IMemeCoinService {
	return &memeCoinService{
		memeCoinRepo: memeCoinRepo,
	}
}

func (s *memeCoinService) Create(command command.CreateMemeCoinCommand) error {
	// transfer command
	entity := adapter_model.NewMemeCoinInfo()
	entity.Name = command.Name
	if command.Description != nil {
		entity.Description = command.Description
	}
	return s.memeCoinRepo.Create(entity)
}

func (s *memeCoinService) Get(id string) (response.MemeCoinInfoResp, error) {
	// transfer id
	idInt, err := util.IDParser(id)
	if err != nil {
		return response.NewMemeCoinInfoResp(), errors.IDParseFailed
	}

	entity, err := s.memeCoinRepo.Get(idInt)
	if err != nil {
		return response.NewMemeCoinInfoResp(), err
	}

	return entity.ToMemeCoinInfoResp(), nil
}

func (s *memeCoinService) Patch(command command.PatchMemeCoinCommand) error {
	idInt, err := util.IDParser(command.ID)
	if err != nil {
		return errors.IDParseFailed
	}

	// get meme coin
	entity, err := s.memeCoinRepo.Get(idInt)
	if err != nil {
		return err
	}

	entity.Description = command.Description

	return s.memeCoinRepo.Update(entity)
}

func (s *memeCoinService) Delete(id string) error {
	// transfer id
	idInt, err := util.IDParser(id)
	if err != nil {
		return errors.IDParseFailed
	}

	// check exist
	if _, err = s.memeCoinRepo.Get(idInt); err != nil {
		return err
	}

	return s.memeCoinRepo.Delete(idInt)
}

func (s *memeCoinService) Poke(id string) error {
	// transfer id
	idInt, err := util.IDParser(id)
	if err != nil {
		return errors.IDParseFailed
	}

	// check exist
	if _, err = s.memeCoinRepo.Get(idInt); err != nil {
		return err
	}

	return s.memeCoinRepo.Poke(idInt)
}
