// Code generated by smithy-go-codegen DO NOT EDIT.

package dlm

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dlm/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateLifecyclePolicy struct {
}

func (*awsRestjson1_serializeOpCreateLifecyclePolicy) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateLifecyclePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateLifecyclePolicyInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/policies")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
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
	if err := awsRestjson1_serializeOpDocumentCreateLifecyclePolicyInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsCreateLifecyclePolicyInput(v *CreateLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateLifecyclePolicyInput(v *CreateLifecyclePolicyInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Description != nil {
		ok := object.Key("Description")
		ok.String(*v.Description)
	}

	if v.ExecutionRoleArn != nil {
		ok := object.Key("ExecutionRoleArn")
		ok.String(*v.ExecutionRoleArn)
	}

	if v.PolicyDetails != nil {
		ok := object.Key("PolicyDetails")
		if err := awsRestjson1_serializeDocumentPolicyDetails(v.PolicyDetails, ok); err != nil {
			return err
		}
	}

	if len(v.State) > 0 {
		ok := object.Key("State")
		ok.String(string(v.State))
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDeleteLifecyclePolicy struct {
}

func (*awsRestjson1_serializeOpDeleteLifecyclePolicy) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteLifecyclePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteLifecyclePolicyInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/policies/{PolicyId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
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

	if err := awsRestjson1_serializeOpHttpBindingsDeleteLifecyclePolicyInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteLifecyclePolicyInput(v *DeleteLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.PolicyId == nil || len(*v.PolicyId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member PolicyId must not be empty")}
	}
	if v.PolicyId != nil {
		if err := encoder.SetURI("PolicyId").String(*v.PolicyId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetLifecyclePolicies struct {
}

func (*awsRestjson1_serializeOpGetLifecyclePolicies) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetLifecyclePolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetLifecyclePoliciesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/policies")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
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

	if err := awsRestjson1_serializeOpHttpBindingsGetLifecyclePoliciesInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetLifecyclePoliciesInput(v *GetLifecyclePoliciesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.PolicyIds != nil {
		for i := range v.PolicyIds {
			encoder.AddQuery("policyIds").String(v.PolicyIds[i])
		}
	}

	if v.ResourceTypes != nil {
		for i := range v.ResourceTypes {
			encoder.AddQuery("resourceTypes").String(string(v.ResourceTypes[i]))
		}
	}

	if len(v.State) > 0 {
		encoder.SetQuery("state").String(string(v.State))
	}

	if v.TagsToAdd != nil {
		for i := range v.TagsToAdd {
			encoder.AddQuery("tagsToAdd").String(v.TagsToAdd[i])
		}
	}

	if v.TargetTags != nil {
		for i := range v.TargetTags {
			encoder.AddQuery("targetTags").String(v.TargetTags[i])
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetLifecyclePolicy struct {
}

func (*awsRestjson1_serializeOpGetLifecyclePolicy) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetLifecyclePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetLifecyclePolicyInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/policies/{PolicyId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
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

	if err := awsRestjson1_serializeOpHttpBindingsGetLifecyclePolicyInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetLifecyclePolicyInput(v *GetLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.PolicyId == nil || len(*v.PolicyId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member PolicyId must not be empty")}
	}
	if v.PolicyId != nil {
		if err := encoder.SetURI("PolicyId").String(*v.PolicyId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListTagsForResource struct {
}

func (*awsRestjson1_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{ResourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
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

	if err := awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member ResourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("ResourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpTagResource struct {
}

func (*awsRestjson1_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{ResourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
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

	if err := awsRestjson1_serializeOpHttpBindingsTagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member ResourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("ResourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpUntagResource struct {
}

func (*awsRestjson1_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{ResourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
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

	if err := awsRestjson1_serializeOpHttpBindingsUntagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member ResourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("ResourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	if v.TagKeys != nil {
		for i := range v.TagKeys {
			encoder.AddQuery("tagKeys").String(v.TagKeys[i])
		}
	}

	return nil
}

type awsRestjson1_serializeOpUpdateLifecyclePolicy struct {
}

func (*awsRestjson1_serializeOpUpdateLifecyclePolicy) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateLifecyclePolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateLifecyclePolicyInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/policies/{PolicyId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PATCH"
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

	if err := awsRestjson1_serializeOpHttpBindingsUpdateLifecyclePolicyInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateLifecyclePolicyInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsUpdateLifecyclePolicyInput(v *UpdateLifecyclePolicyInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.PolicyId == nil || len(*v.PolicyId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member PolicyId must not be empty")}
	}
	if v.PolicyId != nil {
		if err := encoder.SetURI("PolicyId").String(*v.PolicyId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateLifecyclePolicyInput(v *UpdateLifecyclePolicyInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Description != nil {
		ok := object.Key("Description")
		ok.String(*v.Description)
	}

	if v.ExecutionRoleArn != nil {
		ok := object.Key("ExecutionRoleArn")
		ok.String(*v.ExecutionRoleArn)
	}

	if v.PolicyDetails != nil {
		ok := object.Key("PolicyDetails")
		if err := awsRestjson1_serializeDocumentPolicyDetails(v.PolicyDetails, ok); err != nil {
			return err
		}
	}

	if len(v.State) > 0 {
		ok := object.Key("State")
		ok.String(string(v.State))
	}

	return nil
}

func awsRestjson1_serializeDocumentAction(v *types.Action, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CrossRegionCopy != nil {
		ok := object.Key("CrossRegionCopy")
		if err := awsRestjson1_serializeDocumentCrossRegionCopyActionList(v.CrossRegionCopy, ok); err != nil {
			return err
		}
	}

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	return nil
}

func awsRestjson1_serializeDocumentActionList(v []types.Action, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentAction(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentArchiveRetainRule(v *types.ArchiveRetainRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.RetentionArchiveTier != nil {
		ok := object.Key("RetentionArchiveTier")
		if err := awsRestjson1_serializeDocumentRetentionArchiveTier(v.RetentionArchiveTier, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentArchiveRule(v *types.ArchiveRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.RetainRule != nil {
		ok := object.Key("RetainRule")
		if err := awsRestjson1_serializeDocumentArchiveRetainRule(v.RetainRule, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentAvailabilityZoneList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentCreateRule(v *types.CreateRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CronExpression != nil {
		ok := object.Key("CronExpression")
		ok.String(*v.CronExpression)
	}

	if v.Interval != nil {
		ok := object.Key("Interval")
		ok.Integer(*v.Interval)
	}

	if len(v.IntervalUnit) > 0 {
		ok := object.Key("IntervalUnit")
		ok.String(string(v.IntervalUnit))
	}

	if len(v.Location) > 0 {
		ok := object.Key("Location")
		ok.String(string(v.Location))
	}

	if v.Times != nil {
		ok := object.Key("Times")
		if err := awsRestjson1_serializeDocumentTimesList(v.Times, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentCrossRegionCopyAction(v *types.CrossRegionCopyAction, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EncryptionConfiguration != nil {
		ok := object.Key("EncryptionConfiguration")
		if err := awsRestjson1_serializeDocumentEncryptionConfiguration(v.EncryptionConfiguration, ok); err != nil {
			return err
		}
	}

	if v.RetainRule != nil {
		ok := object.Key("RetainRule")
		if err := awsRestjson1_serializeDocumentCrossRegionCopyRetainRule(v.RetainRule, ok); err != nil {
			return err
		}
	}

	if v.Target != nil {
		ok := object.Key("Target")
		ok.String(*v.Target)
	}

	return nil
}

func awsRestjson1_serializeDocumentCrossRegionCopyActionList(v []types.CrossRegionCopyAction, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentCrossRegionCopyAction(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentCrossRegionCopyDeprecateRule(v *types.CrossRegionCopyDeprecateRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Interval != nil {
		ok := object.Key("Interval")
		ok.Integer(*v.Interval)
	}

	if len(v.IntervalUnit) > 0 {
		ok := object.Key("IntervalUnit")
		ok.String(string(v.IntervalUnit))
	}

	return nil
}

func awsRestjson1_serializeDocumentCrossRegionCopyRetainRule(v *types.CrossRegionCopyRetainRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Interval != nil {
		ok := object.Key("Interval")
		ok.Integer(*v.Interval)
	}

	if len(v.IntervalUnit) > 0 {
		ok := object.Key("IntervalUnit")
		ok.String(string(v.IntervalUnit))
	}

	return nil
}

func awsRestjson1_serializeDocumentCrossRegionCopyRule(v *types.CrossRegionCopyRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CmkArn != nil {
		ok := object.Key("CmkArn")
		ok.String(*v.CmkArn)
	}

	if v.CopyTags != nil {
		ok := object.Key("CopyTags")
		ok.Boolean(*v.CopyTags)
	}

	if v.DeprecateRule != nil {
		ok := object.Key("DeprecateRule")
		if err := awsRestjson1_serializeDocumentCrossRegionCopyDeprecateRule(v.DeprecateRule, ok); err != nil {
			return err
		}
	}

	if v.Encrypted != nil {
		ok := object.Key("Encrypted")
		ok.Boolean(*v.Encrypted)
	}

	if v.RetainRule != nil {
		ok := object.Key("RetainRule")
		if err := awsRestjson1_serializeDocumentCrossRegionCopyRetainRule(v.RetainRule, ok); err != nil {
			return err
		}
	}

	if v.Target != nil {
		ok := object.Key("Target")
		ok.String(*v.Target)
	}

	if v.TargetRegion != nil {
		ok := object.Key("TargetRegion")
		ok.String(*v.TargetRegion)
	}

	return nil
}

func awsRestjson1_serializeDocumentCrossRegionCopyRules(v []types.CrossRegionCopyRule, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentCrossRegionCopyRule(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentDeprecateRule(v *types.DeprecateRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Count != nil {
		ok := object.Key("Count")
		ok.Integer(*v.Count)
	}

	if v.Interval != nil {
		ok := object.Key("Interval")
		ok.Integer(*v.Interval)
	}

	if len(v.IntervalUnit) > 0 {
		ok := object.Key("IntervalUnit")
		ok.String(string(v.IntervalUnit))
	}

	return nil
}

func awsRestjson1_serializeDocumentEncryptionConfiguration(v *types.EncryptionConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CmkArn != nil {
		ok := object.Key("CmkArn")
		ok.String(*v.CmkArn)
	}

	if v.Encrypted != nil {
		ok := object.Key("Encrypted")
		ok.Boolean(*v.Encrypted)
	}

	return nil
}

func awsRestjson1_serializeDocumentEventParameters(v *types.EventParameters, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DescriptionRegex != nil {
		ok := object.Key("DescriptionRegex")
		ok.String(*v.DescriptionRegex)
	}

	if len(v.EventType) > 0 {
		ok := object.Key("EventType")
		ok.String(string(v.EventType))
	}

	if v.SnapshotOwner != nil {
		ok := object.Key("SnapshotOwner")
		if err := awsRestjson1_serializeDocumentSnapshotOwnerList(v.SnapshotOwner, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentEventSource(v *types.EventSource, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Parameters != nil {
		ok := object.Key("Parameters")
		if err := awsRestjson1_serializeDocumentEventParameters(v.Parameters, ok); err != nil {
			return err
		}
	}

	if len(v.Type) > 0 {
		ok := object.Key("Type")
		ok.String(string(v.Type))
	}

	return nil
}

func awsRestjson1_serializeDocumentExcludeDataVolumeTagList(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentFastRestoreRule(v *types.FastRestoreRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AvailabilityZones != nil {
		ok := object.Key("AvailabilityZones")
		if err := awsRestjson1_serializeDocumentAvailabilityZoneList(v.AvailabilityZones, ok); err != nil {
			return err
		}
	}

	if v.Count != nil {
		ok := object.Key("Count")
		ok.Integer(*v.Count)
	}

	if v.Interval != nil {
		ok := object.Key("Interval")
		ok.Integer(*v.Interval)
	}

	if len(v.IntervalUnit) > 0 {
		ok := object.Key("IntervalUnit")
		ok.String(string(v.IntervalUnit))
	}

	return nil
}

func awsRestjson1_serializeDocumentParameters(v *types.Parameters, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExcludeBootVolume != nil {
		ok := object.Key("ExcludeBootVolume")
		ok.Boolean(*v.ExcludeBootVolume)
	}

	if v.ExcludeDataVolumeTags != nil {
		ok := object.Key("ExcludeDataVolumeTags")
		if err := awsRestjson1_serializeDocumentExcludeDataVolumeTagList(v.ExcludeDataVolumeTags, ok); err != nil {
			return err
		}
	}

	if v.NoReboot != nil {
		ok := object.Key("NoReboot")
		ok.Boolean(*v.NoReboot)
	}

	return nil
}

func awsRestjson1_serializeDocumentPolicyDetails(v *types.PolicyDetails, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Actions != nil {
		ok := object.Key("Actions")
		if err := awsRestjson1_serializeDocumentActionList(v.Actions, ok); err != nil {
			return err
		}
	}

	if v.EventSource != nil {
		ok := object.Key("EventSource")
		if err := awsRestjson1_serializeDocumentEventSource(v.EventSource, ok); err != nil {
			return err
		}
	}

	if v.Parameters != nil {
		ok := object.Key("Parameters")
		if err := awsRestjson1_serializeDocumentParameters(v.Parameters, ok); err != nil {
			return err
		}
	}

	if len(v.PolicyType) > 0 {
		ok := object.Key("PolicyType")
		ok.String(string(v.PolicyType))
	}

	if v.ResourceLocations != nil {
		ok := object.Key("ResourceLocations")
		if err := awsRestjson1_serializeDocumentResourceLocationList(v.ResourceLocations, ok); err != nil {
			return err
		}
	}

	if v.ResourceTypes != nil {
		ok := object.Key("ResourceTypes")
		if err := awsRestjson1_serializeDocumentResourceTypeValuesList(v.ResourceTypes, ok); err != nil {
			return err
		}
	}

	if v.Schedules != nil {
		ok := object.Key("Schedules")
		if err := awsRestjson1_serializeDocumentScheduleList(v.Schedules, ok); err != nil {
			return err
		}
	}

	if v.TargetTags != nil {
		ok := object.Key("TargetTags")
		if err := awsRestjson1_serializeDocumentTargetTagList(v.TargetTags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentResourceLocationList(v []types.ResourceLocationValues, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentResourceTypeValuesList(v []types.ResourceTypeValues, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentRetainRule(v *types.RetainRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Count != nil {
		ok := object.Key("Count")
		ok.Integer(*v.Count)
	}

	if v.Interval != nil {
		ok := object.Key("Interval")
		ok.Integer(*v.Interval)
	}

	if len(v.IntervalUnit) > 0 {
		ok := object.Key("IntervalUnit")
		ok.String(string(v.IntervalUnit))
	}

	return nil
}

func awsRestjson1_serializeDocumentRetentionArchiveTier(v *types.RetentionArchiveTier, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Count != nil {
		ok := object.Key("Count")
		ok.Integer(*v.Count)
	}

	if v.Interval != nil {
		ok := object.Key("Interval")
		ok.Integer(*v.Interval)
	}

	if len(v.IntervalUnit) > 0 {
		ok := object.Key("IntervalUnit")
		ok.String(string(v.IntervalUnit))
	}

	return nil
}

func awsRestjson1_serializeDocumentSchedule(v *types.Schedule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ArchiveRule != nil {
		ok := object.Key("ArchiveRule")
		if err := awsRestjson1_serializeDocumentArchiveRule(v.ArchiveRule, ok); err != nil {
			return err
		}
	}

	if v.CopyTags != nil {
		ok := object.Key("CopyTags")
		ok.Boolean(*v.CopyTags)
	}

	if v.CreateRule != nil {
		ok := object.Key("CreateRule")
		if err := awsRestjson1_serializeDocumentCreateRule(v.CreateRule, ok); err != nil {
			return err
		}
	}

	if v.CrossRegionCopyRules != nil {
		ok := object.Key("CrossRegionCopyRules")
		if err := awsRestjson1_serializeDocumentCrossRegionCopyRules(v.CrossRegionCopyRules, ok); err != nil {
			return err
		}
	}

	if v.DeprecateRule != nil {
		ok := object.Key("DeprecateRule")
		if err := awsRestjson1_serializeDocumentDeprecateRule(v.DeprecateRule, ok); err != nil {
			return err
		}
	}

	if v.FastRestoreRule != nil {
		ok := object.Key("FastRestoreRule")
		if err := awsRestjson1_serializeDocumentFastRestoreRule(v.FastRestoreRule, ok); err != nil {
			return err
		}
	}

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	if v.RetainRule != nil {
		ok := object.Key("RetainRule")
		if err := awsRestjson1_serializeDocumentRetainRule(v.RetainRule, ok); err != nil {
			return err
		}
	}

	if v.ShareRules != nil {
		ok := object.Key("ShareRules")
		if err := awsRestjson1_serializeDocumentShareRules(v.ShareRules, ok); err != nil {
			return err
		}
	}

	if v.TagsToAdd != nil {
		ok := object.Key("TagsToAdd")
		if err := awsRestjson1_serializeDocumentTagsToAddList(v.TagsToAdd, ok); err != nil {
			return err
		}
	}

	if v.VariableTags != nil {
		ok := object.Key("VariableTags")
		if err := awsRestjson1_serializeDocumentVariableTagsList(v.VariableTags, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentScheduleList(v []types.Schedule, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentSchedule(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentShareRule(v *types.ShareRule, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.TargetAccounts != nil {
		ok := object.Key("TargetAccounts")
		if err := awsRestjson1_serializeDocumentShareTargetAccountList(v.TargetAccounts, ok); err != nil {
			return err
		}
	}

	if v.UnshareInterval != nil {
		ok := object.Key("UnshareInterval")
		ok.Integer(*v.UnshareInterval)
	}

	if len(v.UnshareIntervalUnit) > 0 {
		ok := object.Key("UnshareIntervalUnit")
		ok.String(string(v.UnshareIntervalUnit))
	}

	return nil
}

func awsRestjson1_serializeDocumentShareRules(v []types.ShareRule, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentShareRule(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentShareTargetAccountList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSnapshotOwnerList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentTag(v *types.Tag, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsRestjson1_serializeDocumentTagMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsRestjson1_serializeDocumentTagsToAddList(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentTargetTagList(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentTimesList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentVariableTagsList(v []types.Tag, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentTag(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}
