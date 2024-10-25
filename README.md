## Showcase: Buf + Dependencies

Showcase how to define local Protobuf that depends on Buf Registry package.

Create a [proto/acme.proto](./proto/acme.proto) file with the content:

```proto
syntax = "proto3";

package acme;

import "sf/ethereum/type/v2/type.proto";

option go_package = "github.com/maoueh/showcase-buf-dependencies;pbacme";

message Test {
    sf.ethereum.type.v2.BigInt big_int = 1;
}
```

Create a [buf.yaml](./buf.yaml) with the following content:

```yaml
version: v2
modules:
  - path: proto
    name: buf.build/acme/test
deps:
  - buf.build/streamingfast/firehose-ethereum
```

Create a [buf.gen.yaml](./buf.gen.yaml) with the following content:

```yaml
version: v2
plugins:
  - remote: buf.build/protocolbuffers/go:v1.31.0
    out: pb
    opt: paths=source_relative
```

Run the `buf dep` so dependencies are fetched and locally available:

```bash
buf dep update
```

> [!NOTE]
> This creates a `buf.lock` file, without it present, if you do `buf generate` you get errors like `proto/acme.proto:6:8:import "sf/ethereum/type/v2/type.proto": file does not exist`

Then run Buf generation:

```bash
buf generate
```
