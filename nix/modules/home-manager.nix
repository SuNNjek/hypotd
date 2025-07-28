self: {
  config,
  lib,
  pkgs,
  ...
}: let
  inherit (lib.modules) mkIf;
  inherit (lib.options) mkOption mkEnableOption mkPackageOption;
  inherit (lib.meta) getExe;
  inherit (lib.trivial) importTOML;

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
        After = [ "network-online.target" ];
      };
      Install = {
        WantedBy = [ "graphical-session.target" ];
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