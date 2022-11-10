package main

// redis subscriber will recieve at most 800k message per second
// kafka with 3 cluster will do 2000k and rabit 1400k
// so we can replace it or use redis cluster
