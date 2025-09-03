package dto

type BadRequestResponse struct {
    Error string `json:"error" example:"invalid input"`
}

type InternalServerErrorResponse struct {
	Error string `json:"error" example:"internal server error"`
}

type UnauthorizedResponse struct {
	Error string `json:"error" example:"authentication required"`
}

type ForbiddenResponse struct {
	Error string `json:"error" example:"forbidden"`
}

type NotFoundResponse struct {
	Error string `json:"error" example:"Plan does not exist"`
}