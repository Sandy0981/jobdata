package service

import (
	"context"
	"errors"

	"job-portal-api/internal/models"
)

func (s *Service) GetJobPostingByIDService(ctx context.Context, jid uint64) (models.Jobs, error) {
	if jid < uint64(0) {
		return models.Jobs{}, errors.New("id cannot be 0")
	}
	jobData, err := s.UserRepo.FetchJobPostingByID(ctx, jid)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) GetAllJobPostingsService(ctx context.Context) ([]models.Jobs, error) {
	jobDatas, err := s.UserRepo.FetchAllJobPostings(ctx)
	if err != nil {
		return nil, err
	}
	return jobDatas, nil

}

func (s *Service) CreateJobPostingService(ctx context.Context, jobData models.Jobs, cid uint64) (models.Jobs, error) {
	jobData.Cid = uint(cid)
	jobData, err := s.UserRepo.InsertJobPosting(ctx, jobData)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) ListJobsForCompanyService(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	jobData, err := s.UserRepo.FetchJobsForCompany(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
