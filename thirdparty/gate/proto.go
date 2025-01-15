package gate

type SpotOrderDepthPair struct {
    Price string `json:"current"`
}
type SpotOrderDepthMsg struct {
    Current int64      `json:"current"`
    Update  int64      `json:"update"`
    Asks    [][]string `json:"asks"`
    Bids    [][]string `json:"bids"`
}
