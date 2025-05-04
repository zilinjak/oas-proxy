import http from 'k6/http';
import {check} from 'k6';

export const options = {
    discardResponseBodies: false,
    scenarios: {
        contacts: {
            executor: 'ramping-arrival-rate',
            timeUnit: '1m',
            preAllocatedVUs: 50,
            maxVUs: 100000,
            stages: [
                { target: 1000, duration: '0m' },
                { target: 5000, duration: '2m' },
                { target: 100000, duration: '5m' },
            ],
        },
    },
};
export default function () {
    const res = http.get('http://localhost:8080/contacts.php');
    check(res, {
        'status is 200': (r) => r.status === 200,
    });
}
