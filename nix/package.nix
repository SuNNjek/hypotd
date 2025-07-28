{
    buildGoModule,
    lib,
    ...
}: buildGoModule {
    pname = "hypotd";
    version = "0.1";

    vendorHash = "sha256-KDEKOybdC+Ef4S042feDpAJpI1gfg32BfuGJVpKpRJQ=";

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