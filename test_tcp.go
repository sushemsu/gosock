package main
import (
    "fmt"
    "net"
    "time"
    "os"
);
func helpcontext(error,bin string){
	fmt.Printf("%s\n%s hostname [port]\n",error, bin);
	os.Exit(2);
};
func main(){
	port:="80";
        if len(os.Args) == 1 {
		helpcontext("Expected host/ip", os.Args[0]);
	}
        if len(os.Args) == 3  {
		port=os.Args[2];
	}
        start := time.Now();
        _, error := net.Dial("tcp", os.Args[1]+":"+port);
        spent := time.Now();
	if error != nil {
		fmt.Printf("error: %s\n", error);
		os.Exit(1);
	}
	fmt.Printf("Connect time: %d\n", spent.Sub(start));
};
