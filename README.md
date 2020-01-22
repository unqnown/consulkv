### consulkv

consulkv is a command line utility that simplifies consul kv management.
Exported values are base64 decoded which makes them well readable and editable.
After changing kvs may be imported back to consul with preserving same format.

### installation

```sh
go install github.com/unqnown/consulkv
```

### usage

##### export kv

```
NAME:
   consulkv export - exports kv from consul server.

USAGE:
   consulkv export [command options] [arguments...]

CATEGORY:
   kv

OPTIONS:
   --file FILE, -f FILE       Path to kv FILE
   --address value, -a value  Consul server address
   --prefix value, -p value   KV prefix
   --verbose, -v              Prints progress information
```

##### import kv

```
NAME:
   consulkv import - imports kv to consul server.

USAGE:
   consulkv import [command options] [arguments...]

CATEGORY:
   kv

OPTIONS:
   --file FILE, -f FILE       Path to kv FILE
   --address value, -a value  Consul server address
   --verbose, -v              Prints progress information
```

