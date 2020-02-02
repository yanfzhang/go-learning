/**
The Go standard library comes with excellent support for HTTP clients and servers in the net/http package.
In this example we’ll use it to issue simple HTTP requests.
*/
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Issue an HTTP GET request to a server.
	// http.Get is a convenient shortcut around creating an http.Client object and calling its Get method; it uses the http.DefaultClient object which has useful default settings.
	resp, err := http.Get("http://www.bing.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// output

/**
Response status: 200 OK
<!DOCTYPE html><html lang="zh"><script type="text/javascript" >//<![CDATA[
si_ST=new Date
//]]></script><head><link id="bgLink" rel="preload" href="/th?id=OHR.OberweissbacherBergbahn_ZH-CN1289048050_1920x1080.jpg&amp;rf=LaDigue_1920x1080.jpg&amp;pid=hp" as="image" /><link rel="preload" href="/sa/simg/hpc27.png" as="image" /><meta content="text/html; charset=utf-8" http-equiv="content-type"/><script type="text/javascript">//<![CDATA[
/*!DisableJavascriptProfiler*
0;
*/
