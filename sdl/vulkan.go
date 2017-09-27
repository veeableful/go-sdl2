package sdl

/*
#include "sdl_wrapper.h"

#if !SDL_VERSION_ATLEAST(2,0,6)
#pragma message("SDL_Vulkan_LoadLibrary is not supported before SDL 2.0.6")
static inline int SDL_Vulkan_LoadLibrary(const char *path)
{
	return -2;
}

#pragma message("SDL_Vulkan_GetVkGetInstanceProcAddr is not supported before SDL 2.0.6")
static inline void *SDLCALL SDL_Vulkan_GetVkGetInstanceProcAddr(void)
{
	return NULL;
}

#pragma message("SDL_Vulkan_UnloadLibrary is not supported before SDL 2.0.6")
static inline void SDL_Vulkan_UnloadLibrary()
{
}

static inline void SDL_Vulkan_GetDrawableSize(SDL_Window *window, int *w, int *h)
{
	if (w) *w = 0;
	if (h) *h = 0;
}

#pragma message("VkInstance is not supported before SDL 2.0.6")
typedef struct VkInstance VkInstance;

#pragma message("VkSurfaceKHR is not supported before SDL 2.0.6")
typedef struct VkSurfaceKHR VkSurfaceKHR;

#endif
*/
import "C"
import (
	"errors"
)

type VkInstance C.VkInstance
type VkSurfaceKHR C.VkSurfaceKHR

// VulkanLoadLibrary dynamically loads a Vulkan loader library
func VulkanLoadLibrary(path string) (err error) {
	ret := C.SDL_Vulkan_LoadLibrary(C.CString(path))
	if ret == -1 {
		err = GetError()
	} else if ret == -2 {
		err = errors.New("SDL_VulkanLoadLibrary is not supported before SDL 2.0.6")
	}
	return
}

// TODO
func VulkanGetVkGetInstanceProcAddr() uintptr {
	return uintptr(C.SDL_Vulkan_GetVkGetInstanceProcAddr())
}

// VulkanUnloadLibrary unloads the Vulkan loader library previously loaded by VulkanLoadLibrary()
func VulkanUnloadLibrary() {
	C.SDL_Vulkan_UnloadLibrary()
}

// TODO
func VulkanGetInstanceExtensions(window *Window) (count uint, names []string, err error) {
	return
}

// TODO
func VulkanCreateSurface(window *Window, instance VkInstance, surface *VkSurfaceKHR) (err error) {
	return
}

// Get the size of a window's underlying drawable in pixels (for use with setting viewport, scissor & etc)
func VulkanGetDrawableSize(window *Window) (w, h int) {
	var _w, _h C.int

	C.SDL_Vulkan_GetDrawableSize(window.cptr(), &_w, &_h)
	w = int(_w)
	h = int(_h)

	return
}
