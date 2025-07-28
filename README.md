# hypotd
Download pictures-of-the-day and set them via hyprpaper, inspired by the KDE wallpaper module.
The name is literally just a portmanteau of hyprland + POTD (short for picture of the day) because I'm uncreative ^^

I'm planning to support multiple image providers, for now only Bing is implemented.

## Installation on NixOS

This repository contains a Nix flake you can use to install hypotd with home manager. Simply add the repository as a flake input:

```nix
{
    inputs = {
        # ...

		hypotd.url = "github:SuNNjek/hypotd";
    };

    # ...
}
```

and add it to your home manager config:
```nix
{ inputs, ... }: {
	imports = [
		inputs.hypotd.homeManagerModules.default
	];
	
	programs.hypotd = {
		enable = true;

		config = {
			provider = "bing";
		};
	};

    # ...
}
```