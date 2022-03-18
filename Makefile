gen:
	oapi-codegen -generate types,client -package client ./openapi/fuel_status_api.yml > client/fuel.go