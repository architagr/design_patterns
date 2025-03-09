package routers

import (
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer

type IRouter interface {
	StartApp()
}
