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

