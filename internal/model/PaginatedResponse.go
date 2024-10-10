package model

type PaginatedResponse struct {
    Data       interface{} `json:"data"`      
    TotalCount int         `json:"total_count"` 
    Page       int         `json:"page"`      
    PageSize   int         `json:"page_size"`  
}
