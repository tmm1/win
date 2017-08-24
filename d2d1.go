// Copyright 2017 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

import (
	"syscall"
	"unsafe"
)

type D2D1_FACTORY_TYPE uint32

const (
	D2D1_FACTORY_TYPE_SINGLE_THREADED D2D1_FACTORY_TYPE = 0
	D2D1_FACTORY_TYPE_MULTI_THREADED  D2D1_FACTORY_TYPE = 1
)

type D2D1_DEBUG_LEVEL uint32

const (
	D2D1_DEBUG_LEVEL_NONE        D2D1_DEBUG_LEVEL = 0
	D2D1_DEBUG_LEVEL_ERROR       D2D1_DEBUG_LEVEL = 1
	D2D1_DEBUG_LEVEL_WARNING     D2D1_DEBUG_LEVEL = 2
	D2D1_DEBUG_LEVEL_INFORMATION D2D1_DEBUG_LEVEL = 3
)

type D2D1_RENDER_TARGET_TYPE uint32

const (
	D2D1_RENDER_TARGET_TYPE_DEFAULT  D2D1_RENDER_TARGET_TYPE = 0
	D2D1_RENDER_TARGET_TYPE_SOFTWARE D2D1_RENDER_TARGET_TYPE = 1
	D2D1_RENDER_TARGET_TYPE_HARDWARE D2D1_RENDER_TARGET_TYPE = 2
)

type DXGI_FORMAT uint32

const (
	DXGI_FORMAT_UNKNOWN                    DXGI_FORMAT = 0
	DXGI_FORMAT_R32G32B32A32_TYPELESS      DXGI_FORMAT = 1
	DXGI_FORMAT_R32G32B32A32_FLOAT         DXGI_FORMAT = 2
	DXGI_FORMAT_R32G32B32A32_UINT          DXGI_FORMAT = 3
	DXGI_FORMAT_R32G32B32A32_SINT          DXGI_FORMAT = 4
	DXGI_FORMAT_R32G32B32_TYPELESS         DXGI_FORMAT = 5
	DXGI_FORMAT_R32G32B32_FLOAT            DXGI_FORMAT = 6
	DXGI_FORMAT_R32G32B32_UINT             DXGI_FORMAT = 7
	DXGI_FORMAT_R32G32B32_SINT             DXGI_FORMAT = 8
	DXGI_FORMAT_R16G16B16A16_TYPELESS      DXGI_FORMAT = 9
	DXGI_FORMAT_R16G16B16A16_FLOAT         DXGI_FORMAT = 10
	DXGI_FORMAT_R16G16B16A16_UNORM         DXGI_FORMAT = 11
	DXGI_FORMAT_R16G16B16A16_UINT          DXGI_FORMAT = 12
	DXGI_FORMAT_R16G16B16A16_SNORM         DXGI_FORMAT = 13
	DXGI_FORMAT_R16G16B16A16_SINT          DXGI_FORMAT = 14
	DXGI_FORMAT_R32G32_TYPELESS            DXGI_FORMAT = 15
	DXGI_FORMAT_R32G32_FLOAT               DXGI_FORMAT = 16
	DXGI_FORMAT_R32G32_UINT                DXGI_FORMAT = 17
	DXGI_FORMAT_R32G32_SINT                DXGI_FORMAT = 18
	DXGI_FORMAT_R32G8X24_TYPELESS          DXGI_FORMAT = 19
	DXGI_FORMAT_D32_FLOAT_S8X24_UINT       DXGI_FORMAT = 20
	DXGI_FORMAT_R32_FLOAT_X8X24_TYPELESS   DXGI_FORMAT = 21
	DXGI_FORMAT_X32_TYPELESS_G8X24_UINT    DXGI_FORMAT = 22
	DXGI_FORMAT_R10G10B10A2_TYPELESS       DXGI_FORMAT = 23
	DXGI_FORMAT_R10G10B10A2_UNORM          DXGI_FORMAT = 24
	DXGI_FORMAT_R10G10B10A2_UINT           DXGI_FORMAT = 25
	DXGI_FORMAT_R11G11B10_FLOAT            DXGI_FORMAT = 26
	DXGI_FORMAT_R8G8B8A8_TYPELESS          DXGI_FORMAT = 27
	DXGI_FORMAT_R8G8B8A8_UNORM             DXGI_FORMAT = 28
	DXGI_FORMAT_R8G8B8A8_UNORM_SRGB        DXGI_FORMAT = 29
	DXGI_FORMAT_R8G8B8A8_UINT              DXGI_FORMAT = 30
	DXGI_FORMAT_R8G8B8A8_SNORM             DXGI_FORMAT = 31
	DXGI_FORMAT_R8G8B8A8_SINT              DXGI_FORMAT = 32
	DXGI_FORMAT_R16G16_TYPELESS            DXGI_FORMAT = 33
	DXGI_FORMAT_R16G16_FLOAT               DXGI_FORMAT = 34
	DXGI_FORMAT_R16G16_UNORM               DXGI_FORMAT = 35
	DXGI_FORMAT_R16G16_UINT                DXGI_FORMAT = 36
	DXGI_FORMAT_R16G16_SNORM               DXGI_FORMAT = 37
	DXGI_FORMAT_R16G16_SINT                DXGI_FORMAT = 38
	DXGI_FORMAT_R32_TYPELESS               DXGI_FORMAT = 39
	DXGI_FORMAT_D32_FLOAT                  DXGI_FORMAT = 40
	DXGI_FORMAT_R32_FLOAT                  DXGI_FORMAT = 41
	DXGI_FORMAT_R32_UINT                   DXGI_FORMAT = 42
	DXGI_FORMAT_R32_SINT                   DXGI_FORMAT = 43
	DXGI_FORMAT_R24G8_TYPELESS             DXGI_FORMAT = 44
	DXGI_FORMAT_D24_UNORM_S8_UINT          DXGI_FORMAT = 45
	DXGI_FORMAT_R24_UNORM_X8_TYPELESS      DXGI_FORMAT = 46
	DXGI_FORMAT_X24_TYPELESS_G8_UINT       DXGI_FORMAT = 47
	DXGI_FORMAT_R8G8_TYPELESS              DXGI_FORMAT = 48
	DXGI_FORMAT_R8G8_UNORM                 DXGI_FORMAT = 49
	DXGI_FORMAT_R8G8_UINT                  DXGI_FORMAT = 50
	DXGI_FORMAT_R8G8_SNORM                 DXGI_FORMAT = 51
	DXGI_FORMAT_R8G8_SINT                  DXGI_FORMAT = 52
	DXGI_FORMAT_R16_TYPELESS               DXGI_FORMAT = 53
	DXGI_FORMAT_R16_FLOAT                  DXGI_FORMAT = 54
	DXGI_FORMAT_D16_UNORM                  DXGI_FORMAT = 55
	DXGI_FORMAT_R16_UNORM                  DXGI_FORMAT = 56
	DXGI_FORMAT_R16_UINT                   DXGI_FORMAT = 57
	DXGI_FORMAT_R16_SNORM                  DXGI_FORMAT = 58
	DXGI_FORMAT_R16_SINT                   DXGI_FORMAT = 59
	DXGI_FORMAT_R8_TYPELESS                DXGI_FORMAT = 60
	DXGI_FORMAT_R8_UNORM                   DXGI_FORMAT = 61
	DXGI_FORMAT_R8_UINT                    DXGI_FORMAT = 62
	DXGI_FORMAT_R8_SNORM                   DXGI_FORMAT = 63
	DXGI_FORMAT_R8_SINT                    DXGI_FORMAT = 64
	DXGI_FORMAT_A8_UNORM                   DXGI_FORMAT = 65
	DXGI_FORMAT_R1_UNORM                   DXGI_FORMAT = 66
	DXGI_FORMAT_R9G9B9E5_SHAREDEXP         DXGI_FORMAT = 67
	DXGI_FORMAT_R8G8_B8G8_UNORM            DXGI_FORMAT = 68
	DXGI_FORMAT_G8R8_G8B8_UNORM            DXGI_FORMAT = 69
	DXGI_FORMAT_BC1_TYPELESS               DXGI_FORMAT = 70
	DXGI_FORMAT_BC1_UNORM                  DXGI_FORMAT = 71
	DXGI_FORMAT_BC1_UNORM_SRGB             DXGI_FORMAT = 72
	DXGI_FORMAT_BC2_TYPELESS               DXGI_FORMAT = 73
	DXGI_FORMAT_BC2_UNORM                  DXGI_FORMAT = 74
	DXGI_FORMAT_BC2_UNORM_SRGB             DXGI_FORMAT = 75
	DXGI_FORMAT_BC3_TYPELESS               DXGI_FORMAT = 76
	DXGI_FORMAT_BC3_UNORM                  DXGI_FORMAT = 77
	DXGI_FORMAT_BC3_UNORM_SRGB             DXGI_FORMAT = 78
	DXGI_FORMAT_BC4_TYPELESS               DXGI_FORMAT = 79
	DXGI_FORMAT_BC4_UNORM                  DXGI_FORMAT = 80
	DXGI_FORMAT_BC4_SNORM                  DXGI_FORMAT = 81
	DXGI_FORMAT_BC5_TYPELESS               DXGI_FORMAT = 82
	DXGI_FORMAT_BC5_UNORM                  DXGI_FORMAT = 83
	DXGI_FORMAT_BC5_SNORM                  DXGI_FORMAT = 84
	DXGI_FORMAT_B5G6R5_UNORM               DXGI_FORMAT = 85
	DXGI_FORMAT_B5G5R5A1_UNORM             DXGI_FORMAT = 86
	DXGI_FORMAT_B8G8R8A8_UNORM             DXGI_FORMAT = 87
	DXGI_FORMAT_B8G8R8X8_UNORM             DXGI_FORMAT = 88
	DXGI_FORMAT_R10G10B10_XR_BIAS_A2_UNORM DXGI_FORMAT = 89
	DXGI_FORMAT_B8G8R8A8_TYPELESS          DXGI_FORMAT = 90
	DXGI_FORMAT_B8G8R8A8_UNORM_SRGB        DXGI_FORMAT = 91
	DXGI_FORMAT_B8G8R8X8_TYPELESS          DXGI_FORMAT = 92
	DXGI_FORMAT_B8G8R8X8_UNORM_SRGB        DXGI_FORMAT = 93
	DXGI_FORMAT_BC6H_TYPELESS              DXGI_FORMAT = 94
	DXGI_FORMAT_BC6H_UF16                  DXGI_FORMAT = 95
	DXGI_FORMAT_BC6H_SF16                  DXGI_FORMAT = 96
	DXGI_FORMAT_BC7_TYPELESS               DXGI_FORMAT = 97
	DXGI_FORMAT_BC7_UNORM                  DXGI_FORMAT = 98
	DXGI_FORMAT_BC7_UNORM_SRGB             DXGI_FORMAT = 99
	DXGI_FORMAT_AYUV                       DXGI_FORMAT = 100
	DXGI_FORMAT_Y410                       DXGI_FORMAT = 101
	DXGI_FORMAT_Y416                       DXGI_FORMAT = 102
	DXGI_FORMAT_NV12                       DXGI_FORMAT = 103
	DXGI_FORMAT_P010                       DXGI_FORMAT = 104
	DXGI_FORMAT_P016                       DXGI_FORMAT = 105
	DXGI_FORMAT_420_OPAQUE                 DXGI_FORMAT = 106
	DXGI_FORMAT_YUY2                       DXGI_FORMAT = 107
	DXGI_FORMAT_Y210                       DXGI_FORMAT = 108
	DXGI_FORMAT_Y216                       DXGI_FORMAT = 109
	DXGI_FORMAT_NV11                       DXGI_FORMAT = 110
	DXGI_FORMAT_AI44                       DXGI_FORMAT = 111
	DXGI_FORMAT_IA44                       DXGI_FORMAT = 112
	DXGI_FORMAT_P8                         DXGI_FORMAT = 113
	DXGI_FORMAT_A8P8                       DXGI_FORMAT = 114
	DXGI_FORMAT_B4G4R4A4_UNORM             DXGI_FORMAT = 115
	DXGI_FORMAT_P208                       DXGI_FORMAT = 130
	DXGI_FORMAT_V208                       DXGI_FORMAT = 131
	DXGI_FORMAT_V408                       DXGI_FORMAT = 132
	DXGI_FORMAT_FORCE_UINT                 DXGI_FORMAT = 0xffffffff
)

type D2D1_ALPHA_MODE uint32

const (
	D2D1_ALPHA_MODE_UNKNOWN       D2D1_ALPHA_MODE = 0
	D2D1_ALPHA_MODE_PREMULTIPLIED D2D1_ALPHA_MODE = 1
	D2D1_ALPHA_MODE_STRAIGHT      D2D1_ALPHA_MODE = 2
	D2D1_ALPHA_MODE_IGNORE        D2D1_ALPHA_MODE = 3
)

type D2D1_RENDER_TARGET_USAGE uint32

const (
	D2D1_RENDER_TARGET_USAGE_NONE                  D2D1_RENDER_TARGET_USAGE = 0x00000000
	D2D1_RENDER_TARGET_USAGE_FORCE_BITMAP_REMOTING D2D1_RENDER_TARGET_USAGE = 0x00000001
	D2D1_RENDER_TARGET_USAGE_GDI_COMPATIBLE        D2D1_RENDER_TARGET_USAGE = 0x00000002
)

const (
	D3D10_FEATURE_LEVEL_9_1  = 0x9100
	D3D10_FEATURE_LEVEL_10_0 = 0xa000
)

type D2D1_FEATURE_LEVEL uint32

const (
	D2D1_FEATURE_LEVEL_DEFAULT D2D1_FEATURE_LEVEL = 0
	D2D1_FEATURE_LEVEL_9       D2D1_FEATURE_LEVEL = D3D10_FEATURE_LEVEL_9_1
	D2D1_FEATURE_LEVEL_10      D2D1_FEATURE_LEVEL = D3D10_FEATURE_LEVEL_10_0
)

type D2D1_CAP_STYLE uint32

const (
	D2D1_CAP_STYLE_FLAT     D2D1_CAP_STYLE = 0
	D2D1_CAP_STYLE_SQUARE   D2D1_CAP_STYLE = 1
	D2D1_CAP_STYLE_ROUND    D2D1_CAP_STYLE = 2
	D2D1_CAP_STYLE_TRIANGLE D2D1_CAP_STYLE = 3
)

type D2D1_DASH_STYLE uint32

const (
	D2D1_DASH_STYLE_SOLID        D2D1_DASH_STYLE = 0
	D2D1_DASH_STYLE_DASH         D2D1_DASH_STYLE = 1
	D2D1_DASH_STYLE_DOT          D2D1_DASH_STYLE = 2
	D2D1_DASH_STYLE_DASH_DOT     D2D1_DASH_STYLE = 3
	D2D1_DASH_STYLE_DASH_DOT_DOT D2D1_DASH_STYLE = 4
	D2D1_DASH_STYLE_CUSTOM       D2D1_DASH_STYLE = 5
)

type D2D1_LINE_JOIN uint32

const (
	D2D1_LINE_JOIN_MITER          D2D1_LINE_JOIN = 0
	D2D1_LINE_JOIN_BEVEL          D2D1_LINE_JOIN = 1
	D2D1_LINE_JOIN_ROUND          D2D1_LINE_JOIN = 2
	D2D1_LINE_JOIN_MITER_OR_BEVEL D2D1_LINE_JOIN = 3
)

type D2D1_PRESENT_OPTIONS uint32

const (
	D2D1_PRESENT_OPTIONS_NONE            D2D1_PRESENT_OPTIONS = 0x00000000
	D2D1_PRESENT_OPTIONS_RETAIN_CONTENTS D2D1_PRESENT_OPTIONS = 0x00000001
	D2D1_PRESENT_OPTIONS_IMMEDIATELY     D2D1_PRESENT_OPTIONS = 0x00000002
)

const (
	D2DERR_RECREATE_TARGET = 0x8899000c
)

type D2D1_FACTORY_OPTIONS struct {
	DebugLevel D2D1_DEBUG_LEVEL
}

type D2D1_RENDER_TARGET_PROPERTIES struct {
	Type        D2D1_RENDER_TARGET_TYPE
	PixelFormat D2D1_PIXEL_FORMAT
	DpiX        float32
	DpiY        float32
	Usage       D2D1_RENDER_TARGET_USAGE
	MinLevel    D2D1_FEATURE_LEVEL
}

type D2D1_HWND_RENDER_TARGET_PROPERTIES struct {
	Hwnd           HWND
	PixelSize      D2D1_SIZE_U
	PresentOptions D2D1_PRESENT_OPTIONS
}

type D2D1_PIXEL_FORMAT struct {
	Format    DXGI_FORMAT
	AlphaMode D2D1_ALPHA_MODE
}

type D2D1_TAG uint64

type D2D1_COLOR_F struct {
	R float32
	G float32
	B float32
	A float32
}

type D2D1_SIZE_U struct {
	Width  uint32
	Height uint32
}

type D2D1_RECT_F struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

type D2D1_ROUNDED_RECT struct {
	Rect    D2D1_RECT_F
	RadiusX float32
	RadiusY float32
}

type D2D1_STROKE_STYLE_PROPERTIES struct {
	StartCap   D2D1_CAP_STYLE
	EndCap     D2D1_CAP_STYLE
	DashCap    D2D1_CAP_STYLE
	LineJoin   D2D1_LINE_JOIN
	MiterLimit float32
	DashStyle  D2D1_DASH_STYLE
	DashOffset float32
}

var (
	IID_ID2D1Factory = IID{0x06152247, 0x6f50, 0x465a, [8]byte{0x92, 0x45, 0x11, 0x8b, 0xfd, 0x3b, 0x60, 0x07}}
)

var (
	// Library
	libd2d1 = MustLoadLibrary("d2d1.dll")

	// Functions
	d2D1CreateFactory = MustGetProcAddress(libd2d1, "D2D1CreateFactory")
)

func D2D1CreateFactory(factoryType D2D1_FACTORY_TYPE, riid REFIID, pFactoryOptions *D2D1_FACTORY_OPTIONS, ppIFactory *unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.Syscall6(d2D1CreateFactory, 4,
		uintptr(factoryType),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(pFactoryOptions)),
		uintptr(unsafe.Pointer(ppIFactory)),
		0,
		0)

	return HRESULT(ret)
}

type ID2D1FactoryVtbl struct {
	IUnknownVtbl
	ReloadSystemMetrics            uintptr
	GetDesktopDpi                  uintptr
	CreateRectangleGeometry        uintptr
	CreateRoundedRectangleGeometry uintptr
	CreateEllipseGeometry          uintptr
	CreateGeometryGroup            uintptr
	CreateTransformedGeometry      uintptr
	CreatePathGeometry             uintptr
	CreateStrokeStyle              uintptr
	CreateDrawingStateBlock        uintptr
	CreateWicBitmapRenderTarget    uintptr
	CreateHwndRenderTarget         uintptr
	CreateDxgiSurfaceRenderTarget  uintptr
	CreateDCRenderTarget           uintptr
}

type ID2D1Factory struct {
	Vtbl *ID2D1FactoryVtbl
}

func (f *ID2D1Factory) Release() uint32 {
	ret, _, _ := syscall.Syscall(f.Vtbl.Release, 1,
		uintptr(unsafe.Pointer(f)),
		0,
		0)

	return uint32(ret)
}

func (f *ID2D1Factory) CreateHwndRenderTarget(renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES, hwndRenderTargetProperties *D2D1_HWND_RENDER_TARGET_PROPERTIES, hwndRenderTarget **ID2D1HwndRenderTarget) HRESULT {
	ret, _, _ := syscall.Syscall6(f.Vtbl.CreateHwndRenderTarget, 4,
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(hwndRenderTargetProperties)),
		uintptr(unsafe.Pointer(hwndRenderTarget)),
		0,
		0)

	return HRESULT(ret)
}

func (f *ID2D1Factory) CreateDCRenderTarget(renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES, dcRenderTarget **ID2D1DCRenderTarget) HRESULT {
	ret, _, _ := syscall.Syscall(f.Vtbl.CreateDCRenderTarget, 3,
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(dcRenderTarget)))

	return HRESULT(ret)
}

func (f *ID2D1Factory) CreateStrokeStyle(strokeStyleProperties *D2D1_STROKE_STYLE_PROPERTIES, dashes *float32, dashesCount uint32, strokeStyle **ID2D1StrokeStyle) HRESULT {
	ret, _, _ := syscall.Syscall6(f.Vtbl.CreateStrokeStyle, 5,
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(strokeStyleProperties)),
		uintptr(unsafe.Pointer(dashes)),
		uintptr(dashesCount),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)

	return HRESULT(ret)
}

type ID2D1RenderTargetVtbl struct {
	ID2D1ResourceVtbl
	CreateBitmap                 uintptr
	CreateBitmapFromWicBitmap    uintptr
	CreateSharedBitmap           uintptr
	CreateBitmapBrush            uintptr
	CreateSolidColorBrush        uintptr
	CreateGradientStopCollection uintptr
	CreateLinearGradientBrush    uintptr
	CreateRadialGradientBrush    uintptr
	CreateCompatibleRenderTarget uintptr
	CreateLayer                  uintptr
	CreateMesh                   uintptr
	DrawLine                     uintptr
	DrawRectangle                uintptr
	FillRectangle                uintptr
	DrawRoundedRectangle         uintptr
	FillRoundedRectangle         uintptr
	DrawEllipse                  uintptr
	FillEllipse                  uintptr
	DrawGeometry                 uintptr
	FillGeometry                 uintptr
	FillMesh                     uintptr
	FillOpacityMask              uintptr
	DrawBitmap                   uintptr
	DrawText                     uintptr
	DrawTextLayout               uintptr
	DrawGlyphRun                 uintptr
	SetTransform                 uintptr
	GetTransform                 uintptr
	SetAntialiasMode             uintptr
	GetAntialiasMode             uintptr
	SetTextAntialiasMode         uintptr
	GetTextAntialiasMode         uintptr
	SetTextRenderingParams       uintptr
	GetTextRenderingParams       uintptr
	SetTags                      uintptr
	GetTags                      uintptr
	PushLayer                    uintptr
	PopLayer                     uintptr
	Flush                        uintptr
	SaveDrawingState             uintptr
	RestoreDrawingState          uintptr
	PushAxisAlignedClip          uintptr
	PopAxisAlignedClip           uintptr
	Clear                        uintptr
	BeginDraw                    uintptr
	EndDraw                      uintptr
	GetPixelFormat               uintptr
	SetDpi                       uintptr
	GetDpi                       uintptr
	GetSize                      uintptr
	GetPixelSize                 uintptr
	GetMaximumBitmapSize         uintptr
	IsSupported                  uintptr
}

type ID2D1RenderTarget struct {
	Vtbl *ID2D1RenderTargetVtbl
}

func (rt *ID2D1RenderTarget) Release() uint32 {
	ret, _, _ := syscall.Syscall(rt.Vtbl.Release, 1,
		uintptr(unsafe.Pointer(rt)),
		0,
		0)

	return uint32(ret)
}

func (rt *ID2D1RenderTarget) CreateSolidColorBrush(color *D2D1_COLOR_F, brushProperties *D2D1_BRUSH_PROPERTIES, solidColorBrush **ID2D1SolidColorBrush) HRESULT {
	ret, _, _ := syscall.Syscall6(rt.Vtbl.CreateSolidColorBrush, 4,
		uintptr(unsafe.Pointer(rt)),
		uintptr(unsafe.Pointer(color)),
		uintptr(unsafe.Pointer(brushProperties)),
		uintptr(unsafe.Pointer(solidColorBrush)),
		0,
		0)

	return HRESULT(ret)
}

func (rt *ID2D1RenderTarget) FillRectangle(rect *D2D1_RECT_F, brush *ID2D1Brush) {
	syscall.Syscall(rt.Vtbl.FillRectangle, 3,
		uintptr(unsafe.Pointer(rt)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(brush)))
}

// DrawRoundedRectangle does not work, because syscall does not support float args...
//func (rt *ID2D1RenderTarget) DrawRoundedRectangle(roundedRect *D2D1_ROUNDED_RECT, brush *ID2D1Brush, strokeWidth float32, strokeStyle *ID2D1StrokeStyle) {
//	syscall.Syscall6(rt.Vtbl.DrawRoundedRectangle, 5,
//		uintptr(unsafe.Pointer(rt)),
//		uintptr(unsafe.Pointer(roundedRect)),
//		uintptr(unsafe.Pointer(brush)),
//		uintptr(strokeWidth),
//		uintptr(unsafe.Pointer(strokeStyle)),
//		0)
//}

func (rt *ID2D1RenderTarget) FillRoundedRectangle(roundedRect *D2D1_ROUNDED_RECT, brush *ID2D1Brush) {
	syscall.Syscall(rt.Vtbl.FillRoundedRectangle, 3,
		uintptr(unsafe.Pointer(rt)),
		uintptr(unsafe.Pointer(roundedRect)),
		uintptr(unsafe.Pointer(brush)))
}

func (rt *ID2D1RenderTarget) SetTransform(transform *D2D1_MATRIX_3X2_F) {
	syscall.Syscall(rt.Vtbl.SetTransform, 2,
		uintptr(unsafe.Pointer(rt)),
		uintptr(unsafe.Pointer(transform)),
		0)
}

func (rt *ID2D1RenderTarget) BeginDraw() {
	syscall.Syscall(rt.Vtbl.BeginDraw, 1,
		uintptr(unsafe.Pointer(rt)),
		0,
		0)
}

func (rt *ID2D1RenderTarget) EndDraw(tag1, tag2 *D2D1_TAG) HRESULT {
	ret, _, _ := syscall.Syscall(rt.Vtbl.EndDraw, 3,
		uintptr(unsafe.Pointer(rt)),
		uintptr(unsafe.Pointer(tag1)),
		uintptr(unsafe.Pointer(tag2)))

	return HRESULT(ret)
}

func (rt *ID2D1RenderTarget) Clear(clearColor *D2D1_COLOR_F) {
	syscall.Syscall(rt.Vtbl.Clear, 2,
		uintptr(unsafe.Pointer(rt)),
		uintptr(unsafe.Pointer(clearColor)),
		0)
}

type D2D1_BRUSH_PROPERTIES struct {
	Opacity   float32
	Transform D2D1_MATRIX_3X2_F
}

type D2D1_MATRIX_3X2_F struct {
	M11 float32
	M12 float32
	M21 float32
	M22 float32
	M31 float32
	M32 float32
}

type ID2D1BrushVtbl struct {
	ID2D1ResourceVtbl
	SetOpacity   uintptr
	SetTransform uintptr
	GetOpacity   uintptr
	GetTransform uintptr
}

type ID2D1Brush struct {
	Vtbl *ID2D1BrushVtbl
}

type ID2D1SolidColorBrushVtbl struct {
	ID2D1BrushVtbl
	SetColor uintptr
	GetColor uintptr
}

type ID2D1SolidColorBrush struct {
	Vtbl *ID2D1SolidColorBrushVtbl
}

func (scb *ID2D1SolidColorBrush) Release() uint32 {
	ret, _, _ := syscall.Syscall(scb.Vtbl.Release, 1,
		uintptr(unsafe.Pointer(scb)),
		0,
		0)

	return uint32(ret)
}

type ID2D1StrokeStyleVtbl struct {
	ID2D1ResourceVtbl
	GetStartCap    uintptr
	GetEndCap      uintptr
	GetDashCap     uintptr
	GetMiterLimit  uintptr
	GetLineJoin    uintptr
	GetDashOffset  uintptr
	GetDashStyle   uintptr
	GetDashesCount uintptr
	GetDashes      uintptr
}

type ID2D1StrokeStyle struct {
	Vtbl *ID2D1StrokeStyleVtbl
}

func (ss *ID2D1StrokeStyle) Release() uint32 {
	ret, _, _ := syscall.Syscall(ss.Vtbl.Release, 1,
		uintptr(unsafe.Pointer(ss)),
		0,
		0)

	return uint32(ret)
}

type ID2D1ResourceVtbl struct {
	IUnknownVtbl
	GetFactory uintptr
}

type ID2D1Resource struct {
	Vtbl *ID2D1ResourceVtbl
}

type ID2D1DCRenderTargetVtbl struct {
	ID2D1RenderTargetVtbl
	BindDC uintptr
}

type ID2D1DCRenderTarget struct {
	Vtbl *ID2D1DCRenderTargetVtbl
}

func (dcrt *ID2D1DCRenderTarget) BindDC(hDC HDC, pSubRect *RECT) HRESULT {
	ret, _, _ := syscall.Syscall(dcrt.Vtbl.BindDC, 3,
		uintptr(unsafe.Pointer(dcrt)),
		uintptr(hDC),
		uintptr(unsafe.Pointer(pSubRect)))

	return HRESULT(ret)
}

type ID2D1HwndRenderTargetVtbl struct {
	ID2D1RenderTargetVtbl
	CheckWindowState uintptr
	Resize           uintptr
	GetHwnd          uintptr
}

type ID2D1HwndRenderTarget struct {
	Vtbl *ID2D1HwndRenderTargetVtbl
}

func (hwndrt *ID2D1HwndRenderTarget) Resize(pixelSize *D2D1_SIZE_U) HRESULT {
	ret, _, _ := syscall.Syscall(hwndrt.Vtbl.Resize, 2,
		uintptr(unsafe.Pointer(hwndrt)),
		uintptr(unsafe.Pointer(pixelSize)),
		0)

	return HRESULT(ret)
}
