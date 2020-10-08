package user_type
import (
	userTypeModel "github.com/divisi-developer-poros/poros-web-backend/models/user_type"
	r "github.com/divisi-developer-poros/poros-web-backend/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserTypeHandler struct {
	Model	userTypeModel.User_Type
	Res     r.Response
}

func (usrType *UserTypeHandler) GetAll(c *gin.Context) {
	var userTypes []userTypeModel.User_Type

	if err := userTypeModel.GetAll(&userTypes); err != nil {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
	} else {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "success", "null", http.StatusOK, userTypes)
		return
	}

}

func (usrType *UserTypeHandler) Get(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var userType userTypeModel.User_Type
		if err := userTypeModel.Get(&userType, numId); err != nil {
			usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user type not found", http.StatusNotFound, nil)
		} else {
			usrType.Res.CustomResponse(c, "Content-Type", "application/json", "success", "null", http.StatusOK, userType)
			return
		}
	}
}

func (usrType *UserTypeHandler) Create(c *gin.Context) {
	var userType userTypeModel.User_Type

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

func (usrType *UserTypeHandler) Update(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var userType userTypeModel.User_Type

		if errBind := c.ShouldBindJSON(&userType); errBind != nil {
			usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", errBind.Error(), http.StatusBadRequest, nil)
		} else {
			var existedUserType userTypeModel.User_Type
			if errUserTypeExist := userTypeModel.Get(&existedUserType, numId); errUserTypeExist != nil {
				usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "user type not found", http.StatusNotFound, nil)
			} else {
				if err := userTypeModel.Update(&userType, numId); err != nil {
					usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", err.Error(), http.StatusInternalServerError, nil)
				} else {
					usrType.Res.CustomResponse(c, "Content-Type", "application/json", "success", "user type updated", http.StatusOK, nil)
				}
			}
		}
	}
}

func (usrType *UserTypeHandler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if numId, error := strconv.Atoi(id); error != nil {
		usrType.Res.CustomResponse(c, "Content-Type", "application/json", "error", "ID not valid", http.StatusBadRequest, nil)
	} else {
		var userType userTypeModel.User_Type
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
