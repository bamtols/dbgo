package engines

import (
	"testing"
)

func TestClient(test *testing.T) {
	test.Run("connection", func(t *testing.T) {
		var client, err = NewClient(&INewClient{
			Username: "root",
			Password: "Ba70LgAM3s",
			URL:      "15.165.39.196:3306",
			Params:   ptrString(INewClientParamDefault),
			Database: "test_db",
			DBType:   DBTypeMySQL,
		})

		if err != nil {
			t.Fatal(err)
		}

		stats := client.root.Stats()
		println(stats.InUse)
	})
}
