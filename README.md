# hypotd
Download pictures-of-the-day and set them via hyprpaper, inspired by the KDE wallpaper module.
The name is literally just a portmanteau of hyprland + POTD (short for picture of the day) because I'm uncreative ^^

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
			provider = "pexels";

			pexels = {
				apiKey = "<API_KEY>";
			};
		};
	};

    # ...
}
```

## Configuration

### Providers

#### Bing

The Bing provider has no options, just set provider to "bing":
```toml
provider = "bing"
```

#### Pexels

For the Pexels provider you must provide an API key:
```toml
provider = "pexels"

[pexels]
apiKey = "<API_KEY>" # <-- Replace with your API key
```

#### NASA Astronomy Picture of the Day

For the APOD provider you *can* provide an API key. If you don't, it will use the `DEMO_KEY` key
which is rate-limited (maximum of 30 requests per hour, 50 per day). If you plan to logout/login more
often than that, maybe request an API key ^^

```toml
provider = "apod"

[apod]
apiKey = "<API_KEY>" # <-- Replace with your API key (or don't specify any)
```


### Use custom command to set wallpaper

By default, hypotd uses hyprpaper to set the wallpaper. If you want to use a custom command,
you can configure it like this:

```toml
customCommand = "caelestia wallpaper -f {{.Path}}"
```

This uses a Go template to construct the command. `{{.Path}}` is the path to the wallpaper.