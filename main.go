package main

import (
	"fmt"
	"hotels/service"
	_ "hotels/service"
	"os"
)

func main() {

	router := service.SetupApi()
	// 2023-03-02 22:30:58.536929
	// 2023-03-02 22:35:45.641686257 +0100 WAT m=+4.084937156
	// 2023-03-02T22:35:45.000+07:00
	// 2023-03-03T22:35:45.000+07:00

	//  2023-03-02 23:30:08.0000
	// 2023-03-03 23:30:08.0000
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
