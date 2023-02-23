package app

type successRespone struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func SuccessRespone(data, paging, filter interface{}) *successRespone {
	return &successRespone{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessRespone(data interface{}) *successRespone {
	return SuccessRespone(data, nil, nil)
}
