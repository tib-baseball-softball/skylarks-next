# create hardlinks from submodule
# macOS does not support readable options `--recursive --force --link`
cp -R -f -l "diamond-planner/ui/src/lib/dp" "src/lib/dp"
cp -R -f -l "diamond-planner/ui/src/routes/(dp)" "src/routes/(dp)"
