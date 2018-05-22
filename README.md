<img src="https://storage.googleapis.com/ttp-static/ttp_transparent.png" width="25%">

# Tensor Transmission Protocol (TTP)
## A minimalist yet exstensible protocol for transmitting tensors.

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
The metadata fields are implemented with Protocol Buffer(v3) **Any** types. This allows data to be added to the request, response and tensor fields on the fly or for a given use case. An example of useful metadat would be passing a model version number, or specifying a dimension label. Note:

```
http://ttp.tensortask.org/metadata/<ANY_DEFINITION.proto>
```

