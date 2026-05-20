package request

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type UpdateProfileRequest struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
}

type CreateWorkoutRequest struct {
	Name         string `json:"name"`
	Date         string `json:"date"`
	Duration     int    `json:"duration"`
	Calories     int    `json:"calories"`
	AvgHeartRate int    `json:"avgHeartRate"`
	WeightUnit   string `json:"weightUnit"`
	Notes        string `json:"notes"`
}

type UpdateWorkoutRequest struct {
	Name         string `json:"name"`
	Date         string `json:"date"`
	Duration     int    `json:"duration"`
	Completed    bool   `json:"completed"`
	Calories     int    `json:"calories"`
	AvgHeartRate int    `json:"avgHeartRate"`
	Notes        string `json:"notes"`
}

type AddBodyWeightRequest struct {
	Weight              float64 `json:"weight"`
	BodyFatPercentage   float64 `json:"bodyFatPercentage"`
}

type CreateChatRequest struct {
	Message string `json:"message"`
}

type CreateScheduleRequest struct {
	Date     string `json:"date"`
	Time     string `json:"time"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
	Notes    string `json:"notes"`
}

type UpdateScheduleRequest struct {
	Date     string `json:"date"`
	Time     string `json:"time"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
	Notes    string `json:"notes"`
}

type CreateGoalRequest struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	Target   int    `json:"target"`
	Unit     string `json:"unit"`
	Period   string `json:"period"`
	Deadline string `json:"deadline"`
}

type UpdateGoalRequest struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	Target   int    `json:"target"`
	Current  int    `json:"current"`
	Unit     string `json:"unit"`
	Period   string `json:"period"`
	Deadline string `json:"deadline"`
}

type UpdateGoalProgressRequest struct {
	Current int `json:"current"`
}
