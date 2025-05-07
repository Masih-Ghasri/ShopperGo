package services

import (
	"encoding/json"
	"fmt"
	"github.com/Masih-Ghasri/GolangBackend/config"
	constant "github.com/Masih-Ghasri/GolangBackend/constants"
	"github.com/Masih-Ghasri/GolangBackend/data/cache"
	"github.com/Masih-Ghasri/GolangBackend/pkg/logging"
	"github.com/Masih-Ghasri/GolangBackend/pkg/logging/service_errors"
	"github.com/go-redis/redis/v7"
	"time"
)

type OtpService struct {
	logger logging.Logger
	cfg    *config.Config
	redis  *redis.Client
}

type OtpDto struct {
	Value string `json:"value"`
	Used  bool   `json:"used"`
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{logger: logger, cfg: cfg, redis: redis}
}

func (s *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("otp:%s:%s", constant.RedisOtpDefaultKey, mobileNumber)
	val := &OtpDto{Value: otp, Used: false}

	storedData, err := s.redis.Get(key).Result()
	if err == nil {
		var storedOtp OtpDto
		err = json.Unmarshal([]byte(storedData), &storedOtp)
		if err == nil {
			if !storedOtp.Used {
				return &service_errors.ServiceError{EndUserMessage: service_errors.OptExists}
			} else {
				return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
			}
		}
	}

	jsonData, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = s.redis.Set(key, jsonData, s.cfg.Otp.ExpireTime*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("otp:%s:%s", constant.RedisOtpDefaultKey, mobileNumber)

	storedData, err := s.redis.Get(key).Result()
	if err != nil {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpNotValid}
	}

	var storedOtp OtpDto
	err = json.Unmarshal([]byte(storedData), &storedOtp)
	if err != nil {
		return err
	}

	if storedOtp.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}

	if storedOtp.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpNotValid}
	}

	storedOtp.Used = true

	jsonData, err := json.Marshal(storedOtp)
	if err != nil {
		return err
	}

	err = s.redis.Set(key, jsonData, s.cfg.Otp.ExpireTime*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}
