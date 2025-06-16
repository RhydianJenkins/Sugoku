{
    description = "SuGoKu - A simple Sudoku solver written in Go";

    inputs = {
        nixpkgs = {
            url = "github:nixos/nixpkgs/nixos-25.05";
        };
        flake-utils = {
            url = "github:numtide/flake-utils";
        };
    };

    outputs = { nixpkgs, flake-utils, ... }: flake-utils.lib.eachDefaultSystem (
        system:
        let
            pkgs = nixpkgs.legacyPackages.${system};
        in {
            packages.default = pkgs.buildGoModule {
                pname = "sugoku";
                version = "0.1.0";
                src = ./.;
                vendorHash = null;
            };

            devShells.default = pkgs.mkShell {
                buildInputs = [ pkgs.go_1_24 ];
            };
        }
    );
}
