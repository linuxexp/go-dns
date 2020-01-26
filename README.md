- Simple Authoritative DNS server in Go

```
$ dig example.com @localhost -p1023

; <<>> DiG 9.10.6 <<>> example.com -t MX @localhost -p1023
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 31298
;; flags: qr rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;example.com.                   IN      MX

;; ANSWER SECTION:
example.com.            0       IN      A       127.0.0.1

;; Query time: 0 msec
;; SERVER: ::1#1023(::1)
;; WHEN: Sun Jan 26 08:50:34 IST 2020
;; MSG SIZE  rcvd: 56

```

**TODOs**
- add zone file
- read records from zone file
