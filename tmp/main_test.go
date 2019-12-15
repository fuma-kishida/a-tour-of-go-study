package main

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var (
	input = []map[string][]string{
		map[string][]string{"cdp_id:1": []string{"1", "2", "3"}},
		map[string][]string{"cdp_id:2": []string{"1", "2", "3", "4"}},
		map[string][]string{"cdp_id:3": []string{"1", "2", "3", "4", "5"}},
		map[string][]string{"cdp_id:4": []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}},
		map[string][]string{"cdp_id:5": []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}},
	}
	h handler
)

func TestMain(m *testing.M) {
	// create aws session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Endpoint:    aws.String(dynamoEndpoint),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, token),
	})

	if err != nil {
		log.Fatal(err)
	}

	// create dynamo client
	ddb := dynamo.New(sess, aws.NewConfig().
		WithRegion(region).
		WithMaxRetries(3))

	h = handler{
		faqRecTable: ddb.Table(targetTableName),
	}

	//// delete table
	if err := h.faqRecTable.DeleteTable().Run(); err != nil {
		log.Println(err)
	}

	// create table
	if err := ddb.CreateTable(targetTableName, usersPageView{}).
		Run(); err != nil {
		log.Fatal(err)
	}

	// run test case
	ret := m.Run()

	// exit
	os.Exit(ret)

}

func Test_updateUrls(t *testing.T) {
	/*
		- テスト内容
			- Get => Put or Append => Truncate が目的の形でできているかテスト
		- 入力するパラメータ
			- 5つの (cdp_id,urls) の組合わせを 3周期分 入力する
			- 最終的に、うち2ペアにおいてはdynamoのレコードの上限値を超え、3ペアにおいては超えないとする
		- 期待する結果
			- レコードの上限値を超えない3ペア の場合
				- レコードの長さ => 正しく入力したURLの個数になっているか
				- レコードの内容 => 正しく入力したURLが入っているか
			- レコードの上限値を超える2ペア の場合
				- レコードの長さ => 正しく指定した閾値の個数に切り詰められているか
				- レコードの内容 => 正しく最新のURLたちが入っているか
	*/
	for cnt := 0; cnt < 3; {
		// => cdp_id:1
		if resGetFormer := h.getUrls("cdp_id:1"); len(resGetFormer) == 0 {
			h.putUrls("cdp_id:1", input[0]["cdp_id:1"])
		} else {
			h.appendUrls("cdp_id:1", input[0]["cdp_id:1"])
		}
		h.truncateUrls("cdp_id:1", threshold)

		// => cdp_id:2
		if resGetFormer := h.getUrls("cdp_id:2"); len(resGetFormer) == 0 {
			h.putUrls("cdp_id:2", input[1]["cdp_id:2"])
		} else {
			h.appendUrls("cdp_id:2", input[1]["cdp_id:2"])
		}
		h.truncateUrls("cdp_id:2", threshold)

		//// => cdp_id:3
		if resGetFormer := h.getUrls("cdp_id:3"); len(resGetFormer) == 0 {
			h.putUrls("cdp_id:3", input[2]["cdp_id:3"])
		} else {
			h.appendUrls("cdp_id:3", input[2]["cdp_id:3"])
		}
		h.truncateUrls("cdp_id:3", threshold)

		//// => cdp_id:4
		if resGetFormer := h.getUrls("cdp_id:4"); len(resGetFormer) == 0 {
			h.putUrls("cdp_id:4", input[3]["cdp_id:4"])
		} else {
			h.appendUrls("cdp_id:4", input[3]["cdp_id:4"])
		}
		h.truncateUrls("cdp_id:4", threshold)

		//// => cdp_id:5
		if resGetFormer := h.getUrls("cdp_id:5"); len(resGetFormer) == 0 {
			h.putUrls("cdp_id:5", input[4]["cdp_id:5"])
		} else {
			h.appendUrls("cdp_id:5", input[4]["cdp_id:5"])
		}
		h.truncateUrls("cdp_id:5", threshold)

		cnt += 1
	}

	// actual
	a1 := h.getUrls("cdp_id:1")[0].Urls
	a2 := h.getUrls("cdp_id:2")[0].Urls
	a3 := h.getUrls("cdp_id:3")[0].Urls
	a4 := h.getUrls("cdp_id:4")[0].Urls
	a5 := h.getUrls("cdp_id:5")[0].Urls

	actual := map[string]int{
		strings.Join(a1, ","): len(a1),
		strings.Join(a2, ","): len(a2),
		strings.Join(a3, ","): len(a3),
		strings.Join(a4, ","): len(a4),
		strings.Join(a5, ","): len(a5),
	}

	// expected
	e1 := append(append(input[0]["cdp_id:1"], input[0]["cdp_id:1"]...), input[0]["cdp_id:1"]...)
	e2 := append(append(input[1]["cdp_id:2"], input[1]["cdp_id:2"]...), input[1]["cdp_id:2"]...)
	e3 := append(append(input[2]["cdp_id:3"], input[2]["cdp_id:3"]...), input[2]["cdp_id:3"]...)
	e4 := append(append(input[3]["cdp_id:4"], input[3]["cdp_id:4"]...), input[3]["cdp_id:4"]...)
	e5 := append(append(input[4]["cdp_id:5"], input[4]["cdp_id:5"]...), input[4]["cdp_id:5"]...)

	expected := map[string]int{
		strings.Join(e1, ","):                     len(e1),
		strings.Join(e2, ","):                     len(e2),
		strings.Join(e3, ","):                     len(e3),
		strings.Join(e4[len(e4)-threshold:], ","): len(e4[len(e4)-threshold:]),
		strings.Join(e5[len(e5)-threshold:], ","): len(e5[len(e5)-threshold:]),
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
