package message

import "github.com/aithlete/aithlete-api/pkg/code"

const (
	MsgSuccess = "Success"
	MsgOK      = "OK"
	MsgCreated = "Created"
	MsgDeleted = "Deleted"

	MsgLoginSuccess    = "Login successful"
	MsgLoginFailed     = "Login failed"
	MsgLogoutSuccess   = "Logged out successfully"
	MsgRegisterSuccess = "Account created successfully"
	MsgRegisterFailed  = "Registration failed"
	MsgInvalidCredentials   = "Invalid email or password"
	MsgRefreshTokenExpired  = "Invalid or expired refresh token"
	MsgRegisterEmailExists  = "Email already registered"

	MsgUserNotFound  = "User not found"
	MsgUnauthorized  = "Unauthorized"
	MsgGetUserFailed = "Failed to get user"
	MsgProfileUpdated = "Profile updated successfully"

	MsgWorkoutCreated = "Workout created"
	MsgWorkoutUpdated = "Workout updated"
	MsgWorkoutDeleted = "Workout deleted"

	MsgBodyWeightAdded = "Body weight entry added"

	MsgBadRequest        = "Invalid request body"
	MsgInternalError     = "Internal server error"
	MsgInvalidToken      = "Invalid or expired token"
	MsgInvalidAuthHeader = "Invalid authorization format"
	MsgMissingAuthHeader = "Missing authorization header"
)

func StatusDesc(statusCode string) string {
	switch statusCode {
	case code.Success:
		return "Success"
	case code.BadRequest:
		return "Bad request"
	case code.Unauthorized:
		return "Unauthorized"
	case code.NotFound:
		return "Not found"
	case code.Conflict:
		return "Conflict"
	case code.InternalServerError:
		return "Internal server error"
	default:
		return "Unknown"
	}
}
