import http from 'k6/http';
import { sleep } from 'k6';
import { Counter } from 'k6/metrics';


export const requests = new Counter('http_reqs');

export const options = {
    // Key configurations for spike in this section
    stages: [
        { duration: '3m', target: 3000 }, // fast ramp-up to a high point
        // No plateau
        { duration: '1m', target: 0 }, // quick ramp-down to 0 users
    ],
};

export default () => {
    // Make a GET request to the target URL
    const url = 'http://go_product_api:3000/v1/health';
    http.get(url);

    // Sleep for 1 second to simulate real-world usage
    sleep(1);
};
