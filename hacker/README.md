## address-acl-check
 
A database exists of network address ranges (CIDRs) and user IDs that have access to them.
 
Here is an example of some data in the database.
 
```
user_id     cidr
-------     ----
123         192.168.0.0/24
456         1.0.0.0/24
666         fe80:c001::/48
```
 
We would like to implement an ACL for every new connection that arrives at the
Cloudflare HTTP CDN, validating whether a certain user is allowed to connect to the given edge
network address. This ACL will return a response allowing or denying the connection.
 
A new internal API will be introduced, responsible for implementing this ACL. This API
will be called out to from nginx, the HTTP CDN service running on our edge.
 
The API interface looks like this...
 
```
$ curl -v 'http://localhost:9001/v1/edge-ip-acl?user_id=123&edge_ip=192.168.0.2'
...
HTTP/1.0 200 OK
...
```
 
Extend the given program to implement this API. Some initial parts of
the program have already written, but the code here serves as an example and feel
free to change anything as you see fit.
 
### Golang
 
```
$ cd go
$ go run main.go
...
 
$ curl -v 'http://localhost:9001/v1/edge-ip-acl?user_id=123&edge_ip=192.168.0.2'
```
 
### Python
 
```
$ cd python
$ ./pyenv.sh
$ source env/bin/activate
$ python main.py
...
 
$ curl -v 'http://localhost:9001/v1/edge-ip-acl?user_id=123&edge_ip=192.168.0.2'