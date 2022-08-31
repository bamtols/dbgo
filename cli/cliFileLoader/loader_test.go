package cliFileLoader

import "testing"

func TestLoader(test *testing.T) {
	test.Run("loader", func(t *testing.T) {
		_, err := LoadConfig("/Users/ciaolee/Projects/bamtol/dbgo/config.yml", "1.0")
		if err != nil {
			t.Fatal(err)
		}
	})

}
