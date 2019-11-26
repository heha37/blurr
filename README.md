# blurr

blurr is a declarative HTTP test framework. The requests and responses are represented in a declarative YAML-based form. The test looks like this:

```
envs:
  - TestEnv

testCases:
  - name: "Simple test B"
    method: GET

  - name: "Simple test A"
    method: GET
```

Tests can be run using `go test`.

# Purpose

Each YAML file represents an ordered list of HTTP requests along with the expected responses. A single file represents a process in the API being tested. For example:

* Create resources
* Retrieve resources
* Delete resources
* Retrieve resources to confirm they are gone
