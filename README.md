### consulkv

consulkv is a command line utility that simplifies consul kv management.

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

