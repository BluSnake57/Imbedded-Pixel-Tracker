# Imbedded Pixel Tracker

A basic implementation of an imbedded pixel tracker. It can log IP Addresses along with access time, 
currently the information is not locked once it's opened so it can be overwritten if it's reopened

- Ability to create multiple trackers at once
- Able to view status of each tracker
- Ability to individually delete trackers
- *Really* basic CLI interface

I really made this as a jumping off point for anyone that was like me who really want to know how 
this kind of thing works but don't quite have the know how yet to figure it out

## Prerequisites

Golang (No specific version but I used 1.24.7)

Cloudflare proxy (If planning to run outside of your network)

## Setup/Usage

You can either run this tool from CLI or by sending web requests, the web request method is more fleshed out but both do work

There is also a Dockerfile included that you can use, the image is hosted on this github so you can use that if you know how. 
This would be for a more serious deployment allowing for a reverse proxy.

For everyone else you can just run it in terminal.

In terminal/console enter this directory and then run 

`go run .`

It runs on port 4040

### CLI

The basic commands are
-Help: this will list all commands and uses
-Create \[tracker-name\]: will create a tracking link
-Status \[tracker-name\]: will show the status of a tracker
-Delete \[tracker-name\]: will delete a tracker

### Network Requests

The paths are
-/initalize_tracker
-/get_status
-/kill_tracker

I recommend using curl to send these, I'm not going to be giving that much info on how to do this because it's kinda tedious...pretty much the whole reason I bothered to 
throw together a CLI

[!WARNING]
Please don't use this for anything melicious

[!NOTE]
If you get stuck on anything try using google, that's how I learned all of this so I know the info is out there if you're having a real hard time though feel free to make an 
issue ticket and I'll try and give you a hand.

Also no I didn't put one on this page...probably could've tho I'd be curious to know if that works (Don't do it you could probably get your account banned)

Side note I found some interesting stuff relating to who does and doesn't pre-download image links 
crazy enough Discord was the only one I found to actually do it