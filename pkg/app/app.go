package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-grogramming-tour-book/blog-service/pkg/errcode"
	"net"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// c.Set("X-Trace-ID", traceID)
func (r *Response) getTracID() string {
	if id, ok := r.Ctx.Get("X-Trace-ID"); ok {
		if traceID, isStr := id.(string); isStr {
			return traceID
		}
	}
	return ""
}

func (r *Response) ToResponse(data interface{}) {
	response := gin.H{
		"code":       errcode.Success.Code(),
		"msg":        errcode.Success.Msg(),
		"data":       nil,
		"ip":         GetIpAddr(),
		"x-trace-id": r.getTracID(),
	}

	if data != nil {
		response["data"] = data
	}
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	response := gin.H{
		"code":       errcode.Success.Code(),
		"msg":        errcode.Success.Msg(),
		"data":       nil,
		"ip":         GetIpAddr(),
		"x-trace-id": r.getTracID(),
	}
	response["data"] = map[string]interface{}{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	}

	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code":       err.Code(),
		"msg":        err.Msg(),
		"details":    nil,
		"ip":         GetIpAddr(),
		"x-trace-id": r.getTracID(),
	}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}

func GetIpAddr() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return ""
}
