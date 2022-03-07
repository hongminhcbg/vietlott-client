package vietlott_client

import (
	"context"
	"encoding/json"
	"fmt"
	viettlot_client "github.com/hongminhcbg/viettlot_client/viettlot-client"
	"testing"
)

func Test_main(t *testing.T) {
	cli := viettlot_client.NewVietlottClient()
	result, err := cli.KenoLive(context.Background())
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(result, "", "\t")
	fmt.Println(string(b))
}
