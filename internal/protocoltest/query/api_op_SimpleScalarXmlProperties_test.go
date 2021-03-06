// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_SimpleScalarXmlProperties_awsAwsqueryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *SimpleScalarXmlPropertiesOutput
	}{
		// Serializes simple scalar properties
		"QuerySimpleScalarProperties": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarXmlPropertiesResponse xmlns="https://example.com/">
			    <SimpleScalarXmlPropertiesResult>
			        <stringValue>string</stringValue>
			        <emptyStringValue/>
			        <trueBooleanValue>true</trueBooleanValue>
			        <falseBooleanValue>false</falseBooleanValue>
			        <byteValue>1</byteValue>
			        <shortValue>2</shortValue>
			        <integerValue>3</integerValue>
			        <longValue>4</longValue>
			        <floatValue>5.5</floatValue>
			        <DoubleDribble>6.5</DoubleDribble>
			    </SimpleScalarXmlPropertiesResult>
			</SimpleScalarXmlPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarXmlPropertiesOutput{
				StringValue:       ptr.String("string"),
				EmptyStringValue:  ptr.String(""),
				TrueBooleanValue:  ptr.Bool(true),
				FalseBooleanValue: ptr.Bool(false),
				ByteValue:         ptr.Int8(1),
				ShortValue:        ptr.Int16(2),
				IntegerValue:      ptr.Int32(3),
				LongValue:         ptr.Int64(4),
				FloatValue:        ptr.Float32(5.5),
				DoubleValue:       ptr.Float64(6.5),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params SimpleScalarXmlPropertiesInput
			result, err := client.SimpleScalarXmlProperties(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
