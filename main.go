package main

import "fmt"

func main() {
	fmt.Println(`[12/Nov/2022:04:04:45 +0000] 120.188.5.189 - - GET /api/requestid0/orders/R259747919/max_redeem 789 200 611 89 - "okhttp/4.2.1 dsdsdsd" 10.0.5.23:443 200 0.076 0.004 0.076 0.078 "120.188.5.189" loyalty_service "ID android 3.44 205d2524e9f8f7db" "Root=1-6062a35d-6fbb72d25d455a89583aac99" "92ad011f-0de3-4694-8aee-76b8deec7b80"`)
	fmt.Println(`[12/Nov/2022:04:05:45 +0000] 120.188.5.189 - - GET /api/normal/orders/R259747919/max_redeem 789 200 611 89 - "okhttp/4.2.1 dsdsdsd" 10.0.5.23:443 200 0.076 0.004 0.076 0.078 "120.188.5.189" loyalty_service "ID android 3.44 205d2524e9f8f7db" "Root=1-6062a35d-6fbb72d25d455a89583aac99"`)
	fmt.Println(`[12/Nov/2022:04:06:45 +0000] 120.188.5.189 - - GET /api/status/orders/R259747919/max_redeem 789 200 611 89 - "okhttp/4.2.1 dsdsdsd" 10.0.5.23:443 200`)
	fmt.Println(`[12/Nov/2022:04:04:45 +0000] 120.188.5.189 - - GET /api/requestid1/orders/R259747919/max_redeem 789 200 611 89 - "okhttp/4.2.1 dsdsdsd" 10.0.5.23:443 200 0.076 0.004 0.076 0.078 "120.188.5.189" loyalty_service "ID android 3.44 205d2524e9f8f7db" "Root=1-6062a35d-6fbb72d25d455a89583aac99" "92ad011f-0de3-4694-8aee-76b8deec7b80"`)
	fmt.Println(`[12/Nov/2022:04:04:45 +0000] 120.188.5.189 - - GET /api/trace/orders/R259747919/max_redeem 789 200 611 89 - "okhttp/4.2.1 dsdsdsd" 10.0.5.23:443 200 0.076 0.004 0.076 0.078 "120.188.5.189" loyalty_service "ID android 3.44 205d2524e9f8f7db" "Root=1-6062a35d-6fbb72d25d455a89583aac99" "92ad011f-0de3-4694-8aee-76b8deec7b80" "0000000000000000f4dbb3edd765f620" "43222c2d51a7abe3"`)
	fmt.Println(`[12/Nov/2022:04:04:45 +0000] 120.188.5.189 - - GET /api/hostname/orders/R259747919/max_redeem 789 200 611 89 - "okhttp/4.2.1 dsdsdsd" 10.0.5.23:443 200 0.076 0.004 0.076 0.078 "120.188.5.189" loyalty_service "ID android 3.44 205d2524e9f8f7db" "Root=1-6062a35d-6fbb72d25d455a89583aac99" "92ad011f-0de3-4694-8aee-76b8deec7b80" "0000000000000000f4dbb3edd765f620" "43222c2d51a7abe3" myhost`)
	fmt.Println(`[12/Nov/2022:01:57:36 +0000] 10.0.6.121 - - GET /api/sidecar/H50517227704/payment_eligibility 494 200 454 110 - "rest-client/2.1.0 (linux-gnu x86_64) ruby/2.6.6p146" 127.0.0.1:8080 200 0.028 0.000 0.028 0.025`)
}
