package protocols_test

import (
	"net/http"
	"testing"
	"reflect"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stretchr/testify/assert"
)

func TestProtocols(t *testing.T) {
	Convey("FormRequest", t, func() {
		Convey(".ToValues", func() {
			request := &compliance.SendRequest{
				Source: "Source",
				Sender: "Sender",
				Destination: "Destination",
				Amount: "Amount",
				AssetCode: "AssetCode",
				AssetIssuer: "AssetIssuer",
				SendMax: "SendMax",
				SendAssetCode: "SendAssetCode",
				SendAssetIssuer: "SendAssetIssuer",
				ExtraMemo: "ExtraMemo",
				Path: []protocols.Asset{
					protocols.Asset{"USD", "BLAH"},
					protocols.Asset{},
					protocols.Asset{"EUR", "BLAH2"},
				},
			}

			values := request.ToValues()

			httpRequest := &http.Request{PostForm: values}			

			request2 := &compliance.SendRequest{}
			request2.FromRequest(httpRequest)

			assert.True(t, reflect.DeepEqual(request, request2))
		})
	})
}