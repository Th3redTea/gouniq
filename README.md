# gouiq

`gouiq` is a command-line tool for deduplicating URLs based on unique path and query parameter structures. It's useful for removing redundant URLs with identical structures but different query values or path IDs.

## Usage

before:

```bash
$ cat list.txt
https://google.com
https://google.com/home?qs=value
https://google.com/home?qs=secondValue
https://google.com/home?qs=newValue&secondQs=anotherValue
https://google.com/home?qs=asd&secondQs=das
```
After

```bash
$ cat list.txt | ./gouniq
https://google.com
https://google.com/home?qs=value
https://google.com/home?qs=newValue&secondQs=anotherValue
```
## Features:

The similar flag (-s or --similar) in the gouniq tool enables a mode for deduplicating URLs based on similar paths, rather than on the exact full URL including query strings.

Purpose of the Similar Flag
With the similar flag:

Paths with Variable Segments: The tool treats URLs with variable segments (like IDs or dynamically generated components) as the same. For example:
URLs like /api/users/123 and /api/users/456 would be considered similar.
Instead of treating each as unique, the tool outputs only one representative.
Use Case: This is especially useful for deduplicating API endpoints or URL patterns where the path structure is the same but specific identifiers differ.
How It Works
When the similar flag is set, the tool uses a regex to remove or generalize segments in the path that match certain patterns, such as numerical IDs. For example:

/api/users/123 and /api/users/456 would both be transformed to /api/users/id.
/users/photos/photo.jpg and /users/photos/other_photo.jpg would remain as they are, as they don’t contain ID-like segments.
This mode helps reduce redundancy when the unique aspects of URLs are not relevant to deduplication (e.g., when IDs or specific photo names don’t matter).
### Deduplicate by Query Parameters
