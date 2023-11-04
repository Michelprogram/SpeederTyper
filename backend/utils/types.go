package utils

type RandomWordGenerator struct {
	Data []struct {
		Word struct {
			Value        string `json:"value"`
			NumSyllables int    `json:"numSyllables"`
		} `json:"word"`
	} `json:"data"`
}
