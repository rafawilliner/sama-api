package create

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafawilliner/sama-api/src/api/core/contracts/create"
	"github.com/rafawilliner/sama-api/src/api/core/entities"
	"github.com/rafawilliner/sama-api/src/api/core/errors"
	"github.com/rafawilliner/sama-api/src/api/core/usecases/mocks"
	"github.com/rafawilliner/sama-api/src/api/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockedCreatePetUseCase = new(mocks.MockCreatePetUseCase)
	mockedCreatePetUseCaseInternalError = new(mocks.MockCreatePetUseCase)
	mockCtx         = func(ctx context.Context) bool {
		return true
	}
	router = prepareRouter()
)

func prepareRouter() *gin.Engine {

	handler := Pet {
		CreateUseCase: mockedCreatePetUseCase,
	}

	router := utils.GetTestRouter()
	creditsGroup := router.Group("/sama")
	creditsGroup.POST("/pet", handler.Handle)
	return router
}

func TestCreatingPetWhenRequestIsValidShouldReturnNewPet(t *testing.T) {

	
	
	// Given
	rightBody := getMockedRequest(3)

	jsonBody, _ := json.Marshal(rightBody)

	var mockedPetRequest = getMockedRequest(3)
	var mockedPet = getEntity()

	request := httptest.NewRequest(http.MethodPost, "/sama/pet", bytes.NewBuffer(jsonBody))
	
	mockedCreatePetUseCase.On("Execute", mock.MatchedBy(mockCtx), mockedPetRequest).
		Return(mockedPet, nil)

	// When
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	bodyResponse := w.Body.String()

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	mockedCreatePetUseCase.AssertNumberOfCalls(t, "Execute", 1)
	assert.Equal(t, getResponseJSON(), bodyResponse)
}

func TestCreatingPetWhenRequestIsInvalidShouldReturnBadRequest(t *testing.T) {
	
	// Given		
	tests := []struct {
		name string
		body map[string]interface{}
	}{
		{
			name: "Request without name",
			body: map[string]interface{}{
				"age":           3,
				"specie":           "gato",
			},
		},
		{
			name: "Request age lower than 0",
			body: map[string]interface{}{
				"name": "Manchi",
				"age":           -2,
				"specie":           "gato",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)
			request := httptest.NewRequest(http.MethodPost, "/sama/pet", bytes.NewBuffer(jsonBody))
			
			// When
			w := httptest.NewRecorder()
			router.ServeHTTP(w, request)
		
			// Then
			assert.Equal(t, http.StatusBadRequest, w.Code)
			mockedCreatePetUseCase.AssertNotCalled(t, "Execute")
		})
	}
}

func TestCreatingPetWithErrorShouldReturnInternalServerError(t *testing.T) {
	
	// Given	

	handler2 := Pet {
		CreateUseCase: mockedCreatePetUseCaseInternalError,
	}
	gin.SetMode(gin.TestMode)
	gin.Default()
	router:= utils.GetTestRouter()
	creditsGroup := router.Group("/sama")
	creditsGroup.POST("/pet", handler2.Handle)


	rightBody := getMockedRequest(3)
	jsonBody, _ := json.Marshal(rightBody)

	var mockedPetRequest = getMockedRequest(3)
	request := httptest.NewRequest(http.MethodPost, "/sama/pet", bytes.NewBuffer(jsonBody))
	
	mockedCreatePetUseCaseInternalError.On("Execute", mock.MatchedBy(mockCtx), mockedPetRequest).
		Return(nil, errors.NewRepositoryError("Error in database"))

	// When
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	// Then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockedCreatePetUseCaseInternalError.AssertNumberOfCalls(t, "Execute", 1)
}

func getMockedRequest(age int32) *create.PetRequest {
	return &create.PetRequest{
		Name: "Firu",
		Age: &age,
		
	}
}

func getEntity() *entities.Pet {
	return &entities.Pet{
		Name: "Firu",
	}
}

func getResponseJSON() string {
	petJSON := `{"id":0,"description":"Firu","family":null,"race":null,"age":null,"specie":"","weight":null}`
	return 	petJSON
}
