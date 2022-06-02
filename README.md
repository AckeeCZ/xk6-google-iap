# xk6-googleIap

**⚠️This is proof of concept**

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download `xk6`:
  ```bash
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```bash
  $ xk6 build --with github.com/AckeeDevOps/xk6-google-iap
  ```

## Example:

```javascript
import googleIap from 'k6/x/googleIap';
import { check } from 'k6';
import http from 'k6/http';

const client_id = __ENV.CLIENT_ID;
const sa_key = __ENV.SA_KEY;
// URL protected by Google IAP
const url = __ENV.TARGET_URL; 

export default function () {
    const token = googleIap.getToken(client_id, sa_key);

    const params = {
        headers: {
            'Proxy-Authorization': `Bearer ${token}`,
        },

    };

    const response = http.get(url, params);
    console.log(response.url)

    check(response, { "URL doesn't begin with accounts.google.com": response => !response.url.startsWith("https://accounts.google.com") })
}


```

Run with
```bash
./k6 -e SA_KEY="SERVICE_ACCOUNT_KEY" -e CLIENT_ID="OAUTH_CLIENT_ID" -e TARGET_URL="YOUR_WEBSITE_URL" script.js
```

