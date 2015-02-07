# Lookup

Demonstrate how to use the relaxed placeholder (notation `#paramName`).
This placeholder matches everything until the first `/`, including `.`

curl demo:
```
curl -i http://127.0.0.1:8080/lookup/google.com
curl -i http://127.0.0.1:8080/lookup/notadomain
```
