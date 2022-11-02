package handler

import (
	"latihan3/usecase"

	"github.com/gofiber/fiber/v2"

	"latihan3/entity/response"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(usercase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: usercase,
	}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := response.CreateUserRequest{}

	users := h.userUsecase.CreateUser(userRequest)

	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Massage: "invalid req body",
			Data:    nil,
		})
	}

	if err := h.userUsecase.CreateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Massage: "Failed to create user",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Massage: "user created successfully",
		Data:    users,
	})
}

func (u UserHandler) GetList(ctx *fiber.Ctx) error {

	users, err := u.userUsecase.GetList()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Massage: "error",
			Data:    nil,
		})
	}

	if len(users) < 0 {
		return ctx.Status(fiber.StatusNoContent).JSON(response.BaseResponse{
			Code:    fiber.StatusNoContent,
			Massage: "no data found",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Massage: "success",
		Data:    nil,
	})
}

// fungsi UpdateUser
func (u UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	userRequest := response.UpdateUserRequest{}

	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Massage: "invalid req body",
			Data:    nil,
		})
	}

	if err := u.userUsecase.UpdateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Massage: "Failed to update user",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Massage: "user updated successfully",
		Data:    nil,
	})
}

// fungsi DeleteUser
func (u UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("ID")

	if err := u.userUsecase.DeleteUser(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Massage: "Failed to delete user",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Massage: "user deleted successfully",
		Data:    nil,
	})
}
