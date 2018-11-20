<img src="https://storage.googleapis.com/ttp-static/ttp_transparent.png" width="25%">

# Tensor Transport Protocol (TTP)
## A minimalist yet exstensible protocol for transporting tensors.

The goal of TTP is to create a standard minimum that enables basic conversion between protocols.

### Per Request/Response
* Target (where the tensor should be sent and how it should be processed) e.g. `/ip4/127.0.0.1/udp/1234:predict`
* Map of Alias' and Tensors (these are the inputs/outputs)

### Per Tensor
* Data Type (specifies the type of data for the given tensor)
* Tensor Dimensions (the size/dimensionality of the given tensor)
* Tensor Contents (raw bytes, baby)

## JSON Support
Protocol Buffers automatically map to JSON and vice versa. Examples of standard and extended JSON tensors, requests, and responses are on the way!
