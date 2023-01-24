package create

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafawilliner/sama-api/src/api/core/contracts/create"
	createUseCase "github.com/rafawilliner/sama-api/src/api/core/usecases/create"
)

type Pet struct {
	CreateUseCase createUseCase.PetUseCase
}

func (handler *Pet) Handle(c *gin.Context) {
	handler.handle(c)
}

func (handler *Pet) handle(c *gin.Context) error {
	ctx := c.Request.Context()

	var err error

	var request create.PetRequest
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	if err = request.ValidateRequest(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	pet, err := handler.CreateUseCase.Execute(ctx, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	response := create.NewResponse(pet)

	c.JSON(http.StatusOK, response)
	return nil
}
