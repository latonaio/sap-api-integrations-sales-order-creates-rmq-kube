package requests

type HeaderItem struct {
	Header
	ToItem `json:"to_Item"`
}

type ToItem struct {
	Results []Item `json:"results"`
}
