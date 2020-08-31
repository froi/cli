// +build !windows

package iostreams

type colorableWriter struct{}

func (*colorableWriter) Write(_ []byte) (int, error) {
	return 0, nil
}
