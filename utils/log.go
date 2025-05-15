package utils

import "go.uber.org/zap"

func InitializeLogger() (*zap.Logger, error) {
	Logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return Logger, nil
}
