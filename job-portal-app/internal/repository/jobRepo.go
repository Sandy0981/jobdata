package repository

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"job-portal-api/internal/models"
)

func (r *Repo) FetchJobPostingByID(ctx context.Context, jid uint64) (models.Jobs, error) {
	var job models.Jobs
	result := r.DB.Where("id = ?", jid).Find(&job)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Jobs{}, errors.New("failed to create job postings")
	}
	return job, nil
}

func (r *Repo) InsertJobPosting(ctx context.Context, jobData models.Jobs) (models.Jobs, error) {
	result := r.DB.Create(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Jobs{}, errors.New("failed to insert job posting")
	}
	return jobData, nil
}

func (r *Repo) FetchAllJobPostings(ctx context.Context) ([]models.Jobs, error) {
	var jobDatas []models.Jobs
	result := r.DB.Find(&jobDatas)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("failed to fetch all job postings")
	}
	return jobDatas, nil

}

func (r *Repo) FetchJobsForCompany(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	var jobData []models.Jobs
	result := r.DB.Where("cid = ?", cid).Find(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("failed to fetch jobs for company")
	}
	return jobData, nil
}
