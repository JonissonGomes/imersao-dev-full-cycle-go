package usecase

import (
	"github.com/google/uuid"
	"imersao-dev-full-cycle-go/entity"
)

type CreateCourse struct {
	Repository entity.CourseRepository
}

func (c CreateCourse) Execute(input CreateCourseInputDto) (CreateCourseOutputDto, error) {

	// Instanciando a entidade
	course := entity.Course{}

	// Criando os dados vindo via Input
	course.ID = uuid.New().String()
	course.Name = input.Name
	course.Description = input.Description
	course.Status = input.Status

	// Criando variável de erro passando o resultado do repository
	err := c.Repository.Insert(course)

	// Validando se houve erro
	if err != nil {
		return CreateCourseOutputDto{}, err
	}

	// Instanciando o Dto
	output := CreateCourseOutputDto{}

	// Retornando o Output com os dados validados
	output.ID = course.ID
	output.Name = course.Name
	output.Description = course.Description
	output.Status = course.Status

	return output, nil
}
