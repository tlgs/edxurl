# edxurl

`edxurl` is a botched attempt at creating a utility to download the videos of edX courses.
I needed some source video files and both leading projects in the space
([coursera-dl/edx-dl](https://github.com/coursera-dl/edx-dl), and
[rehmatworks/edx-downloader](https://github.com/rehmatworks/edx-downloader))
were broken.

After a couple hours of revisiting Go, looking through those projects' source code,
and some investigating it became apparent that the problem would take more than one
weekend's worth of coding.
Arriving at the realization that the better solution is to directly inspect HTML tags
using the browser's developer tools is always an enriching experience.
