## Nanog 90 Golang for network engineers

This is a repo for the Golang for network engineers that was presented within Charlotte, NC for NANOG90.

### Start the lab
```
sudo containerlab -t clab.yaml deploy
```

The result should look as follows
```
+---+---------------------+--------------+----------------+------+---------+----------------+----------------------+
| # |        Name         | Container ID |     Image      | Kind |  State  |  IPv4 Address  |     IPv6 Address     |
+---+---------------------+--------------+----------------+------+---------+----------------+----------------------+
| 1 | clab-nanog-90-ceos1 | fc113e784a6d | ceoslab:4.31.1 | ceos | running | 172.20.20.2/24 | 2001:172:20:20::2/64 |
| 2 | clab-nanog-90-ceos2 | b1f3c3fd2614 | ceoslab:4.31.1 | ceos | running | 172.20.20.3/24 | 2001:172:20:20::3/64 |
+---+---------------------+--------------+----------------+------+---------+----------------+----------------------+
```

Within the root of the directory lies the main.go file.  This has very basic eapi commands to r1.

```go
go run main.go
```

Within the examples directory there are two examples.
- gnmi
- restclient

The gnmi directory has a streaming gNMI example.

```
cd examples/gnmi
```

```go
go run main.go

[2024-02-08T01:46:56.035724102Z] (172.20.20.2:6030) Update /interfaces/interface[name=Management0]/state/counters/in-octets = 20333
[2024-02-08T01:46:56.035724102Z] (172.20.20.2:6030) Update /interfaces/interface[name=Management0]/state/counters/in-unicast-pkts = 156
[2024-02-08T01:46:56.035724102Z] (172.20.20.2:6030) Update /interfaces/interface[name=Management0]/state/counters/out-octets = 37337
[2024-02-08T01:46:56.035724102Z] (172.20.20.2:6030) Update /interfaces/interface[name=Management0]/state/counters/out-unicast-pkts = 104
[2024-02-08T01:46:56.035724102Z] (172.20.20.2:6030) Update /interfaces/interface[name=Management0]/state/counters/in-pkts = 156
[2024-02-08T01:46:56.035724102Z] (172.20.20.2:6030) Update /interfaces/interface[name=Management0]/state/counters/out-pkts = 104
```

Within the restclient folder it is a small rest api get to the icanhazdadjokes api.

```
cd examples/restclient
```

```go
go run main.go
This joke is hilairous: It was raining cats and dogs the other day. I almost stepped in a poodle.
```

```go
go run main.go
This joke is hilairous: Dad I’m hungry’ … ‘Hi hungry I’m dad
```

This will give random responses from the api for each time. 

