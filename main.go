package main

import (
	install_packages "hypr_conf/src/install_package"
	"hypr_conf/src/read_packages"

	"fmt"
)

func main() {
	fmt.Println("Установка пакетов....")
	install_packages.InstallPackages(read_packages.ReadPackages("packages/package_list.txt"))
	fmt.Println("Пакеты установлены")
}
