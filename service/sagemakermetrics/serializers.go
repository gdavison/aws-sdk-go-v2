// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakermetrics

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sagemakermetrics/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"math"
)

type awsRestjson1_serializeOpBatchPutMetrics struct {
}

func (*awsRestjson1_serializeOpBatchPutMetrics) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpBatchPutMetrics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchPutMetricsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/BatchPutMetrics")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PUT"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentBatchPutMetricsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsBatchPutMetricsInput(v *BatchPutMetricsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentBatchPutMetricsInput(v *BatchPutMetricsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MetricData != nil {
		ok := object.Key("MetricData")
		if err := awsRestjson1_serializeDocumentRawMetricDataList(v.MetricData, ok); err != nil {
			return err
		}
	}

	if v.TrialComponentName != nil {
		ok := object.Key("TrialComponentName")
		ok.String(*v.TrialComponentName)
	}

	return nil
}

func awsRestjson1_serializeDocumentRawMetricData(v *types.RawMetricData, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MetricName != nil {
		ok := object.Key("MetricName")
		ok.String(*v.MetricName)
	}

	if v.Step != nil {
		ok := object.Key("Step")
		ok.Integer(*v.Step)
	}

	if v.Timestamp != nil {
		ok := object.Key("Timestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.Timestamp))
	}

	if v.Value != nil {
		ok := object.Key("Value")
		switch {
		case math.IsNaN(*v.Value):
			ok.String("NaN")

		case math.IsInf(*v.Value, 1):
			ok.String("Infinity")

		case math.IsInf(*v.Value, -1):
			ok.String("-Infinity")

		default:
			ok.Double(*v.Value)

		}
	}

	return nil
}

func awsRestjson1_serializeDocumentRawMetricDataList(v []types.RawMetricData, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentRawMetricData(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}
