self: {
  config,
  lib,
  pkgs,
  ...
}: let
  inherit (lib.modules) mkIf;
  inherit (lib.options) mkEnableOption mkPackageOption;
  inherit (lib.meta) getExe;

  cfg = config.programs.hypotd;
in {
  options = {
    programs.hypotd = {
      enable = mkEnableOption "hypotd";
      
      package = mkPackageOption self.packages.${pkgs.system} "hypotd" {
        default = "default";
        pkgsText = "hypotd.packages.\${pkgs.system}";
      };
    };
  };

  config = mkIf cfg.enable {
    home.packages = [cfg.package];

    systemd.user.services.hypotd = {
      Unit.Description = "hypotd - Picture of the day for hyprpaper";
      Install.WantedBy = [ "graphical-session.target" ];
      Service = {
        Type = "oneshot";
        Restart = "on-failure";
        ExecStart = "${getExe cfg.package}";
      };
    };
  };
}