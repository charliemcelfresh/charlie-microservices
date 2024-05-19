gen:
    # Generate twirp code
	protoc --go_out=. --twirp_out=. rpc/charlie-microservices/transaction_service.proto

