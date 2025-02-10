import http from 'k6/http';
import { sleep } from 'k6';
import { Counter } from 'k6/metrics';


export const requests = new Counter('http_reqs');

export const options = {
    // Key configurations for Stress in this section
    stages: [
        { duration: '10m', target: 200 }, // traffic ramp-up from 1 to a higher 200 users over 10 minutes.
        { duration: '25m', target: 400 }, // stay at higher 400 users for 25 minutes
        { duration: '5m', target: 0 }, // ramp-down to 0 users
    ],
};

export default function () {
    // Make a GET request to the target URL
    const url = 'http://go_product_api:3000/v1/health';
    http.get(url);

    // Sleep for 1 second to simulate real-world usage
    sleep(1);
};
