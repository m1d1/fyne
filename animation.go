package fyne

import (
	"math"
	"time"
)

// AnimationCurve represents an animation algorithm for calculating the progress through a timeline.
// Custom animations can be provided by implementing the "func(float32) float32" definition.
// The input parameter will start at 0.0 when an animation starts and travel up to 1.0 at which point it will end.
// A linear animation would return the same output value as is passed in.
type AnimationCurve func(float32) float32

// AnimationRepeatForever is an AnimationCount value that indicates it should not stop looping.
//
// Since: 2.0
const AnimationRepeatForever = -1

var (
	// AnimationEaseInOut is the default easing, it starts slowly, accelerates to the middle and slows to the end.
	//
	// Since: 2.0
	AnimationEaseInOut = animationEaseInOut
	// AnimationEaseIn starts slowly and accelerates to the end.
	//
	// Since: 2.0
	AnimationEaseIn = animationEaseIn
	// AnimationEaseOut starts at speed and slows to the end.
	//
	// Since: 2.0
	AnimationEaseOut = animationEaseOut
	// AnimationLinear is a linear mapping for animations that progress uniformly through their duration.
	//
	// Since: 2.0
	AnimationLinear = animationLinear

	// These animation types are adapted from Robert Penner's easing functions.
	// http://robertpenner.com/easing/
	AnimationBackEaseInOut = animationBackEaseInOut
	AnimationBackEaseIn    = animationBackEaseIn
	AnimationBackEaseOut   = animationBackEaseOut

	AnimationBounceEaseInOut = animationBounceEaseInOut
	AnimationBounceEaseIn    = animationBounceEaseIn
	AnimationBounceEaseOut   = animationBounceEaseOut

	AnimationCircularEaseInOut = animationCircularEaseInOut
	AnimationCircularEaseIn    = animationCircularEaseIn
	AnimationCircularEaseOut   = animationCircularEaseOut

	AnimationCubicEaseInOut = animationCubicEaseInOut
	AnimationCubicEaseIn    = animationCubicEaseIn
	AnimationCubicEaseOut   = animationCubicEaseOut

	AnimationElasticEaseInOut = animationElasticEaseInOut
	AnimationElasticEaseIn    = animationElasticEaseIn
	AnimationElasticEaseOut   = animationElasticEaseOut

	AnimationExponentialEaseInOut = animationExponentialEaseInOut
	AnimationExponentialEaseIn    = animationExponentialEaseIn
	AnimationExponentialEaseOut   = animationExponentialEaseOut

	AnimationQuadraticEaseInOut = animationQuadaticEaseInOut
	AnimationQuadraticEaseIn    = animationQuadaticEaseIn
	AnimationQuadraticEaseOut   = animationQuadaticEaseOut

	AnimationQuarticEaseInOut = animationQuarticEaseInOut
	AnimationQuarticEaseIn    = animationQuarticEaseIn
	AnimationQuarticEaseOut   = animationQuarticEaseOut

	AnimationQuinticEaseInOut = animationQuinticEaseInOut
	AnimationQuinticEaseIn    = animationQuinticEaseIn
	AnimationQuinticEaseOut   = animationQuinticEaseOut

	AnimationSineEaseInOut = animationSineEaseInOut
	AnimationSineEaseIn    = animationSineEaseIn
	AnimationSineEaseOut   = animationSineEaseOut
)

// Animation represents an animated element within a Fyne canvas.
// These animations may control individual objects or entire scenes.
//
// Since: 2.0
type Animation struct {
	AutoReverse bool
	Curve       AnimationCurve
	Duration    time.Duration
	RepeatCount int
	Tick        func(float32)
}

// NewAnimation creates a very basic animation where the callback function will be called for every
// rendered frame between time.Now() and the specified duration. The callback values start at 0.0 and
// will be 1.0 when the animation completes.
//
// Since: 2.0
func NewAnimation(d time.Duration, fn func(float32)) *Animation {
	return &Animation{Duration: d, Tick: fn}
}

// Start registers the animation with the application run-loop and starts its execution.
func (a *Animation) Start() {
	CurrentApp().Driver().StartAnimation(a)
}

// Stop will end this animation and remove it from the run-loop.
func (a *Animation) Stop() {
	CurrentApp().Driver().StopAnimation(a)
}

func animationEaseIn(val float32) float32 {
	return val * val
}

func animationEaseInOut(val float32) float32 {
	if val <= 0.5 {
		return val * val * 2
	}
	return -1 + (4-val*2)*val
}

func animationEaseOut(val float32) float32 {
	return val * (2 - val)
}

func animationLinear(val float32) float32 {
	return val
}

// Back
func animationBackEaseIn(val float32) float32 {
	var s float32 = 1.70158
	return val * val * ((s+1)*val - s)
}

func animationBackEaseOut(val float32) float32 {
	var s float32 = 1.70158
	val -= 1
	return val*val*((s+1)*val+s) + 1
}

func animationBackEaseInOut(val float32) float32 {
	var s float32 = 1.70158
	var r float32 = 1.525
	//return ((ratio *= 2) < 1) ? 0.5*(ratio*ratio*((s*1.525+1)*ratio-s*1.525)) : 0.5*((ratio -= 2)*ratio*((s*1.525+1)*ratio+s*1.525)+2);
	val *= 2
	if val < 1 {
		return 0.5 * (val * val * ((s*r+1)*val - s*r))
	}
	val -= 2
	return 0.5 * ((val)*val*((s*r+1)*val+s*r) + 2)
}

// Bounce
func animationBounceEaseIn(val float32) float32 {
	return 1 - animationBounceEaseOut(1-val)
}

func animationBounceEaseOut(val float32) float32 {
	var s float32 = 7.5625
	var r float32 = 2.75
	if val < 1/r {
		return s * val * val
	} else if val < 2/r {
		val -= 1.5 / r
		return s*val*val + 0.75
	} else if val < 2.5/r {
		val -= 2.25 / r
		return s*val*val + 0.9375
	}
	val -= 2.625 / r
	return s*val*val + 0.984375
}

func animationBounceEaseInOut(val float32) float32 {
	val *= 2
	if val < 1 {
		return 0.5 * animationBounceEaseIn(val)
	}
	return 0.5*animationBounceEaseOut(val-1) + 0.5
}

// Circular
func animationCircularEaseIn(val float32) float32 {
	val2 := float64(val)
	return float32(-(math.Sqrt(1-val2*val2) - 1))
}

func animationCircularEaseOut(val float32) float32 {
	val2 := float64(val)
	return float32(math.Sqrt(1 - (val2-1)*(val2-1)))
}

func animationCircularEaseInOut(val float32) float32 {
	val *= 2
	val2 := float64(val)
	if val2 < 1 {
		return float32(-0.5 * (math.Sqrt(1-val2*val2) - 1))
	}
	val2 -= 2
	return float32(0.5 * (math.Sqrt(1-val2*val2) + 1))
}

// Cubic
func animationCubicEaseIn(val float32) float32 {
	return val * val * val
}
func animationCubicEaseOut(val float32) float32 {
	val -= 1
	return val*val*val + 1
}
func animationCubicEaseInOut(val float32) float32 {
	if val < 0.5 {
		return 4 * val * val * val
	}
	val -= 1
	return 4*val*val*val + 1
}

// Elastic
func animationElasticEaseIn(val float32) float32 {
	var a float64 = 1
	var p float64 = 0.3
	var s float64 = p / 4

	if val == 0 || val == 1 {
		return val
	}
	val -= 1
	val2 := float64(val)
	return float32(-(a * math.Pow(2, 10*val2) * math.Sin((val2-s)*(2*math.Pi)/p)))

}
func animationElasticEaseOut(val float32) float32 {
	var a float64 = 1
	var p float64 = 0.3
	var s float64 = p / 4

	if val == 0 || val == 1 {
		return val
	}
	val2 := float64(val)
	return float32(a*math.Pow(2, -10*val2)*math.Sin((val2-s)*(2*math.Pi)/p) + 1)
}
func animationElasticEaseInOut(val float32) float32 {

	var a float64 = 1
	var p float64 = 0.3
	var s float64 = p / 4

	if val == 0 || val == 1 {
		return val
	}
	val = val*2 - 1
	val2 := float64(val)
	if val < 0 {
		return float32(-0.5 * (a * math.Pow(2, 10*val2) * math.Sin((val2-s*1.5)*(2*math.Pi)/(p*1.5))))
	}
	return float32(0.5*a*math.Pow(2, -10*val2)*math.Sin((val2-s*1.5)*(2*math.Pi)/(p*1.5)) + 1)
}

// Exponential
func animationExponentialEaseIn(val float32) float32 {
	if val == 0 {
		return 0
	}
	val2 := float64(val)
	return float32(math.Pow(2, 10*(val2-1)))
}

func animationExponentialEaseOut(val float32) float32 {
	if val == 1 {
		return 1
	}
	val2 := float64(val)
	return float32(1 - math.Pow(2, -10*val2))
}

func animationExponentialEaseInOut(val float32) float32 {
	if val == 0 || val == 1 {
		return val
	}
	val = val*2 - 1
	val2 := float64(val)
	if 0 > val {
		return float32(0.5 * math.Pow(2, 10*val2))
	}
	return float32(1 - 0.5*math.Pow(2, -10*val2))
}

// // Quadratic
func animationQuadaticEaseIn(val float32) float32 {
	return val * val
}
func animationQuadaticEaseOut(val float32) float32 {
	return -val * (val - 2)
}
func animationQuadaticEaseInOut(val float32) float32 {
	if val < 0.5 {
		return 2 * val * val
	}
	return -2*val*(val-2) - 1
}

// // Quartic
func animationQuarticEaseIn(val float32) float32 {
	return val * val * val * val
}
func animationQuarticEaseOut(val float32) float32 {
	val -= 1
	return 1 - val*val*val*val
}
func animationQuarticEaseInOut(val float32) float32 {
	if val < 0.5 {
		return 8 * val * val * val * val
	}
	val -= 1
	return -8*val*val*val*val + 1

}

// // Quintic
func animationQuinticEaseIn(val float32) float32 {
	return val * val * val * val * val
}
func animationQuinticEaseOut(val float32) float32 {
	val -= 1
	return 1 + val*val*val*val*val
}
func animationQuinticEaseInOut(val float32) float32 {
	if val < 0.5 {
		return 16 * val * val * val * val * val
	}
	val -= 1
	return 16*val*val*val*val*val + 1
}

// Sine
func animationSineEaseIn(val float32) float32 {
	val2 := float64(val)
	return float32(1 - math.Cos(val2*(math.Pi/2)))
}
func animationSineEaseOut(val float32) float32 {
	val2 := float64(val)
	return float32(math.Sin(val2 * (math.Pi / 2)))
}
func animationSineEaseInOut(val float32) float32 {
	val2 := float64(val)
	return float32(-0.5 * (math.Cos(val2*math.Pi) - 1))
}
