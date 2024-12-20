package mantis

type Job struct {
	ID          string
	Name        string
	Description string
	Status      string
}

func NewJob(id, name, description, status string) *Job {
	return &Job{
		ID:          id,
		Name:        name,
		Description: description,
		Status:      status,
	}
}
