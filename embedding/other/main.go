package main

import "GoStudyProject/embedding/wrongEmbedding"

func main() {
	inMem := wrongEmbedding.New()
	// inMem.M // M is exported, while m is not!
	inMem.Lock() // Lock is visible here
}
