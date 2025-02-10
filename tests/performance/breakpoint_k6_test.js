import http from 'k6/http';
import { sleep } from 'k6';
import { Counter } from 'k6/metrics';


export const requests = new Counter('http_reqs');

export const options = {
    // Key configurations for breakpoint in this section
    executor: 'ramping-arrival-rate', //Assure load increase if the system slows
    stages: [
        { duration: '45m', target: 20000 }, // just slowly ramp-up to a HUGE load
    ],
};

export default () => {
    // Make a GET request to the target URL
    const url = 'http://go_product_api:3000/v1/health';
    http.get(url);

    // Sleep for 1 second to simulate real-world usage
    sleep(1);
};
