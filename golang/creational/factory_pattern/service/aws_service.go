package service

type AwsService struct {
	Apikey string
}

func InitAwsService() *AwsService {
	return &AwsService{
		Apikey: "AWSKey",
	}
}
