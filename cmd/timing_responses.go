package cmd

type GetProjectsResponse struct {
	s string
}
type GetAllProjectsReponseData struct {
	SelfTag            string `json:"self"`
	Title              string `json:"title"`
	Color              string `json:"color"`
	Productivity_score int    `json:"productivity_score"`
	Is_archived        bool   `json:"is_archived"`
	//    "team_id"
	//		"title_chain": ["Efforts"],
}
type GetAllProjectsReponse struct {
	Data []GetAllProjectsReponseData `json:"data"`
}
