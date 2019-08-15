package overwatch

// Response harbours the generic response returned in collections. This ought
// to be embedded in other response structs to supply these fields surplus.
type Response struct {
	Total int    `json:"total"`
	First string `json:"first"`
	Last  string `json:"last"`
}
