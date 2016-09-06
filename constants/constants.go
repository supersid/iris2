package constants

//import "time"

const (
	DEVELOPMENT_ENV = "DEVELOPMENT_ENV"
	TEST_ENV 	= "TEST_ENV"

	//WORKER_HEARTBEAT_INTERVAL = 2 * time.Second
	WORKER_READY    = "WORKER_READY"
	CLIENT_REQUEST  = "CLIENT_REQUEST"
	WORKER_RESPONSE = "WORKER_RESPONSE"
	CLIENT_REQUEST_TO_WORKER = "CLIENT_REQUEST_TO_WORKER"
	WORKER_RESPONSE_RELAY = "WORKER_RESPONSE_RELAY"
	WORKER_RECV_HWM = 1000
)