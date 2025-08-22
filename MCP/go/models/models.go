package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetDataProtectionPolicyResponse represents the GetDataProtectionPolicyResponse schema from the OpenAPI specification
type GetDataProtectionPolicyResponse struct {
	Dataprotectionpolicy interface{} `json:"DataProtectionPolicy,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tags interface{} `json:"Tags"`
}

// VerifySMSSandboxPhoneNumberResult represents the VerifySMSSandboxPhoneNumberResult schema from the OpenAPI specification
type VerifySMSSandboxPhoneNumberResult struct {
}

// GetPlatformApplicationAttributesInput represents the GetPlatformApplicationAttributesInput schema from the OpenAPI specification
type GetPlatformApplicationAttributesInput struct {
	Platformapplicationarn interface{} `json:"PlatformApplicationArn"`
}

// PublishBatchInput represents the PublishBatchInput schema from the OpenAPI specification
type PublishBatchInput struct {
	Publishbatchrequestentries interface{} `json:"PublishBatchRequestEntries"`
	Topicarn interface{} `json:"TopicArn"`
}

// ListSubscriptionsByTopicInput represents the ListSubscriptionsByTopicInput schema from the OpenAPI specification
type ListSubscriptionsByTopicInput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Topicarn interface{} `json:"TopicArn"`
}

// SetSMSAttributesInput represents the SetSMSAttributesInput schema from the OpenAPI specification
type SetSMSAttributesInput struct {
	Attributes interface{} `json:"attributes"`
}

// GetEndpointAttributesResponse represents the GetEndpointAttributesResponse schema from the OpenAPI specification
type GetEndpointAttributesResponse struct {
	Attributes interface{} `json:"Attributes,omitempty"`
}

// SetSubscriptionAttributesInput represents the SetSubscriptionAttributesInput schema from the OpenAPI specification
type SetSubscriptionAttributesInput struct {
	Attributename interface{} `json:"AttributeName"`
	Attributevalue interface{} `json:"AttributeValue,omitempty"`
	Subscriptionarn interface{} `json:"SubscriptionArn"`
}

// ListSubscriptionsInput represents the ListSubscriptionsInput schema from the OpenAPI specification
type ListSubscriptionsInput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SetSMSAttributesResponse represents the SetSMSAttributesResponse schema from the OpenAPI specification
type SetSMSAttributesResponse struct {
}

// GetSMSAttributesResponse represents the GetSMSAttributesResponse schema from the OpenAPI specification
type GetSMSAttributesResponse struct {
	Attributes interface{} `json:"attributes,omitempty"`
}

// DeleteTopicInput represents the DeleteTopicInput schema from the OpenAPI specification
type DeleteTopicInput struct {
	Topicarn interface{} `json:"TopicArn"`
}

// GetSMSAttributesInput represents the GetSMSAttributesInput schema from the OpenAPI specification
type GetSMSAttributesInput struct {
	Attributes interface{} `json:"attributes,omitempty"`
}

// BatchResultErrorEntry represents the BatchResultErrorEntry schema from the OpenAPI specification
type BatchResultErrorEntry struct {
	Code interface{} `json:"Code"`
	Id interface{} `json:"Id"`
	Message interface{} `json:"Message,omitempty"`
	Senderfault interface{} `json:"SenderFault"`
}

// MessageAttributeMap represents the MessageAttributeMap schema from the OpenAPI specification
type MessageAttributeMap struct {
}

// CreatePlatformEndpointInput represents the CreatePlatformEndpointInput schema from the OpenAPI specification
type CreatePlatformEndpointInput struct {
	Attributes interface{} `json:"Attributes,omitempty"`
	Customuserdata interface{} `json:"CustomUserData,omitempty"`
	Platformapplicationarn interface{} `json:"PlatformApplicationArn"`
	Token interface{} `json:"Token"`
}

// GetSMSSandboxAccountStatusInput represents the GetSMSSandboxAccountStatusInput schema from the OpenAPI specification
type GetSMSSandboxAccountStatusInput struct {
}

// ListPlatformApplicationsResponse represents the ListPlatformApplicationsResponse schema from the OpenAPI specification
type ListPlatformApplicationsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Platformapplications interface{} `json:"PlatformApplications,omitempty"`
}

// PublishBatchResponse represents the PublishBatchResponse schema from the OpenAPI specification
type PublishBatchResponse struct {
	Failed interface{} `json:"Failed,omitempty"`
	Successful interface{} `json:"Successful,omitempty"`
}

// PutDataProtectionPolicyInput represents the PutDataProtectionPolicyInput schema from the OpenAPI specification
type PutDataProtectionPolicyInput struct {
	Dataprotectionpolicy interface{} `json:"DataProtectionPolicy"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// ListSMSSandboxPhoneNumbersInput represents the ListSMSSandboxPhoneNumbersInput schema from the OpenAPI specification
type ListSMSSandboxPhoneNumbersInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// PublishBatchRequestEntry represents the PublishBatchRequestEntry schema from the OpenAPI specification
type PublishBatchRequestEntry struct {
	Messagededuplicationid interface{} `json:"MessageDeduplicationId,omitempty"`
	Messagegroupid interface{} `json:"MessageGroupId,omitempty"`
	Messagestructure interface{} `json:"MessageStructure,omitempty"`
	Subject interface{} `json:"Subject,omitempty"`
	Id interface{} `json:"Id"`
	Message interface{} `json:"Message"`
	Messageattributes interface{} `json:"MessageAttributes,omitempty"`
}

// Subscription represents the Subscription schema from the OpenAPI specification
type Subscription struct {
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
	Subscriptionarn interface{} `json:"SubscriptionArn,omitempty"`
	Topicarn interface{} `json:"TopicArn,omitempty"`
}

// ListSMSSandboxPhoneNumbersResult represents the ListSMSSandboxPhoneNumbersResult schema from the OpenAPI specification
type ListSMSSandboxPhoneNumbersResult struct {
	Phonenumbers interface{} `json:"PhoneNumbers"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteSMSSandboxPhoneNumberResult represents the DeleteSMSSandboxPhoneNumberResult schema from the OpenAPI specification
type DeleteSMSSandboxPhoneNumberResult struct {
}

// CheckIfPhoneNumberIsOptedOutResponse represents the CheckIfPhoneNumberIsOptedOutResponse schema from the OpenAPI specification
type CheckIfPhoneNumberIsOptedOutResponse struct {
	Isoptedout interface{} `json:"isOptedOut,omitempty"`
}

// SetPlatformApplicationAttributesInput represents the SetPlatformApplicationAttributesInput schema from the OpenAPI specification
type SetPlatformApplicationAttributesInput struct {
	Attributes interface{} `json:"Attributes"`
	Platformapplicationarn interface{} `json:"PlatformApplicationArn"`
}

// CreateEndpointResponse represents the CreateEndpointResponse schema from the OpenAPI specification
type CreateEndpointResponse struct {
	Endpointarn interface{} `json:"EndpointArn,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// PlatformApplication represents the PlatformApplication schema from the OpenAPI specification
type PlatformApplication struct {
	Attributes interface{} `json:"Attributes,omitempty"`
	Platformapplicationarn interface{} `json:"PlatformApplicationArn,omitempty"`
}

// ListOriginationNumbersRequest represents the ListOriginationNumbersRequest schema from the OpenAPI specification
type ListOriginationNumbersRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateSMSSandboxPhoneNumberInput represents the CreateSMSSandboxPhoneNumberInput schema from the OpenAPI specification
type CreateSMSSandboxPhoneNumberInput struct {
	Phonenumber interface{} `json:"PhoneNumber"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
}

// CreateTopicResponse represents the CreateTopicResponse schema from the OpenAPI specification
type CreateTopicResponse struct {
	Topicarn interface{} `json:"TopicArn,omitempty"`
}

// ListPhoneNumbersOptedOutInput represents the ListPhoneNumbersOptedOutInput schema from the OpenAPI specification
type ListPhoneNumbersOptedOutInput struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// SMSSandboxPhoneNumber represents the SMSSandboxPhoneNumber schema from the OpenAPI specification
type SMSSandboxPhoneNumber struct {
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// UnsubscribeInput represents the UnsubscribeInput schema from the OpenAPI specification
type UnsubscribeInput struct {
	Subscriptionarn interface{} `json:"SubscriptionArn"`
}

// GetPlatformApplicationAttributesResponse represents the GetPlatformApplicationAttributesResponse schema from the OpenAPI specification
type GetPlatformApplicationAttributesResponse struct {
	Attributes interface{} `json:"Attributes,omitempty"`
}

// PublishInput represents the PublishInput schema from the OpenAPI specification
type PublishInput struct {
	Subject interface{} `json:"Subject,omitempty"`
	Messagegroupid interface{} `json:"MessageGroupId,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Topicarn interface{} `json:"TopicArn,omitempty"`
	Messagededuplicationid interface{} `json:"MessageDeduplicationId,omitempty"`
	Targetarn interface{} `json:"TargetArn,omitempty"`
	Message interface{} `json:"Message"`
	Messageattributes interface{} `json:"MessageAttributes,omitempty"`
	Messagestructure interface{} `json:"MessageStructure,omitempty"`
}

// Endpoint represents the Endpoint schema from the OpenAPI specification
type Endpoint struct {
	Attributes interface{} `json:"Attributes,omitempty"`
	Endpointarn interface{} `json:"EndpointArn,omitempty"`
}

// PublishBatchResultEntry represents the PublishBatchResultEntry schema from the OpenAPI specification
type PublishBatchResultEntry struct {
	Id interface{} `json:"Id,omitempty"`
	Messageid interface{} `json:"MessageId,omitempty"`
	Sequencenumber interface{} `json:"SequenceNumber,omitempty"`
}

// ListEndpointsByPlatformApplicationResponse represents the ListEndpointsByPlatformApplicationResponse schema from the OpenAPI specification
type ListEndpointsByPlatformApplicationResponse struct {
	Endpoints interface{} `json:"Endpoints,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// ListPhoneNumbersOptedOutResponse represents the ListPhoneNumbersOptedOutResponse schema from the OpenAPI specification
type ListPhoneNumbersOptedOutResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Phonenumbers interface{} `json:"phoneNumbers,omitempty"`
}

// CreatePlatformApplicationResponse represents the CreatePlatformApplicationResponse schema from the OpenAPI specification
type CreatePlatformApplicationResponse struct {
	Platformapplicationarn interface{} `json:"PlatformApplicationArn,omitempty"`
}

// CreateTopicInput represents the CreateTopicInput schema from the OpenAPI specification
type CreateTopicInput struct {
	Attributes interface{} `json:"Attributes,omitempty"`
	Dataprotectionpolicy interface{} `json:"DataProtectionPolicy,omitempty"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListTopicsInput represents the ListTopicsInput schema from the OpenAPI specification
type ListTopicsInput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SetEndpointAttributesInput represents the SetEndpointAttributesInput schema from the OpenAPI specification
type SetEndpointAttributesInput struct {
	Attributes interface{} `json:"Attributes"`
	Endpointarn interface{} `json:"EndpointArn"`
}

// MapStringToString represents the MapStringToString schema from the OpenAPI specification
type MapStringToString struct {
}

// CreateSMSSandboxPhoneNumberResult represents the CreateSMSSandboxPhoneNumberResult schema from the OpenAPI specification
type CreateSMSSandboxPhoneNumberResult struct {
}

// VerifySMSSandboxPhoneNumberInput represents the VerifySMSSandboxPhoneNumberInput schema from the OpenAPI specification
type VerifySMSSandboxPhoneNumberInput struct {
	Onetimepassword interface{} `json:"OneTimePassword"`
	Phonenumber interface{} `json:"PhoneNumber"`
}

// ListSubscriptionsResponse represents the ListSubscriptionsResponse schema from the OpenAPI specification
type ListSubscriptionsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Subscriptions interface{} `json:"Subscriptions,omitempty"`
}

// DeleteEndpointInput represents the DeleteEndpointInput schema from the OpenAPI specification
type DeleteEndpointInput struct {
	Endpointarn interface{} `json:"EndpointArn"`
}

// GetSubscriptionAttributesInput represents the GetSubscriptionAttributesInput schema from the OpenAPI specification
type GetSubscriptionAttributesInput struct {
	Subscriptionarn interface{} `json:"SubscriptionArn"`
}

// GetSMSSandboxAccountStatusResult represents the GetSMSSandboxAccountStatusResult schema from the OpenAPI specification
type GetSMSSandboxAccountStatusResult struct {
	Isinsandbox interface{} `json:"IsInSandbox"`
}

// ConfirmSubscriptionResponse represents the ConfirmSubscriptionResponse schema from the OpenAPI specification
type ConfirmSubscriptionResponse struct {
	Subscriptionarn interface{} `json:"SubscriptionArn,omitempty"`
}

// SubscriptionAttributesMap represents the SubscriptionAttributesMap schema from the OpenAPI specification
type SubscriptionAttributesMap struct {
}

// OptInPhoneNumberInput represents the OptInPhoneNumberInput schema from the OpenAPI specification
type OptInPhoneNumberInput struct {
	Phonenumber interface{} `json:"phoneNumber"`
}

// ListTopicsResponse represents the ListTopicsResponse schema from the OpenAPI specification
type ListTopicsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Topics interface{} `json:"Topics,omitempty"`
}

// GetSubscriptionAttributesResponse represents the GetSubscriptionAttributesResponse schema from the OpenAPI specification
type GetSubscriptionAttributesResponse struct {
	Attributes interface{} `json:"Attributes,omitempty"`
}

// GetDataProtectionPolicyInput represents the GetDataProtectionPolicyInput schema from the OpenAPI specification
type GetDataProtectionPolicyInput struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// CheckIfPhoneNumberIsOptedOutInput represents the CheckIfPhoneNumberIsOptedOutInput schema from the OpenAPI specification
type CheckIfPhoneNumberIsOptedOutInput struct {
	Phonenumber interface{} `json:"phoneNumber"`
}

// ListPlatformApplicationsInput represents the ListPlatformApplicationsInput schema from the OpenAPI specification
type ListPlatformApplicationsInput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MessageAttributeValue represents the MessageAttributeValue schema from the OpenAPI specification
type MessageAttributeValue struct {
	Binaryvalue interface{} `json:"BinaryValue,omitempty"`
	Datatype interface{} `json:"DataType"`
	Stringvalue interface{} `json:"StringValue,omitempty"`
}

// AddPermissionInput represents the AddPermissionInput schema from the OpenAPI specification
type AddPermissionInput struct {
	Actionname interface{} `json:"ActionName"`
	Label interface{} `json:"Label"`
	Topicarn interface{} `json:"TopicArn"`
	Awsaccountid interface{} `json:"AWSAccountId"`
}

// OptInPhoneNumberResponse represents the OptInPhoneNumberResponse schema from the OpenAPI specification
type OptInPhoneNumberResponse struct {
}

// GetTopicAttributesResponse represents the GetTopicAttributesResponse schema from the OpenAPI specification
type GetTopicAttributesResponse struct {
	Attributes interface{} `json:"Attributes,omitempty"`
}

// ListEndpointsByPlatformApplicationInput represents the ListEndpointsByPlatformApplicationInput schema from the OpenAPI specification
type ListEndpointsByPlatformApplicationInput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Platformapplicationarn interface{} `json:"PlatformApplicationArn"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// ListOriginationNumbersResult represents the ListOriginationNumbersResult schema from the OpenAPI specification
type ListOriginationNumbersResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Phonenumbers interface{} `json:"PhoneNumbers,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tagkeys interface{} `json:"TagKeys"`
}

// SubscribeResponse represents the SubscribeResponse schema from the OpenAPI specification
type SubscribeResponse struct {
	Subscriptionarn interface{} `json:"SubscriptionArn,omitempty"`
}

// RemovePermissionInput represents the RemovePermissionInput schema from the OpenAPI specification
type RemovePermissionInput struct {
	Label interface{} `json:"Label"`
	Topicarn interface{} `json:"TopicArn"`
}

// CreatePlatformApplicationInput represents the CreatePlatformApplicationInput schema from the OpenAPI specification
type CreatePlatformApplicationInput struct {
	Platform interface{} `json:"Platform"`
	Attributes interface{} `json:"Attributes"`
	Name interface{} `json:"Name"`
}

// ListSubscriptionsByTopicResponse represents the ListSubscriptionsByTopicResponse schema from the OpenAPI specification
type ListSubscriptionsByTopicResponse struct {
	Subscriptions interface{} `json:"Subscriptions,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// PublishResponse represents the PublishResponse schema from the OpenAPI specification
type PublishResponse struct {
	Messageid interface{} `json:"MessageId,omitempty"`
	Sequencenumber interface{} `json:"SequenceNumber,omitempty"`
}

// DeleteSMSSandboxPhoneNumberInput represents the DeleteSMSSandboxPhoneNumberInput schema from the OpenAPI specification
type DeleteSMSSandboxPhoneNumberInput struct {
	Phonenumber interface{} `json:"PhoneNumber"`
}

// PhoneNumberInformation represents the PhoneNumberInformation schema from the OpenAPI specification
type PhoneNumberInformation struct {
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Iso2countrycode interface{} `json:"Iso2CountryCode,omitempty"`
	Numbercapabilities interface{} `json:"NumberCapabilities,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Routetype interface{} `json:"RouteType,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// SetTopicAttributesInput represents the SetTopicAttributesInput schema from the OpenAPI specification
type SetTopicAttributesInput struct {
	Topicarn interface{} `json:"TopicArn"`
	Attributename interface{} `json:"AttributeName"`
	Attributevalue interface{} `json:"AttributeValue,omitempty"`
}

// GetTopicAttributesInput represents the GetTopicAttributesInput schema from the OpenAPI specification
type GetTopicAttributesInput struct {
	Topicarn interface{} `json:"TopicArn"`
}

// TopicAttributesMap represents the TopicAttributesMap schema from the OpenAPI specification
type TopicAttributesMap struct {
}

// GetEndpointAttributesInput represents the GetEndpointAttributesInput schema from the OpenAPI specification
type GetEndpointAttributesInput struct {
	Endpointarn interface{} `json:"EndpointArn"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// SubscribeInput represents the SubscribeInput schema from the OpenAPI specification
type SubscribeInput struct {
	Topicarn interface{} `json:"TopicArn"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Protocol interface{} `json:"Protocol"`
	Returnsubscriptionarn interface{} `json:"ReturnSubscriptionArn,omitempty"`
}

// ConfirmSubscriptionInput represents the ConfirmSubscriptionInput schema from the OpenAPI specification
type ConfirmSubscriptionInput struct {
	Authenticateonunsubscribe interface{} `json:"AuthenticateOnUnsubscribe,omitempty"`
	Token interface{} `json:"Token"`
	Topicarn interface{} `json:"TopicArn"`
}

// DeletePlatformApplicationInput represents the DeletePlatformApplicationInput schema from the OpenAPI specification
type DeletePlatformApplicationInput struct {
	Platformapplicationarn interface{} `json:"PlatformApplicationArn"`
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Topicarn interface{} `json:"TopicArn,omitempty"`
}
