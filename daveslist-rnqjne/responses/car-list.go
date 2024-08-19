package responses

type CarCategoriesResp struct {
	CategoryName string     `json:"category_name"`
	CarList      []ListResp `json:"car_list"`
}

type ListResp struct {
	Title   string     `json:"title"`
	Content string     `json:"content"`
	FileImg []FileResp `json:"file_img"`
	Reply   []ReplyResp `json:"reply"`
}

type FileResp struct {
	URL string `json:"url"`
}

type ReplyResp struct {
	Text string `json:"text"`
}
