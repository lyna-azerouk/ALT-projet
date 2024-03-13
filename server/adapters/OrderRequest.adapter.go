package adapters

import (
	"serveur/server/models"
	"strconv"
)

/** Convert OrderDetails to OrderDetailsRequest
 * Get order details
 */
func OrderDetailsToOrderRequestMapper(order models.OrderDetails) models.OrderDetailsRequest {
	var adaptedOrderRequest models.OrderDetailsRequest
	adaptedOrderRequest.Id = strconv.FormatUint(order.Id, 10)
	adaptedOrderRequest.ClientId = strconv.FormatUint(order.ClientId, 10)
	adaptedOrderRequest.RestaurantId = strconv.FormatUint(order.RestaurantId, 10)
	adaptedOrderRequest.Status = order.Status
	adaptedOrderRequest.Price = strconv.FormatFloat(order.Price, 'f', -1, 64)
	adaptedOrderRequest.Date = order.Date
	adaptedOrderRequest.OrderItems = orderItemsToOrderRequestMapper(order.OrderItems)
	return adaptedOrderRequest
}

func orderItemsToOrderRequestMapper(orderItem []models.OrderItem) []models.OrderItemRequest {
	var adaptedOrderItems []models.OrderItemRequest
	for _, item := range orderItem {
		adaptedOrderItem := models.OrderItemRequest{
			MenuId: strconv.FormatUint(item.MenuId, 10),
			Count:  strconv.Itoa(item.Count),
		}
		adaptedOrderItems = append(adaptedOrderItems, adaptedOrderItem)
	}
	return adaptedOrderItems
}

func orderItemsRequestToOrderMapper(orderItem []models.OrderItemRequest) []models.OrderItem {
	var adaptedOrderItems []models.OrderItem
	for _, item := range orderItem {
		menuId, _ := strconv.ParseUint(item.MenuId, 10, 64)
		count, _ := strconv.Atoi(item.Count)
		adaptedOrderItem := models.OrderItem{
			MenuId: menuId,
			Count:  count,
		}
		adaptedOrderItems = append(adaptedOrderItems, adaptedOrderItem)
	}
	return adaptedOrderItems
}

/** Convert OrderDetailsRequestV2 to OrderDetailsRequest
 * OrderDetailsRequestV2 field are string.
 * .*Id are uint64
 */
func OrderRequestToOrderDetailsMapper(orderRequest models.OrderDetailsRequest) models.OrderDetails {
	var adaptedOrderRequest models.OrderDetails
	adaptedOrderRequest.Id, _ = strconv.ParseUint(orderRequest.Id, 10, 64)
	adaptedOrderRequest.ClientId, _ = strconv.ParseUint(orderRequest.ClientId, 10, 64)
	adaptedOrderRequest.RestaurantId, _ = strconv.ParseUint(orderRequest.RestaurantId, 10, 64)
	adaptedOrderRequest.Status = orderRequest.Status
	adaptedOrderRequest.Price, _ = strconv.ParseFloat(orderRequest.Price, 64)
	adaptedOrderRequest.Date = orderRequest.Date
	adaptedOrderRequest.OrderItems = orderItemsRequestToOrderMapper(orderRequest.OrderItems)
	return adaptedOrderRequest
}
