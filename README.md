# edxurl

`edxurl` is a botched attempt at creating a utility to the videos of edX courses.
Its need originated from me needing the source video files of lectures and
both leading projects in the space (
[coursera-dl/edx-dl](https://github.com/coursera-dl/edx-dl), and
[rehmatworks/edx-downloader](https://github.com/rehmatworks/edx-downloader))
being broken.

After a couple hours of revisiting Go, looking through those projects' source code,
and some investigating it became apparent that the problem would take more than one
weekend's worth of coding.
Arriving at the realization that the better solution is directly inspecting HTML tags
using a browser's developer tools is always fun. :)
