// auto generate code, dont edit it
package order

import (
	"fmt"
	"gateway/pkg/monitor"
	"gateway/proto/order"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

type createOrderHandler struct {
}

func NewCreateOrderHandler() *createOrderHandler {
	return &createOrderHandler{}
}

// @Summary permission:
// @Tags OrderService
// @Produce json
// @Param AccountEmail  body  string false "Missing comment! you should fill its for long-term"
// @Param ShopItems  body  []order.ShopItem false "Missing comment! you should fill its for long-term"
// @Param DeliveryAddress  body  string false "Missing comment! you should fill its for long-term"
// @Param body  body  order.CreateOrderReq true "Body example"
// @Success 200 {object} order.CreateOrderRes
// @Router /api/orders/ [post]
func (handler *createOrderHandler) Handle(ctx *gin.Context) (interface{}, error) {
	monitor.SetApmContext(apm.DetachedContext(ctx.Request.Context()))
	data := order.CreateOrderReq{}
	if err := ctx.BindJSON(&data); err != nil {
		return nil, err
	}
	fmt.Println("asdasdasd => ", data)
	return &data, nil
}
