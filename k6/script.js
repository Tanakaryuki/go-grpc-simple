import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    vus: 10,
    duration: '30s',
};

export default function () {
    let url = 'http://127.0.0.1:59789';
    let res = http.get(`${url}/hello`);
    check(res, {
        'status is 200': (r) => r.status === 200,
        'body is "Hello World\\n"': (r) => r.body === 'Hello World\n',
    });

    res = http.get(`${url}/health`);
    check(res, {
        'status is 200': (r) => r.status === 200,
        'body is "hello\\n"': (r) => r.body === 'hello\n',
    });

    sleep(1);
}
