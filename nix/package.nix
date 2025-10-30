{
  buildGoModule,
  lib,

  version,
  ...
}: buildGoModule {
  pname = "hypotd";
  inherit version;

  vendorHash = "sha256-vjlEgaC62t5EZkkAuk3qXA5KkadWN674nXorVDtrkjI=";

  src = ../.;
  subPackages = ["cmd/hypotd"];

  meta = {
    mainProgram = "hypotd";
    platforms = lib.platforms.linux;
    description = "Download pictures-of-the-day and set them via hyprpaper";
    homepage = "https://github.com/SuNNjek/hypotd";
    license = lib.licenses.gpl3Plus;
    maintainers = [
      {
        name = "Sunner";
        email = "sunnerlp@gmail.com";
        github = "SuNNjek";
        githubId = 20151081;
      }
    ];
  };
}