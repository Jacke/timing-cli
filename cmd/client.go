package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/tidwall/pretty"
)

type TimingClient struct {
	api_key string
	client  *resty.Client
}

func getClient() *TimingClient {
	timingClient := TimingClient{
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiYmNmY2VmNDQ0MTk5NWYwY2U2NjMyYmMzZmQ1ZmFmYzI4ODM5MjU3OGM0NTA5OGFjNDk0ZTA2OWI3MWViNmNmYjYzODhmYWI5N2Q3MTgxOTAiLCJpYXQiOjE2NDc4ODg2MzEuMjE1Mjk3LCJuYmYiOjE2NDc4ODg2MzEuMjE1MywiZXhwIjoxOTYzNTA3ODMxLjIxMDIzNywic3ViIjoiMzQ3MTI1NTAwMTQzODQ4MTkyMCIsInNjb3BlcyI6W119.U_fgUBP0O7Ynzz-yXB40HN7tC1hW4KahjpfusI9PO9CpnbvTB-eWZmnTnqXZSpkMnOzPO_0S39THLIUyUC1VvxwFwn7ce6EMBpYOfoJT8qbZu6npIKAvHA7VnA3rPtwKSfF6EDz-l-LW_QeTQ4aGWstdsfkXX51_mrVZuNUkCf9225cQTaiSma_6UiZDFxhhXYYJuLnHYdJLgJYCohmbnxKNYBjTJiN1ylACOBmcenChDo0LjLlyPYvFalbRCb7C8A3jtkASlyib2qw11KOAMZnefz6znCGyr7qFyQLyWgpfVzp9cGP63uDylMS05TBn8_jDuvUYRPS6pCxQO2LDA0DiVvdeq_h0dCGAEnoP0qta7hOx6nTwXIERALx7BRzJpFFPvLDLjof5TihroQpdT-md4NnWBg2YF_DejZJfDLZ1SX1iljEC4vyV9WDGLjXqaWfaGU_Nfuy7e-gA9gde_Myifk848NTL0WVB8juN6iAPX2o8j4Fvo043x4NbALcl9dUwJuORb8CBhZRy7FTR4L8GuLDpHA7i2MxkpQ1u1WaWTgwPw_x4jdoIlh1y8vBZ8asrjuvZNuothYtXGXk16reEI-7ZvfQn-EmteZ8HV8_xHAn_fgw4QQNPY1GQJnnkqMbR5MUDscX049phN21l_RlW5KWVH5jPMWHBp-n4R1w",
		resty.New(),
	}
	return &timingClient
}

func getProjects() (*resty.Response, error) {
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Get("https://web.timingapp.com/api/v1/projects/hierarchy")
	jsonResponse := pretty.Pretty(resp.Body()) //string(_)[:]
	getAllProjectsReponse := GetAllProjectsReponse{}
	errResp := json.Unmarshal(jsonResponse, &getAllProjectsReponse)
	if err != nil {
		fmt.Println(errResp)
		return resp, errResp
	}

	var data = [][]string{}
	for _, s := range getAllProjectsReponse.Data {
		toAdd := []string{s.Self, s.Title, strconv.Itoa(s.ProductivityScore)}
		data = append(data, toAdd)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Productivity"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return resp, err
}

func getAllProjects() (*resty.Response, error) {
	// Return a list containing all projects.
	// title	optional	Filter for projects whose title contains all words in this parameter. The search is case-insensitive but diacritic-sensitive.
	// hide_archived	optional	If set to true, archived projects and their children will not be included in the result.
	// team_id	optional	The ID of the team to list projects for. Can be omitted to list the user's private projects. See Return a list containing all the teams you are a member of. for obtaining a team ID to provide here.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Get("https://web.timingapp.com/api/v1/projects")
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func createNewProject() (*resty.Response, error) {
	// Create a new project.
	// title	string	required	The project's title.
	// parent	project	optional	A reference to an existing project. The new project will be appended to the parent's children. Can be a project reference in the form "/projects/1", a project title (e.g. "Project at root level"), or an array with the project's entire title chain (e.g. ["Project at root level", "Unproductive child project"]).
	// color	color	optional	The project's color, in hexadecimal format (#RRGGBB). If omitted, a color with random hue, 70% saturation and 100% value will be used.
	// productivity_score	float	optional	The project's productivity rating, between -1 (very unproductive) and 1 (very productive). Defaults to 1.
	// is_archived	boolean	optional	Whether the project has been archived. Defaults to false.
	// team_id	The	optional	ID of the team to add the project to. See Return a list containing all the teams you are a member of. for obtaining a team ID to provide here.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Post("https://web.timingapp.com/api/v1/projects")
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func getProject(id int) (*resty.Response, error) {
	// Display the specified project.
	// team_id	optional	The ID of the team to list projects for. Can be omitted to list the user's private projects. See Return a list containing all the teams you are a member of. for obtaining a team ID to provide here.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Get("https://web.timingapp.com/api/v1/projects/" + strconv.Itoa(id))
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func updateProject(id int) (*resty.Response, error) {
	// Update the specified project.
	// title	string	optional	The project's title.
	// color	color	optional	The project's color, in hexadecimal format (#RRGGBB).
	// productivity_score	float	optional	The project's productivity rating, between -1 (very unproductive) and 1 (very productive).
	// is_archived	boolean	optional	Whether the project has been archived.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Post("https://web.timingapp.com/api/v1/projects/" + strconv.Itoa(id))
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func deleteProject(id int) (*resty.Response, error) {
	// Delete the specified project and all of its children.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Delete("https://web.timingapp.com/api/v1/projects/" + strconv.Itoa(id))
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func startTimer() (*resty.Response, error) {
	// Start a new timer.
	// start_date	date	optional	The date this timer should have started at. Defaults to "now". Example:
	// project	project	optional	The project this timer is associated with. Can be a project reference in the form "/projects/1", a project title (e.g. "Project at root level"), or an array with the project's entire title chain (e.g. ["Project at root level", "Unproductive child project"]).
	// title	string	optional	The timer's title.
	// notes	string	optional	The timer's notes.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Post("https://web.timingapp.com/api/v1/time-entries/start")
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func stopTimer() (*resty.Response, error) {
	// Stop the currently running timer.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Put("https://web.timingapp.com/api/v1/time-entries/stop")
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func getTimers(arguments GetTimersArguments) (*AllTimersResponse, error) {
	// Return a list of time entries.
	// start_date_min	optional	Restricts the query to time entries whose start date is equal to or later than this parameter.
	// start_date_max	optional	Restricts the query to time entries whose start date is equal to or earlier than this parameter.
	// projects[]	optional	Restricts the query to time entries associated with the given project. Can be repeated to include time entries from several projects. If you would like to include time entries that are not assigned to any project, you can provide an empty string, i.e. projects[]=
	// include_child_projects	optional	If true, the response will also contain time entries that belong to any child projects of the ones provided in projects[].
	// search_query	optional	Restricts the query to time entries whose title and/or notes contain all words in this parameter. The search is case-insensitive but diacritic-sensitive.
	// is_running	optional	If provided, returns only time entries that are either running or not running.
	// include_project_data	optional	If true, the properties of the time entry's project will be included in the response.
	// include_team_members	optional	If true, the response will also contain time entries that belong to other team members, provided the current user has permission to view them.
	// team_members[]	optional	Restricts the query to data associated with the given user. Can be repeated to include time entries from several users.
	timingClient := getClient()
	queryArguments := make(map[string]string)

	if (arguments.IsRunning != nil && *arguments.IsRunning) {
		queryArguments["is_running"] = strconv.FormatBool(*arguments.IsRunning)
	}
	queryParams := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		SetQueryParams(queryArguments)

	resp, err := queryParams.
		Get("https://web.timingapp.com/api/v1/time-entries")

	jsonResponse := pretty.Pretty(resp.Body()) //string(_)[:]
	allTimersResponse := AllTimersResponse{}
	errResp := json.Unmarshal(jsonResponse, &allTimersResponse)
	if err != nil {
		fmt.Println(errResp)
		return nil, errResp
	}
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	for i, s := range allTimersResponse.Data {
		fmt.Println(strconv.Itoa(i+1)+". Timer: ", s.Title, "Is running", s.IsRunning, "Start & End time", s.StartDate, s.EndDate, s.Duration, time.Duration(s.Duration).Minutes())
	}
	return &allTimersResponse, err
}

func createTimeEntry() (*resty.Response, error) {
	// Create a new time entry.
	// start_date	date	required	The time entry's start date and time.
	// end_date	date	required	The time entry's end date and time.
	// project	project	optional	The project this time entry is associated with. Can be a project reference in the form "/projects/1", a project title (e.g. "Project at root level"), or an array with the project's entire title chain (e.g. ["Project at root level", "Unproductive child project"]).
	// title	string	optional	The time entry's title.
	// notes	string	optional	The time entry's notes.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Post("https://web.timingapp.com/api/v1/time-entries")
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func getTimeEntry(id int) (*resty.Response, error) {
	// Display the specified time entry.
	// start_date	date	optional	The date this timer should have started at. Defaults to "now". Example:
	// project	project	optional	The project this timer is associated with. Can be a project reference in the form "/projects/1", a project title (e.g. "Project at root level"), or an array with the project's entire title chain (e.g. ["Project at root level", "Unproductive child project"]).
	// title	string	optional	The timer's title.
	// notes	string	optional	The timer's notes.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Get("https://web.timingapp.com/api/v1/time-entries/" + strconv.Itoa(id))
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func updateTimeEntry(id int) (*resty.Response, error) {
	// Update the specified time entry.
	// start_date	date	optional	The time entry's start date and time.
	// end_date	date	optional	The time entry's start date and time.
	// project	project	optional	The project this time entry is associated with. Can be a project reference in the form "/projects/1", a project title (e.g. "Project at root level"), or an array with the project's entire title chain (e.g. ["Project at root level", "Unproductive child project"]).
	// title	string	optional	The time entry's title.
	// notes	string	optional	The time entry's notes.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Put("https://web.timingapp.com/api/v1/time-entries/" + strconv.Itoa(id))
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func deleteTimeEntry(id int) (*resty.Response, error) {
	// Delete the specified time entry.
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+timingClient.api_key).
		Delete("https://web.timingapp.com/api/v1/projects/" + strconv.Itoa(id))
	fmt.Println("  Body       :\n", string(pretty.Pretty(resp.Body()))[:])
	return resp, err
}

func Client() (*resty.Response, error) {
	timingClient := getClient()
	resp, err := timingClient.client.R().
		EnableTrace().
		Get("https://httpbin.org/get")
	// Explore response object
	// fmt.Println("Response Info:")
	// fmt.Println("  Error      :", err)
	// fmt.Println("  Status Code:", resp.StatusCode())
	// fmt.Println("  Status     :", resp.Status())
	// fmt.Println("  Proto      :", resp.Proto())
	// fmt.Println("  Time       :", resp.Time())
	// fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// // Explore trace info
	// fmt.Println("Request Trace Info:")
	// ti := resp.Request.TraceInfo()
	// fmt.Println("  DNSLookup     :", ti.DNSLookup)
	// fmt.Println("  ConnTime      :", ti.ConnTime)
	// fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	// fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	// fmt.Println("  ServerTime    :", ti.ServerTime)
	// fmt.Println("  ResponseTime  :", ti.ResponseTime)
	// fmt.Println("  TotalTime     :", ti.TotalTime)
	// fmt.Println("  IsConnReused  :", ti.IsConnReused)
	// fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	// fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	// fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	// fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
	return resp, err
}
