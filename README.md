# ens feed

## Clone
```sh
git clone https://github.com/hbdgr/ens_feed
```

## Build
```sh
cd end_feed && go build
```

## Config
The defult configuration sets local http server on localhost:8000 and use Goerli testnet ens contract: `0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e`

For proper funcioning `InfuraURL` must be provided, both `https` and `wss` works.

See `config.yml`

## Run
```
./ens_feed
```

### Example requests:

For Goerli testnet:
```sh
# test
curl "http://localhost:8000/test"

# resolve
curl "http://localhost:8000/resolve/eth"   # root
curl "http://localhost:8000/resolve/hbdgr1234.eth"

# reverse-resolve
curl "http://localhost:8000/reverse-resolve/0x23aa8FB48d26AB69361300A953D2dF9e7C9d19b6"
```
