package tutorials

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var a0 *fyne.Animation

func makeAnimationScreen(_ fyne.Window) fyne.CanvasObject {
	curves := makeAnimationCurves()
	curves.Move(fyne.NewPos(0, 140+theme.Padding()))
	return fyne.NewContainerWithoutLayout(makeAnimationCanvas(), curves)
}

func makeAnimationCanvas() fyne.CanvasObject {
	rect := canvas.NewRectangle(color.Black)
	rect.Resize(fyne.NewSize(410, 140))

	a := canvas.NewColorRGBAAnimation(theme.PrimaryColorNamed(theme.ColorBlue), theme.PrimaryColorNamed(theme.ColorGreen),
		time.Second*3, func(c color.Color) {
			rect.FillColor = c
			canvas.Refresh(rect)
		})
	a.RepeatCount = fyne.AnimationRepeatForever
	a.AutoReverse = true
	a.Start()

	i := widget.NewIcon(theme.CheckButtonCheckedIcon())
	a0 = canvas.NewPositionAnimation(fyne.NewPos(0, 0), fyne.NewPos(350, 80), time.Second*3, func(p fyne.Position) {
		i.Move(p)

		width := 10 + (p.X / 7)
		i.Resize(fyne.NewSize(width, width))
	})
	a0.RepeatCount = fyne.AnimationRepeatForever
	a0.AutoReverse = true
	a0.Curve = fyne.AnimationLinear
	a0.Start()

	running := true
	var toggle *widget.Button
	toggle = widget.NewButton("Stop", func() {
		if running {
			a.Stop()
			a0.Stop()
			toggle.SetText("Start")
		} else {
			a.Start()
			a0.Start()
			toggle.SetText("Stop")
		}
		running = !running
	})
	toggle.Resize(toggle.MinSize())
	toggle.Move(fyne.NewPos(152, 54))
	return fyne.NewContainerWithoutLayout(rect, i, toggle)
}

func makeAnimationCurves() fyne.CanvasObject {
	var (
		yOff1 float32 = 50
		yOff2 float32 = 80 + theme.Padding()
		yOff3 float32 = 110 + theme.Padding()*2
	)

	label1, box1, a1 := makeAnimationCurveItem("EaseInOut", fyne.AnimationLinear, yOff1)
	label2, box2, a2 := makeAnimationCurveItem("EaseIn", fyne.AnimationLinear, yOff2)
	label3, box3, a3 := makeAnimationCurveItem("EaseOut", fyne.AnimationLinear, yOff3)

	combo := widget.NewSelect([]string{"Linear", "Ease", "Back", "Bounce", "Circular", "Cubic", "Elastic", "Exponential", "Quadratic", "Quartic", "Quintic", "Sine"}, func(value string) {
		if a0 == nil {
			// catch early combo.SetSelectedIndex(0)
			return
		}

		switch value {
		case "Linear":
			a0.Curve = fyne.AnimationLinear
			a1.Curve = fyne.AnimationLinear
			a2.Curve = fyne.AnimationLinear
			a3.Curve = fyne.AnimationLinear
		case "Ease":
			a0.Curve = fyne.AnimationEaseOut
			a1.Curve = fyne.AnimationEaseInOut
			a2.Curve = fyne.AnimationEaseIn
			a3.Curve = fyne.AnimationEaseOut
		case "Back":
			a0.Curve = fyne.AnimationBackEaseOut
			a1.Curve = fyne.AnimationBackEaseInOut
			a2.Curve = fyne.AnimationBackEaseIn
			a3.Curve = fyne.AnimationBackEaseOut
		case "Bounce":
			a0.Curve = fyne.AnimationBounceEaseOut
			a1.Curve = fyne.AnimationBounceEaseInOut
			a2.Curve = fyne.AnimationBounceEaseIn
			a3.Curve = fyne.AnimationBounceEaseOut
		case "Circular":
			a0.Curve = fyne.AnimationCircularEaseOut
			a1.Curve = fyne.AnimationCircularEaseInOut
			a2.Curve = fyne.AnimationCircularEaseIn
			a3.Curve = fyne.AnimationCircularEaseOut
		case "Cubic":
			a0.Curve = fyne.AnimationCubicEaseOut
			a1.Curve = fyne.AnimationCubicEaseInOut
			a2.Curve = fyne.AnimationCubicEaseIn
			a3.Curve = fyne.AnimationCubicEaseOut
		case "Elastic":
			a0.Curve = fyne.AnimationElasticEaseOut
			a1.Curve = fyne.AnimationElasticEaseInOut
			a2.Curve = fyne.AnimationElasticEaseIn
			a3.Curve = fyne.AnimationElasticEaseOut
		case "Exponential":
			a0.Curve = fyne.AnimationExponentialEaseOut
			a1.Curve = fyne.AnimationExponentialEaseInOut
			a2.Curve = fyne.AnimationExponentialEaseIn
			a3.Curve = fyne.AnimationExponentialEaseOut
		case "Quadratic":
			a0.Curve = fyne.AnimationQuadraticEaseOut
			a1.Curve = fyne.AnimationQuadraticEaseInOut
			a2.Curve = fyne.AnimationQuadraticEaseIn
			a3.Curve = fyne.AnimationQuadraticEaseOut
		case "Quartic":
			a0.Curve = fyne.AnimationQuarticEaseOut
			a1.Curve = fyne.AnimationQuarticEaseInOut
			a2.Curve = fyne.AnimationQuarticEaseIn
			a3.Curve = fyne.AnimationQuarticEaseOut
		case "Quintic":
			a0.Curve = fyne.AnimationQuinticEaseOut
			a1.Curve = fyne.AnimationQuinticEaseInOut
			a2.Curve = fyne.AnimationQuinticEaseIn
			a3.Curve = fyne.AnimationQuinticEaseOut
		case "Sine":
			a0.Curve = fyne.AnimationQuinticEaseOut
			a1.Curve = fyne.AnimationSineEaseInOut
			a2.Curve = fyne.AnimationSineEaseIn
			a3.Curve = fyne.AnimationSineEaseOut

		}
	})

	combo.Resize(combo.MinSize())
	combo.Move(fyne.NewPos(0, 0+theme.Padding()))
	combo.SetSelectedIndex(0)
	start := widget.NewButton("Compare", func() {
		a1.Start()
		a2.Start()
		a3.Start()
	})
	start.Resize(start.MinSize())
	start.Move(fyne.NewPos(0, 150+theme.Padding()))
	return fyne.NewContainerWithoutLayout(label1, label2, label3, box1, box2, box3, start, combo)
}

func makeAnimationCurveItem(label string, curve fyne.AnimationCurve, yOff float32) (
	text *widget.Label, box fyne.CanvasObject, anim *fyne.Animation) {
	text = widget.NewLabel(label)
	text.Alignment = fyne.TextAlignCenter
	text.Resize(fyne.NewSize(380, 30))
	text.Move(fyne.NewPos(0, yOff))
	box = newThemedBox()
	box.Resize(fyne.NewSize(30, 30))
	box.Move(fyne.NewPos(0, yOff))

	anim = canvas.NewPositionAnimation(
		fyne.NewPos(0, yOff), fyne.NewPos(380, yOff), time.Millisecond*1500, func(p fyne.Position) {
			box.Move(p)
			box.Refresh()
		})
	anim.Curve = curve
	anim.AutoReverse = false
	anim.RepeatCount = 0
	return
}

// themedBox is a simple box that change its background color according
// to the selected theme
type themedBox struct {
	widget.BaseWidget
}

func newThemedBox() *themedBox {
	b := &themedBox{}
	b.ExtendBaseWidget(b)
	return b
}

func (b *themedBox) CreateRenderer() fyne.WidgetRenderer {
	b.ExtendBaseWidget(b)
	bg := canvas.NewRectangle(theme.ForegroundColor())
	return &themedBoxRenderer{bg: bg, objects: []fyne.CanvasObject{bg}}
}

type themedBoxRenderer struct {
	bg      *canvas.Rectangle
	objects []fyne.CanvasObject
}

func (r *themedBoxRenderer) Destroy() {
}

func (r *themedBoxRenderer) Layout(size fyne.Size) {
	r.bg.Resize(size)
}

func (r *themedBoxRenderer) MinSize() fyne.Size {
	return r.bg.MinSize()
}

func (r *themedBoxRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *themedBoxRenderer) Refresh() {
	r.bg.FillColor = theme.ForegroundColor()
	r.bg.Refresh()
}
