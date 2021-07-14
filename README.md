# edxurl

`edxurl` is a botched attempt at creating a utility to download the videos of edX courses.
I needed some source video files and both leading projects in the space
([coursera-dl/edx-dl](https://github.com/coursera-dl/edx-dl), and
[rehmatworks/edx-downloader](https://github.com/rehmatworks/edx-downloader))
were broken.

After a couple hours of revisiting Go, looking through those projects' source code,
and some investigating it became apparent that the problem would take more than one
weekend's worth of coding.
The realization you're better off manually inspecting HTML tags is bittersweet.

## Usage

```console
$ ./edxurl -email $EDX_EMAIL -password $EDX_PASSWORD -course course-v1:HarvardX+CS50+X
2021/07/14 15:29:28 csrftoken: oB09s3x0fbYUFnhLN4NqjPGgGGMYEbCTFEIyvsaCw9d9vwWIryp7hmaTrOWuYs0I
2021/07/14 15:29:30 authentication successful
{ ... }
```
