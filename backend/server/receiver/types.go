package receiver

type JSONIdRoom struct {
	Id string `json:"id"`
}

type JSONUsername struct {
	Username string `json:"username"`
}

type JSONSetReady struct {
	JSONIdRoom
	JSONUsername
	Ready bool `json:"ready"`
}

type JSONTyper struct {
	JSONIdRoom
	JSONUsername
}
