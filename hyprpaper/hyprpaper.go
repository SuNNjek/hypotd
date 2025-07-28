package hyprpaper

import (
	"context"
	"fmt"
	"os/exec"
)

func SetWallpaper(ctx context.Context, path string) error {
	// Unload all old wallpapers
	unloadCmd := exec.CommandContext(ctx, "hyprctl", "hyprpaper", "unload", "all")
	if err := unloadCmd.Run(); err != nil {
		return err
	}

	// Preload given wallpaper
	preloadCmd := exec.CommandContext(ctx, "hyprctl", "hyprpaper", "preload", path)
	if err := preloadCmd.Run(); err != nil {
		return err
	}

	// Set the wallpaper for all monitors (for now)
	wallpaperCmd := exec.CommandContext(ctx, "hyprctl", "hyprpaper", "wallpaper", fmt.Sprintf(",%s", path))
	if err := wallpaperCmd.Run(); err != nil {
		return err
	}

	return nil
}
