package server

import "regexp"

func addressAndPortGenerator(port string) string {
	var re = regexp.MustCompile(`(?m)([\w\.]*):?(\d{2,5})`)
	result := re.FindStringSubmatch(port)
	var address = ""

	if value := result[1]; value != "" {
		address += value
	}

	address += result[2]

	return result[0]
}
