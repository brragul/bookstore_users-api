package users

import (
	"github.com/brragul/bookstore_users-api/domain/users"
	services "github.com/brragul/bookstore_users-api/service"
	"github.com/brragul/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func CreateUser(c *gin.Context){
	var user users.User
/* -- This whole block can be replaced by ShouldBindJSON
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO
		return
	}
	if err := json.Unmarshal(bytes,&user); err != nil{
		fmt.Println(err.Error())
		return
		return
	}

 */

	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Code, restErr)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil{
		c.JSON(saveError.Code, saveError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil{
		err:= errors.NewBadRequestError("User ID should be a number")
		c.JSON(err.Code,err)
		return
	}

	result, getError := services.GetUser(userId)
	if getError != nil{
		c.JSON(getError.Code, getError)
		return
	}
	c.JSON(http.StatusCreated, result)
}