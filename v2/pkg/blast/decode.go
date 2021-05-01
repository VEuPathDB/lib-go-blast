package blast

import (
	"errors"

	"github.com/francoispqt/gojay"
)

func DecodeConfig(tool Tool, b []byte) (BlastConfig, error) {
	switch tool {
	case ToolBlastN:
		out := new(BlastN)
		return out, gojay.Unmarshal(b, out)
	case ToolBlastP:
		out := new(BlastP)
		return out, gojay.Unmarshal(b, out)
	case ToolBlastX:
		out := new(BlastX)
		return out, gojay.Unmarshal(b, out)
	case ToolDeltaBlast:
		out := new(DeltaBlast)
		return out, gojay.Unmarshal(b, out)
	case ToolPSIBlast:
		out := new(PSIBlast)
		return out, gojay.Unmarshal(b, out)
	case ToolRPSBlast:
		out := new(RPSBlast)
		return out, gojay.Unmarshal(b, out)
	case ToolRPSTBlastN:
		out := new(RPSTBlastN)
		return out, gojay.Unmarshal(b, out)
	case ToolTBlastN:
		out := new(TBlastN)
		return out, gojay.Unmarshal(b, out)
	case ToolTBlastX:
		out := new(TBlastX)
		return out, gojay.Unmarshal(b, out)
	default:
		return nil, errors.New("Invalid blast tool.")
	}
}
