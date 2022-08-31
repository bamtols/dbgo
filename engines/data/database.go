package data

type (
	DatabaseMetadata struct {
		Nm      string
		CharSet CharSet
	}

	CharSet string
)

const (
	CharSetUTF8MB4 CharSet = "UTF8MB4"
)
