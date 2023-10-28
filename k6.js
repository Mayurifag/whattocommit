import { check } from "k6";
import http from "k6/http";

// define configuration
export const options = {
  // define thresholds
  thresholds: {
    http_req_failed: [{ threshold: "rate<0.01", abortOnFail: true }], // availability threshold for error rate
    http_req_duration: ["p(99)<1000"], // Latency threshold for percentile
  },
  // define scenarios
  scenarios: {
    breaking: {
      executor: "ramping-vus",
      stages: [
        { duration: "5s", target: 2000 },
        { duration: "5s", target: 4000 },
        { duration: "5s", target: 6000 },
        { duration: "5s", target: 8000 },
        { duration: "5s", target: 10000 },
        { duration: "5s", target: 12000 },
        { duration: "5s", target: 14000 },
        //....
      ],
    },
  },
};

export default function () {
  const res = http.get("http://localhost:9321")

  // check that response is 200
  check(res, {
    "response code was 200": (res) => res.status == 200,
  });
}
