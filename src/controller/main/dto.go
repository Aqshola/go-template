package controller_main

type GetMainResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetDetailMainResponse struct {
	Id     int                     `json:"id"`
	Name   string                  `json:"name"`
	Detail ChildDetailMainResponse `json:"detail"`
}

type ChildDetailMainResponse struct {
	Id       int    `json:"id"`
	IsDetail string `json:"isDetail"`
}
