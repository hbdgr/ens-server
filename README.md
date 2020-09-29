# ens feed

### Clone
```sh
git clone https://github.com/hbdgr/ens_feed
```

### Build
```sh
cd end_feed && go build
```

### Config
<<<<<<< HEAD
The defult configuration sets local http server on localhost:8000 and uses Goerli testnet ens contract: `0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e`
=======
The defult configuration sets local http server on localhost:8000 and use Goerli testnet ens contract: `0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e`
>>>>>>> 488a9b3 (update readme)

For proper funcioning `InfuraURL` must be provided, both `https` and `wss` work.

See `config.yml`

### Run
```
./ens_feed
```

### Example requests:

For Goerli testnet:
```sh
# TEST
curl "http://localhost:8000/test"
# {"eth_addr":"0x23aa8FB48d26AB69361300A953D2dF9e7C9d19b6"}


# RESOLVE
curl "http://localhost:8000/resolve/eth"   # root
# {"eth_addr":"0x000000000000000000000000000000000000000E"}

curl "http://localhost:8000/resolve/hbdgr1234.eth"
# {"eth_addr":"0x23aa8FB48d26AB69361300A953D2dF9e7C9d19b6"}

curl "http://localhost:8000/resolve/app.hbdgr1234.eth"
# {"eth_addr":"0x07CAbC42bE55020ba2E971dA12a53BAE81972889"}


# REVERSE-RESOLVE
curl "http://localhost:8000/reverse-resolve/0x23aa8FB48d26AB69361300A953D2dF9e7C9d19b6"
# {"name":"hbdgr1234.eth"}

# '0x' can be omitted
curl "http://localhost:8000/reverse-resolve/07CAbC42bE55020ba2E971dA12a53BAE81972889"
# {"name":"app.hbdgr1234.eth"}


# SUBDOMAINS
<<<<<<< HEAD
curl "http://localhost:8000/subdomains/hbdgr1234.eth"
=======
curl "http://localhost:8000/subdomains/hbdgr1234.eth" | json_reformat
>>>>>>> 488a9b3 (update readme)
# [
#     "app.hbdgr1234.eth",
#     "chain.hbdgr1234.eth",
#     "test5.hbdgr1234.eth"
# ]


# SUBDOMAINS/INFO
<<<<<<< HEAD
curl "http://localhost:8000/subdomains/hbdgr1234.eth/info"
=======
curl "http://localhost:8000/subdomains/hbdgr1234.eth/info" | json_reformat
>>>>>>> 488a9b3 (update readme)
# {
#     "parent_name": "hbdgr1234.eth",
#     "subnames": [
#         {
#             "label_hash": "0xbd4c2bf3814af0934e4deccab96a872647bbadbc4b89baea4c2f403bf476e37f",
#             "node": "0x776413d36c418eb77f3de22b9274a8ceb18996de6169705a7958202737f0ec93",
#             "owner": "0x23aa8FB48d26AB69361300A953D2dF9e7C9d19b6",
#             "resolver": "0xDaaF96c344f63131acadD0Ea35170E7892d3dfBA",
#             "name": "",
#             "error_message": "not a resolver"                       # reverse-resolver not set
#         },
#         {
#             "label_hash": "0xd6f028ca0e8edb4a8c9757ca4fdccab25fa1e0317da1188108f7d2dee14902fb",
#             "node": "0x615e76671f99e7d505ea1d79542d5e0084d1b679b778afde3320901c6c493b20",
#             "owner": "0x23aa8FB48d26AB69361300A953D2dF9e7C9d19b6",
#             "resolver": "0x4B1488B7a6B320d2D721406204aBc3eeAa9AD329",
#             "name": "app.hbdgr1234.eth"
#         },
#         {
#             "label_hash": "0x9e199fafc1079dfb2b375cdac741cefb6c51d5f471f8afffa517442b6160463c",
#             "node": "0xfc1c1089018043890bd4f20814e73004adfd9c9705e955ef1461125638530bf0",
#             "owner": "0x23aa8FB48d26AB69361300A953D2dF9e7C9d19b6",
#             "resolver": "0xDaaF96c344f63131acadD0Ea35170E7892d3dfBA",
#             "name": "test5.hbdgr1234.eth"
#         }
#     ]
# }

# works for nested names
<<<<<<< HEAD
curl "http://localhost:8000/subdomains/test5.hbdgr1234.eth/info"
=======
curl "http://localhost:8000/subdomains/test5.hbdgr1234.eth/info" | json_reformat
>>>>>>> 488a9b3 (update readme)
# {
#     "parent_name": "test5.hbdgr1234.eth",
#     "subnames": [
#         {
#             "label_hash": "0x23dc111d7c3ad1df9806ce1e8eb4f55f57dba117339c545e7593d1f6c3b02662",
#             "node": "0x2057443eb407f709682300e0bf27ca4a178598dfbf7d654632f7a61ed920821b",
#             "owner": "0x23aa8FB48d26AB69361300A953D2dF9e7C9d19b6",
#             "resolver": "0xDaaF96c344f63131acadD0Ea35170E7892d3dfBA",
#             "name": "one.test5.hbdgr1234.eth"
#         },
#         {
#             "label_hash": "0x332c39dcd398ea34a48b871898d589f55fc4c7bce00562fb670c972e7e1b0720",
#             "node": "0x14b51feffe9fb72a18b25ce8fe04d6eaedd2d4e6b053c191eb794f89d0510115",
#             "owner": "0x23aa8FB48d26AB69361300A953D2dF9e7C9d19b6",
#             "resolver": "0xDaaF96c344f63131acadD0Ea35170E7892d3dfBA",
#             "name": "two.test5.hbdgr1234.eth"
#         }
#     ]
# }
```


## Notes

<<<<<<< HEAD
* Not every subdomain can be reverse-resolved. In fact by default (in `app.ens.domains`) they can't. Reverse-resolver must be set for having this ability, e.g. with transaction to `ReverseRegistrar` function:  
=======
* Not every subdomain can be reverse-resolved. In fact by default (in `app.ens.domains`) they can't. To have this ability, reverse-resolver must be set. It can be acomplished e.g. with transaction to `ReverseRegistrar` function:  
>>>>>>> 488a9b3 (update readme)
doc: [https://docs.ens.domains/contract-api-reference/reverseregistrar#set-name](https://docs.ens.domains/contract-api-reference/reverseregistrar#set-name)
```
Function: setName(string name)
MethodID: 0xc47f0027
```

* This solution does not check expiration date of the names
