self: {
  config,
  lib,
  pkgs,
  ...
}:
with lib;
let
  tomlFormat = pkgs.formats.toml {};

  cfg = config.programs.hypotd;
in {
  options = {
    programs.hypotd = {
      enable = mkEnableOption "hypotd";
      
      package = mkPackageOption self.packages.${pkgs.system} "hypotd" {
        default = "default";
        pkgsText = "hypotd.packages.\${pkgs.system}";
      };

      target = mkOption {
        type = types.str;
        default = config.wayland.systemd.target;
        defaultText = "config.wayland.systemd.target";
        description = ''
          systemd target after which to start the service
        '';
      };

      config = mkOption {
        inherit (tomlFormat) type;
        default = importTOML ../../config.toml;
        defaultText = "importTOML ../../config.toml";
        description = ''
          Configuration written to {file}`$XDG_CONFIG_HOME/hypotd/config.toml`.
        '';
      };
    };
  };

  config = mkIf cfg.enable {
    home.packages = [cfg.package];

    xdg.configFile."hypotd/config.toml".source = mkIf (cfg.config != {})
      (tomlFormat.generate "hypotd-config.toml" cfg.config);

    systemd.user.services.hypotd = {
      Unit = {
        Description = "hypotd - Picture of the day for hyprpaper";
        After = [ "network-online.target" cfg.target ];
      };

      Install = {
        WantedBy = [ cfg.target ];
      };

      Service = {
        Type = "oneshot";
        Restart = "on-failure";
        RestartSec = "5";
        ExecStart = "${getExe cfg.package}";
      };
    };
  };
}