package apimanagement

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiDiagnostic -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apis/api1/diagnostics/diagnostic1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiManagement -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ApiVersionSet -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/apiVersionSets/apiVersionSet1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=CustomDomain -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/customDomains/customdomain
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Diagnostic -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/diagnostics/diagnostic1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Logger -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/loggers/logger1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Policy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/policies/policy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Product -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/products/product1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ProductApi -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/products/product1/apis/api1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ProductGroup -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/products/product1/groups/group1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ProductPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/products/product1/policies/policy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Property -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/namedValues/namedvalue1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Subscription -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/subscriptions/subscription1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=User -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/users/user1
