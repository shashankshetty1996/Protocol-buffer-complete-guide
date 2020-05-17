# Data types

## oneof

Only one of the fields can have valid. Here in the example given below only `my_string` or `my_bool` can have value.

_NOTE:_

- oneof can be repeated
- If we write both the field then only `my_bool` value will be set

```proto
message MyMessage {
  int32 id = 1;
  oneof example_of {
    string my_string = 2;
    bool my_bool = 3;
  }
}
```

## maps

Maps can be used to map scalars (expect float/double) to values of any type. Maps can't be repeated or there will be any ordering since it is a map

```proto
message MyMessage {
  int32 id = 1;
  map<string, Result> result = 2;
}
```

## timestamp

This is one of the standard types. It stores seconds and nano seconds in UTC.

```proto
syntax = "proto3"

import "google/protobuf/timestamp.proto";

message MyMessage {
  google.protobuf.Timestamp my_field = 1;
}
```

## Duration

Represents timestamp between two timestamps. This also stores in seconds and nano seconds

```proto
syntax = "proto3"

import "google/protobuf/timestamp.proto";
import "google/protobuf/duration.proto";

message MyMessage {
  google.protobuf.Timestamp my_date = 1;
  google.protobuf.Duration validate = 2;
}
```

# Protocol Buffer Options

This are the options which we can add in proto file which will help protoc complier to complier and provide different option/packages.

```proto
option csharp_namespace = "Google.Protobuf.WellKnownTypes";
option cc_enable_arenas = true;
option go_package = "github.com/golang/protobuf/ptypes/duration";
option java_package = "com.google.protobuf";
option java_outer_classname = "DurationProto";
option java_multiple_files = true;
option objc_class_prefix = "GPB";
```

# Naming conventions

- Message name _CamelCasing_
- field name _UnderscoreCasing_
- enum for naming we use _CamelCasing_ and value we use _UpperCase_

# Services in Protocol buffers

Protocol Buffers can define services on top of messages. A service is a set of endpoints your application can be accessible from.

```proto
service SearchService {
  rpc Search (SearchRequest) return (SearchResponse);
}
```

Service need to interpreted by a framework to generate associated code. The main framework is gRPC. Example proto file is also attached

# Protocol Buffer Encoding

- This has same serialization and deserialization for all the language.
- Serialization means transforming object into bytes.
- Deserialization means getting object from bytes.
