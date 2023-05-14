package services

type ActivityDependencies struct {
	AR ActivityResource
}

type Activity struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GetActivityByID struct {
	ID int `json:"id"`
}

type CreateActivity struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type UpdateActivity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type DeleteActivity struct {
	ID int `json:"id"`
}
