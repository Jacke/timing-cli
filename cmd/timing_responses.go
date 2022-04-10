package cmd

import "time"

type GetAllProjectsReponse struct {
	Data []struct {
		Self              string      `json:"self"`
		TeamID            interface{} `json:"team_id"`
		Title             string      `json:"title"`
		TitleChain        []string    `json:"title_chain"`
		Color             string      `json:"color"`
		ProductivityScore int         `json:"productivity_score"`
		IsArchived        bool        `json:"is_archived"`
		Children          []struct {
			Self              string        `json:"self"`
			TeamID            interface{}   `json:"team_id"`
			Title             string        `json:"title"`
			TitleChain        []string      `json:"title_chain"`
			Color             string        `json:"color"`
			ProductivityScore int           `json:"productivity_score"`
			IsArchived        bool          `json:"is_archived"`
			Children          []interface{} `json:"children"`
			Parent            struct {
				Self string `json:"self"`
			} `json:"parent"`
		} `json:"children"`
		Parent interface{} `json:"parent"`
	} `json:"data"`
}

type GetProjects struct {
	Data []struct {
		Self              string      `json:"self"`
		TeamID            interface{} `json:"team_id"`
		Title             string      `json:"title"`
		TitleChain        []string    `json:"title_chain"`
		Color             string      `json:"color"`
		ProductivityScore int         `json:"productivity_score"`
		IsArchived        bool        `json:"is_archived"`
		Children          []struct {
			Self string `json:"self"`
		} `json:"children"`
		Parent interface{} `json:"parent"`
	} `json:"data"`
}

type NewProjectResponse struct {
	Data struct {
		Self              string        `json:"self"`
		TeamID            interface{}   `json:"team_id"`
		Title             string        `json:"title"`
		TitleChain        []string      `json:"title_chain"`
		Color             string        `json:"color"`
		ProductivityScore int           `json:"productivity_score"`
		IsArchived        bool          `json:"is_archived"`
		Children          []interface{} `json:"children"`
		Parent            struct {
			Self string `json:"self"`
		} `json:"parent"`
	} `json:"data"`
	Links struct {
		TimeEntries string `json:"time-entries"`
	} `json:"links"`
}

type GetProjectResponse struct {
	Data struct {
		Self              string      `json:"self"`
		TeamID            interface{} `json:"team_id"`
		Title             string      `json:"title"`
		TitleChain        []string    `json:"title_chain"`
		Color             string      `json:"color"`
		ProductivityScore int         `json:"productivity_score"`
		IsArchived        bool        `json:"is_archived"`
		Children          []struct {
			Self string `json:"self"`
		} `json:"children"`
		Parent interface{} `json:"parent"`
	} `json:"data"`
	Links struct {
		TimeEntries string `json:"time-entries"`
	} `json:"links"`
}

type UpdateProjectResponse struct {
	Data struct {
		Self              string      `json:"self"`
		TeamID            interface{} `json:"team_id"`
		Title             string      `json:"title"`
		TitleChain        []string    `json:"title_chain"`
		Color             string      `json:"color"`
		ProductivityScore int         `json:"productivity_score"`
		IsArchived        bool        `json:"is_archived"`
		Children          []struct {
			Self string `json:"self"`
		} `json:"children"`
		Parent interface{} `json:"parent"`
	} `json:"data"`
	Links struct {
		TimeEntries string `json:"time-entries"`
	} `json:"links"`
}

type TimerResponse struct {
	Data struct {
		Self      string    `json:"self"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Duration  int       `json:"duration"`
		Project   struct {
			Self string `json:"self"`
		} `json:"project"`
		Title       string `json:"title"`
		Notes       string `json:"notes"`
		IsRunning   bool   `json:"is_running"`
		CreatorName string `json:"creator_name"`
	} `json:"data"`
	Message string `json:"message"`
}

type AllTimersResponse struct {
	Data []struct {
		Self      string    `json:"self"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Duration  int       `json:"duration"`
		Project   struct {
			Self              string      `json:"self"`
			TeamID            interface{} `json:"team_id"`
			Title             string      `json:"title"`
			TitleChain        []string    `json:"title_chain"`
			Color             string      `json:"color"`
			ProductivityScore int         `json:"productivity_score"`
			IsArchived        bool        `json:"is_archived"`
			Parent            interface{} `json:"parent"`
		} `json:"project"`
		Title       string `json:"title"`
		Notes       string `json:"notes"`
		IsRunning   *bool   `json:"is_running"`
		CreatorName string `json:"creator_name"`
	} `json:"data"`
	Links struct {
		First string      `json:"first"`
		Last  string      `json:"last"`
		Prev  interface{} `json:"prev"`
		Next  interface{} `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int `json:"current_page"`
		From        int `json:"from"`
		LastPage    int `json:"last_page"`
		Links       []struct {
			URL    interface{} `json:"url"`
			Label  string      `json:"label"`
			Active bool        `json:"active"`
		} `json:"links"`
		Path    string `json:"path"`
		PerPage int    `json:"per_page"`
		To      int    `json:"to"`
		Total   int    `json:"total"`
	} `json:"meta"`
}

type TimeEntryResponse struct {
	Data struct {
		Self      string    `json:"self"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Duration  int       `json:"duration"`
		Project   struct {
			Self string `json:"self"`
		} `json:"project"`
		Title       string `json:"title"`
		Notes       string `json:"notes"`
		IsRunning   bool   `json:"is_running"`
		CreatorName string `json:"creator_name"`
	} `json:"data"`
}

type ReportsResponse struct {
	Data []struct {
		Duration int `json:"duration"`
		Project  struct {
			Self              string      `json:"self"`
			TeamID            interface{} `json:"team_id"`
			Title             string      `json:"title"`
			TitleChain        []string    `json:"title_chain"`
			Color             string      `json:"color"`
			ProductivityScore int         `json:"productivity_score"`
			IsArchived        bool        `json:"is_archived"`
			Parent            interface{} `json:"parent"`
		} `json:"project"`
	} `json:"data"`
}

type TeamMembersResponse struct {
	Data []struct {
		Self  string `json:"self"`
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"data"`
}

type TeamsResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}
