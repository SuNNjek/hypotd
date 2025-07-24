{
    buildGoModule,
    lib,
    ...
}: buildGoModule {
    pname = "hypotd";
    version = "0.1";

    vendorHash = null;

    src = ../.;
    subPackages = ["cmd/hypotd"];

    meta = {
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
        platforms = lib.platforms.linux;
        mainProgram = "hypotd";
    };
}