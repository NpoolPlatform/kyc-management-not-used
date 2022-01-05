# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/kyc-management.proto](#npool/kyc-management.proto)
    - [CreateKycRecordRequest](#kyc.management.v1.CreateKycRecordRequest)
    - [CreateKycRecordResponse](#kyc.management.v1.CreateKycRecordResponse)
    - [GetAllKycInfosRequest](#kyc.management.v1.GetAllKycInfosRequest)
    - [GetAllKycInfosResponse](#kyc.management.v1.GetAllKycInfosResponse)
    - [GetKycImgRequest](#kyc.management.v1.GetKycImgRequest)
    - [GetKycImgResponse](#kyc.management.v1.GetKycImgResponse)
    - [KycInfo](#kyc.management.v1.KycInfo)
    - [UpdateKycRequest](#kyc.management.v1.UpdateKycRequest)
    - [UpdateKycResponse](#kyc.management.v1.UpdateKycResponse)
    - [UpdateKycStatusRequest](#kyc.management.v1.UpdateKycStatusRequest)
    - [UpdateKycStatusResponse](#kyc.management.v1.UpdateKycStatusResponse)
    - [UploadKycImgRequest](#kyc.management.v1.UploadKycImgRequest)
    - [UploadKycImgResponse](#kyc.management.v1.UploadKycImgResponse)
    - [VersionResponse](#kyc.management.v1.VersionResponse)
  
    - [KycManagement](#kyc.management.v1.KycManagement)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/kyc-management.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/kyc-management.proto



<a name="kyc.management.v1.CreateKycRecordRequest"></a>

### CreateKycRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [KycInfo](#kyc.management.v1.KycInfo) |  |  |






<a name="kyc.management.v1.CreateKycRecordResponse"></a>

### CreateKycRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [KycInfo](#kyc.management.v1.KycInfo) |  |  |






<a name="kyc.management.v1.GetAllKycInfosRequest"></a>

### GetAllKycInfosRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| KycIDs | [string](#string) | repeated |  |






<a name="kyc.management.v1.GetAllKycInfosResponse"></a>

### GetAllKycInfosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [KycInfo](#kyc.management.v1.KycInfo) | repeated |  |






<a name="kyc.management.v1.GetKycImgRequest"></a>

### GetKycImgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| ImgID | [string](#string) |  |  |






<a name="kyc.management.v1.GetKycImgResponse"></a>

### GetKycImgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  | return ImgBase64 |






<a name="kyc.management.v1.KycInfo"></a>

### KycInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| FirstName | [string](#string) |  |  |
| LastName | [string](#string) |  |  |
| Region | [string](#string) |  |  |
| CardType | [string](#string) |  |  |
| CardID | [string](#string) |  |  |
| FrontCardImg | [string](#string) |  |  |
| BackCardImg | [string](#string) |  |  |
| UserHandlingCardImg | [string](#string) |  |  |
| ReviewStatus | [string](#string) |  |  |
| CreateAT | [uint32](#uint32) |  |  |
| UpdateAT | [uint32](#uint32) |  |  |






<a name="kyc.management.v1.UpdateKycRequest"></a>

### UpdateKycRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [KycInfo](#kyc.management.v1.KycInfo) |  |  |






<a name="kyc.management.v1.UpdateKycResponse"></a>

### UpdateKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [KycInfo](#kyc.management.v1.KycInfo) |  |  |






<a name="kyc.management.v1.UpdateKycStatusRequest"></a>

### UpdateKycStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Status | [int32](#int32) |  | 1: pass audit 2: refuse audit other: wait audit |
| UserID | [string](#string) |  |  |
| KycID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="kyc.management.v1.UpdateKycStatusResponse"></a>

### UpdateKycStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="kyc.management.v1.UploadKycImgRequest"></a>

### UploadKycImgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| ImgType | [string](#string) |  |  |
| ImgBase64 | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="kyc.management.v1.UploadKycImgResponse"></a>

### UploadKycImgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  | return ImgID |






<a name="kyc.management.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="kyc.management.v1.KycManagement"></a>

### KycManagement
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#kyc.management.v1.VersionResponse) | Method Version |
| CreateKycRecord | [CreateKycRecordRequest](#kyc.management.v1.CreateKycRecordRequest) | [CreateKycRecordResponse](#kyc.management.v1.CreateKycRecordResponse) |  |
| GetAllKycInfos | [GetAllKycInfosRequest](#kyc.management.v1.GetAllKycInfosRequest) | [GetAllKycInfosResponse](#kyc.management.v1.GetAllKycInfosResponse) |  |
| UpdateKycStatus | [UpdateKycStatusRequest](#kyc.management.v1.UpdateKycStatusRequest) | [UpdateKycStatusResponse](#kyc.management.v1.UpdateKycStatusResponse) | update kyc status |
| UpdateKyc | [UpdateKycRequest](#kyc.management.v1.UpdateKycRequest) | [UpdateKycResponse](#kyc.management.v1.UpdateKycResponse) |  |
| UploadKycImg | [UploadKycImgRequest](#kyc.management.v1.UploadKycImgRequest) | [UploadKycImgResponse](#kyc.management.v1.UploadKycImgResponse) |  |
| GetKycImg | [GetKycImgRequest](#kyc.management.v1.GetKycImgRequest) | [GetKycImgResponse](#kyc.management.v1.GetKycImgResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

