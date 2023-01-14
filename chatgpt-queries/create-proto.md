Query: write a proto spec of a service called shop that has a method called listBananas with size in the request and a second method called getBananaColor

Response:
```
syntax = "proto3";

service Shop {
    rpc ListBananas(ListBananasRequest) returns (ListBananasResponse);
    rpc GetBananaColor(GetBananaColorRequest) returns (GetBananaColorResponse);
}

message ListBananasRequest {
    int32 size = 1;
}

message ListBananasResponse {
    repeated Banana bananas = 1;
}

message GetBananaColorRequest {
    string banana_id = 1;
}

message GetBananaColorResponse {
    string color = 1;
}

message Banana {
    string id = 1;
    string color = 2;
}
```


This is a basic proto3 specification for a service called "Shop" that has two methods: "ListBananas" and "GetBananaColor".
The ListBananas method takes a request with a single field "size" of type int32, and returns a ListBananasResponse with a repeated field "bananas" of type Banana.
The GetBananaColor method takes a single field "banana_id" of type string and returns a GetBananaColorResponse with a single field "color" of type string.
The Banana message has two fields "id" and "color" which are both strings.