package usertype

import (
	"net/http"
	"strconv"

	userTypeModel "github.com/divisi-developer-poros/poros-web-backend/models/usertype"
	r "github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
)

// UserTypeHandler ... User type handler struct declaration
type UserTypeHandler struct {
	Model userTypeModel.UserType
	Res   r.Response
}

// GetAll ... Get all user type
func (usrType *UserTypeHandler) GetAll(c *gin.Context) {
	var userTypes []userTypeModel.UserType

	if err := userTypeModel.GetAll(&userTypes); err != nil {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
	} else {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "success", "null", http.StatusOK, userTypes)
		return
	}

}

// Get ... Get single user type
func (usrType *UserTypeHandler) Get(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var userType userTypeModel.UserType
		if err := userTypeModel.Get(&userType, numId); err != nil {
			usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
		} else {
			usrType.Res.CustomResponse(c, "Content-Type", "application/json", "success", "null", http.StatusOK, userType)
			return
		}
	}
}

// Create ... create single user type
func (usrType *UserTypeHandler) Create(c *gin.Context) {
	var userType userTypeModel.UserType

	if errBind := c.ShouldBindJSON(&userType); errBind != nil {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", errBind.Error(), http.StatusBadRequest, nil)
	} else {
		if err := userTypeModel.Create(&userType); err != nil {
			usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
		} else {
			usrType.Res.CustomResponse(c, "Content-Type", "application/json", "sucess", "user type created", http.StatusOK, userType)
		}
	}
}

// Update ... update single user type
func (usrType *UserTypeHandler) Update(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var userType userTypeModel.UserType

		if errBind := c.ShouldBindJSON(&userType); errBind != nil {
			usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", errBind.Error(), http.StatusBadRequest, nil)
		} else {
			var existedUserType userTypeModel.UserType
			if errUserTypeExist := userTypeModel.Get(&existedUserType, numId); errUserTypeExist != nil {
				usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user not found", http.StatusNotFound, nil)
			} else {
				if err := userTypeModel.Update(&userType, numId); err != nil {
					usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
				} else {
					userType.ID = existedUserType.ID
					usrType.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user type updated", http.StatusOK, userType)
				}
			}
		}
	}
}

// Delete ... Delete single user type
func (usrType *UserTypeHandler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var userType userTypeModel.UserType
		if errUserTypeExist := userTypeModel.Get(&userType, numId); errUserTypeExist != nil {
			usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user not found", http.StatusNotFound, nil)
		} else {
			if err := userTypeModel.Delete(&userType, numId); err != nil {
				usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
			} else {
				usrType.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user type deleted", http.StatusOK, nil)
			}
		}
	}
}
