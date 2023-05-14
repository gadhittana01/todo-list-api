package domain

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

type Todo struct {
	ID              int    `json:"id"`
	ActivityGroupID int    `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
	CreateAt        string `json:"createAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type GetTodoByID struct {
	ID int `json:"id"`
}

type CreateTodo struct {
	Title           string `json:"title"`
	ActivityGroupID int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
}

type UpdateTodo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool   `json:"is_active"`
}

type DeleteTodo struct {
	ID int `json:"id"`
}
