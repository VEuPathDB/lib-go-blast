package bval

type Defaulter interface {
	Flag() string
	IsDefault() bool
}

type F64Defaulter interface {
	Defaulter
	Get() float64
}

type I32Defaulter interface {
	Defaulter
	Get() int32
}

type U8Defaulter interface {
	Defaulter
	Get() uint8
}

type U16Defaulter interface {
	Defaulter
	Get() uint16
}

type U32Defaulter interface {
	Defaulter
	Get() uint32
}
