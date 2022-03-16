gen:
	oapi-codegen -package spec ./openapi/fuel_status_api.yml > client/fuel.go
	oapi-codegen -package spec ./openapi/pay_as_you_drive_insurance_api.yml > client/pay_as_you_drive_insurance_api.go
	oapi-codegen -package spec ./openapi/vehicle_lock_status_api.yml > client/vehicle_lock_status_api.go
	oapi-codegen -package spec ./openapi/vehicle_status_api.yml > client/vehicle_status_api.go