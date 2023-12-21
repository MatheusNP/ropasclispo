package domain

type FightRequest struct {
	Value uint8 `json:"value" validate:"required,gte=0,lte=4"`
}

type FightResponse struct {
	Enemy  Choice `json:"enemy"`
	Point  int8   `json:"point"`
	Result string `json:"result"`
}
