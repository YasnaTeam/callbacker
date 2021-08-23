# Callbacker :calling:

Callbacker is a utility for creating fake :hear_no_evil: callback URLs. Suppose you are working on a service which needs
some callback calls from other services (like github or google). Then the main problem is getting callbacks on
the `localhost` without any valid IP or URL :no_entry:. Callbacker is answer to that problem :ok_hand:.

## Installing :computer:

### Direct method :arrow_double_down:

Get executable package from [latest release](https://github.com/YasnaTeam/callbacker/releases/latest) page and make it
executable on your system.

### Via GO! :1st_place_medal:

Run this command

```bash
go get https://github.com/YasnaTeam/callbacker
```

## Terminology :arrow_heading_down:

Callbacker contains two part: `server` and `client`. `server` is one which register user connections, get their requests
for callback definition and finally forward callback requests to `client`. `client` is the end user tools which can
register callback, get forwarded request and send them to final destination.

### Running a `server` instance :construction_worker:

Suppose your final domain is `http://domain.ltd`, first you must define an endpoint with `/callback` URI. On `Nginx` you
can do that with some configurations like this:

```nginx
server {
	server_name domain.ltd;

	location /callback {
		proxy_pass http://127.0.0.1:1616;
	}
}
```

then run the `server`:

```bash
./callbacker -v -s -u http://domain.ltd/callback
```

#### Defining a `systemd` service for server instance

For definition of a `systemd` service, you can create a file on `~/.config/systemd/user` and put this lines on it:

```bash
[Unit]
Description=Callbacker
After=network.target

[Service]
Type=simple
NonBlocking=true

ExecStart=/absolute/path/to/callbacker -s -u http://domain.ltd/callback
Restart=on-failure
RestartSec=5s
```

then reload daemons with `systemctl daemon-reload` and run the service with `systemctl start callbacker`.

With `./callbacker -h` you can see all available flags.

### Running a `client` instance :post_office:

In your local command line, just run this:

```bash
./callbacker -c domain.ltd:2424
```

on first run of callbacker you must choose your username:

```bash
$ ./callbacker -c domain.ltd:2424
Enter your username: 
```

with submitting username, you can view all options be pressing `h`:

```bash
$ ./callbacker -c domain.ltd:2424
Enter your username: meysampg
Please select a command (Press h for help): h
Use this commands:
    a (or add)	Add a route
    l (or list)	Print route table
    t (or truncate)	Truncate route table and configuration file
    x (or exit)	Exit from program
```

by pressing `a` you can add a local route to callbacker. Local route is an endpoint on your software which wanna receive
information (for example from a github webhook or after logging in on google):

```bash
$ ./callbacker -c domain.ltd:2424
Please select a command (Press h for help): a

In this section you can add a local callback url. After adding the route, you
will recieve a unique callback url which you can use it as a valid callback.
Enter route: http://mysite.mpg/api/v1/my_endpoint
```

Now by pressing `l` you can view route and assigned callback to it:

```bash
Please select a command (Press h for help): l

 #	Route					Callback
 --	-----					--------
 1	http://mysite.mpg/api/v1/my_endpoint 	http://domain.ltd/callback/b5b1c30a-57e2-428e-861e-813d9a92fc0f
```

Now you can use `http://domain.ltd/callback/b5b1c30a-57e2-428e-861e-813d9a92fc0f` as a valid URL and pass it to origin (
e.g. github webhook, ...). Each time a callback has been called to this URL, all information of request will be sent to
your running client and a callback will be simulated.

## Contribution :love_letter:

Fork from this project :fork_and_knife:, do your modification and send a pull request!

---------
Made with love in [YasnaTeam](https://yasna.team) :heart: