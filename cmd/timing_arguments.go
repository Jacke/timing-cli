package cmd

import "time"

// start_date_min	optional	Restricts the query to time entries whose start date is equal to or later than this parameter.
// start_date_max	optional	Restricts the query to time entries whose start date is equal to or earlier than this parameter.
// projects[]	optional	Restricts the query to time entries associated with the given project. Can be repeated to include time entries from several projects. If you would like to include time entries that are not assigned to any project, you can provide an empty string, i.e. projects[]=
// include_child_projects	optional	If true, the response will also contain time entries that belong to any child projects of the ones provided in projects[].
// search_query	optional	Restricts the query to time entries whose title and/or notes contain all words in this parameter. The search is case-insensitive but diacritic-sensitive.
// is_running	optional	If provided, returns only time entries that are either running or not running.
// include_project_data	optional	If true, the properties of the time entry's project will be included in the response.
// include_team_members	optional	If true, the response will also contain time entries that belong to other team members, provided the current user has permission to view them.
// team_members[]	optional	Restricts the query to data associated with the given user. Can be repeated to include time entries from several users.
type GetTimersArguments struct {
		StartDateMin *time.Time
		EndDateMin   *time.Time
		Projects		 *[]string
		TeamMembers  *[]string
		SearchQuery *string
		IsRunning *bool
		IncludeChildProjects *bool
		IncludeProjectData *bool
		IncludeTeamMembers *bool
	}

type StartTimerArguments struct {
	//	The title and project fields can not both be empty.
	// start_date	date	required	The time entry's start date and time.
	StartDate *time.Time
	// end_date	date	required	The time entry's end date and time.
	EndDate   *time.Time
	// project	project	optional	The project this time entry is associated with. Can be a project reference in the form "/projects/1", a project title (e.g. "Project at root level"), or an array with the project's entire title chain (e.g. ["Project at root level", "Unproductive child project"]).
	Project *string
	// title	string	optional	The time entry's title.
	Title *string
	// notes	string	optional	The time entry's notes.
	Notes *string
}