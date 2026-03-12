package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is the standard JSON envelope returned by all API endpoints.
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// PaginatedResponse wraps Response with pagination metadata.
type PaginatedResponse struct {
	Response
	Meta PaginationMeta `json:"meta"`
}

// PaginationMeta holds pagination details for list responses.
type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// SuccessResponse writes a 200 JSON response with data.
func SuccessResponse(ctx *gin.Context, message string, data any) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// CreatedResponse writes a 201 JSON response with data.
func CreatedResponse(ctx *gin.Context, message string, data any) {
	ctx.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse writes a JSON error response with the given status code.
func ErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	response := Response{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	ctx.JSON(statusCode, response)
}

// BadRequestResponse writes a 400 JSON error response.
func BadRequestResponse(ctx *gin.Context, message string, err error) {
	ErrorResponse(ctx, http.StatusBadRequest, message, err)
}

// UnauthorizedResponse writes a 401 JSON error response.
func UnauthorizedResponse(ctx *gin.Context, message string) {
	ErrorResponse(ctx, http.StatusUnauthorized, message, nil)
}

// ForbiddenResponse writes a 403 JSON error response.
func ForbiddenResponse(ctx *gin.Context, message string) {
	ErrorResponse(ctx, http.StatusForbidden, message, nil)
}

// NotFoundResponse writes a 404 JSON error response.
func NotFoundResponse(ctx *gin.Context, message string) {
	ErrorResponse(ctx, http.StatusNotFound, message, nil)
}

// InternalServerErrorResponse writes a 500 JSON error response.
func InternalServerErrorResponse(ctx *gin.Context, message string, err error) {
	ErrorResponse(ctx, http.StatusInternalServerError, message, err)
}

// PaginatedSuccessResponse writes a 200 JSON response with data and pagination metadata.
func PaginatedSuccessResponse(ctx *gin.Context, message string, data any, meta PaginationMeta) {
	ctx.JSON(http.StatusOK, PaginatedResponse{
		Response: Response{
			Success: true,
			Message: message,
			Data:    data,
		},
		Meta: meta,
	})
}
