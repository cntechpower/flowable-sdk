# AttachmentResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**UserId** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Type_** | **string** | Can be any arbitrary value. When a valid formatted media-type (e.g. application/xml, text/plain) is included, the binary content HTTP response content-type will be set the given value. | [optional] [default to null]
**TaskUrl** | **string** |  | [optional] [default to null]
**ProcessInstanceUrl** | **string** |  | [optional] [default to null]
**ExternalUrl** | **string** | contentUrl:In case the attachment is a link to an external resource, the externalUrl contains the URL to the external content. If the attachment content is present in the Flowable engine, the contentUrl will contain an URL where the binary content can be streamed from. | [optional] [default to null]
**ContentUrl** | **string** |  | [optional] [default to null]
**Time** | [**time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

