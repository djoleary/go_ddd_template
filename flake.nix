{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    outrig.url = "github:outrigdev/outrig";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      ...
    }@inputs:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            inputs.outrig.overlays.${system}.default
          ];
        };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputts = with pkgs; [
            # Go
            go
            gopls
            delve

            # Build
            docker-compose

            # Dev
            go-task
            golangci-lint
            outrig
          ];
        };
      }
    );
}
