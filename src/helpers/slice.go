package helpers

/*
Deletes an element from a slice at the index specified and returns the new slice
*/
func Delete[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
