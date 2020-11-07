package host

// Host ... active host
var (
	Host   string
	Scheme string
)

// GetURL ... Serve url from Scheme and Host
func GetURL() string {
	return Scheme + "://" + Host + "/"
}
