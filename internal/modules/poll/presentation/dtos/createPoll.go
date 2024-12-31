package dtos

type OptionTexts struct {
    Text string `json:"text"`
}

type Poll struct {
    Name string `json:"name"`
    Options []OptionTexts `json:"options"`
}

type CreatePollRequest struct {
    Poll Poll `json:"poll"`
}