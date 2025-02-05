import http from 'k6/http';
import { check } from 'k6';

export let options = {
    scenarios: {
        constant_request_rate: {
            executor: 'constant-arrival-rate',
            rate: 100,
            timeUnit: '1s',
            duration: '30s',
            preAllocatedVUs: 10,
        },
    },
};

export default function () {
    let url = 'http://127.0.0.1:58256';

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
}
