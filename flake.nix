{
    description = ''
        Download pictures-of-the-day and set them via hyprpaper
    '';

    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
        flake-utils.url = "github:numtide/flake-utils";
    };

    outputs = {
        self,
        nixpkgs,
        flake-utils,
        ...
    }: flake-utils.lib.eachDefaultSystem (system:
        let
            pkgs = nixpkgs.legacyPackages.${system};
        in {
            devShells.default = pkgs.mkShell {
                packages = with pkgs; [
                    go
                ];
            };

            packages = {
                default = self.packages.${system}.hypotd;
                hypotd = pkgs.callPackage ./nix/package.nix {};
            };
        }
    ) // flake-utils.lib.eachDefaultSystemPassThrough (system: {
        # TODO
    });
}
