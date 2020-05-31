package util

var excludedURLs = []string{"http://tech-challenge-golang.s3-website.eu-central-1.amazonaws.com/post/", "http://tech-challenge-golang.s3-website.eu-central-1.amazonaws.com/"}

func Contains(value string) bool {
	for _, url := range excludedURLs {
		if url == value {
			return true
		}
	}
	return false
}

func RemoveValuFromIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
