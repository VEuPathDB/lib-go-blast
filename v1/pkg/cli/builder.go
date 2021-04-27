package cli

func NewBuilder() *Builder {
	tmp := make(Builder, 0, 32)
	return &tmp
}

type Builder []string

func (b *Builder) Append(val Appender) *Builder {
	if val != nil && !val.IsDefault() {
		b.AppendFlag(val.FlagString())
	}
	return b
}

func (b *Builder) AppendFlag(flag string) *Builder {
	*b = append(*b, flag)
	return b
}

func (b *Builder) AppendArgFlag(flag, arg string) *Builder {
	return b.AppendFlag(flag + "=" + arg)
}

type Appender interface {
	Flag() string
	IsDefault() bool
	FlagString() string
}
