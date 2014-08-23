// +build !solution

// Leave an empty line above this comment.
package config

func LoadGobConfig(file string) (conf *Configuration, err error) {
	// Create a dummy Configuration object into which to decode it
	conf = &Configuration{}
	err = nil
	// TODO: Open file using os.Open()
	// TODO: Decode using a gob decoder.
	return
}
