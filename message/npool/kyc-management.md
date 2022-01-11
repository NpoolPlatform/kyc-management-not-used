# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/kyc-management.proto](#npool/kyc-management.proto)
    - [CreateKycRequest](#kyc.management.v1.CreateKycRequest)
    - [CreateKycResponse](#kyc.management.v1.CreateKycResponse)
    - [GetAllKycRequest](#kyc.management.v1.GetAllKycRequest)
    - [GetAllKycResponse](#kyc.management.v1.GetAllKycResponse)
    - [GetKycByAppIDRequest](#kyc.management.v1.GetKycByAppIDRequest)
    - [GetKycByAppIDResponse](#kyc.management.v1.GetKycByAppIDResponse)
    - [GetKycByUserIDRequest](#kyc.management.v1.GetKycByUserIDRequest)
    - [GetKycByUserIDResponse](#kyc.management.v1.GetKycByUserIDResponse)
    - [GetKycImageRequest](#kyc.management.v1.GetKycImageRequest)
    - [GetKycImageResponse](#kyc.management.v1.GetKycImageResponse)
    - [KycInfo](#kyc.management.v1.KycInfo)
    - [UpdateKycRequest](#kyc.management.v1.UpdateKycRequest)
    - [UpdateKycResponse](#kyc.management.v1.UpdateKycResponse)
    - [UploadKycImageRequest](#kyc.management.v1.UploadKycImageRequest)
    - [UploadKycImageResponse](#kyc.management.v1.UploadKycImageResponse)
    - [VersionResponse](#kyc.management.v1.VersionResponse)
  
    - [KycManagement](#kyc.management.v1.KycManagement)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/kyc-management.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/kyc-management.proto



<a name="kyc.management.v1.CreateKycRequest"></a>

### CreateKycRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| FirstName | [string](#string) |  |  |
| LastName | [string](#string) |  |  |
| Region | [string](#string) |  |  |
| CardType | [string](#string) |  |  |
| CardID | [string](#string) |  |  |
| FrontCardImg | [string](#string) |  |  |
| BackCardImg | [string](#string) |  |  |
| UserHandlingCardImg | [string](#string) |  |  |






<a name="kyc.management.v1.CreateKycResponse"></a>

### CreateKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [KycInfo](#kyc.management.v1.KycInfo) |  |  |






<a name="kyc.management.v1.GetAllKycRequest"></a>

### GetAllKycRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Limit | [int32](#int32) |  |  |
| Offset | [int32](#int32) |  |  |






<a name="kyc.management.v1.GetAllKycResponse"></a>

### GetAllKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [KycInfo](#kyc.management.v1.KycInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="kyc.management.v1.GetKycByAppIDRequest"></a>

### GetKycByAppIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Limit | [int32](#int32) |  |  |
| Offset | [int32](#int32) |  |  |






<a name="kyc.management.v1.GetKycByAppIDResponse"></a>

### GetKycByAppIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [KycInfo](#kyc.management.v1.KycInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="kyc.management.v1.GetKycByUserIDRequest"></a>

### GetKycByUserIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="kyc.management.v1.GetKycByUserIDResponse"></a>

### GetKycByUserIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [KycInfo](#kyc.management.v1.KycInfo) |  |  |






<a name="kyc.management.v1.GetKycImageRequest"></a>

### GetKycImageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| ImageS3Key | [string](#string) |  |  |






<a name="kyc.management.v1.GetKycImageResponse"></a>

### GetKycImageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






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
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |






<a name="kyc.management.v1.UpdateKycRequest"></a>

### UpdateKycRequest



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






<a name="kyc.management.v1.UpdateKycResponse"></a>

### UpdateKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [KycInfo](#kyc.management.v1.KycInfo) |  |  |






<a name="kyc.management.v1.UploadKycImageRequest"></a>

### UploadKycImageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| ImageType | [string](#string) |  |  |
| ImageBase64 | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="kyc.management.v1.UploadKycImageResponse"></a>

### UploadKycImageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






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
| CreateKycRecord | [CreateKycRequest](#kyc.management.v1.CreateKycRequest) | [CreateKycResponse](#kyc.management.v1.CreateKycResponse) |  |
| GetKycByUserID | [GetKycByUserIDRequest](#kyc.management.v1.GetKycByUserIDRequest) | [GetKycByUserIDResponse](#kyc.management.v1.GetKycByUserIDResponse) |  |
| GetKycByAppID | [GetKycByAppIDRequest](#kyc.management.v1.GetKycByAppIDRequest) | [GetKycByAppIDResponse](#kyc.management.v1.GetKycByAppIDResponse) |  |
| GetAllKyc | [GetAllKycRequest](#kyc.management.v1.GetAllKycRequest) | [GetAllKycResponse](#kyc.management.v1.GetAllKycResponse) |  |
| UpdateKyc | [UpdateKycRequest](#kyc.management.v1.UpdateKycRequest) | [UpdateKycResponse](#kyc.management.v1.UpdateKycResponse) |  |
| UploadKycImage | [UploadKycImageRequest](#kyc.management.v1.UploadKycImageRequest) | [UploadKycImageResponse](#kyc.management.v1.UploadKycImageResponse) |  |
| GetKycImage | [GetKycImageRequest](#kyc.management.v1.GetKycImageRequest) | [GetKycImageResponse](#kyc.management.v1.GetKycImageResponse) |  |

 



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

