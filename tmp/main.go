package main

import (
	"fmt"
	"log"
	"os"

	"github.com/guregu/dynamo"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	region             = "us-west-2"
	dynamoEndpoint     = "http://localhost:8000"
	awsAccessKeyId     = os.Getenv("AWS_ID")
	awsSecretAccessKey = os.Getenv("AWS_SECRET")
	token              = "dummy"
	targetTableName    = "faq-recommend-users-viewed-urls"
	cdpId              = "test"
	urls               = []string{"https://www.dmm.com/", "https://book.dmm.com/", "https://book.dmm.com/otherbooks/", "https://book.dmm.co.jp/"}
	//urls = []string{
	//	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	//	"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
	//	"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
	//	"31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
	//	"41", "42", "43", "44", "45", "46", "47", "48", "49", "50",
	//}
	threshold = 30
)

type handler struct {
	faqRecTable dynamo.Table
}

type usersPageView struct {
	CdpID string   `dynamo:"cdp_id,hash"`
	Urls  []string `dynamo:"urls"`
}

func (h *handler) getUrls(cdpId string) []usersPageView {
	fmt.Println("------> Get")
	var urls []usersPageView
	h.faqRecTable.Get("cdp_id", cdpId).All(&urls)
	return urls
}

func (h *handler) putUrls(cdpId string, urls []string) {
	fmt.Println("------> Put")
	h.faqRecTable.Put(usersPageView{
		CdpID: cdpId,
		Urls:  urls,
	}).Run()
}

func (h *handler) appendUrls(cdpId string, urls []string) {
	fmt.Println("------> Append")
	req := h.faqRecTable.Update("cdp_id", cdpId)
	req.Append("urls", urls).Run()
}

func (h *handler) truncateUrls(cdpId string, threshold int) {
	if resGetLatter := h.getUrls(cdpId); len(resGetLatter[0].Urls) > threshold {
		fmt.Println("------> Truncate")
		req := h.faqRecTable.Update("cdp_id", cdpId)
		req.Set("urls", resGetLatter[0].Urls[len(resGetLatter[0].Urls)-threshold:]).Run()
	} else {
		fmt.Println("There is no need to truncate because len(resGetLatter) : ", len(resGetLatter[0].Urls))
	}
}

func main() {
	// create dynamo config
	ddbConfig := aws.NewConfig().
		WithRegion(region).
		WithMaxRetries(3)

	// create aws session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Endpoint:    aws.String(dynamoEndpoint),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, token),
	})

	if err != nil {
		log.Fatalf("failed to initialize aws session: %v", err)
	}

	// create dynamo client
	ddb := dynamo.New(sess, ddbConfig)

	h := &handler{
		faqRecTable: ddb.Table(targetTableName),
	}
	fmt.Println(h)

	// confirm if requested cdp_id exists at table
	if resGetFormer := h.getUrls(cdpId); len(resGetFormer) == 0 {
		h.putUrls(cdpId, urls) // if not exists => put
	} else {
		h.appendUrls(cdpId, urls) // if not exists => append
	}

	// truncate length of record to the threshold
	h.truncateUrls(cdpId, threshold)

}
