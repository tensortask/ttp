<img src="https://storage.googleapis.com/ttp-static/ttp_transparent.png" width="25%">

# Tensor Transport Protocol (TTP)
## A minimalist yet exstensible protocol for transporting tensors.

The goal of TTP is to create a standard minimum that enables basic conversion between protocols while still providing use case dependent extensibility.

## Required Fields
#### Per Request/Response
* Target Computation
* Map of Alias' and Tensors 

#### Per Tensor
* Data Type
* Tensor Shape
* Tensor Contents

## Extensible Fields
#### Metadata
The metadata fields are implemented with Protocol Buffer(v3) **Any** types. This allows data to be added to the request, response and tensor fields on the fly or for a given use case. An example of useful metadata would be passing a model version number, or specifying a dimension label. 

**To add metadata type definitions create a pull request with your use case specific proto file added to the metadata directory. The added proto will automatically be added to a type resolver hosted at:**

```
http://ttp.tensortask.org/metadata/<PACKAGE>.<MESSAGE>
```

## JSON Support
Protocol Buffers automatically map to JSON and vice versa. Examples of standard and extended JSON tensors, requests, and responses are on the way!
