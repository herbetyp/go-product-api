import http from 'k6/http';
import { sleep } from 'k6';
import { Counter } from 'k6/metrics';


export const requests = new Counter('http_reqs');

export const options = {
    // Key configurations for avg load test in this section
    stages: [
        { duration: '5m', target: 100 }, // traffic ramp-up from 1 to 100 users over 5 minutes.
        { duration: '30m', target: 100 }, // stay at 100 users for 30 minutes
        { duration: '5m', target: 0 }, // ramp-down to 0 users
    ],
};

export default () => {
    // Make a GET request to the target URL
    const url = 'http://go_product_api:3000/v1/health';
    http.get(url);

    // Sleep for 1 second to simulate real-world usage
    sleep(1);
};
