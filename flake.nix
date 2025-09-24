{
    description = ''
        Download pictures-of-the-day and set them via hyprpaper
    '';

    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    };

    outputs = {
        self,
        nixpkgs,
        ...
    }: let
        allSystems = builtins.attrNames nixpkgs.legacyPackages;

        forAllSystems = (f:
            nixpkgs.lib.genAttrs allSystems (system:
                f nixpkgs.legacyPackages.${system}
            )
        );
    in {
        packages = forAllSystems (pkgs: {
            hypotd = pkgs.callPackage ./nix/package.nix {};
            default = self.packages.${pkgs.system}.hypotd;
        });

        devShells = forAllSystems (pkgs: {
            default = pkgs.mkShell {
                packages = with pkgs; [
                    go
                ];
            };
        });

        homeManagerModules = {
            default = self.homeManagerModules.hypotd;
            hypotd = import ./nix/modules/home-manager.nix self;
        };
    };
}
