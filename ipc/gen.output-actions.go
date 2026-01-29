package ipc

import "github.com/frizinak/goniri/ipc/types"

func OutputActionOff(output string) Request {
	return newOutputAction(output, "Off")
}

func OutputActionOn(output string) Request {
	return newOutputAction(output, "On")
}

func OutputActionMode(
	output string,
	mode types.ModeToSet,
) Request {
	return newOutputAction(
		output,
		map[string]map[string]types.ModeToSet{
			"Mode": {
				"mode": mode,
			},
		},
	)
}

func OutputActionCustomMode(
	output string,
	mode types.ConfiguredMode,
) Request {
	return newOutputAction(
		output,
		map[string]map[string]types.ConfiguredMode{
			"CustomMode": {
				"mode": mode,
			},
		},
	)
}

func OutputActionModeline(
	output string,
	clock float64,
	hdisplay uint16,
	hsyncStart uint16,
	hsyncEnd uint16,
	htotal uint16,
	vdisplay uint16,
	vsyncStart uint16,
	vsyncEnd uint16,
	vtotal uint16,
	hsyncPolarity types.HSyncPolarity,
	vsyncPolarity types.VSyncPolarity,
) Request {
	return newOutputAction(
		output,
		map[string]map[string]any{
			"Modeline": {
				"clock": clock,
				"hdisplay": hdisplay,
				"hsync_start": hsyncStart,
				"hsync_end": hsyncEnd,
				"htotal": htotal,
				"vdisplay": vdisplay,
				"vsync_start": vsyncStart,
				"vsync_end": vsyncEnd,
				"vtotal": vtotal,
				"hsync_polarity": hsyncPolarity,
				"vsync_polarity": vsyncPolarity,
			},
		},
	)
}

func OutputActionScale(
	output string,
	scale types.ScaleToSet,
) Request {
	return newOutputAction(
		output,
		map[string]map[string]types.ScaleToSet{
			"Scale": {
				"scale": scale,
			},
		},
	)
}

func OutputActionTransform(
	output string,
	transform types.Transform,
) Request {
	return newOutputAction(
		output,
		map[string]map[string]types.Transform{
			"Transform": {
				"transform": transform,
			},
		},
	)
}

func OutputActionPosition(
	output string,
	position types.PositionToSet,
) Request {
	return newOutputAction(
		output,
		map[string]map[string]types.PositionToSet{
			"Position": {
				"position": position,
			},
		},
	)
}

func OutputActionVRR(
	output string,
	vrr types.VRRToSet,
) Request {
	return newOutputAction(
		output,
		map[string]map[string]types.VRRToSet{
			"Vrr": {
				"vrr": vrr,
			},
		},
	)
}
