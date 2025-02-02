# Go Map Access with Non-String Key

This example demonstrates a common error in Go when working with maps: attempting to access a map element using a non-string key.  Go maps require string keys; using an integer key (or any other type besides string) will result in a runtime panic. The solution involves either using a string key or implementing a different data structure if a non-string key is necessary.