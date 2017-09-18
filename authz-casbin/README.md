# Casbin Authz

Demonstrate how to setup an [Casbin](https://github.com/casbin/casbin) authorization (AuthZ) middleware based on AuthBasicMiddleware.

curl demo:
```
curl -i -u alice:123 http://127.0.0.1:8080/dataset1/1
curl -i -u alice:123 http://127.0.0.1:8080/dataset2/1
```

## How to control the access

The authorization determines a request based on ``{subject, object, action}``, which means what ``subject`` can perform what ``action`` on what ``object``. In this plugin, the meanings are:

1. ``subject``: the logged-on user name
2. ``object``: the URL path for the web resource like "dataset1/item1"
3. ``action``: HTTP method like GET, POST, PUT, DELETE, or the high-level actions you defined like "read-file", "write-blog"


For how to write authorization policy and other details, please refer to [the Casbin's documentation](https://github.com/casbin/casbin).
