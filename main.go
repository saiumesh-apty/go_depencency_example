package main

import "fmt"

type RepoInterface interface {
	GetNamesDB() string
}

type Repo struct {
	repo string
}

func (repo *Repo) GetNamesDB() string {
	return repo.repo
}

type ServiceRepo interface {
	GetNamesService() string
}

type Service struct {
	repo *Repo
}

func (service *Service) GetNamesService() string {
	return service.repo.GetNamesDB()
}

func getService() *Service {
	repo := &Repo{repo: "Hey!!"}
	var _ RepoInterface = repo
	service := &Service{repo: repo}
	var _ ServiceRepo = service
	return service
}

func main() {
	service := getService()
	name := service.GetNamesService()
	fmt.Println(fmt.Sprintf("Hello %s ", name))
}
